package layout

import (
	"github.com/pattyshack/chickadee/platform/layout"
)

var (
	ArchitectureLayout = layout.ArchitectureConfig{
		MergeContentThreshold: 1024 * 1024, // 1 MB (arbitrary choice)
		RegisterAlignment:     16,          // SSE2 xmm registers
		MemoryPageSize:        4096,
		Relocator:             NewRelocator(),
	}

	LinuxLayout = layout.Config{
		Architecture:             ArchitectureLayout,
		ExecutableImageStartPage: 1,
		InstructionPadding:       []byte{0xcc}, // int3 instruction
		DataPadding:              []byte{0},
		InitSymbol:               "_init",
		InitEpilogue:             []byte{0xc3}, // ret instruction
		EntryPointSymbolPrefix:   "_start_",
	}
)
