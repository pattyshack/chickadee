package instructions

import (
	"fmt"
	"math"

	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

func isMISupportedImmediate(src ir.Value) bool {
	immediate, ok := src.(*ir.Immediate)
	if !ok {
		return false
	}

	switch value := immediate.Value.(type) {
	case int8:
	case int16:
	case int32:
	case int64:
		return math.MinInt32 <= value && value <= math.MaxInt32
	case uint8:
	case uint16:
	case uint32:
	case uint64:
		return value <= math.MaxInt32
	case float32:
		return false
	case float64:
		return false
	default:
		panic(fmt.Sprintf("unsupported immediate value type: %#V", immediate.Value))
	}

	return true
}

type imulRMIOperation struct {
	*ir.Definition

	immediate interface{}

	architecture.InstructionConstraints
}

func (op imulRMIOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op imulRMIOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op imulRMIOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	dest := op.RegisterDestinations[0].RegisterConstraint
	src := op.RegisterSources[0].RegisterConstraint
	mulIntImmediate(
		builder,
		op.Type,
		selectedRegisters[dest],
		selectedRegisters[src],
		op.immediate)
}

type encodeMIFunc func(
	*layout.SegmentBuilder,
	ir.Type,
	*architecture.Register,
	interface{})

// NOTE: This handles both MI and MI8 Op/En
type binaryMIOperation struct {
	*ir.Definition

	immediate interface{}

	architecture.InstructionConstraints

	encodeMI encodeMIFunc
}

func (op binaryMIOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op binaryMIOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op binaryMIOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	constraint := op.RegisterSources[0].RegisterConstraint
	register := selectedRegisters[constraint]
	op.encodeMI(builder, op.Type, register, op.immediate)
}

type encodeRMFunc func(
	*layout.SegmentBuilder,
	ir.Type,
	*architecture.Register,
	*architecture.Register)

type binaryRMOperation struct {
	*ir.Definition

	architecture.InstructionConstraints

	encodeRM encodeRMFunc
}

func (op binaryRMOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op binaryRMOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op binaryRMOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	if len(op.RegisterSources) == 1 {
		constraint := op.RegisterSources[0].RegisterConstraint
		register := selectedRegisters[constraint]
		op.encodeRM(builder, op.Type, register, register)
	} else {
		registers := make([]*architecture.Register, 2)
		for idx, source := range op.RegisterSources {
			registers[idx] = selectedRegisters[source.RegisterConstraint]
		}
		op.encodeRM(builder, op.Type, registers[0], registers[1])
	}
}

type encodeMFunc func(
	*layout.SegmentBuilder,
	ir.Type,
	*architecture.Register)

// NOTE: This handles both M and MC Op/En
type mOperation struct {
	*ir.Definition

	architecture.InstructionConstraints

	encodeM encodeMFunc
}

func (op mOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op mOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op mOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	constraint := op.RegisterSources[0].RegisterConstraint
	op.encodeM(builder, op.Type, selectedRegisters[constraint])
}

type divRemOperation struct {
	*ir.Definition

	architecture.InstructionConstraints
}

func (op divRemOperation) Instruction() ir.Instruction {
	return op.Definition
}

func (op divRemOperation) Constraints() architecture.InstructionConstraints {
	return op.InstructionConstraints
}

func (op divRemOperation) EmitTo(
	builder *layout.SegmentBuilder,
	selectedRegisters map[*architecture.RegisterConstraint]*architecture.Register,
) {
	constraint := op.RegisterSources[len(op.RegisterSources)-1].RegisterConstraint
	divRemInt(builder, op.Type, selectedRegisters[constraint])
}

// Common binary operation of the form (<dest> = <op> <dest> <src>) with
// optional immediate specialization (<dest> = <op> <dest> <immediate>)
type commonBinaryOperationSelector struct {
	isFloat bool

	isSymmetric bool // a <op> b == b <op> a

	encodeMI encodeMIFunc // optional
	encodeRM encodeRMFunc
}

func (selector commonBinaryOperationSelector) Select(
	config architecture.Config,
	def *ir.Definition,
	binaryOp *ir.BinaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	instruction := selector.maybeNewRMIOperation(def, binaryOp, hint)
	if instruction != nil {
		return instruction
	}

	instruction = selector.maybeNewBinaryMIOperation(def, binaryOp, hint)
	if instruction != nil {
		return instruction
	}

	return selector.newBinaryRMOperation(binaryOp, def, hint)
}

func (selector commonBinaryOperationSelector) maybeNewRMIOperation(
	def *ir.Definition,
	binaryOp *ir.BinaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	if binaryOp.Kind != ir.Mul {
		return nil
	}

	src := binaryOp.Src1
	immediate := binaryOp.Src2
	if isMISupportedImmediate(binaryOp.Src2) {
		// do nothing
	} else if selector.isSymmetric && isMISupportedImmediate(binaryOp.Src1) {
		src = binaryOp.Src2
		immediate = binaryOp.Src1
	} else {
		return nil
	}

	srcChunk := src.Def().Chunks[0]
	_, reuseRegister := hint.CheapRegisterSources[srcChunk]

	if !reuseRegister {
		reuseRegister = hint.NumFreeGeneralRegisters == 0
	}

	destChunk := def.Chunks[0]
	if !reuseRegister {
		preferred := hint.PreferredRegisterDestination[destChunk]
		reuseRegister = preferred == srcChunk
	}

	constraints := architecture.InstructionConstraints{}
	if reuseRegister {
		register := &architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
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
					AnyGeneral: true,
					AnyFloat:   false,
				},
				DefinitionChunk: srcChunk,
			},
		}
		constraints.RegisterDestinations = []architecture.RegisterMapping{
			{
				RegisterConstraint: &architecture.RegisterConstraint{
					Clobbered:  true,
					AnyGeneral: true,
					AnyFloat:   false,
				},
				DefinitionChunk: destChunk,
			},
		}
	}

	return imulRMIOperation{
		Definition:             def,
		immediate:              immediate.(*ir.Immediate).Value,
		InstructionConstraints: constraints,
	}
}

