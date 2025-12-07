package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeUint64(t *testing.T) {
	// cmp rdi, rbx
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Uint64, registers.Rdi, registers.Rbx)
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

func TestJeImmediateUint64(t *testing.T) {
	// cmp rbp, 0x12345678
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jeIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.Rbp,
		uint64(0x12345678))
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

func TestJneUint64(t *testing.T) {
	// cmp rdx, rcx
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Uint64, registers.Rdx, registers.Rcx)
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

func TestJneImmediateUint64(t *testing.T) {
	// cmp rbx, 0x23456789
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jneIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.Rbx,
		uint64(0x23456789))
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

func TestJltUint64(t *testing.T) {
	// cmp rax, r10
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Uint64, registers.Rax, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x49, 0x3b, 0xc2, // cmp
			0x0f, 0x82, 0, 0, 0, 0, // jb
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

func TestJltImmediateUint64(t *testing.T) {
	// cmp rax, 0x34567890
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jltIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.Rax,
		uint64(0x34567890))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xf8, 0x90, 0x78, 0x56, 0x34, // cmp
			0x0f, 0x82, 0, 0, 0, 0, // jb
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

func TestJleUint64(t *testing.T) {
	// cmp r11, rcx
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Uint64, registers.R11, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x4c, 0x3b, 0xd9, // cmp
			0x0f, 0x86, 0, 0, 0, 0, // jbe
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

func TestJleImmediateUint64(t *testing.T) {
	// cmp rdx, 0x7ffffffff
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jleIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.Rdx,
		uint64(0x7fffffff))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfa, 0xff, 0xff, 0xff, 0x7f, // cmp
			0x0f, 0x86, 0, 0, 0, 0, // jbe
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

func TestJgtUint64(t *testing.T) {
	// cmp rdx, rbp
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Uint64, registers.Rdx, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xd5, // cmp
			0x0f, 0x87, 0, 0, 0, 0, // ja
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

func TestJgtImmediateUint64(t *testing.T) {
	// cmp rcx, 0x4567890a
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgtIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.Rcx,
		uint64(0x4567890a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xf9, 0x0a, 0x89, 0x67, 0x45, // cmp
			0x0f, 0x87, 0, 0, 0, 0, // ja
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

func TestJgeUint64(t *testing.T) {
	// cmp rbx, rax
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Uint64, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x3b, 0xd8, // cmp
			0x0f, 0x83, 0, 0, 0, 0, // jae
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

func TestJgeImmediateUint64(t *testing.T) {
	// cmp r13, 0x567890ab
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgeIntImmediate(
		builder,
		"jump-label",
		ir.Uint64,
		registers.R13,
		uint64(0x567890ab))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x49, 0x81, 0xfd, 0xab, 0x90, 0x78, 0x56, // cmp
			0x0f, 0x83, 0, 0, 0, 0, // jae
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
