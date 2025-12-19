package architecture

import (
	"fmt"

	"github.com/pattyshack/chickadee/ir"
)

// The instruction selector is free to ignore the provided hints.
//
// NOTE: hints are not provided for call / return instructions since their
// behavior is governed entirely by call conventions.
type SelectorHint struct {
	// The instruction selector may choose to use more registers (e.g.,
	// 3-address instruction instead of 2-address instruction) when the
	// register pressure is low, and vice versa.
	//
	// NOTE: The number of free registers may include occupied registers that can
	// be cheaply rematerialized.
	NumFreeGeneralRegisters int
	NumFreeFloatRegisters   int

	// Source chunks that can be cheaply clobbered (either because there are
	// multiple register copies, or the definition is dead after use)
	CheapRegisterSources map[*ir.DefinitionChunk]struct{}

	// Destination definition chunk -> Source definition chunk (or nil)
	//
	// Whenever possible:
	// 1. If source chunk is not nil, the instruction selector should use the
	//    same register for the (destination, source) pair.
	// 2. If source chunk is nil, the instruction selector should allocate a new
	//    register for the destination chunk.
	PreferredRegisterDestination map[*ir.DefinitionChunk]*ir.DefinitionChunk
}

type JumpSelector interface {
	Select(*ir.Jump, SelectorHint) MachineInstruction
}

type ConditionalJumpSelector interface {
	Select(*ir.ConditionalJump, SelectorHint) MachineInstruction
}

type UnaryOperationSelector interface {
	Select(*ir.Definition, *ir.UnaryOperation, SelectorHint) MachineInstruction
}

type BinaryOperationSelector interface {
	Select(*ir.Definition, *ir.BinaryOperation, SelectorHint) MachineInstruction
}

// The set of machine instructions
type InstructionSet struct {
	Jump JumpSelector

	JeqUint  ConditionalJumpSelector
	JeqInt   ConditionalJumpSelector
	JeqFloat ConditionalJumpSelector

	JneUint  ConditionalJumpSelector
	JneInt   ConditionalJumpSelector
	JneFloat ConditionalJumpSelector

	JltUint  ConditionalJumpSelector
	JltInt   ConditionalJumpSelector
	JltFloat ConditionalJumpSelector

	JleUint  ConditionalJumpSelector
	JleInt   ConditionalJumpSelector
	JleFloat ConditionalJumpSelector

	JgtUint  ConditionalJumpSelector
	JgtInt   ConditionalJumpSelector
	JgtFloat ConditionalJumpSelector

	JgeUint  ConditionalJumpSelector
	JgeInt   ConditionalJumpSelector
	JgeFloat ConditionalJumpSelector

	// Unary operations

	NotUint UnaryOperationSelector
	NotInt  UnaryOperationSelector

	NegInt   UnaryOperationSelector
	NegFloat UnaryOperationSelector

	UintToUint  UnaryOperationSelector
	IntToUint   UnaryOperationSelector
	FloatToUint UnaryOperationSelector

	UintToInt  UnaryOperationSelector
	IntToInt   UnaryOperationSelector
	FloatToInt UnaryOperationSelector

	UintToFloat  UnaryOperationSelector
	IntToFloat   UnaryOperationSelector
	FloatToFloat UnaryOperationSelector

	// Binary operations

	AddUint  BinaryOperationSelector
	AddInt   BinaryOperationSelector
	AddFloat BinaryOperationSelector

	MulUint  BinaryOperationSelector
	MulInt   BinaryOperationSelector
	MulFloat BinaryOperationSelector

	SubUint  BinaryOperationSelector
	SubInt   BinaryOperationSelector
	SubFloat BinaryOperationSelector

	DivUint  BinaryOperationSelector
	DivInt   BinaryOperationSelector
	DivFloat BinaryOperationSelector

	RemUint BinaryOperationSelector
	RemInt  BinaryOperationSelector

	ShlUint BinaryOperationSelector
	ShlInt  BinaryOperationSelector

	ShrUint BinaryOperationSelector
	ShrInt  BinaryOperationSelector

	AndUint BinaryOperationSelector
	AndInt  BinaryOperationSelector

	OrUint BinaryOperationSelector
	OrInt  BinaryOperationSelector

	XorUint BinaryOperationSelector
	XorInt  BinaryOperationSelector
}

func SelectInstruction(
	set InstructionSet,
	inst ir.Instruction,
	hint SelectorHint,
) MachineInstruction {
	switch instruction := inst.(type) {
	case *ir.Definition:
		return selectOperation(set, instruction, hint)
	case *ir.Jump:
		return set.Jump.Select(instruction, hint)
	case *ir.ConditionalJump:
		return selectConditionalJump(set, instruction, hint)
	case *ir.Terminal:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported instruction: %#v", instruction))
	}
}

