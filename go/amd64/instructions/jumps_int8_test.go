package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeInt8(t *testing.T) {
	// cmp dil, bl
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Int8, registers.Rdi, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x3a, 0xfb, // cmp
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

func TestJeImmediateInt8(t *testing.T) {
	// cmp bpl, 5
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jeIntImmediate(builder, "jump-label", ir.Int8, registers.Rbp, int8(5))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x80, 0xfd, 0x05, // cmp
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJneInt8(t *testing.T) {
	// cmp dl, cl
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Int8, registers.Rdx, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x3a, 0xd1, // cmp
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
					Offset: 4,
				},
			},
		},
		segment.Relocations)
}

func TestJneImmediateInt8(t *testing.T) {
	// cmp bl, 5
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jneIntImmediate(builder, "jump-label", ir.Int8, registers.Rbx, int8(5))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xfb, 0x05, // cmp
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

func TestJltInt8(t *testing.T) {
	// cmp al, r10b
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Int8, registers.Rax, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x41, 0x3a, 0xc2, // cmp
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
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJltImmediateInt8(t *testing.T) {
	// cmp al, 15
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jltIntImmediate(builder, "jump-label", ir.Int8, registers.Rax, int8(15))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xf8, 0x0f, // cmp
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
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJleInt8(t *testing.T) {
	// cmp r11b, cl
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Int8, registers.R11, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x44, 0x3a, 0xd9, // cmp
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
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJleImmediateInt8(t *testing.T) {
	// cmp dl, -1
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jleIntImmediate(builder, "jump-label", ir.Int8, registers.Rdx, int8(-1))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xfa, 0xff, // cmp
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
					Offset: 5,
				},
			},
		},
		segment.Relocations)
}

func TestJgtInt8(t *testing.T) {
	// cmp dl, bpl
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Int8, registers.Rdx, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x3a, 0xd5, // cmp
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

func TestJgtImmediateInt8(t *testing.T) {
	// cmp cl, 17
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgtIntImmediate(builder, "jump-label", ir.Int8, registers.Rcx, int8(17))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xf9, 0x11, // cmp
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

func TestJgeInt8(t *testing.T) {
	// cmp bl, al
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Int8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x3a, 0xd8, // cmp
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
					Offset: 4,
				},
			},
		},
		segment.Relocations)
}

func TestJgeImmediateInt8(t *testing.T) {
	// cmp r13b, 127
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgeIntImmediate(builder, "jump-label", ir.Int8, registers.R13, int8(127))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x41, 0x80, 0xfd, 0x7f, // cmp
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}
