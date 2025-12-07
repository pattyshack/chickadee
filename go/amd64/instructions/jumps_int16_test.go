package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeInt16(t *testing.T) {
	// cmp di, bx
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Int16, registers.Rdi, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x3b, 0xfb, // cmp
			0x0f, 0x84, 0, 0, 0, 0, // je
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJeImmediateInt16(t *testing.T) {
	// cmp bp, 0x1234
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jeIntImmediate(builder, "jump-label", ir.Int16, registers.Rbp, int16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x81, 0xfd, 0x34, 0x12, // cmp
			0x0f, 0x84, 0, 0, 0, 0, // je
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJneInt16(t *testing.T) {
	// cmp dx, cx
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Int16, registers.Rdx, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x3b, 0xd1, // cmp
			0x0f, 0x85, 0, 0, 0, 0, // jne
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJneImmediateInt16(t *testing.T) {
	// cmp bx, 0x2345
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jneIntImmediate(builder, "jump-label", ir.Int16, registers.Rbx, int16(0x2345))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x81, 0xfb, 0x45, 0x23, // cmp
			0x0f, 0x85, 0, 0, 0, 0, // jne
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJltInt16(t *testing.T) {
	// cmp ax, r10w
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Int16, registers.Rax, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x41, 0x3b, 0xc2, // cmp
			0x0f, 0x8c, 0, 0, 0, 0, // jl
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJltImmediateInt16(t *testing.T) {
	// cmp ax, 0x3456
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jltIntImmediate(builder, "jump-label", ir.Int16, registers.Rax, int16(0x3456))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x81, 0xf8, 0x56, 0x34, // cmp
			0x0f, 0x8c, 0, 0, 0, 0, // jl
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJleInt16(t *testing.T) {
	// cmp r11w, cx
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Int16, registers.R11, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x44, 0x3b, 0xd9, // cmp
			0x0f, 0x8e, 0, 0, 0, 0, // jle
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJleImmediateInt16(t *testing.T) {
	// cmp dx, -1
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jleIntImmediate(builder, "jump-label", ir.Int16, registers.Rdx, int16(-1))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x81, 0xfa, 0xff, 0xff, // cmp
			0x0f, 0x8e, 0, 0, 0, 0, // jle
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJgtInt16(t *testing.T) {
	// cmp dx, bp
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Int16, registers.Rdx, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x3b, 0xd5, // cmp
			0x0f, 0x8f, 0, 0, 0, 0, // jg
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJgtImmediateInt16(t *testing.T) {
	// cmp cx, 0x4567
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgtIntImmediate(builder, "jump-label", ir.Int16, registers.Rcx, int16(0x4567))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x81, 0xf9, 0x67, 0x45, // cmp
			0x0f, 0x8f, 0, 0, 0, 0, // jg
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJgeInt16(t *testing.T) {
	// cmp bx, ax
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Int16, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x3b, 0xd8, // cmp
			0x0f, 0x8d, 0, 0, 0, 0, // jge
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJgeImmediateInt16(t *testing.T) {
	// cmp r13w, 0x5678
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgeIntImmediate(builder, "jump-label", ir.Int16, registers.R13, int16(0x5678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x41, 0x81, 0xfd, 0x78, 0x56, // cmp
			0x0f, 0x8d, 0, 0, 0, 0, // jge
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 8,
				},
			},
		},
		segment.Relocations)
}
