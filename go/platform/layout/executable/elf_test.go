package executable_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"

	"github.com/pattyshack/chickadee/amd64"
	"github.com/pattyshack/chickadee/platform/layout"
	. "github.com/pattyshack/chickadee/platform/layout/executable"
)

type ElfSuite struct{}

func TestElf(t *testing.T) {
	suite.RunTests(t, &ElfSuite{})
}

func (ElfSuite) TestStringTable(t *testing.T) {
	table := NewElfStringTable()

	idx, inserted := table.MaybeInsert("")
	expect.False(t, inserted)
	expect.Equal(t, 0, idx)

	idx, inserted = table.MaybeInsert("foo")
	expect.True(t, inserted)
	expect.Equal(t, 1, idx)

	idx, inserted = table.MaybeInsert("asdf")
	expect.True(t, inserted)
	expect.Equal(t, 5, idx)

	idx, inserted = table.MaybeInsert("foo")
	expect.False(t, inserted)
	expect.Equal(t, 1, idx)

	idx, inserted = table.MaybeInsert("012345")
	expect.True(t, inserted)
	expect.Equal(t, 10, idx)

	expect.Equal(t, 17, table.Size)

	buffer := &bytes.Buffer{}

	n, err := table.WriteTo(buffer)
	expect.Nil(t, err)
	expect.Equal(t, 17, n)
	expect.Equal(t, "\x00foo\x00asdf\x00012345\x00", string(buffer.Bytes()))
}

func (ElfSuite) TestStringTableFromSymbols(t *testing.T) {
	table := NewElfStringTableFromSymbols(
		[]*layout.Symbol{
			{
				Name: "foo",
			},
			{
				Name: "asdf",
			},
			{
				Name: "012345",
			},
		})

	expect.Equal(t, 17, table.Size)
	expect.Equal(
		t,
		map[string]uint32{"": 0, "foo": 1, "asdf": 5, "012345": 10},
		table.Indices)
	expect.Equal(t, []string{"foo", "asdf", "012345"}, table.Entries)

	buffer := &bytes.Buffer{}

	n, err := table.WriteTo(buffer)
	expect.Nil(t, err)
	expect.Equal(t, 17, n)
	expect.Equal(t, "\x00foo\x00asdf\x00012345\x00", string(buffer.Bytes()))
}

// This is approximately:
//
// .global start
//
// .section .text
//
// add:
//
//	movq rovalue(%rip), %rax
//	movq value(%rip), %rdi
//	addq %rdi, %rax
//	ret
//
// start:
//
//	call _init
//	movq $60, %rax # exit sys call id
//	movq exit_code(%rip), %rdi
//	syscall
//
// save_exit_code:
//
//	movq %rax, exit_code(%rip)
//	ret
//
// _init:
//
//	call add
//	call save_exit_code
//	ret
//
// .section .rodata
//
// rovalue: .long 31
//
// .section .data  <- uses .bss instead
//
// value: .long 11
//
// exit_code: .long 0
func (ElfSuite) newWriter(t *testing.T) ElfWriter {
	builder := layout.NewObjectFileBuilder()

	builder.Text.AppendData(
		[]byte{
			0x48, 0x8b, 0x05, 0, 0, 0, 0, // movq rovalue(%rip), %rax - (relocation)
			0x48, 0x8b, 0x3d, 0, 0, 0, 0, // movq value(%rip), %rax - (relocation)
			0x48, 0x01, 0xf8, // add %rdi, %rax
			0xc3, // ret
		},
		layout.Definitions{
			Symbols: []*layout.Symbol{
				{
					Kind:    layout.FunctionKind,
					Section: layout.TextSection,
					Name:    "add",
					Offset:  0,
					Size:    18,
				},
			},
		},
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "rovalue",
					Offset: 3,
				},
				{
					Name:   "value",
					Offset: 10,
				},
			},
		})

	builder.Text.AppendData(
		[]byte{
			0xe8, 0, 0, 0, 0, // call init - (relocation)
			0x48, 0xc7, 0xc0, 0x3c, 0, 0, 0, // movq $60, %rax
			0x48, 0x8b, 0x3d, 0, 0, 0, 0, // movq exit_code(%rip), %rdi - (relocation)
			0x0f, 0x05, // syscall
		},
		layout.Definitions{
			Symbols: []*layout.Symbol{
				{
					Kind:    layout.FunctionKind,
					Section: layout.TextSection,
					Name:    "start",
					Offset:  0,
					Size:    21,
				},
			},
		},
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "_init",
					Offset: 1,
				},
				{
					Name:   "exit_code",
					Offset: 15,
				},
			},
		})

	builder.Text.AppendData(
		[]byte{
			0x48, 0x89, 0x05, 0, 0, 0, 0, // movq %rax, exit_code(%rip) - (relocation)
			0xc3, // ret
		},
		layout.Definitions{
			Symbols: []*layout.Symbol{
				{
					Kind:    layout.FunctionKind,
					Section: layout.TextSection,
					Name:    "save_exit_code",
					Offset:  0,
					Size:    8,
				},
			},
		},
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "exit_code",
					Offset: 3,
				},
			},
		})

	// NOTE: the _init function symbol and the init epilogue []byte{0xc3} will
	// be added to the executable image,
	builder.Init.AppendData(
		[]byte{
			0xe8, 0, 0, 0, 0, // call add (relocation)
			0xe8, 0, 0, 0, 0, // call save_exit_code (relocation)
		},
		layout.Definitions{},
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "add",
					Offset: 1,
				},
				{
					Name:   "save_exit_code",
					Offset: 6,
				},
			},
		})

	builder.ReadOnlyData.AppendData(
		[]byte{31, 0, 0, 0, 0, 0, 0, 0},
		layout.Definitions{
			Symbols: []*layout.Symbol{
				{
					Kind:    layout.ObjectKind,
					Section: layout.ReadOnlyDataSection,
					Name:    "rovalue",
					Offset:  0,
					Size:    8,
				},
			},
		},
		layout.Relocations{})

	builder.Data.AppendData(
		[]byte{11, 0, 0, 0, 0, 0, 0, 0},
		layout.Definitions{
			Symbols: []*layout.Symbol{
				{
					Kind:    layout.ObjectKind,
					Section: layout.ReadWriteDataSection,
					Name:    "value",
					Offset:  0,
					Size:    8,
				},
			},
		},
		layout.Relocations{})

	builder.BSS.AppendObject("exit_code", 8)

	file, err := builder.Finalize(amd64.Linux.Layout)
	expect.Nil(t, err)

	image, err := file.ToExecutableImage(amd64.Linux.Layout, "start")
	expect.Nil(t, err)

	writer, err := NewElfWriter(amd64.Linux.ExecutableFormat, image)
	expect.Nil(t, err)

	return writer
}

