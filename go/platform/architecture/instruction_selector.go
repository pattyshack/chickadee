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

type OperationSelector interface {
	Select(*ir.Definition, SelectorHint) MachineInstruction
}

// The set of machine instructions
type InstructionSet struct {

	// Unary operations

	NotUint OperationSelector
	NotInt  OperationSelector

	NegInt   OperationSelector
	NegFloat OperationSelector

	UintToUint8  OperationSelector
	IntToUint8   OperationSelector
	FloatToUint8 OperationSelector

	UintToUint16  OperationSelector
	IntToUint16   OperationSelector
	FloatToUint16 OperationSelector

	UintToUint32  OperationSelector
	IntToUint32   OperationSelector
	FloatToUint32 OperationSelector

	UintToUint64  OperationSelector
	IntToUint64   OperationSelector
	FloatToUint64 OperationSelector

	UintToInt8  OperationSelector
	IntToInt8   OperationSelector
	FloatToInt8 OperationSelector

	UintToInt16  OperationSelector
	IntToInt16   OperationSelector
	FloatToInt16 OperationSelector

	UintToInt32  OperationSelector
	IntToInt32   OperationSelector
	FloatToInt32 OperationSelector

	UintToInt64  OperationSelector
	IntToInt64   OperationSelector
	FloatToInt64 OperationSelector

	UintToFloat32  OperationSelector
	IntToFloat32   OperationSelector
	FloatToFloat32 OperationSelector

	UintToFloat64  OperationSelector
	IntToFloat64   OperationSelector
	FloatToFloat64 OperationSelector

	// Binary operations

	AddUint  OperationSelector
	AddInt   OperationSelector
	AddFloat OperationSelector

	MulUint  OperationSelector
	MulInt   OperationSelector
	MulFloat OperationSelector

	SubUint  OperationSelector
	SubInt   OperationSelector
	SubFloat OperationSelector

	DivUint  OperationSelector
	DivInt   OperationSelector
	DivFloat OperationSelector

	RemUint OperationSelector
	RemInt  OperationSelector

	ShlUint OperationSelector
	ShlInt  OperationSelector

	ShrUint OperationSelector
	ShrInt  OperationSelector

	AndUint OperationSelector
	AndInt  OperationSelector

	OrUint OperationSelector
	OrInt  OperationSelector

	XorUint OperationSelector
	XorInt  OperationSelector
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
		panic("TODO")
	case *ir.ConditionalJump:
		panic("TODO")
	case *ir.Terminal:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported instruction: %#v", instruction))
	}
}

func selectOperation(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch operation := instruction.Operation.(type) {
	case *ir.UnaryOperation:
		return selectUnaryOperation(operation.Kind, set, instruction, hint)
	case *ir.BinaryOperation:
		return selectBinaryOperation(operation.Kind, set, instruction, hint)
	case *ir.FuncCall:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported operation: %#v", instruction.Operation))
	}
}

func selectUnaryOperation(
	kind ir.UnaryOperationKind,
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch kind {
	case ir.Neg:
		return selectNeg(set, instruction, hint)
	case ir.Not:
		return selectNot(set, instruction, hint)
	case ir.ToInt8:
		return selectToInt8(set, instruction, hint)
	case ir.ToInt16:
		return selectToInt16(set, instruction, hint)
	case ir.ToInt32:
		return selectToInt32(set, instruction, hint)
	case ir.ToInt64:
		return selectToInt64(set, instruction, hint)
	case ir.ToUint8:
		return selectToUint8(set, instruction, hint)
	case ir.ToUint16:
		return selectToUint16(set, instruction, hint)
	case ir.ToUint32:
		return selectToUint32(set, instruction, hint)
	case ir.ToUint64:
		return selectToUint64(set, instruction, hint)
	case ir.ToFloat32:
		return selectToFloat32(set, instruction, hint)
	case ir.ToFloat64:
		return selectToFloat64(set, instruction, hint)
	default:
		panic("unsupported unary operation: " + kind)
	}
}

