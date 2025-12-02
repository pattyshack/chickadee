package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestOrInt8(t *testing.T) {
	// or eax, ecx
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0b, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt16(t *testing.T) {
	// or edx, ebx
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt32(t *testing.T) {
	// or ebp, esi
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt64(t *testing.T) {
	// or rdi, r8
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x0b, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint8(t *testing.T) {
	// or r9d, r10d
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint16(t *testing.T) {
	// or r11d, r12d
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0b, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint32(t *testing.T) {
	// or r13d, r14d
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint64(t *testing.T) {
	// or r15, rax
	builder := layout.NewSegmentBuilder()
	or(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0b, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt8Immediate(t *testing.T) {
	// or dl, 0x12
	builder := layout.NewSegmentBuilder()
	orIntImmediate(builder, ir.Int8, registers.Rdx, []byte{0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xca, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt16Immediate(t *testing.T) {
	// or dx, 0x1234
	builder := layout.NewSegmentBuilder()
	orIntImmediate(builder, ir.Int16, registers.Rdx, []byte{0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xca, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt32Immediate(t *testing.T) {
	// or edx, 0x12345678
	builder := layout.NewSegmentBuilder()
	orIntImmediate(
		builder,
		ir.Int32,
		registers.Rdx,
		[]byte{0x78, 0x56, 0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xca, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrInt64Immediate(t *testing.T) {
	// or rdx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	orIntImmediate(
		builder,
		ir.Int64,
		registers.Rdx,
		[]byte{0x9a, 0x78, 0x56, 0x34})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xca, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint8Immediate(t *testing.T) {
	// or dil, 0x12
	builder := layout.NewSegmentBuilder()
	orIntImmediate(builder, ir.Uint8, registers.Rdi, []byte{0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x80, 0xcf, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint16Immediate(t *testing.T) {
	// or di, 0x1234
	builder := layout.NewSegmentBuilder()
	orIntImmediate(builder, ir.Uint16, registers.Rdi, []byte{0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xcf, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint32Immediate(t *testing.T) {
	// or edi, 0x12345678
	builder := layout.NewSegmentBuilder()
	orIntImmediate(
		builder,
		ir.Uint32,
		registers.Rdi,
		[]byte{0x78, 0x56, 0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xcf, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestOrUint64Immediate(t *testing.T) {
	// or rdi, 0x3456789a
	builder := layout.NewSegmentBuilder()
	orIntImmediate(
		builder,
		ir.Uint64,
		registers.Rdi,
		[]byte{0x9a, 0x78, 0x56, 0x34})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xcf, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
