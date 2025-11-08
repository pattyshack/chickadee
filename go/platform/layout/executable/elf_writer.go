package executable

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pattyshack/chickadee/platform/layout"
)

type ElfWriter struct {
	Parameters

	Header Elf64Header

	ProgramHeader []Elf64ProgramHeaderEntry

	layout.ExecutableImage

	TextIndex         uint16
	InitIndex         uint16
	ReadOnlyDataIndex uint16
	DataIndex         uint16
	BSSIndex          uint16

	// NOTE: The symbol table is implicitly generated from []*layout.Symbol
	// during writing.
	SymbolTableStart int64
	SymbolTableIndex uint16

	StringTableStart int64
	StringTableIndex uint16
	StringTable      ElfStringTable

	SectionStringTableStart int64
	SectionStringTableIndex uint16
	SectionStringTable      ElfStringTable

	SectionHeaderStart int64
	SectionHeader      []Elf64SectionHeaderEntry
}

func NewElfWriter(
	parameters Parameters,
	image layout.ExecutableImage,
) (
	ElfWriter,
	error,
) {
	dataEncoding := ELFDATA2LSB
	if parameters.ByteOrder == binary.BigEndian {
		dataEncoding = ELFDATA2MSB
	}

	sectionStringTableIndex := uint16(1)
	entryPoint := parameters.VirtualAddressStart + uint64(image.EntryPoint)

	// NOTE: We'll assign string tables and symbol table to fixed section header
	// locations at the front of the list to simplify file generation.
	//
	// The section header order:
	//   (NULL) - required by elf specification
	//   .shstrtab
	//   .strtab
	//   .symtab
	//   .text
	//   .init
	//   .rodata
	//   .data
	//   .bss
	//
	// The program header order:
	//   PHDR             (r-- memory page aligned)
	//   LOAD (header)    (r-- memory page aligned)
	//   LOAD .text .init (r-e memory page aligned)
	//   LOAD .rodata     (r-- memory page aligned)
	//   LOAD .data .bss  (rw- memory page aligned)
	//   GNU_STACK        (rw- 8-byte aligned)
	//
	// NOTE: the GNU_STACK entry tells linux that the stack is not executable
	//
	// TODO INTERP for dynamic linking / PIC executable
	writer := ElfWriter{
		Parameters: parameters,

		// NOTE: SectionHeaderStart, NumProgramHeaderEntries and
		// NumSectionHeaderEntries are defer populated.
		Header: Elf64Header{
			ElfIdentifier: ElfIdentifier{
				Magic:              [4]byte{0x7f, 'E', 'L', 'F'},
				Class:              2, // ELFCLASS64 (we don't support elf32 format)
				DataEncoding:       dataEncoding,
				IdentifierVersion:  1, // EI_CURRENT (only valid value)
				OperatingSystemABI: 0, // ELFOSABI_NONE (aka System V ABI)
				ABIVersion:         0, // (only valid value)
			},
			// TODO switch to position independent executable (ET_DYN)
			FileType:                ET_EXEC,
			MachineArchitecture:     parameters.ElfMachineArchitecture,
			FormatVersion:           1, // EV_CURRENT (only valid value)
			EntryPointAddress:       entryPoint,
			ProgramHeaderOffset:     Elf64HeaderSize,
			ArchitectureFlags:       0, // (only valid value)
			ElfHeaderSize:           Elf64HeaderSize,
			ProgramHeaderEntrySize:  Elf64ProgramHeaderEntrySize,
			SectionHeaderEntrySize:  Elf64SectionHeaderEntrySize,
			SectionStringTableIndex: sectionStringTableIndex,
		},

		ProgramHeader:    make([]Elf64ProgramHeaderEntry, 3, 6),
		ExecutableImage:  image,
		SymbolTableIndex: 3,
		StringTableIndex: 2,
		StringTable: NewElfStringTableFromSymbols(
			image.Definitions.Symbols),
		SectionStringTableIndex: sectionStringTableIndex,
		SectionStringTable:      NewElfStringTable(),
		SectionHeader:           make([]Elf64SectionHeaderEntry, 4, 10),
	}

	writer.addExecutableSegmentHeaderEntries()
	writer.maybeAddReadOnlySegmentHeaderEntries()
	writer.maybeAddReadWriteSegmentHeaderEntries()
	writer.updateMetadataHeaderEntries()

	writer.addSymbolTableHeaderEntry()
	writer.addStringTableHeaderEntry()
	writer.addSectionStringTableHeaderEntry()

	return writer, nil
}