func selectNeg(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.SignedIntType:
		return set.NegInt.Select(instruction, hint)
	case ir.FloatType:
		return set.NegFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported neg type: %v", instruction.Type))
	}
}

func selectNot(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.NotUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.NotInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported not type: %v", instruction.Type))
	}
}

func selectToInt8(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToInt8.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToInt8.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToInt8.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toInt8 type: %v", instruction.Type))
	}
}

func selectToInt16(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToInt16.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToInt16.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToInt16.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toInt16 type: %v", instruction.Type))
	}
}

func selectToInt32(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToInt32.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToInt32.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToInt32.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toInt32 type: %v", instruction.Type))
	}
}

func selectToInt64(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToInt64.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToInt64.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToInt64.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toInt64 type: %v", instruction.Type))
	}
}

func selectToUint8(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToUint8.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToUint8.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToUint8.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toUint8 type: %v", instruction.Type))
	}
}

func selectToUint16(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToUint16.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToUint16.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToUint16.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toUint16 type: %v", instruction.Type))
	}
}

func selectToUint32(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToUint32.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToUint32.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToUint32.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toUint32 type: %v", instruction.Type))
	}
}

func selectToUint64(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToUint64.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToUint64.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToUint64.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toUint64 type: %v", instruction.Type))
	}
}

func selectToFloat32(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToFloat32.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToFloat32.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToFloat32.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toFloat32 type: %v", instruction.Type))
	}
}

func selectToFloat64(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.UintToFloat64.Select(instruction, hint)
	case ir.SignedIntType:
		return set.IntToFloat64.Select(instruction, hint)
	case ir.FloatType:
		return set.FloatToFloat64.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported toFloat64 type: %v", instruction.Type))
	}
}

func selectBinaryOperation(
	kind ir.BinaryOperationKind,
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch kind {
	case ir.Add:
		return selectAdd(set, instruction, hint)
	case ir.Mul:
		return selectMul(set, instruction, hint)
	case ir.Sub:
		return selectSub(set, instruction, hint)
	case ir.Div:
		return selectDiv(set, instruction, hint)
	case ir.Rem:
		return selectRem(set, instruction, hint)
	case ir.Shl:
		return selectShl(set, instruction, hint)
	case ir.Shr:
		return selectShr(set, instruction, hint)
	case ir.And:
		return selectAnd(set, instruction, hint)
	case ir.Or:
		return selectOr(set, instruction, hint)
	case ir.Xor:
		return selectXor(set, instruction, hint)
	default:
		panic("unsupported binary operation: " + kind)
	}
}

func selectAdd(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.AddUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.AddInt.Select(instruction, hint)
	case ir.FloatType:
		return set.AddFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported add type: %v", instruction.Type))
	}
}

func selectMul(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.MulUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.MulInt.Select(instruction, hint)
	case ir.FloatType:
		return set.MulFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported mul type: %v", instruction.Type))
	}
}

func selectSub(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.SubUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.SubInt.Select(instruction, hint)
	case ir.FloatType:
		return set.SubFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported sub type: %v", instruction.Type))
	}
}

func selectDiv(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.DivUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.DivInt.Select(instruction, hint)
	case ir.FloatType:
		return set.DivFloat.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported div type: %v", instruction.Type))
	}
}

func selectRem(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.RemUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.RemInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported rem type: %v", instruction.Type))
	}
}

func selectShl(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.ShlUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.ShlInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported shl type: %v", instruction.Type))
	}
}

func selectShr(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.ShrUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.ShrInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported shr type: %v", instruction.Type))
	}
}

func selectAnd(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.AndUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.AndInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported and type: %v", instruction.Type))
	}
}

func selectOr(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.OrUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.OrInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported or type: %v", instruction.Type))
	}
}

func selectXor(
	set InstructionSet,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return set.XorUint.Select(instruction, hint)
	case ir.SignedIntType:
		return set.XorInt.Select(instruction, hint)
	default:
		panic(fmt.Sprintf("supported xor type: %v", instruction.Type))
	}
}
