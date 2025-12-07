package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestNegInt8(t *testing.T) {
	// neg bpl
	builder := layout.NewSegmentBuilder()
	negSignedInt(builder, ir.Int8, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0xf6, 0xdd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestNegInt16(t *testing.T) {
	// neg bx
	builder := layout.NewSegmentBuilder()
	negSignedInt(builder, ir.Int16, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xf7, 0xdb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestNegInt32(t *testing.T) {
	// neg edx
	builder := layout.NewSegmentBuilder()
	negSignedInt(builder, ir.Int32, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf7, 0xda}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestNegInt64(t *testing.T) {
	// neg r8
	builder := layout.NewSegmentBuilder()
	negSignedInt(builder, ir.Int64, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xf7, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestNegFloat32(t *testing.T) {
	// xor ecx, 0x80000000
	builder := layout.NewSegmentBuilder()
	negFloat32(builder, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xf1, 0x00, 0x00, 0x00, 0x80},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestNegFloat64(t *testing.T) {
	// mov rcx, 0x80000000_00000000
	// xor rcx, rdx
	builder := layout.NewSegmentBuilder()
	negFloat64(builder, registers.Rcx, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0xb9, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, // mov
			0x48, 0x33, 0xca, // xor
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
