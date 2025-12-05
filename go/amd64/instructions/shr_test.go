package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestShrInt8(t *testing.T) {
	// sar r14b, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int8, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd2, 0xfe}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt16(t *testing.T) {
	// sar si, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int16, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xfe}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt32(t *testing.T) {
	// sar ebx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int32, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xfb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt64(t *testing.T) {
	// sar rcx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xd3, 0xf9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint8(t *testing.T) {
	// shr al, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint8, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd2, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint16(t *testing.T) {
	// shr dx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint16, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xea}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint32(t *testing.T) {
	// shr r10d, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint32, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd3, 0xea}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint64(t *testing.T) {
	// shr r8, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint64, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xd3, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt8Immediate(t *testing.T) {
	// sar r14b, 1
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int8, registers.R14, 1)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc0, 0xfe, 1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt16Immediate(t *testing.T) {
	// sar si, 2
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int16, registers.Rsi, 2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xc1, 0xfe, 2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt32Immediate(t *testing.T) {
	// sar ebx, 3
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int32, registers.Rbx, 3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xfb, 3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt64Immediate(t *testing.T) {
	// sar rcx, 4
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int64, registers.Rcx, 4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xc1, 0xf9, 4}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint8Immediate(t *testing.T) {
	// shr al, 5
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint8, registers.Rax, 5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc0, 0xe8, 5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint16Immediate(t *testing.T) {
	// shr dx, 6
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint16, registers.Rdx, 6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xc1, 0xea, 6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint32Immediate(t *testing.T) {
	// shr r10d, 7
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint32, registers.R10, 7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc1, 0xea, 7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint64Immediate(t *testing.T) {
	// shr r8, 8
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint64, registers.R8, 8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xc1, 0xe8, 8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