// LOAD .text .init
func (elf *ElfWriter) addExecutableSegmentHeaderEntries() {
	start := uint64(elf.ExecutableSegmentStart)
	startAddress := elf.VirtualAddressStart + start
	fileSize := uint64(elf.Text.Size + elf.Init.Size)

	// NOTE: This may not be the final start location if there are additional
	// executable image file segments.
	elf.SymbolTableStart = int64(start + fileSize)

	elf.ProgramHeader[2] = Elf64ProgramHeaderEntry{
		Type:            PT_LOAD,
		Flags:           0b101, // r-e
		ContentOffset:   start,
		VirtualAddress:  startAddress,
		PhysicalAddress: startAddress,
		FileImageSize:   fileSize,
		MemoryImageSize: fileSize,
		Alignment:       uint64(elf.MemoryPageSize),
	}

	nameIdx, _ := elf.SectionStringTable.MaybeInsert(".text")
	elf.TextIndex = uint16(len(elf.SectionHeader))
	elf.SectionHeader = append(
		elf.SectionHeader,
		Elf64SectionHeaderEntry{
			NameIndex:        nameIdx,
			Type:             SHT_PROGBITS,
			Flags:            SHF_EXECINSTR | SHF_ALLOC,
			Address:          startAddress,
			Offset:           start,
			Size:             uint64(elf.Text.Size),
			Link:             0,
			Info:             0,
			AddressAlignment: uint64(elf.RegisterAlignment),
			EntrySize:        0,
		})

	if elf.Init.Size > 0 {
		start += uint64(elf.Text.Size)
		startAddress += uint64(elf.Text.Size)
		nameIdx, _ = elf.SectionStringTable.MaybeInsert(".init")
		elf.InitIndex = uint16(len(elf.SectionHeader))
		elf.SectionHeader = append(
			elf.SectionHeader,
			Elf64SectionHeaderEntry{
				NameIndex:        nameIdx,
				Type:             SHT_PROGBITS,
				Flags:            SHF_EXECINSTR | SHF_ALLOC,
				Address:          startAddress,
				Offset:           start,
				Size:             uint64(elf.Init.Size),
				Link:             0,
				Info:             0,
				AddressAlignment: uint64(elf.RegisterAlignment),
				EntrySize:        0,
			})
	}
}

// LOAD .rodata
func (elf *ElfWriter) maybeAddReadOnlySegmentHeaderEntries() {
	if elf.ReadOnlyData.Size == 0 {
		return
	}

	start := uint64(elf.ReadOnlySegmentStart)
	startAddress := elf.VirtualAddressStart + start
	fileSize := uint64(elf.ReadOnlyData.Size)

	// NOTE: This may not be the final start location if there are additional
	// executable image file segments.
	elf.SymbolTableStart = int64(start + fileSize)

	elf.ProgramHeader = append(
		elf.ProgramHeader,
		Elf64ProgramHeaderEntry{
			Type:            PT_LOAD,
			Flags:           0b100, // r--
			ContentOffset:   start,
			VirtualAddress:  startAddress,
			PhysicalAddress: startAddress,
			FileImageSize:   fileSize,
			MemoryImageSize: fileSize,
			Alignment:       uint64(elf.MemoryPageSize),
		})

	nameIdx, _ := elf.SectionStringTable.MaybeInsert(".rodata")
	elf.ReadOnlyDataIndex = uint16(len(elf.SectionHeader))
	elf.SectionHeader = append(
		elf.SectionHeader,
		Elf64SectionHeaderEntry{
			NameIndex:        nameIdx,
			Type:             SHT_PROGBITS,
			Flags:            SHF_ALLOC,
			Address:          startAddress,
			Offset:           start,
			Size:             fileSize,
			Link:             0,
			Info:             0,
			AddressAlignment: uint64(elf.RegisterAlignment),
			EntrySize:        0,
		})
}

