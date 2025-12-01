package amd64

import (
	"github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/platform"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout/executable"
)

var (
	Linux = platform.Config{
		OperatingSystem: platform.Linux,
		Architecture: architecture.Config{
			Name:        "amd64",
			RegisterSet: registers.Registers,
		},
		Layout: layout.LinuxLayout,
		ExecutableFormat: executable.Config{
			VirtualAddressStart:    0x400000,
			ElfMachineArchitecture: executable.EM_X86_64,
		},
	}
)