func selectConditionalJump(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Kind {
	case ir.Jeq:
		return selectJeq(set, instruction, hint)
	case ir.Jne:
		return selectJne(set, instruction, hint)
	case ir.Jlt:
		return selectJlt(set, instruction, hint)
	case ir.Jle:
		return selectJle(set, instruction, hint)
	case ir.Jgt:
		return selectJgt(set, instruction, hint)
	case ir.Jge:
		return selectJge(set, instruction, hint)
	default:
		panic("unsupported conditional jump kind: " + instruction.Kind)
	}
}

func selectJeq(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JeqInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JeqUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JeqFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jeq type: %v", instruction.Src1.Type()))
	}
}

func selectJne(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JneInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JneUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JneFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jne type: %v", instruction.Src1.Type()))
	}
}

func selectJlt(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JltInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JltUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JltFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jlt type: %v", instruction.Src1.Type()))
	}
}

func selectJle(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JleInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JleUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JleFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jle type: %v", instruction.Src1.Type()))
	}
}

func selectJgt(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JgtInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JgtUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JgtFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jgt type: %v", instruction.Src1.Type()))
	}
}

func selectJge(
	set InstructionSet,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return set.JgeInt.Select(instruction, hint)
	case ir.UnsignedIntType:
		return set.JgeUint.Select(instruction, hint)
	case ir.FloatType:
		return set.JgeFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported jge type: %v", instruction.Src1.Type()))
	}
}

func selectOperation(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch operation := instruction.Operation.(type) {
	case *ir.UnaryOperation:
		return selectUnaryOperation(set, instruction, operation, hint)
	case *ir.BinaryOperation:
		return selectBinaryOperation(set, instruction, operation, hint)
	case *ir.FuncCall:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported operation: %#v", instruction.Operation))
	}
}

func selectUnaryOperation(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Kind {
	case ir.Neg:
		return selectNeg(set, instruction, operation, hint)
	case ir.Not:
		return selectNot(set, instruction, operation, hint)
	case ir.ToInt8, ir.ToInt16, ir.ToInt32, ir.ToInt64:
		return selectToInt(set, instruction, operation, hint)
	case ir.ToUint8, ir.ToUint16, ir.ToUint32, ir.ToUint64:
		return selectToUint(set, instruction, operation, hint)
	case ir.ToFloat32, ir.ToFloat64:
		return selectToFloat(set, instruction, operation, hint)
	default:
		panic("unsupported unary operation: " + operation.Kind)
	}
}

func selectNeg(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.SignedIntType:
		return set.NegInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.NegFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported neg type: %v", instruction.Type))
	}
}

func selectNot(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.NotUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.NotInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported not type: %v", instruction.Type))
	}
}

func selectToInt(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToInt.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.IntToInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.FloatToInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toInt type: %v", instruction.Type))
	}
}

func selectToUint(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.IntToUint.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.FloatToUint.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toUint type: %v", instruction.Type))
	}
}

func selectToFloat(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToFloat.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.IntToFloat.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.FloatToFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toFloat type: %v", instruction.Type))
	}
}

func selectBinaryOperation(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Kind {
	case ir.Add:
		return selectAdd(set, instruction, operation, hint)
	case ir.Mul:
		return selectMul(set, instruction, operation, hint)
	case ir.Sub:
		return selectSub(set, instruction, operation, hint)
	case ir.Div:
		return selectDiv(set, instruction, operation, hint)
	case ir.Rem:
		return selectRem(set, instruction, operation, hint)
	case ir.Shl:
		return selectShl(set, instruction, operation, hint)
	case ir.Shr:
		return selectShr(set, instruction, operation, hint)
	case ir.And:
		return selectAnd(set, instruction, operation, hint)
	case ir.Or:
		return selectOr(set, instruction, operation, hint)
	case ir.Xor:
		return selectXor(set, instruction, operation, hint)
	default:
		panic("unsupported binary operation: " + operation.Kind)
	}
}

func selectAdd(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.AddUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.AddInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.AddFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported add type: %v", instruction.Type))
	}
}

func selectMul(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.MulUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.MulInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.MulFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported mul type: %v", instruction.Type))
	}
}

func selectSub(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.SubUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.SubInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.SubFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported sub type: %v", instruction.Type))
	}
}

func selectDiv(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.DivUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.DivInt.Select(instruction, operation, hint)
	case ir.FloatType:
		return set.DivFloat.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported div type: %v", instruction.Type))
	}
}

func selectRem(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.RemUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.RemInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported rem type: %v", instruction.Type))
	}
}

func selectShl(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.ShlUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.ShlInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported shl type: %v", instruction.Type))
	}
}

func selectShr(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.ShrUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.ShrInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported shr type: %v", instruction.Type))
	}
}

func selectAnd(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.AndUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.AndInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported and type: %v", instruction.Type))
	}
}

func selectOr(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.OrUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.OrInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported or type: %v", instruction.Type))
	}
}

func selectXor(
	set InstructionSet,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.XorUint.Select(instruction, operation, hint)
	case ir.SignedIntType:
		return set.XorInt.Select(instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported xor type: %v", instruction.Type))
	}
}
