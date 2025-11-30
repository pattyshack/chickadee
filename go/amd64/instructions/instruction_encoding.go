package instructions

import (
	"fmt"

	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// Instruction selection simplifying assumptions:
//
// - Data loading/storing are handled by distinct mov instructions.  All other
// instructions will only use register-direct addressing mode
// - Address must be loaded into register before use.  We won't mix address
// computation with other operations.

const (
	int16OperandPrefix   = 0x66 // aka operand size prefix in intel manual
	float32OperandPrefix = 0xf3
	float64OperandPrefix = 0xf2
	rexPrefix            = byte(0x40)
	rexWBit              = byte(0x08) // int 64 operand

	// encoding = (mode, reg, r/m)
	indirectDisp0ModRMMode  = 0b00_000_000 // [r/m] or [SIB]
	indirectDisp8ModRMMode  = 0b01_000_000 // [r/m + disp8]
	indirectDisp32ModRMMode = 0b10_000_000 // [r/m + disp32]
	directModRMMode         = 0b11_000_000
)

func modRMInstruction(
	builder *layout.SegmentBuilder,
	isFloat bool,
	operandSize int,
	// baseRex is normally rexPrefix, rexPrefix|rexWBit for float<->int64
	// conversion
	baseRex byte,
	opCode []byte,
	modRMMode int,
	regXReg int, // either 1. X.Reg, or 2. /0 - /7 op code extension
	rmXReg int, // always X.Reg
	immediateOrSib []byte, // nil / ib|iw|id|io (1|2|4|8 byte) / sib
) {
	// +3 for 16-bit/float prefix, rex prefix, and modRM suffix
	instruction := make([]byte, 0, len(opCode)+3+len(immediateOrSib))

	requireRex := false
	rex := baseRex
	if isFloat {
		switch operandSize {
		case 4:
			instruction = append(instruction, float32OperandPrefix)
		case 8:
			instruction = append(instruction, float64OperandPrefix)
		default:
			panic("should never happen")
		}
	} else {
		switch operandSize {
		case 1:
			// NOTE: rex makes AH/CH/DH/BH inaccessible for 8-bit operand
			requireRex = (4 <= regXReg && regXReg < 7) || (4 <= rmXReg && rmXReg < 7)
		case 2:
			instruction = append(instruction, int16OperandPrefix)
		case 4:
		case 8:
			rex |= rexWBit
		default:
			panic("should never happen")
		}
	}

	// reg's rex extension bit (R-bit) and modR/M reg bits
	rexRegX := (regXReg & 0x08) >> 1
	modRMReg := (regXReg & 0x07) << 3

	// rm's rex extension bit (B-bit) and modR/M rm bits
	rexRmX := (rmXReg & 0x08) >> 3
	modRMRm := rmXReg & 0x07

	rex |= byte(rexRegX | rexRmX)

	if requireRex || rex != rexPrefix {
		instruction = append(instruction, rex)
	}

	instruction = append(instruction, opCode...)
	instruction = append(instruction, byte(modRMMode|modRMReg|modRMRm))

	instruction = append(instruction, immediateOrSib...)

	builder.AppendBasicData(instruction)
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) RM Op/En: <opCode> <ModRM:reg (r, w)>, <ModRM:r/m (r)>
// (SSE2) A Op/En:     <opCode> <ModRM:reg (r, w)>, <ModRM:r/m (r)>
func rmInstruction(
	builder *layout.SegmentBuilder,
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register,
) {
	if isFloat {
		if !reg.AllowFloatOperations || !rm.AllowFloatOperations {
			panic("invalid register")
		}
	} else {
		if !reg.AllowGeneralOperations || !rm.AllowGeneralOperations {
			panic("invalid register")
		}
	}

	modRMInstruction(
		builder,
		isFloat,
		operandSize,
		rexPrefix,
		opCode,
		directModRMMode,
		reg.Encoding,
		rm.Encoding,
		nil)
}

// indirect addressing ModRM instruction of the form:
//
// (general) RM Op/En: <opCode> <ModRM:reg (r, w)>, [<ModRM:r/m (r)>]
// (general) MR Op/En: <opCode> [<ModRM:r/m (r, w)>], <ModRM:reg (r)>
// (SSE2) A Op/En: <opCode> <ModRM:reg (r, w)>, [<ModRM:r/m (r)>]
// (SSE2) B Op/En: <opCode> [<ModRM:r/m (r, w)>], <ModRM:reg (r)>
func indirectModRMInstruction(
	builder *layout.SegmentBuilder,
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register, // address
) {
	rex := rexPrefix
	if isFloat {
		if !reg.AllowFloatOperations {
			panic("invalid register")
		}

		switch operandSize {
		case 4:
		case 8:
			rex |= rexWBit
		default:
			panic(fmt.Sprintf("unsupported size: %d", operandSize))
		}

		// float mov uses int16 (operand size prefixed) encoding
		operandSize = 2
	} else {
		if !reg.AllowGeneralOperations {
			panic("invalid register")
		}
	}

	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	rexRmX := rm.Encoding & 0x08
	modRMRm := rm.Encoding & 0x07

	addressMode := indirectDisp0ModRMMode
	var immediateOrSib []byte

	switch modRMRm {
	case 4: // either rsp or r12
		// NOTE: we must use an alternative encoding for rsp/r12 since the default
		// encoding refers to [SIB] rather than [r/m].

		// SIB byte = (SIB.scale, SIB.index, SIB.base) where
		//
		// SIB.scale = 00 (factor s = 1)
		//  - We can use any scale factor since it's ignore
		//
		// SIB.index = 0.100 (rsp)
		//  - rsp index mode ignores index and scale: i.e., address = [base]
		//
		// SIB.base = <rexRMX>.100 (either rsp or r12)
		//  - the upper bit is in REX.B
		immediateOrSib = []byte{0b00_100_100}

	case 5: // either rbp or r13
		// NOTE: we must use an alternative encoding for rbp/r13 since the default
		// encoding refers to [RIP + disp32] rather than [r/m].

		addressMode = indirectDisp8ModRMMode // use [<r/m> + disp8] encoding
		immediateOrSib = []byte{0}           // immediate
	}

	modRMInstruction(
		builder,
		false,
		operandSize,
		rex,
		opCode,
		addressMode,
		reg.Encoding,
		rexRmX|modRMRm,
		immediateOrSib)
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) M Op/En: <opCode> </digit> <ModRM:r/m (r, w)>
func mInstruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	opCode []byte,
	opCodeExtension int, // instead of reg's X.Reg
	rm *architecture.Register,
) {
	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	modRMInstruction(
		builder,
		false, // isFloat
		operandSize,
		rexPrefix,
		opCode,
		directModRMMode,
		opCodeExtension,
		rm.Encoding,
		nil)
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) MC Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <RCX>
func mcInstruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	opCode []byte,
	opCodeExtension int, // instead of reg's X.Reg
	rm *architecture.Register,
) {
	// mc has same byte encoding as m instruction, but different instruction
	// constraints (extra hardcoded RCX)
	mInstruction(builder, operandSize, opCode, opCodeExtension, rm)
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) MI Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <ib|iw|id immediate>
func miInstruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	opCode []byte,
	opCodeExtension int, // instead of reg's X.Reg
	rm *architecture.Register,
	immediate []byte,
) {
	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	// NOTE: In general, 64 bit operand support id (4 byte) immediate, but not
	// io (8 byte) immediate.
	expectedLength := operandSize
	if operandSize == 8 {
		expectedLength = 4
	}

	if len(immediate) != expectedLength {
		panic(fmt.Sprintf(
			"incorrect immediate length (%d != %d)",
			len(immediate),
			expectedLength))
	}

	modRMInstruction(
		builder,
		false, // isFloat
		operandSize,
		rexPrefix,
		opCode,
		directModRMMode,
		opCodeExtension,
		rm.Encoding,
		immediate)
}

