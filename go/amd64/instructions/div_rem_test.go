package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestDivRemUint8(t *testing.T) {
	// movzx eax, al
	// movzx ebp, bpl
	// xor edx, edx
	// div ebp
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint8, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xb6, 0xc0, // movzx
			0x40, 0x0f, 0xb6, 0xed, // movzx
			0x33, 0xd2, // xor
			0xf7, 0xf5, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint16(t *testing.T) {
	// xor edx, edx
	// div di
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint16, registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x66, 0xf7, 0xf7, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint32(t *testing.T) {
	// xor edx, edx
	// div r13d
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint32, registers.R13)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x41, 0xf7, 0xf5, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint64(t *testing.T) {
	// xor edx, edx
	// div rcx
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x48, 0xf7, 0xf1, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt8(t *testing.T) {
	// movsx eax, al
	// movsx ebp, bpl
	// cdq
	// idiv ebp
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int8, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xbe, 0xc0, // movsx
			0x40, 0x0f, 0xbe, 0xed, // movsx
			0x99,       // cdq
			0xf7, 0xfd, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt16(t *testing.T) {
	// cwd
	// idiv di
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int16, registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x99, // cwd
			0x66, 0xf7, 0xff, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt32(t *testing.T) {
	// cdq
	// idiv r13d
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int32, registers.R13)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x99,             // cdq
			0x41, 0xf7, 0xfd, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt64(t *testing.T) {
	// cqo
	// idiv rcx
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x99, // cqo
			0x48, 0xf7, 0xf9, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivFloat32(t *testing.T) {
	// divss xmm3, xmm7
	builder := layout.NewSegmentBuilder()
	divFloat(builder, ir.Float32, registers.Xmm3, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf3, 0x0f, 0x5e, 0xdf,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivFloat64(t *testing.T) {
	// divsd xmm10, xmm5
	builder := layout.NewSegmentBuilder()
	divFloat(builder, ir.Float64, registers.Xmm10, registers.Xmm5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf2, 0x44, 0x0f, 0x5e, 0xd5,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
