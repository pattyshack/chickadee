package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestShlInt8(t *testing.T) {
	// shl r14d, cl (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Int8, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd3, 0xe6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt16(t *testing.T) {
	// shl esi, cl (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Int16, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xe6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt32(t *testing.T) {
	// shl ebx, cl
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Int32, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xe3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt64(t *testing.T) {
	// shl rcx, cl
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Int64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xd3, 0xe1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint8(t *testing.T) {
	// shl eax, cl (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Uint8, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xe0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint16(t *testing.T) {
	// shl edx, cl (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Uint16, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xe2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint32(t *testing.T) {
	// shl r10d, cl
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Uint32, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd3, 0xe2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint64(t *testing.T) {
	// shl r8, cl
	builder := layout.NewSegmentBuilder()
	shl(builder, ir.Uint64, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xd3, 0xe0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt8Immediate(t *testing.T) {
	// shl r14d, 1 (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Int8, registers.R14, 1)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc1, 0xe6, 1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt16Immediate(t *testing.T) {
	// shl esi, 2 (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Int16, registers.Rsi, 2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xe6, 2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt32Immediate(t *testing.T) {
	// shl ebx, 3
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Int32, registers.Rbx, 3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xe3, 3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlInt64Immediate(t *testing.T) {
	// shl rcx, 4
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Int64, registers.Rcx, 4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xc1, 0xe1, 4}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint8Immediate(t *testing.T) {
	// shl eax, 5 (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Uint8, registers.Rax, 5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xe0, 5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint16Immediate(t *testing.T) {
	// shl edx, 6 (32-bit variant)
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Uint16, registers.Rdx, 6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xe2, 6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint32Immediate(t *testing.T) {
	// shl r10d, 7
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Uint32, registers.R10, 7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc1, 0xe2, 7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShlUint64Immediate(t *testing.T) {
	// shl r8, 8
	builder := layout.NewSegmentBuilder()
	shlIntImmediate(builder, ir.Uint64, registers.R8, 8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xc1, 0xe0, 8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
