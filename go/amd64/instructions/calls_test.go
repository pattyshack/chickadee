package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestSyscall(t *testing.T) {
	// syscall
	builder := layout.NewSegmentBuilder()
	syscall(builder)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0x05}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCallAddress(t *testing.T) {
	// call rdx
	builder := layout.NewSegmentBuilder()
	callAddress(builder, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xff, 0xd2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// call r11
	builder = layout.NewSegmentBuilder()
	callAddress(builder, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xff, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCallSymbol(t *testing.T) {
	// call 0x00000000
	builder := layout.NewSegmentBuilder()
	callSymbol(builder, "function-symbol")
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xe8, 0, 0, 0, 0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "function-symbol",
					Offset: 1,
				},
			},
		},
		segment.Relocations)
}

func TestRet(t *testing.T) {
	// ret
	builder := layout.NewSegmentBuilder()
	ret(builder)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
