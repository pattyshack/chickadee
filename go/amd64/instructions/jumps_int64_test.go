package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeInt64(t *testing.T) {
	// cmp rdi, rbx
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Int64, registers.Rdi, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xfb, // cmp
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

func TestJeImmediateInt64(t *testing.T) {
	// cmp rbp, 0x12345678
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jeIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.Rbp,
		int64(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfd, 0x78, 0x56, 0x34, 0x12, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJneInt64(t *testing.T) {
	// cmp rdx, rcx
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Int64, registers.Rdx, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xd1, // cmp
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

func TestJneImmediateInt64(t *testing.T) {
	// cmp rbx, 0x23456789
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jneIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.Rbx,
		int64(0x23456789))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfb, 0x89, 0x67, 0x45, 0x23, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJltInt64(t *testing.T) {
	// cmp rax, r10
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Int64, registers.Rax, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x49, 0x3b, 0xc2, // cmp
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

func TestJltImmediateInt64(t *testing.T) {
	// cmp rax, 0x34567890
	// jl (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jltIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.Rax,
		int64(0x34567890))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xf8, 0x90, 0x78, 0x56, 0x34, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJleInt64(t *testing.T) {
	// cmp r11, rcx
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Int64, registers.R11, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x4c, 0x3b, 0xd9, // cmp
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

func TestJleImmediateInt64(t *testing.T) {
	// cmp rdx, -1
	// jle (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jleIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.Rdx,
		int64(-1))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfa, 0xff, 0xff, 0xff, 0xff, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJgtInt64(t *testing.T) {
	// cmp rdx, rbp
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Int64, registers.Rdx, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xd5, // cmp
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

func TestJgtImmediateInt64(t *testing.T) {
	// cmp rcx, 0x4567890a
	// jg (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgtIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.Rcx,
		int64(0x4567890a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xf9, 0x0a, 0x89, 0x67, 0x45, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJgeInt64(t *testing.T) {
	// cmp rbx, rax
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Int64, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xd8, // cmp
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

func TestJgeImmediateInt64(t *testing.T) {
	// cmp r13, 0x567890ab
	// jge (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgeIntImmediate(
		builder,
		"jump-label",
		ir.Int64,
		registers.R13,
		int64(0x567890ab))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x49, 0x81, 0xfd, 0xab, 0x90, 0x78, 0x56, // cmp
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
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}
