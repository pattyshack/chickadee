package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestSelectJump(t *testing.T) {
	jump := &ir.Jump{
		Label: "jump-label",
	}

	instruction := architecture.SelectInstruction(
		InstructionSet,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(jumpInstruction)
	expect.True(t, ok)

	expect.Equal(
		t,
		architecture.InstructionConstraints{},
		instruction.Constraints())

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(builder, nil)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xe9, 0, 0, 0, 0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jump-label",
					Offset: 1,
				},
			},
		},
		segment.Relocations)
}
