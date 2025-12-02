package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestSubInt8(t *testing.T) {
	// sub bl, al
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2a, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub bpl, cl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rbp, registers.Rcx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xe9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub dl, sil
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rdx, registers.Rsi)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9b, dl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x2a, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub cl, r12b
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x2a, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10b, r11b
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x2a, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt16(t *testing.T) {
	// sub bp, si
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9w, dx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x44, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub cx, r12w
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x41, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10w, r11w
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x45, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt32(t *testing.T) {
	// sub ebp, esi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9d, edx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub ecx, r12d
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10d, r11d
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt64(t *testing.T) {
	// sub rbp, rsi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9, rdx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub rcx, r12
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10, r11
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4d, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint8(t *testing.T) {
	// sub bl, al
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2a, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint16(t *testing.T) {
	// sub bp, si
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint16, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint32(t *testing.T) {
	// sub ebp, esi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint64(t *testing.T) {
	// sub rbp, rsi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint64, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubFloat32(t *testing.T) {
	// subss xmm1, xmm7
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm1, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x5c, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm1, xmm12
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm1, registers.Xmm12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x41, 0x0f, 0x5c, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm10, xmm3
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm10, registers.Xmm3)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x5c, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm9, xmm14
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm9, registers.Xmm14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x45, 0x0f, 0x5c, 0xce},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubFloat64(t *testing.T) {
	// subsd xmm5, xmm6
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Float64, registers.Xmm5, registers.Xmm6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x5c, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

// TODO test subIntImmediate