// LOAD .data .bss
func (elf *ElfWriter) maybeAddReadWriteSegmentHeaderEntries() {
	if elf.Data.Size == 0 && elf.BSSSize == 0 {
		return
	}

	start := uint64(elf.ReadWriteSegmentStart)
	startAddress := elf.VirtualAddressStart + start
	fileSize := uint64(elf.Data.Size)

	// NOTE: This is the symbol table's final start location since there are no
	// more executable image file segments.
	elf.SymbolTableStart = int64(start + fileSize)

	elf.ProgramHeader = append(
		elf.ProgramHeader,
		Elf64ProgramHeaderEntry{
			Type:            PT_LOAD,
			Flags:           0b110, // rw-
			ContentOffset:   start,
			VirtualAddress:  startAddress,
			PhysicalAddress: startAddress,
			FileImageSize:   fileSize,
			MemoryImageSize: uint64(elf.Data.Size + elf.BSSSize),
			Alignment:       uint64(elf.MemoryPageSize),
		})

	if elf.Data.Size > 0 {
		nameIdx, _ := elf.SectionStringTable.MaybeInsert(".data")
		elf.DataIndex = uint16(len(elf.SectionHeader))
		elf.SectionHeader = append(
			elf.SectionHeader,
			Elf64SectionHeaderEntry{
				NameIndex:        nameIdx,
				Type:             SHT_PROGBITS,
				Flags:            SHF_ALLOC | SHF_WRITE,
				Address:          startAddress,
				Offset:           start,
				Size:             fileSize,
				Link:             0,
				Info:             0,
				AddressAlignment: uint64(elf.RegisterAlignment),
				EntrySize:        0,
			})
	}

	if elf.BSSSize > 0 {
		start += fileSize
		startAddress += fileSize
		nameIdx, _ := elf.SectionStringTable.MaybeInsert(".bss")
		elf.BSSIndex = uint16(len(elf.SectionHeader))
		elf.SectionHeader = append(
			elf.SectionHeader,
			Elf64SectionHeaderEntry{
				NameIndex:        nameIdx,
				Type:             SHT_NOBITS,
				Flags:            SHF_ALLOC | SHF_WRITE,
				Address:          startAddress,
				Offset:           start,
				Size:             uint64(elf.BSSSize),
				Link:             0,
				Info:             0,
				AddressAlignment: uint64(elf.RegisterAlignment),
				EntrySize:        0,
			})
	}
}

// PHDR
// LOAD (elf header page)
// ...
// GNU_STACK
func (elf *ElfWriter) updateMetadataHeaderEntries() {
	// GNU_STACK
	elf.ProgramHeader = append(
		elf.ProgramHeader,
		Elf64ProgramHeaderEntry{
			Type:            PT_GNU_STACK,
			Flags:           0b110, // rw-
			ContentOffset:   0,
			VirtualAddress:  0,
			PhysicalAddress: 0,
			FileImageSize:   0,
			MemoryImageSize: 0,
			Alignment:       8,
		})

	// LOAD (elf header page)
	memoryPageSize := uint64(elf.MemoryPageSize)
	elf.ProgramHeader[1] = Elf64ProgramHeaderEntry{
		Type:            PT_LOAD,
		Flags:           0b100, // r--
		ContentOffset:   0,
		VirtualAddress:  elf.VirtualAddressStart,
		PhysicalAddress: elf.VirtualAddressStart,
		FileImageSize:   memoryPageSize,
		MemoryImageSize: memoryPageSize,
		Alignment:       memoryPageSize,
	}

	// PHDR
	size := uint64(len(elf.ProgramHeader)) * Elf64ProgramHeaderEntrySize
	elf.ProgramHeader[0] = Elf64ProgramHeaderEntry{
		Type:            PT_PHDR,
		Flags:           0b100, // r--
		ContentOffset:   Elf64HeaderSize,
		VirtualAddress:  elf.VirtualAddressStart + Elf64HeaderSize,
		PhysicalAddress: elf.VirtualAddressStart + Elf64HeaderSize,
		FileImageSize:   size,
		MemoryImageSize: size,
		Alignment:       memoryPageSize,
	}

	elf.Header.NumProgramHeaderEntries = uint16(len(elf.ProgramHeader))
	elf.Header.NumSectionHeaderEntries = uint16(len(elf.SectionHeader))
}

