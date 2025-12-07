package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJeFloat32(t *testing.T) {
	// comiss xmm0, xmm7
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Float32, registers.Xmm0, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0x2f, 0xc7, // comiss
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

func TestJneFloat32(t *testing.T) {
	// comiss xmm8, xmm6
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Float32, registers.Xmm8, registers.Xmm6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x44, 0x0f, 0x2f, 0xc6, // comiss
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJltFloat32(t *testing.T) {
	// comiss xmm5, xmm9
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Float32, registers.Xmm5, registers.Xmm9)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x41, 0x0f, 0x2f, 0xe9, // comiss
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJleFloat32(t *testing.T) {
	// comiss xmm10, xmm12
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Float32, registers.Xmm10, registers.Xmm12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x45, 0x0f, 0x2f, 0xd4, // comiss
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJgtFloat32(t *testing.T) {
	// comiss xmm2, xmm4
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Float32, registers.Xmm2, registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0x2f, 0xd4, // comiss
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

func TestJgeFloat32(t *testing.T) {
	// comiss xmm1, xmm3
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Float32, registers.Xmm1, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0x2f, 0xcb, // comiss
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

func TestJeFloat64(t *testing.T) {
	// comisd xmm0, xmm7
	// je (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	je(builder, "jump-label", ir.Float64, registers.Xmm0, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x0f, 0x2f, 0xc7, // comisd
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

func TestJneFloat64(t *testing.T) {
	// comisd xmm8, xmm6
	// jne (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jne(builder, "jump-label", ir.Float64, registers.Xmm8, registers.Xmm6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x44, 0x0f, 0x2f, 0xc6, // comisd
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

func TestJltFloat64(t *testing.T) {
	// comisd xmm5, xmm9
	// jb (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jlt(builder, "jump-label", ir.Float64, registers.Xmm5, registers.Xmm9)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x41, 0x0f, 0x2f, 0xe9, // comisd
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
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJleFloat64(t *testing.T) {
	// comisd xmm10, xmm12
	// jbe (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jle(builder, "jump-label", ir.Float64, registers.Xmm10, registers.Xmm12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x45, 0x0f, 0x2f, 0xd4, // comisd
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
					Offset: 7,
				},
			},
		},
		segment.Relocations)
}

func TestJgtFloat64(t *testing.T) {
	// comisd xmm2, xmm4
	// ja (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jgt(builder, "jump-label", ir.Float64, registers.Xmm2, registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x0f, 0x2f, 0xd4, // comisd
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
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}

func TestJgeFloat64(t *testing.T) {
	// comisd xmm1, xmm3
	// jae (4 byte placeholder)
	builder := layout.NewSegmentBuilder()
	jge(builder, "jump-label", ir.Float64, registers.Xmm1, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x0f, 0x2f, 0xcb, // comisd
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
