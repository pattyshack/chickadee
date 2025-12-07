package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestComputeSymbolAddress(t *testing.T) {
	// lea rdx, [rip + 0x12345678]
	builder := layout.NewSegmentBuilder()
	computeSymbolAddress(builder, registers.Rdx, "symbol", int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x8d, 0x15, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "symbol",
					Offset: 3,
				},
			},
		},
		segment.Relocations)

	// lea rbp, [rip + 0x01020304]
	builder = layout.NewSegmentBuilder()
	computeSymbolAddress(builder, registers.Rbp, "symbol2", int32(0x01020304))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x8d, 0x2d, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name:   "symbol2",
					Offset: 3,
				},
			},
		},
		segment.Relocations)
}

func TestComputeStackAddress(t *testing.T) {
	// lea rcx, [rsp + 36]
	builder := layout.NewSegmentBuilder()
	computeStackAddress(builder, registers.Rcx, int32(36))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x8d, 0x8c, 0x24, 0x24, 0x00, 0x00, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// lea r10, [rsp + 0x01020304]
	builder = layout.NewSegmentBuilder()
	computeStackAddress(builder, registers.R10, int32(0x01020304))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x4c, 0x8d, 0x94, 0x24, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAllocateStackFrame(t *testing.T) {
	// add rsp, -16
	builder := layout.NewSegmentBuilder()
	allocateStackFrame(builder, int32(16))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xc4, 0xf0, 0xff, 0xff, 0xff},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDeallocateStackFrame(t *testing.T) {
	// add rsp, 0x12345678
	builder := layout.NewSegmentBuilder()
	deallocateStackFrame(builder, int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xc4, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
