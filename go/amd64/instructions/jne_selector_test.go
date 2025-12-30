package instructions

import (
	"math"
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestJneUint(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Uint32,
	}
	src1.(*ir.LocalReference).UseDef = src1Def
	src1Chunk := src1Def.Chunks()[0]

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Uint32,
	}
	src2.(*ir.LocalReference).UseDef = src2Def
	src2Chunk := src2Def.Chunks()[0]

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  src1,
		Src2:  src2,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 2, len(constraints.RegisterSources))
	expect.Equal(t, src1Chunk, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, src2Chunk, constraints.RegisterSources[1].DefinitionChunk)

	src1Register := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		src1Register)

	src2Register := constraints.RegisterSources[1].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		src2Register)

	expect.True(t, src1Register != src2Register)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			src1Register: registers.Rdx,
			src2Register: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x3b, 0xd1,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 4,
				},
			},
		},
		segment.Relocations)
}

func TestJneUintRightImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(uint64(0))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Uint64,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  src,
		Src2:  imm,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpImmediateInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		srcRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfb, 0x00, 0x00, 0x00, 0x00,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJneUintLeftImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(uint64(math.MaxInt32))
	immDef := &ir.Definition{
		Name: "imm",
		Type: imm.Type(),
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  imm,
		Src2:  src,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpImmediateInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		srcRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfb, 0xff, 0xff, 0xff, 0x7f,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJneInt(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Int32,
	}
	src1.(*ir.LocalReference).UseDef = src1Def
	src1Chunk := src1Def.Chunks()[0]

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Int32,
	}
	src2.(*ir.LocalReference).UseDef = src2Def
	src2Chunk := src2Def.Chunks()[0]

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  src1,
		Src2:  src2,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 2, len(constraints.RegisterSources))
	expect.Equal(t, src1Chunk, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, src2Chunk, constraints.RegisterSources[1].DefinitionChunk)

	src1Register := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		src1Register)

	src2Register := constraints.RegisterSources[1].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		src2Register)

	expect.True(t, src1Register != src2Register)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			src1Register: registers.Rdx,
			src2Register: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x3b, 0xd1,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 4,
				},
			},
		},
		segment.Relocations)
}

func TestJneIntRightImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Int64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(int64(math.MinInt32))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Int64,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  src,
		Src2:  imm,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpImmediateInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		srcRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfb, 0x00, 0x00, 0x00, 0x80,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJneIntLeftImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Int64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(int64(math.MaxInt32))
	immDef := &ir.Definition{
		Name: "imm",
		Type: imm.Type(),
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  imm,
		Src2:  src,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpImmediateInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
		},
		srcRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x81, 0xfb, 0xff, 0xff, 0xff, 0x7f,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 9,
				},
			},
		},
		segment.Relocations)
}

func TestJneFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Float64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(float64(3))
	immDef := &ir.Definition{
		Name: "imm",
		Type: imm.Type(),
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef
	immChunk := immDef.Chunks()[0]

	jump := &ir.ConditionalJump{
		Kind:  ir.Jne,
		Label: "jne-label",
		Src1:  src,
		Src2:  imm,
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		jump,
		architecture.SelectorHint{})

	_, ok := instruction.(conditionalJumpInstruction)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)
	expect.Nil(t, constraints.RegisterDestinations)

	expect.Equal(t, 2, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, immChunk, constraints.RegisterSources[1].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: false,
			AnyFloat:   true,
		},
		srcRegister)

	immRegister := constraints.RegisterSources[1].RegisterConstraint
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: false,
			AnyFloat:   true,
		},
		immRegister)

	expect.True(t, srcRegister != immRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Xmm1,
			immRegister: registers.Xmm2,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x0f, 0x2f, 0xca,
			0x0f, 0x85, 0, 0, 0, 0,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(
		t,
		layout.Relocations{
			Labels: []*layout.Relocation{
				{
					Name:   "jne-label",
					Offset: 6,
				},
			},
		},
		segment.Relocations)
}
