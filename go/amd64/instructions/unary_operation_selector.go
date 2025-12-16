package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

type encodeConversionFunc func(
	*layout.SegmentBuilder,
	ir.Type,
	*architecture.Register,
	ir.Type,
	*architecture.Register)

type conversionOperation struct {
	*ir.Definition

	srcType ir.Type

	architecture.InstructionConstraints

	encodeConversion encodeConversionFunc
}

func (op conversionOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op conversionOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op conversionOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	dest := op.RegisterDestinations[0].RegisterConstraint
	src := op.RegisterSources[0].RegisterConstraint
	op.encodeConversion(
		builder,
		op.Type,
		selectedRegisters[dest],
		op.srcType,
		selectedRegisters[src])
}

type uint64ToFloatOperation struct {
	*ir.Definition

	architecture.InstructionConstraints
}

func (op uint64ToFloatOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op uint64ToFloatOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op uint64ToFloatOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	dest := op.RegisterDestinations[0].RegisterConstraint
	src := op.RegisterSources[0].RegisterConstraint
	scratch := op.RegisterSources[1].RegisterConstraint
	convertUint64ToFloat(
		builder,
		op.Type,
		selectedRegisters[dest],
		selectedRegisters[src],
		selectedRegisters[scratch])
}

type conversionSelector struct {
	srcIsFloat  bool
	destIsFloat bool

	encodeConversion encodeConversionFunc
}

func (selector conversionSelector) Select(
	def *ir.Definition,
	unaryOp *ir.UnaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	destChunk := def.Chunks[0]
	srcChunk := unaryOp.Src.Def().Chunks[0]

	reuseRegister := false
	if selector.srcIsFloat == selector.destIsFloat {
		// NOTE: certain conversions are no-op if both source and destination are
		// the same location.
		_, reuseRegister = hint.CheapRegisterSources[srcChunk]

		if !reuseRegister {
			preferred := hint.PreferredRegisterDestination[destChunk]
			reuseRegister = preferred == srcChunk
		}

		if !reuseRegister {
			if selector.srcIsFloat {
				reuseRegister = hint.NumFreeFloatRegisters == 0
			} else {
				reuseRegister = hint.NumFreeGeneralRegisters == 0
			}
		}
	}

	constraints := architecture.InstructionConstraints{}
	if reuseRegister {
		register := &architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: !selector.srcIsFloat,
			AnyFloat:   selector.srcIsFloat,
		}
		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: register,
				DefinitionChunk:    srcChunk,
			},
		}
		constraints.RegisterDestinations = []architecture.RegisterMapping{
			{
				RegisterConstraint: register,
				DefinitionChunk:    destChunk,
			},
		}
	} else {
		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: &architecture.RegisterConstraint{
					Clobbered:  false,
					AnyGeneral: !selector.srcIsFloat,
					AnyFloat:   selector.srcIsFloat,
				},
				DefinitionChunk: srcChunk,
			},
		}
		constraints.RegisterDestinations = []architecture.RegisterMapping{
			{
				RegisterConstraint: &architecture.RegisterConstraint{
					Clobbered:  true,
					AnyGeneral: !selector.destIsFloat,
					AnyFloat:   selector.destIsFloat,
				},
				DefinitionChunk: destChunk,
			},
		}
	}

	return conversionOperation{
		Definition:             def,
		srcType:                unaryOp.Src.Def().Type,
		InstructionConstraints: constraints,
		encodeConversion:       selector.encodeConversion,
	}
}

type uintToFloatSelector struct {
	smallUint conversionSelector
}

func (selector uintToFloatSelector) Select(
	def *ir.Definition,
	unaryOp *ir.UnaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	if unaryOp.Src.Def().Size() != 8 {
		return selector.smallUint.Select(def, unaryOp, hint)
	}

	destChunk := def.Chunks[0]
	srcChunk := unaryOp.Src.Def().Chunks[0]
	return uint64ToFloatOperation{
		Definition: def,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: &architecture.RegisterConstraint{
						Clobbered:  true,
						AnyGeneral: true,
						AnyFloat:   false,
					},
					DefinitionChunk: srcChunk,
				},
				{
					RegisterConstraint: &architecture.RegisterConstraint{
						Clobbered:  true,
						AnyGeneral: true,
						AnyFloat:   false,
					},
					DefinitionChunk: nil, // scratch register
				},
			},
			RegisterDestinations: []architecture.RegisterMapping{
				{
					RegisterConstraint: &architecture.RegisterConstraint{
						Clobbered:  true,
						AnyGeneral: false,
						AnyFloat:   true,
					},
					DefinitionChunk: destChunk,
				},
			},
		},
	}
}

type unaryMSelector struct {
	encodeM encodeMFunc
}

func (selector unaryMSelector) Select(
	def *ir.Definition,
	unaryOp *ir.UnaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	register := &architecture.RegisterConstraint{
		Clobbered:  true,
		AnyGeneral: true,
		AnyFloat:   false,
	}

	destChunk := def.Chunks[0]
	srcChunk := unaryOp.Src.Def().Chunks[0]
	return mOperation{
		Definition: def,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    srcChunk,
				},
			},
			RegisterDestinations: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    destChunk,
				},
			},
		},
		encodeM: selector.encodeM,
	}
}

type negFloatSelector struct {
	f32 unaryMSelector
}

func (selector negFloatSelector) Select(
	def *ir.Definition,
	unaryOp *ir.UnaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	if def.Type.Size() == 4 {
		return selector.f32.Select(def, unaryOp, hint)
	}

	dest := &architecture.RegisterConstraint{
		Clobbered:  true,
		AnyGeneral: true,
		AnyFloat:   false,
	}

	destChunk := def.Chunks[0]
	srcChunk := unaryOp.Src.Def().Chunks[0]
	return binaryRMOperation{
		Definition: def,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: dest,
					DefinitionChunk:    nil, // scratch
				},
				{
					RegisterConstraint: &architecture.RegisterConstraint{
						Clobbered:  false,
						AnyGeneral: true,
						AnyFloat:   false,
					},
					DefinitionChunk: srcChunk,
				},
			},
			RegisterDestinations: []architecture.RegisterMapping{
				{
					RegisterConstraint: dest,
					DefinitionChunk:    destChunk,
				},
			},
		},
		encodeRM: negFloat64,
	}
}