func (s ElfSuite) TestWriterInitialization(t *testing.T) {
	writer := s.newWriter(t)

	expect.Equal(t, ELFDATA2LSB, writer.Header.DataEncoding)

	expect.Equal(
		t,
		writer.VirtualAddressStart+uint64(writer.ExecutableImage.EntryPoint),
		writer.Header.EntryPointAddress)

	// check section header

	expect.Equal(t, 9, writer.Header.NumSectionHeaderEntries)
	expect.Equal(t, 9, len(writer.SectionHeader))

	expect.Equal(t, 1, writer.SectionStringTableIndex)
	expect.Equal(t, 2, writer.StringTableIndex)
	expect.Equal(t, 3, writer.SymbolTableIndex)
	expect.Equal(t, 4, writer.TextIndex)
	expect.Equal(t, 5, writer.InitIndex)
	expect.Equal(t, 6, writer.ReadOnlyDataIndex)
	expect.Equal(t, 7, writer.DataIndex)
	expect.Equal(t, 8, writer.BSSIndex)

	prevEnd := uint64(writer.MemoryPageSize)
	// .text .init .rodata .data .bss .symtab .strtab .shstrtab
	for _, idx := range []int{4, 5, 6, 7, 8, 3, 2, 1} {
		entry := writer.SectionHeader[idx]
		expect.True(t, entry.Offset >= prevEnd)
		expect.True(t, entry.Size > 0)

		if idx > 3 { // .text .init .rodata .data .bss
			expect.Equal(
				t,
				writer.VirtualAddressStart+entry.Offset,
				entry.Address)
		} else { // .symtab .strtab .shstrtab
			expect.Equal(t, 0, entry.Address)
		}

		if idx != 8 { // .bss (bss doesn't take up file space)
			prevEnd = entry.Offset + entry.Size
		}
	}

	shstrtabEnd := writer.SectionStringTableStart +
		int64(writer.SectionStringTable.Size)
	expect.True(t, writer.SectionHeaderStart >= shstrtabEnd)

	// check program header

	expect.Equal(t, Elf64HeaderSize, writer.Header.ProgramHeaderOffset)
	expect.Equal(t, 6, writer.Header.NumProgramHeaderEntries)
	expect.Equal(t, 6, len(writer.ProgramHeader))

	prevFileEnd := uint64(0)
	prevMemoryEnd := writer.VirtualAddressStart
	for idx, entry := range writer.ProgramHeader {
		if idx == 0 {
			expect.Equal(t, PT_PHDR, entry.Type)
			continue
		} else if idx == 5 { // last
			expect.Equal(t, PT_GNU_STACK, entry.Type)
			continue
		}

		expect.Equal(t, PT_LOAD, entry.Type)
		expect.True(t, entry.ContentOffset >= prevFileEnd)
		expect.True(t, entry.VirtualAddress >= prevMemoryEnd)
		expect.True(t, entry.FileImageSize > 0)
		expect.True(t, entry.MemoryImageSize >= entry.FileImageSize)

		prevFileEnd = entry.ContentOffset + entry.FileImageSize
		prevMemoryEnd = entry.VirtualAddress + entry.MemoryImageSize
	}

	// check section string table

	expect.Equal(
		t,
		[]string{
			".text",
			".init",
			".rodata",
			".data",
			".bss",
			".symtab",
			".strtab",
			".shstrtab",
		},
		writer.SectionStringTable.Entries)

	// check string table

	expect.Equal(
		t,
		[]string{
			"add",
			"start",
			"save_exit_code",
			"_init",
			"rovalue",
			"value",
			"exit_code",
		},
		writer.StringTable.Entries)
}