func (selector commonBinaryOperationSelector) maybeNewBinaryMIOperation(
	def *ir.Definition,
	binaryOp *ir.BinaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	if selector.encodeMI == nil {
		return nil
	}

	src := binaryOp.Src1
	immediate := binaryOp.Src2
	if isMISupportedImmediate(binaryOp.Src2) {
		// do nothing
	} else if selector.isSymmetric && isMISupportedImmediate(binaryOp.Src1) {
		src = binaryOp.Src2
		immediate = binaryOp.Src1
	} else {
		return nil
	}

	register := &architecture.RegisterConstraint{
		Clobbered:  true,
		AnyGeneral: true,
	}

	return binaryMIOperation{
		Definition: def,
		immediate:  immediate.(*ir.Immediate).Value,
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    src.Def().Chunks[0],
				},
			},
			RegisterDestinations: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    def.Chunks[0],
				},
			},
		},
		encodeMI: selector.encodeMI,
	}
}

func (selector commonBinaryOperationSelector) newBinaryRMOperation(
	binaryOp *ir.BinaryOperation,
	def *ir.Definition,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	clobberedRegister := &architecture.RegisterConstraint{
		Clobbered:  true,
		AnyGeneral: !selector.isFloat,
		AnyFloat:   selector.isFloat,
	}
	nonClobberedRegister := &architecture.RegisterConstraint{
		AnyGeneral: !selector.isFloat,
		AnyFloat:   selector.isFloat,
	}

	constraints := architecture.InstructionConstraints{
		RegisterDestinations: []architecture.RegisterMapping{
			{
				RegisterConstraint: clobberedRegister,
				DefinitionChunk:    def.Chunks[0],
			},
		},
	}

	src1Chunk := binaryOp.Src1.Def().Chunks[0]
	src2Chunk := binaryOp.Src2.Def().Chunks[0]
	if src1Chunk == src2Chunk {
		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: clobberedRegister,
				DefinitionChunk:    src1Chunk,
			},
		}
	} else {
		clobberedSrc := src1Chunk
		nonClobberedSrc := src2Chunk

		if selector.isSymmetric {
			_, src1IsCheap := hint.CheapRegisterSources[src1Chunk]
			_, src2IsCheap := hint.CheapRegisterSources[src2Chunk]
			if src1IsCheap == src2IsCheap {
				preferred := hint.PreferredRegisterDestination[def.Chunks[0]]
				if src2Chunk == preferred {
					clobberedSrc = src2Chunk
					nonClobberedSrc = src1Chunk
				}
			} else if !src1IsCheap {
				clobberedSrc = src2Chunk
				nonClobberedSrc = src1Chunk
			}
		}

		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: clobberedRegister,
				DefinitionChunk:    clobberedSrc,
			},
			{
				RegisterConstraint: nonClobberedRegister,
				DefinitionChunk:    nonClobberedSrc,
			},
		}
	}

	return binaryRMOperation{
		Definition:             def,
		InstructionConstraints: constraints,
		encodeRM:               selector.encodeRM,
	}
}

