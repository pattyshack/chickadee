package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeUint8(t *testing.T) {
	// cmp dil, bl
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Uint8, registers.Rdi, registers.Rbx)
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

func TestJeImmediateUint8(t *testing.T) {
	// cmp bpl, 5
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jeIntImmediate(builder, "jump-label", ir.Uint8, registers.Rbp, uint8(5))
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

func TestJneUint8(t *testing.T) {
	// cmp dl, cl
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Uint8, registers.Rdx, registers.Rcx)
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

func TestJneImmediateUint8(t *testing.T) {
	// cmp bl, 5
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jneIntImmediate(builder, "jump-label", ir.Uint8, registers.Rbx, uint8(5))
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

func TestJltUint8(t *testing.T) {
	// cmp al, r10b
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Uint8, registers.Rax, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x41, 0x3a, 0xc2, // cmp
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

func TestJltImmediateUint8(t *testing.T) {
	// cmp al, 15
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jltIntImmediate(builder, "jump-label", ir.Uint8, registers.Rax, uint8(15))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xf8, 0x0f, // cmp
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

func TestJleUint8(t *testing.T) {
	// cmp r11b, cl
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Uint8, registers.R11, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x44, 0x3a, 0xd9, // cmp
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

func TestJleImmediateUint8(t *testing.T) {
	// cmp dl, 1
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jleIntImmediate(builder, "jump-label", ir.Uint8, registers.Rdx, uint8(1))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xfa, 0x01, // cmp
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

func TestJgtUint8(t *testing.T) {
	// cmp dl, bpl
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Uint8, registers.Rdx, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x3a, 0xd5, // cmp
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

func TestJgtImmediateUint8(t *testing.T) {
	// cmp cl, 17
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgtIntImmediate(builder, "jump-label", ir.Uint8, registers.Rcx, uint8(17))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x80, 0xf9, 0x11, // cmp
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

func TestJgeUint8(t *testing.T) {
	// cmp bl, al
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Uint8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x3a, 0xd8, // cmp
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
					Offset: 4,
				},
			},
		},
		segment.Relocations)
}

func TestJgeImmediateUint8(t *testing.T) {
	// cmp r13b, 127
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgeIntImmediate(builder, "jump-label", ir.Uint8, registers.R13, uint8(127))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x41, 0x80, 0xfd, 0x7f, // cmp
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}
