package layout

// OS independent architecture config
type ArchitectureConfig struct {
	// On Finalize, merge content whenever the accumulated size is larger than
	// the threshold.
	MergeContentThreshold int64

	// The architecture's largest register size in bytes (e.g., 32 bytes ymm
	// on amd64)
	RegisterAlignment int64

	// 4KB on most architectures
	MemoryPageSize int64

	Relocator
}

// OS/architecture dependent config
type Config struct {
	Architecture ArchitectureConfig

	// Page(s) prior to the start page is reserved for executable file format's
	// metadata
	ExecutableImageStartPage int64

	// Usually the SIGTRAP software interrupt instruction (e.g., int3 on amd64)
	InstructionPadding []byte

	// Usually []byte{0}
	DataPadding []byte

	// Usually "_init"
	InitSymbol string

	// Usually the return instruction or equivalent (e.g., ret on amd64), used
	// for finalizing the init function for the executable image.  Assumption:
	// the .init section is a sequence of call instructions to modules' init
	// functions.
	InitEpilogue []byte

	// Usually "_start_"
	EntryPointSymbolPrefix string
}
