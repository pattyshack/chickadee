package amd64

import (
	"encoding/binary"

	"github.com/pattyshack/chickadee/platform"
	"github.com/pattyshack/chickadee/platform/layout"
	"github.com/pattyshack/chickadee/platform/layout/executable"
)

var (
	Linux = platform.Parameters{
		Architecture:    platform.Amd64,
		OperatingSystem: platform.Linux,
		Layout: layout.Parameters{
			MergeContentThreshold:    1024 * 1024, // 1 MB (arbitrary choice)
			RegisterAlignment:        32,
			MemoryPageSize:           4096,
			ExecutableImageStartPage: 1,
			InstructionPadding:       []byte{0xcc}, // int3 instruction
			DataPadding:              []byte{0},
			InitSymbol:               "_init",
			InitEpilogue:             []byte{0xc3}, // ret instruction
			Relocator:                NewRel32Relocator(),
		},
		ExecutableFormat: executable.Parameters{
			ByteOrder:              binary.LittleEndian,
			VirtualAddressStart:    0x400000,
			ElfMachineArchitecture: executable.EM_X86_64,
		},
	}
)