// Register-direct addressing ModRM instruction of the form:
//
// (shift) MI Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <ib immediate>
func mi8Instruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	opCode []byte,
	opCodeExtension int, // instead of reg's X.Reg
	rm *architecture.Register,
	immediate []byte,
) {
	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	if len(immediate) != 1 {
		panic(fmt.Sprintf("incorrect immediate length (%d != 1)", len(immediate)))
	}

	modRMInstruction(
		builder,
		false, // isFloat
		operandSize,
		rexPrefix,
		opCode,
		directModRMMode,
		opCodeExtension,
		rm.Encoding,
		immediate)
}

// Register-direct addressing ModRM instruction of the form:
//
// (imul) RMI Op/En:
// <opCode> <ModRM:reg (r, w)> <ModRM:r/m (r)> <ib|iw|id immediate>
func rmiInstruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register,
	immediate []byte,
) {
	if !reg.AllowGeneralOperations || !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	// NOTE: In general, 64 bit operand support id (4 byte) immediate, but not
	// io (8 byte) immediate.
	expectedLength := operandSize
	if operandSize == 8 {
		expectedLength = 4
	}

	if len(immediate) != expectedLength {
		panic(fmt.Sprintf(
			"incorrect immediate length (%d != %d)",
			len(immediate),
			expectedLength))
	}

	modRMInstruction(
		builder,
		false, // isFloat
		operandSize,
		rexPrefix,
		opCode,
		directModRMMode,
		reg.Encoding,
		rm.Encoding,
		immediate)
}

