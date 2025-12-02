package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestMulInt8(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul eax, ecx
	expect.Equal(t, []byte{0x0f, 0xaf, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt16(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul edx, ebx
	expect.Equal(t, []byte{0x0f, 0xaf, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt32(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul ebp, esi
	expect.Equal(t, []byte{0x0f, 0xaf, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt64(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul rdi, r8
	expect.Equal(t, []byte{0x49, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint8(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul r9d, r10d
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint16(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul r11d, r12d
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint32(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul r13d, r14d
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint64(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// imul r15, rax
	expect.Equal(t, []byte{0x4c, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulFloat32(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Float32, registers.Xmm0, registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// mulss xmm0, xmm2
	expect.Equal(t, []byte{0xf3, 0x0f, 0x59, 0xc2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulFloat64(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Float64, registers.Xmm1, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)

	// mulsd xmm1, xmm3
	expect.Equal(t, []byte{0xf2, 0x0f, 0x59, 0xcb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

// TODO test mulIntImmediate
