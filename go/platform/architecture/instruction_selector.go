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
	Select(Config, *ir.Jump, SelectorHint) MachineInstruction
}

type ConditionalJumpSelector interface {
	Select(Config, *ir.ConditionalJump, SelectorHint) MachineInstruction
}

type UnaryOperationSelector interface {
	Select(
		Config,
		*ir.Definition,
		*ir.UnaryOperation,
		SelectorHint,
	) MachineInstruction
}

type BinaryOperationSelector interface {
	Select(
		Config,
		*ir.Definition,
		*ir.BinaryOperation,
		SelectorHint,
	) MachineInstruction
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
	config Config,
	inst ir.Instruction,
	hint SelectorHint,
) MachineInstruction {
	switch instruction := inst.(type) {
	case *ir.Definition:
		return selectOperation(config, instruction, hint)
	case *ir.Jump:
		return config.Jump.Select(config, instruction, hint)
	case *ir.ConditionalJump:
		return selectConditionalJump(config, instruction, hint)
	case *ir.Terminal:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported instruction: %#v", instruction))
	}
}

func selectConditionalJump(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Kind {
	case ir.Jeq:
		return selectJeq(config, instruction, hint)
	case ir.Jne:
		return selectJne(config, instruction, hint)
	case ir.Jlt:
		return selectJlt(config, instruction, hint)
	case ir.Jle:
		return selectJle(config, instruction, hint)
	case ir.Jgt:
		return selectJgt(config, instruction, hint)
	case ir.Jge:
		return selectJge(config, instruction, hint)
	default:
		panic("unsupported conditional jump kind: " + instruction.Kind)
	}
}

func selectJeq(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JeqInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JeqUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JeqFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jeq type: %v", instruction.Src1.Type()))
	}
}

func selectJne(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JneInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JneUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JneFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jne type: %v", instruction.Src1.Type()))
	}
}

func selectJlt(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JltInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JltUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JltFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jlt type: %v", instruction.Src1.Type()))
	}
}

func selectJle(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JleInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JleUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JleFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jle type: %v", instruction.Src1.Type()))
	}
}

func selectJgt(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JgtInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JgtUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JgtFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jgt type: %v", instruction.Src1.Type()))
	}
}

func selectJge(
	config Config,
	instruction *ir.ConditionalJump,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Src1.Type().(type) {
	case ir.SignedIntType:
		return config.JgeInt.Select(config, instruction, hint)
	case ir.UnsignedIntType:
		return config.JgeUint.Select(config, instruction, hint)
	case ir.FloatType:
		return config.JgeFloat.Select(config, instruction, hint)
	default:
		panic(fmt.Sprintf("supported jge type: %v", instruction.Src1.Type()))
	}
}

func selectOperation(
	config Config,
	instruction *ir.Definition,
	hint SelectorHint,
) MachineInstruction {
	switch operation := instruction.Operation.(type) {
	case *ir.UnaryOperation:
		return selectUnaryOperation(config, instruction, operation, hint)
	case *ir.BinaryOperation:
		return selectBinaryOperation(config, instruction, operation, hint)
	case *ir.FuncCall:
		panic("TODO")
	default:
		panic(fmt.Sprintf("unsupported operation: %#v", instruction.Operation))
	}
}

func selectUnaryOperation(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Kind {
	case ir.Neg:
		return selectNeg(config, instruction, operation, hint)
	case ir.Not:
		return selectNot(config, instruction, operation, hint)
	case ir.ToInt8, ir.ToInt16, ir.ToInt32, ir.ToInt64:
		return selectToInt(config, instruction, operation, hint)
	case ir.ToUint8, ir.ToUint16, ir.ToUint32, ir.ToUint64:
		return selectToUint(config, instruction, operation, hint)
	case ir.ToFloat32, ir.ToFloat64:
		return selectToFloat(config, instruction, operation, hint)
	default:
		panic("unsupported unary operation: " + operation.Kind)
	}
}

func selectNeg(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.SignedIntType:
		return config.NegInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.NegFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported neg type: %v", instruction.Type))
	}
}

func selectNot(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.NotUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.NotInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported not type: %v", instruction.Type))
	}
}

func selectToInt(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return config.UintToInt.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.IntToInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.FloatToInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toInt type: %v", instruction.Type))
	}
}

func selectToUint(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return config.UintToUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.IntToUint.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.FloatToUint.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toUint type: %v", instruction.Type))
	}
}

func selectToFloat(
	config Config,
	instruction *ir.Definition,
	operation *ir.UnaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Src.Def().Type.(type) {
	case ir.UnsignedIntType:
		return config.UintToFloat.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.IntToFloat.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.FloatToFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported toFloat type: %v", instruction.Type))
	}
}

func selectBinaryOperation(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch operation.Kind {
	case ir.Add:
		return selectAdd(config, instruction, operation, hint)
	case ir.Mul:
		return selectMul(config, instruction, operation, hint)
	case ir.Sub:
		return selectSub(config, instruction, operation, hint)
	case ir.Div:
		return selectDiv(config, instruction, operation, hint)
	case ir.Rem:
		return selectRem(config, instruction, operation, hint)
	case ir.Shl:
		return selectShl(config, instruction, operation, hint)
	case ir.Shr:
		return selectShr(config, instruction, operation, hint)
	case ir.And:
		return selectAnd(config, instruction, operation, hint)
	case ir.Or:
		return selectOr(config, instruction, operation, hint)
	case ir.Xor:
		return selectXor(config, instruction, operation, hint)
	default:
		panic("unsupported binary operation: " + operation.Kind)
	}
}

func selectAdd(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.AddUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.AddInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.AddFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported add type: %v", instruction.Type))
	}
}

func selectMul(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.MulUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.MulInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.MulFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported mul type: %v", instruction.Type))
	}
}

func selectSub(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.SubUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.SubInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.SubFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported sub type: %v", instruction.Type))
	}
}

func selectDiv(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.DivUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.DivInt.Select(config, instruction, operation, hint)
	case ir.FloatType:
		return config.DivFloat.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported div type: %v", instruction.Type))
	}
}

func selectRem(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.RemUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.RemInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported rem type: %v", instruction.Type))
	}
}

func selectShl(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.ShlUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.ShlInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported shl type: %v", instruction.Type))
	}
}

func selectShr(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.ShrUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.ShrInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported shr type: %v", instruction.Type))
	}
}

func selectAnd(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.AndUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.AndInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported and type: %v", instruction.Type))
	}
}

func selectOr(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.OrUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.OrInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported or type: %v", instruction.Type))
	}
}

func selectXor(
	config Config,
	instruction *ir.Definition,
	operation *ir.BinaryOperation,
	hint SelectorHint,
) MachineInstruction {
	switch instruction.Type.(type) {
	case ir.UnsignedIntType:
		return config.XorUint.Select(config, instruction, operation, hint)
	case ir.SignedIntType:
		return config.XorInt.Select(config, instruction, operation, hint)
	default:
		panic(fmt.Sprintf("supported xor type: %v", instruction.Type))
	}
}
