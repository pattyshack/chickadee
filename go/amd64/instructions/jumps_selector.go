package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

type jumpInstruction struct {
	*ir.Jump
}

func (inst jumpInstruction) Instruction() ir.Instruction {
	return inst.Jump
}

func (jumpInstruction) Constraints() architecture.InstructionConstraints {
	return architecture.InstructionConstraints{}
}

func (inst jumpInstruction) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	jump(builder, inst.Label)
}

type encodeConditionalJumpFunc func(
	*layout.SegmentBuilder,
	string,
	ir.Type,
	*architecture.Register,
	*architecture.Register)

type conditionalJumpInstruction struct {
	*ir.ConditionalJump

	architecture.InstructionConstraints

	encode encodeConditionalJumpFunc
}

func (inst conditionalJumpInstruction) Instruction() ir.Instruction {
	return inst.ConditionalJump
}

func (inst conditionalJumpInstruction) Constraints() architecture.InstructionConstraints {
	return inst.InstructionConstraints
}

func (inst conditionalJumpInstruction) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	src1 := inst.InstructionConstraints.RegisterSources[0].RegisterConstraint
	src2 := src1
	if len(inst.InstructionConstraints.RegisterSources) == 2 {
		src2 = inst.InstructionConstraints.RegisterSources[1].RegisterConstraint
	}

	inst.encode(
		builder,
		inst.Label,
		inst.Src1.Type(),
		selectedRegisters[src1],
		selectedRegisters[src2])
}

type encodeConditionalJumpImmediateFunc func(
	*layout.SegmentBuilder,
	string,
	ir.Type,
	*architecture.Register,
	interface{})

type conditionalJumpImmediateInstruction struct {
	*ir.ConditionalJump

	immediate interface{}

	architecture.InstructionConstraints

	encode encodeConditionalJumpImmediateFunc
}

func (inst conditionalJumpImmediateInstruction) Instruction() ir.Instruction {
	return inst.ConditionalJump
}

func (inst conditionalJumpImmediateInstruction) Constraints() architecture.InstructionConstraints {
	return inst.InstructionConstraints
}

func (inst conditionalJumpImmediateInstruction) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	src1 := inst.InstructionConstraints.RegisterSources[0].RegisterConstraint

	inst.encode(
		builder,
		inst.Label,
		inst.Src1.Type(),
		selectedRegisters[src1],
		inst.immediate)
}

type jumpSelector struct{}

func (jumpSelector) Select(
	config architecture.Config,
	jump *ir.Jump,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	return jumpInstruction{
		Jump: jump,
	}
}

type conditionalJumpSelector struct {
	isFloat bool

	encodeRightImmediate encodeConditionalJumpImmediateFunc
	encodeLeftImmediate  encodeConditionalJumpImmediateFunc
	encode               encodeConditionalJumpFunc
}

func (selector conditionalJumpSelector) Select(
	config architecture.Config,
	jump *ir.ConditionalJump,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	instruction := selector.maybeNewImmediateJump(jump, hint)
	if instruction != nil {
		return instruction
	}

	return selector.newJump(jump, hint)
}

func (selector conditionalJumpSelector) maybeNewImmediateJump(
	jump *ir.ConditionalJump,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	if selector.isFloat {
		return nil
	}

	encode := selector.encodeRightImmediate
	src := jump.Src1
	immediate := jump.Src2
	if isMISupportedImmediate(jump.Src2) {
		// do nothing
	} else if isMISupportedImmediate(jump.Src1) {
		encode = selector.encodeLeftImmediate
		src = jump.Src2
		immediate = jump.Src1
	} else {
		return nil
	}

	return conditionalJumpImmediateInstruction{
		ConditionalJump: jump,
		immediate:       immediate.(*ir.Immediate).Value,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: &architecture.RegisterConstraint{
						Clobbered:  false,
						AnyGeneral: !selector.isFloat,
						AnyFloat:   selector.isFloat,
					},
					DefinitionChunk: src.Def().Chunks()[0],
				},
			},
		},
		encode: encode,
	}
}

func (selector conditionalJumpSelector) newJump(
	jump *ir.ConditionalJump,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	sources := []architecture.RegisterMapping{
		{
			RegisterConstraint: &architecture.RegisterConstraint{
				Clobbered:  false,
				AnyGeneral: !selector.isFloat,
				AnyFloat:   selector.isFloat,
			},
			DefinitionChunk: jump.Src1.Def().Chunks()[0],
		},
	}

	if jump.Src1.Def() != jump.Src2.Def() {
		sources = append(
			sources,
			architecture.RegisterMapping{
				RegisterConstraint: &architecture.RegisterConstraint{
					Clobbered:  false,
					AnyGeneral: !selector.isFloat,
					AnyFloat:   selector.isFloat,
				},
				DefinitionChunk: jump.Src2.Def().Chunks()[0],
			})
	}

	return conditionalJumpInstruction{
		ConditionalJump: jump,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: sources,
		},
		encode: selector.encode,
	}
}