// Register encoded op code instruction of the form:
//
// (mov) OI Op/En: <opCode + rd (w)> <ib|iw|id|io>
func oiInstruction(
	builder *layout.SegmentBuilder,
	operandSize int,
	baseOpCode byte,
	register *architecture.Register,
	immediate []byte,
) {
	if !register.AllowGeneralOperations {
		panic("invalid register")
	}

	switch operandSize {
	case 1, 2, 4, 8:
		// ok
	default:
		panic("unexpected operand size")
	}

	if len(immediate) != operandSize {
		panic("unexpected immediate length")
	}

	// +3 for 16-bit prefix, rex prefix, opCode
	instruction := make([]byte, 0, 3+len(immediate))

	requireRex := false
	rex := rexPrefix
	switch operandSize {
	case 1:
		// NOTE: rex makes AH/CH/DH/BH inaccessible for 8-bit operand
		requireRex = 4 <= register.Encoding && register.Encoding < 7
	case 2:
		instruction = append(instruction, int16OperandPrefix)
	case 4:
	case 8:
		rex |= rexWBit
	default:
		panic("should never happen")
	}

	rex |= byte((register.Encoding & 0x08) >> 3) // REX.B bit

	if requireRex || rex != rexPrefix {
		instruction = append(instruction, rex)
	}

	opCode := baseOpCode | byte(register.Encoding&0x07)
	instruction = append(instruction, opCode)

	instruction = append(instruction, immediate...)

	builder.AppendBasicData(instruction)
}

// Instruction of the form:
//
// (general) D Op/En: <op code> <rel32>
func d32Instruction(
	builder *layout.SegmentBuilder,
	opCode []byte,
	kind layout.SymbolKind,
	symbol string,
) {
	bytes := make([]byte, len(opCode)+4)
	copy(bytes, opCode)

	entry := &layout.Relocation{
		Name:   symbol,
		Offset: int64(len(opCode)),
	}

	relocations := layout.Relocations{}
	switch kind {
	case layout.BasicBlockKind:
		relocations.Labels = []*layout.Relocation{entry}
	case layout.FunctionKind, layout.ObjectKind:
		relocations.Symbols = []*layout.Relocation{entry}
	default:
		panic("unsupported symbol kind " + kind)
	}

	builder.AppendData(bytes, layout.Definitions{}, relocations)
}
