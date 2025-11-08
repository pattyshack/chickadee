// Based on linux's man page, elf.h, golang's debug/elf package,
// and the elf 1.2 spec.
package executable

const (
	Elf64HeaderSize             = 64
	Elf64SectionHeaderEntrySize = 64
	Elf64ProgramHeaderEntrySize = 56
	Elf64SymbolEntrySize        = 24

	ELFDATA2LSB = byte(1) // little endian
	ELFDATA2MSB = byte(1) // big endian

	EM_X86_64 = 62 // aka amd64

	ET_EXEC = 2
	ET_DYN  = 3 // shared library or position independent executable

	PT_LOAD      = 1
	PT_PHDR      = 6
	PT_GNU_STACK = 0x6474e551

	SHT_PROGBITS = 1
	SHT_SYMTAB   = 2
	SHT_STRTAB   = 3
	SHT_NOBITS   = 8

	SHF_WRITE     = 0x1
	SHF_ALLOC     = 0x2
	SHF_EXECINSTR = 0x4

	STB_GLOBAL = byte(1)

	STT_OBJECT = byte(1)
	STT_FUNC   = byte(2)

	functionSymbolInfo = (STB_GLOBAL << 4) | STT_FUNC
	objectSymbolInfo   = (STB_GLOBAL << 4) | STT_OBJECT

	STV_DEFAULT = byte(0)
)

// Header structs matching c's elf64 header definitions.  These are only used
// for (de-)serialization.

// e_ident
type ElfIdentifier struct {
	Magic              [4]byte // EI_MAG0 ... EI_MAG3
	Class              byte    // EI_CLASS
	DataEncoding       byte    // EI_DATA
	IdentifierVersion  byte    // EI_VERSION
	OperatingSystemABI byte    // EI_OSABI
	ABIVersion         byte    // EI_ABIVERSION
	Padding            [7]byte // EI_PAD
}

// Elf64_Ehdr
type Elf64Header struct {
	ElfIdentifier                  // e_ident[EI_NIDENT]
	FileType                uint16 // e_type
	MachineArchitecture     uint16 // e_machine
	FormatVersion           uint32 // e_version
	EntryPointAddress       uint64 // e_entry
	ProgramHeaderOffset     uint64 // e_phoff
	SectionHeaderOffset     uint64 // e_shoff
	ArchitectureFlags       uint32 // e_flags
	ElfHeaderSize           uint16 // e_ehsize
	ProgramHeaderEntrySize  uint16 // e_phentsize
	NumProgramHeaderEntries uint16 // e_phnum
	SectionHeaderEntrySize  uint16 // e_shentsize
	NumSectionHeaderEntries uint16 // e_shnum
	SectionStringTableIndex uint16 // e_shstrndx
}

// Elf64_Phdr
type Elf64ProgramHeaderEntry struct {
	Type            uint32 // p_type
	Flags           uint32 // p_flags
	ContentOffset   uint64 // p_offset
	VirtualAddress  uint64 // p_vaddr
	PhysicalAddress uint64 // p_paddr
	FileImageSize   uint64 // filesz
	MemoryImageSize uint64 // p_memsz
	Alignment       uint64 // p_align
}

// Elf64_Shdr
type Elf64SectionHeaderEntry struct {
	NameIndex        uint32 // sh_name
	Type             uint32 // sh_type
	Flags            uint64 // sh_flags
	Address          uint64 // sh_addr
	Offset           uint64 // sh_offset
	Size             uint64 // sh_size
	Link             uint32 // sh_link
	Info             uint32 // sh_info
	AddressAlignment uint64 // sh_addralign
	EntrySize        uint64 // sh_entsize
}

// Elf64_Sym
type Elf64SymbolEntry struct {
	NameIndex    uint32 // st_name
	Info         byte   // st_info.  (4 bits st_bind, 4 bits st_type)
	Visibility   byte   // st_other
	SectionIndex uint16 // st_shndx
	Value        uint64 // st_value
	Size         uint64 // st_size
}