func (s ElfSuite) TestWrite(t *testing.T) {
	writer := s.newWriter(t)

	buffer := &bytes.Buffer{}

	numWritten, err := writer.WriteTo(buffer)
	expect.Nil(t, err)

	content := buffer.Bytes()
	expect.Equal(t, int64(len(content)), numWritten)

	// check elf header

	header := &Elf64Header{}
	n, err := binary.Decode(content, binary.LittleEndian, header)
	expect.Nil(t, err)
	expect.Equal(t, Elf64HeaderSize, n)

	expect.Equal(t, [4]byte{0x7f, 'E', 'L', 'F'}, header.ElfIdentifier.Magic)
	expect.Equal(t, 6, header.NumProgramHeaderEntries)
	expect.Equal(t, 9, header.NumSectionHeaderEntries)

	// check program header and segment contents

	start := header.ProgramHeaderOffset
	end := start + 6*Elf64ProgramHeaderEntrySize
	programHeaderBytes := content[start:end]
	programHeader := [6]Elf64ProgramHeaderEntry{}
	n, err = binary.Decode(
		programHeaderBytes,
		binary.LittleEndian,
		&programHeader)
	expect.Nil(t, err)
	expect.Equal(t, len(programHeaderBytes), n)

	prevFileEnd := uint64(0)
	prevMemoryEnd := writer.VirtualAddressStart
	for idx, entry := range programHeader {
		if idx == 0 {
			expect.Equal(t, PT_PHDR, entry.Type)
			continue
		} else if idx == 5 { // last
			expect.Equal(t, PT_GNU_STACK, entry.Type)
			continue
		}

		expect.Equal(t, PT_LOAD, entry.Type)
		expect.True(t, entry.ContentOffset >= prevFileEnd)
		expect.True(t, entry.VirtualAddress >= prevMemoryEnd)
		expect.True(t, entry.FileImageSize > 0)

		if idx < 4 {
			expect.Equal(t, entry.MemoryImageSize, entry.FileImageSize)
		} else { // .bss .data
			expect.True(t, entry.MemoryImageSize > entry.FileImageSize)
		}

		start := entry.ContentOffset
		end := start + entry.FileImageSize
		segment := content[start:end]

		switch idx {
		case 1: // metadata segement
			expect.True(t, bytes.HasPrefix(segment, []byte{0x7f, 'E', 'L', 'F'}))
		case 2: // .text .init
			expect.True(t, int64(len(segment)) >= writer.Text.Size)
			textContent := segment[:writer.Text.Size]

			// prefix of: movq rovalue(%rip), %rax
			expect.True(t, bytes.HasPrefix(textContent, []byte{0x48, 0x8b, 0x05}))

			expect.Equal(t, writer.Text.Flatten(), textContent)

			// movq $60, %rax
			expect.True(t, bytes.Contains(
				textContent,
				[]byte{0x48, 0xc7, 0xc0, 0x3c, 0, 0, 0}))

			// prefix of: movq %rax, exit_code(%rip)
			expect.True(t, bytes.Contains(textContent, []byte{0x48, 0x89, 0x05}))

			initContent := segment[writer.Text.Size:]
			expect.True(t, int64(len(initContent)) >= writer.Init.Size)

			expect.True(t, len(initContent) >= 11)
			expect.Equal(t, 0xe8, initContent[0])  // prefix of call
			expect.Equal(t, 0xe8, initContent[5])  // prefix of call
			expect.Equal(t, 0xc3, initContent[10]) // ret from init epilogue

			expect.Equal(t, writer.Init.Flatten(), initContent[:writer.Init.Size])

		case 3: // .rodata
			expect.True(t, bytes.HasPrefix(segment, writer.ReadOnlyData.Flatten()))
		case 4: // .data .bss
			expect.True(t, bytes.HasPrefix(segment, writer.Data.Flatten()))
		}

		prevFileEnd = entry.ContentOffset + entry.FileImageSize
		prevMemoryEnd = entry.VirtualAddress + entry.MemoryImageSize

	}

	// check section header / contents

	start = header.SectionHeaderOffset
	end = start + 9*Elf64SectionHeaderEntrySize
	sectionHeaderBytes := content[start:end]
	sectionHeader := [9]Elf64SectionHeaderEntry{}
	n, err = binary.Decode(
		sectionHeaderBytes,
		binary.LittleEndian,
		&sectionHeader)
	expect.Nil(t, err)
	expect.Equal(t, len(sectionHeaderBytes), n)

	expect.Equal(t, Elf64SectionHeaderEntry{}, sectionHeader[0])

	entry := sectionHeader[writer.SectionStringTableIndex]
	expect.Equal(t, SHT_STRTAB, entry.Type)

	// .shstrtab
	shstrtabContent := content[entry.Offset : entry.Offset+entry.Size]
	for _, name := range []string{
		".text",
		".init",
		".rodata",
		".data",
		".bss",
		".symtab",
		".strtab",
		".shstrtab",
	} {
		expect.True(t, bytes.Contains(shstrtabContent, []byte(name)))
	}

	// .strtab
	entry = sectionHeader[writer.StringTableIndex]
	expect.Equal(t, SHT_STRTAB, entry.Type)

	strtabContent := content[entry.Offset : entry.Offset+entry.Size]
	for _, symbol := range writer.Definitions.Symbols {
		expect.True(t, bytes.Contains(strtabContent, []byte(symbol.Name)))
	}

	// .symtab
	entry = sectionHeader[writer.SymbolTableIndex]
	expect.Equal(t, SHT_SYMTAB, entry.Type)

	expect.Equal(t, Elf64SymbolEntrySize, entry.EntrySize)
	expect.Equal(t, 7, entry.Size/entry.EntrySize)
	expect.Equal(t, 0, entry.Size%entry.EntrySize)

	elfSymbols := [7]Elf64SymbolEntry{}
	n, err = binary.Decode(
		content[entry.Offset:entry.Offset+entry.Size],
		binary.LittleEndian,
		&elfSymbols)
	expect.Nil(t, err)
	expect.Equal(t, entry.Size, uint64(n))

	for idx, symbol := range writer.Definitions.Symbols {
		elfSymbol := elfSymbols[idx]
		expect.True(
			t,
			bytes.HasPrefix(strtabContent[elfSymbol.NameIndex:], []byte(symbol.Name)))

		kind := STT_FUNC
		if symbol.Kind == layout.ObjectKind {
			kind = STT_OBJECT
		}
		expect.Equal(t, kind, elfSymbol.Info&0xf)
		expect.Equal(t, STV_DEFAULT, elfSymbol.Visibility)
		expect.Equal(
			t,
			writer.VirtualAddressStart+uint64(symbol.Offset),
			elfSymbol.Value)
		expect.Equal(t, uint64(symbol.Size), elfSymbol.Size)
	}

	// .text
	entry = sectionHeader[writer.TextIndex]
	expect.Equal(t, SHT_PROGBITS, entry.Type)
	expect.Equal(
		t,
		content[entry.Offset:entry.Offset+entry.Size],
		writer.Text.Flatten())

	// .init
	entry = sectionHeader[writer.InitIndex]
	expect.Equal(t, SHT_PROGBITS, entry.Type)
	expect.Equal(
		t,
		content[entry.Offset:entry.Offset+entry.Size],
		writer.Init.Flatten())

	// .rodata
	entry = sectionHeader[writer.ReadOnlyDataIndex]
	expect.Equal(t, SHT_PROGBITS, entry.Type)
	expect.Equal(
		t,
		content[entry.Offset:entry.Offset+entry.Size],
		writer.ReadOnlyData.Flatten())

	// .data
	entry = sectionHeader[writer.DataIndex]
	expect.Equal(t, SHT_PROGBITS, entry.Type)
	expect.Equal(
		t,
		content[entry.Offset:entry.Offset+entry.Size],
		writer.Data.Flatten())

	// .bss
	entry = sectionHeader[writer.BSSIndex]
	expect.Equal(t, SHT_NOBITS, entry.Type)
	expect.Equal(t, uint64(writer.BSSSize), entry.Size)
}