func (elf *ElfWriter) addSymbolTableHeaderEntry() {
	nameIndex, _ := elf.SectionStringTable.MaybeInsert(".symtab")
	size := int64(len(elf.Definitions.Symbols)) * Elf64SymbolEntrySize
	elf.SectionHeader[int(elf.SymbolTableIndex)] = Elf64SectionHeaderEntry{
		NameIndex: nameIndex,
		Type:      SHT_SYMTAB,
		Flags:     0,
		Address:   0,
		Offset:    uint64(elf.SymbolTableStart),
		Size:      uint64(size),
		Link:      uint32(elf.StringTableIndex),
		// Reference: Elf Book III Figure 1-1. sh_link and sh_info Interpretation
		//
		// One greater than the symbol table index of the last local symbol
		// (binding STB_LOCAL).
		//
		// This is always zero since we only generate global symbols.
		Info:             0,
		AddressAlignment: 8,
		EntrySize:        Elf64SymbolEntrySize,
	}

	elf.StringTableStart = elf.SymbolTableStart + size
}

func (elf *ElfWriter) addStringTableHeaderEntry() {
	nameIndex, _ := elf.SectionStringTable.MaybeInsert(".strtab")
	size := elf.StringTable.Size
	elf.SectionHeader[int(elf.StringTableIndex)] = Elf64SectionHeaderEntry{
		NameIndex:        nameIndex,
		Type:             SHT_STRTAB,
		Flags:            0,
		Address:          0,
		Offset:           uint64(elf.StringTableStart),
		Size:             uint64(size),
		Link:             0,
		Info:             0,
		AddressAlignment: 1,
		EntrySize:        0,
	}

	elf.SectionStringTableStart = elf.StringTableStart + int64(size)
}

func (elf *ElfWriter) addSectionStringTableHeaderEntry() {
	nameIndex, _ := elf.SectionStringTable.MaybeInsert(".shstrtab")
	size := elf.SectionStringTable.Size
	elf.SectionHeader[int(elf.SectionStringTableIndex)] = Elf64SectionHeaderEntry{
		NameIndex:        nameIndex,
		Type:             SHT_STRTAB,
		Flags:            0,
		Address:          0,
		Offset:           uint64(elf.SectionStringTableStart),
		Size:             uint64(size),
		Link:             0,
		Info:             0,
		AddressAlignment: 1,
		EntrySize:        0,
	}

	sectionHeaderStart := elf.SectionStringTableStart + int64(size)
	elf.SectionHeaderStart = ((sectionHeaderStart + 7) / 8) * 8
	elf.Header.SectionHeaderOffset = uint64(elf.SectionHeaderStart)
}

func (elf ElfWriter) WriteTo(writer io.Writer) (int64, error) {
	numWritten, err := elf.writeHeader(writer)
	if err != nil {
		return 0, err
	}

	numWritten, err = elf.writeExecutableImage(writer, numWritten)
	if err != nil {
		return 0, err
	}

	numWritten, err = elf.writeSymbolTable(writer, numWritten)
	if err != nil {
		return 0, err
	}

	numWritten, err = elf.writeStringTable(writer, numWritten)
	if err != nil {
		return 0, err
	}

	numWritten, err = elf.writeSectionStringTable(writer, numWritten)
	if err != nil {
		return 0, err
	}

	// NOTE: By convention, elf section header are at the end of the file.
	numWritten, err = elf.writeSectionHeader(writer, numWritten)
	if err != nil {
		return 0, err
	}

	return numWritten, nil
}

