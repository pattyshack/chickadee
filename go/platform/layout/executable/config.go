package executable

type Config struct {
	// Where the compiled executable should be loaded (e.g., 4MB on amd64 linux)
	VirtualAddressStart uint64

	// Elf's EM_* machine architecture constant (defined in elf_header.go)
	ElfMachineArchitecture uint16
}