type shiftSelector struct {
	encodeMI8 encodeMIFunc
	encodeMC  encodeMFunc
}

func (selector shiftSelector) Select(
	config architecture.Config,
	def *ir.Definition,
	binaryOp *ir.BinaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	instruction := selector.maybeNewBinaryMI8Operation(binaryOp, def, hint)
	if instruction != nil {
		return instruction
	}

	return selector.newBinaryMCOperation(binaryOp, def, hint)
}

func (selector shiftSelector) maybeNewBinaryMI8Operation(
	binaryOp *ir.BinaryOperation,
	def *ir.Definition,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	immediate, ok := binaryOp.Src2.(*ir.Immediate)
	if !ok {
		return nil
	}

	src := binaryOp.Src1

	register := &architecture.RegisterConstraint{
		Clobbered:  true,
		AnyGeneral: true,
	}

	return binaryMIOperation{
		Definition: def,
		immediate:  immediate.Value, // always an uint8
		InstructionConstraints: architecture.InstructionConstraints{
			RegisterSources: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    src.Def().Chunks[0],
				},
			},
			RegisterDestinations: []architecture.RegisterMapping{
				{
					RegisterConstraint: register,
					DefinitionChunk:    def.Chunks[0],
				},
			},
		},
		encodeMI: selector.encodeMI8,
	}
}

func (selector shiftSelector) newBinaryMCOperation(
	binaryOp *ir.BinaryOperation,
	def *ir.Definition,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	constraints := architecture.InstructionConstraints{}
	src1Chunk := binaryOp.Src1.Def().Chunks[0]
	src2Chunk := binaryOp.Src2.Def().Chunks[0]
	if src1Chunk == src2Chunk {
		register := &architecture.RegisterConstraint{
			Clobbered: true,
			Require:   registers.Rcx,
		}

		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: register,
				DefinitionChunk:    src1Chunk,
			},
		}
		constraints.RegisterDestinations = []architecture.RegisterMapping{
			{
				RegisterConstraint: register,
				DefinitionChunk:    def.Chunks[0],
			},
		}
	} else {
		dest := &architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
		}

		count := &architecture.RegisterConstraint{
			Require: registers.Rcx,
		}

		constraints.RegisterSources = []architecture.RegisterMapping{
			{
				RegisterConstraint: dest,
				DefinitionChunk:    src1Chunk,
			},
			{
				RegisterConstraint: count,
				DefinitionChunk:    src2Chunk,
			},
		}
		constraints.RegisterDestinations = []architecture.RegisterMapping{
			{
				RegisterConstraint: dest,
				DefinitionChunk:    def.Chunks[0],
			},
		}
	}

	return mOperation{
		Definition:             def,
		InstructionConstraints: constraints,
		encodeM:                selector.encodeMC,
	}
}

type divRemSelector struct {
	isRem bool
}

func (selector divRemSelector) Select(
	config architecture.Config,
	def *ir.Definition,
	binaryOp *ir.BinaryOperation,
	hint architecture.SelectorHint,
) architecture.MachineInstruction {
	rax := &architecture.RegisterConstraint{
		Clobbered: true,
		Require:   registers.Rax,
	}
	rdx := &architecture.RegisterConstraint{
		Clobbered: true,
		Require:   registers.Rdx,
	}

	dest := rax
	if selector.isRem {
		dest = rdx
	}

	src1Chunk := binaryOp.Src1.Def().Chunks[0]
	constraints := architecture.InstructionConstraints{
		RegisterSources: []architecture.RegisterMapping{
			{ // scratch space for dividend upper bytes
				RegisterConstraint: rdx,
				DefinitionChunk:    nil,
			},
			{
				RegisterConstraint: rax,
				DefinitionChunk:    src1Chunk,
			},
		},
		RegisterDestinations: []architecture.RegisterMapping{
			{
				RegisterConstraint: dest,
				DefinitionChunk:    def.Chunks[0],
			},
		},
	}

	src2Chunk := binaryOp.Src2.Def().Chunks[0]
	if src1Chunk != src2Chunk {
		constraints.RegisterSources = append(
			constraints.RegisterSources,
			architecture.RegisterMapping{
				RegisterConstraint: &architecture.RegisterConstraint{
					AnyGeneral: true,
				},
				DefinitionChunk: src2Chunk,
			})
	}

	return divRemOperation{
		Definition:             def,
		InstructionConstraints: constraints,
	}
}