func (elf ElfWriter) writeHeader(writer io.Writer) (int64, error) {
	headerPage := make([]byte, elf.MemoryPageSize)
	remaining := headerPage

	n, err := binary.Encode(remaining, elf.ByteOrder, elf.Header)
	if err != nil || n != Elf64HeaderSize {
		panic("should never happen")
	}
	remaining = remaining[n:]

	for _, entry := range elf.ProgramHeader {
		n, err := binary.Encode(remaining, elf.ByteOrder, entry)
		if err != nil || n != Elf64ProgramHeaderEntrySize {
			panic("should never happen")
		}
		remaining = remaining[n:]
	}

	n, err = writer.Write(headerPage)
	if err != nil {
		return 0, fmt.Errorf("failed to write elf header: %w", err)
	}

	return int64(n), nil
}

func (elf ElfWriter) writeExecutableImage(
	writer io.Writer,
	numWritten int64,
) (
	int64,
	error,
) {
	padding := elf.ExecutableSegmentStart - numWritten
	if padding > 0 {
		n, err := writer.Write(make([]byte, padding))
		if err != nil {
			return 0, fmt.Errorf("failed to write executable image: %w", err)
		}
		numWritten += int64(n)
	}

	for _, chunk := range elf.Text.DataChunks {
		n, err := writer.Write(chunk)
		if err != nil {
			return 0, fmt.Errorf("failed to write executable image: %w", err)
		}
		numWritten += int64(n)
	}

	for _, chunk := range elf.Init.DataChunks {
		n, err := writer.Write(chunk)
		if err != nil {
			return 0, fmt.Errorf("failed to write executable image: %w", err)
		}
		numWritten += int64(n)
	}

	if elf.ReadOnlyData.Size > 0 {
		padding = elf.ReadOnlySegmentStart - numWritten
		if padding > 0 {
			n, err := writer.Write(make([]byte, padding))
			if err != nil {
				return 0, fmt.Errorf("failed to write executable image: %w", err)
			}
			numWritten += int64(n)
		}

		for _, chunk := range elf.ReadOnlyData.DataChunks {
			n, err := writer.Write(chunk)
			if err != nil {
				return 0, fmt.Errorf("failed to write executable image: %w", err)
			}
			numWritten += int64(n)
		}
	}

	if elf.Data.Size > 0 || elf.BSSSize > 0 {
		padding := elf.ReadWriteSegmentStart - numWritten
		if padding > 0 {
			n, err := writer.Write(make([]byte, padding))
			if err != nil {
				return 0, fmt.Errorf("failed to write executable image: %w", err)
			}
			numWritten += int64(n)
		}

		for _, chunk := range elf.Data.DataChunks {
			n, err := writer.Write(chunk)
			if err != nil {
				return 0, fmt.Errorf("failed to write executable image: %w", err)
			}
			numWritten += int64(n)
		}
	}

	return numWritten, nil
}

func (elf ElfWriter) convertSymbol(
	symbol *layout.Symbol,
) (
	Elf64SymbolEntry,
	error,
) {
	nameIdx, ok := elf.StringTable.Indices[symbol.Name]
	if !ok {
		panic("should never happen")
	}

	var symbolInfo byte
	switch symbol.Kind {
	case layout.FunctionKind:
		symbolInfo = functionSymbolInfo
	case layout.ObjectKind:
		symbolInfo = objectSymbolInfo
	default:
		return Elf64SymbolEntry{}, fmt.Errorf(
			"unsupported symbol kind (%s)",
			symbol.Kind)
	}

	sectionIdx := uint16(0)
	switch symbol.Section {
	case layout.UnknownSection:
		// do nothing
	case layout.TextSection:
		sectionIdx = elf.TextIndex
	case layout.InitSection:
		sectionIdx = elf.InitIndex
	case layout.ReadOnlyDataSection:
		sectionIdx = elf.ReadOnlyDataIndex
	case layout.ReadWriteDataSection:
		sectionIdx = elf.DataIndex
	case layout.BSSSection:
		sectionIdx = elf.BSSIndex
	default:
		return Elf64SymbolEntry{}, fmt.Errorf(
			"unsupported symbol section (%s)",
			symbol.Section)
	}

	return Elf64SymbolEntry{
		NameIndex:    nameIdx,
		Info:         symbolInfo,
		Visibility:   STV_DEFAULT,
		SectionIndex: sectionIdx,
		Value:        elf.VirtualAddressStart + uint64(symbol.Offset),
		Size:         uint64(symbol.Size),
	}, nil
}

func (elf ElfWriter) writeSymbolTable(
	writer io.Writer,
	numWritten int64,
) (
	int64,
	error,
) {
	padding := elf.SymbolTableStart - numWritten
	if padding > 0 {
		n, err := writer.Write(make([]byte, padding))
		if err != nil {
			return 0, fmt.Errorf("failed to write symbols: %w", err)
		}
		numWritten += int64(n)
	}

	buffer := make([]byte, Elf64SymbolEntrySize)
	for _, symbol := range elf.Definitions.Symbols {
		elfSymbol, err := elf.convertSymbol(symbol)
		if err != nil {
			return 0, err
		}

		n, err := binary.Encode(buffer, elf.ByteOrder, elfSymbol)
		if err != nil || n != Elf64SymbolEntrySize {
			panic("should never happen")
		}

		n, err = writer.Write(buffer)
		if err != nil {
			return 0, fmt.Errorf("failed to write symbols: %w", err)
		}
		numWritten += int64(n)
	}

	return numWritten, nil
}

func (elf ElfWriter) writeStringTable(
	writer io.Writer,
	numWritten int64,
) (
	int64,
	error,
) {
	padding := elf.StringTableStart - numWritten
	if padding > 0 {
		n, err := writer.Write(make([]byte, padding))
		if err != nil {
			return 0, fmt.Errorf("failed to write string table: %w", err)
		}
		numWritten += int64(n)
	}

	n, err := elf.StringTable.WriteTo(writer)
	if err != nil {
		return 0, fmt.Errorf("failed to write string table: %w", err)
	}
	numWritten += n

	return numWritten, nil
}

func (elf ElfWriter) writeSectionStringTable(
	writer io.Writer,
	numWritten int64,
) (
	int64,
	error,
) {
	padding := elf.SectionStringTableStart - numWritten
	if padding > 0 {
		n, err := writer.Write(make([]byte, padding))
		if err != nil {
			return 0, fmt.Errorf("failed to write section string table: %w", err)
		}
		numWritten += int64(n)
	}

	n, err := elf.SectionStringTable.WriteTo(writer)
	if err != nil {
		return 0, fmt.Errorf("failed to write section string table: %w", err)
	}
	numWritten += n

	return numWritten, nil
}

func (elf ElfWriter) writeSectionHeader(
	writer io.Writer,
	numWritten int64,
) (
	int64,
	error,
) {
	padding := elf.SectionHeaderStart - numWritten
	if padding > 0 {
		n, err := writer.Write(make([]byte, padding))
		if err != nil {
			return 0, fmt.Errorf("failed to write section header: %w", err)
		}
		numWritten += int64(n)
	}

	buffer := make([]byte, Elf64SectionHeaderEntrySize)
	for _, entry := range elf.SectionHeader {
		n, err := binary.Encode(buffer, elf.ByteOrder, entry)
		if err != nil || n != Elf64SectionHeaderEntrySize {
			panic("should never happen")
		}

		n, err = writer.Write(buffer)
		if err != nil {
			return 0, fmt.Errorf("failed to write section header: %w", err)
		}
		numWritten += int64(n)
	}

	return numWritten, nil
}
