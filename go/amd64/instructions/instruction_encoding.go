package instructions

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// Instruction selection simplifying assumptions:
//
// - Data loading/storing are handled by distinct mov instructions.  All other
// instructions will only use register-direct addressing mode
// - Address must be loaded into register before use.  We won't mix address
// computation with other operations.

// Resources:
//
// https://wiki.osdev.org/X86-64_Instruction_Encoding
// https://www.felixcloutier.com/x86/
// https://defuse.ca/online-x86-assembler.htm

// NOTE: we'll use intel assembly syntax in comments since online assembler
// all uses intel syntax.

const (
	operandSizePrefix    = 0x66
	addressSizePrefix    = 0x67
	float32OperandPrefix = 0xf3
	float64OperandPrefix = 0xf2

	rexPrefix = byte(0b0100_0000)
	rexWBit   = byte(0b0000_1_000)
	rexRBit   = byte(0b00000_1_00)
	rexXBit   = byte(0b000000_1_0)
	rexBBit   = byte(0b0000000_1)

	// encoding = (mode, reg, r/m)
	indirectDisp0ModRMMode  = byte(0b00_000_000) // [r/m] or [SIB]
	indirectDisp8ModRMMode  = byte(0b01_000_000) // [r/m + disp8]
	indirectDisp32ModRMMode = byte(0b10_000_000) // [r/m + disp32]
	directModRMMode         = byte(0b11_000_000)
)

type modRMSpec struct {
	requireOperandSizePrefix bool // (0x66) 16-bit int (and some float) operation
	requireAddressSizePrefix bool // (0x67)
	requireFloat32Prefix     bool // (0xf3) SSE2 float32 operations
	requireFloat64Prefix     bool // (0xf2) SSE2 float64 operations

	requireRexPrefix bool // make AH/CH/DH/BH registers inaccessible
	requireRexWBit   bool // 64-bit operation
	requireRexRBit   bool // MODRM.reg's extension
	requireRexXBit   bool // SIB.index's extension
	requireRexBBit   bool // MODRM.rm or SIB.base's extension

	opCode []byte

	// MODRM
	mode byte
	reg  byte // MODRM.reg or op code extension
	rm   byte // MODRM.rm

	// Could be nil.  sib byte and/or ib|iw|id|io immediate (1|2|4|8 bytes)
	sibAndOrImmediate []byte
}

func (spec *modRMSpec) maybeSetRexPrefix(xReg int) {
	if 4 <= xReg && xReg <= 7 {
		spec.requireRexPrefix = true
	}
}

func (spec modRMSpec) encode(builder *layout.SegmentBuilder) {
	// 6 for prefixes and modRM suffix
	instruction := make([]byte, 0, 6+len(spec.opCode)+len(spec.sibAndOrImmediate))

	if spec.requireOperandSizePrefix {
		instruction = append(instruction, operandSizePrefix)
	}

	if spec.requireAddressSizePrefix {
		instruction = append(instruction, addressSizePrefix)
	}

	if spec.requireFloat32Prefix {
		instruction = append(instruction, float32OperandPrefix)
	}

	if spec.requireFloat64Prefix {
		instruction = append(instruction, float64OperandPrefix)
	}

	rex := rexPrefix
	if spec.requireRexWBit {
		rex |= rexWBit
	}
	if spec.requireRexRBit {
		rex |= rexRBit
	}
	if spec.requireRexXBit {
		rex |= rexXBit
	}
	if spec.requireRexBBit {
		rex |= rexBBit
	}

	if spec.requireRexPrefix || rex != rexPrefix {
		instruction = append(instruction, rex)
	}

	instruction = append(instruction, spec.opCode...)
	instruction = append(instruction, spec.mode|(spec.reg<<3)|spec.rm)
	instruction = append(instruction, spec.sibAndOrImmediate...)

	builder.AppendBasicData(instruction)
}

func _newRMI(
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register,
	immediate []byte,
) modRMSpec {
	spec := modRMSpec{
		requireRexRBit:    (reg.Encoding & 0x08) != 0,
		requireRexBBit:    (rm.Encoding & 0x08) != 0,
		opCode:            opCode,
		mode:              directModRMMode,
		reg:               byte(reg.Encoding & 0x07),
		rm:                byte(rm.Encoding & 0x07),
		sibAndOrImmediate: immediate,
	}

	if isFloat {
		switch operandSize {
		case 4:
			spec.requireFloat32Prefix = true
		case 8:
			spec.requireFloat64Prefix = true
		default:
			panic("should never happen")
		}
	} else {
		switch operandSize {
		case 1:
			spec.maybeSetRexPrefix(reg.Encoding)
			spec.maybeSetRexPrefix(rm.Encoding)
		case 2:
			spec.requireOperandSizePrefix = true
		case 4:
		case 8:
			spec.requireRexWBit = true
		default:
			panic("should never happen")
		}
	}

	return spec
}

// Register-direct addressing ModRM instruction of the form:
//
// (imul) RMI Op/En:   <opCode> <ModRM:reg (r, w)> <ModRM:r/m (r)> <ib|iw|id>
func newRMI(
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register,
	immediate []byte,
) modRMSpec {
	if isFloat {
		if !reg.AllowFloatOperations || !rm.AllowFloatOperations {
			panic("invalid register")
		}
	} else {
		if !reg.AllowGeneralOperations || !rm.AllowGeneralOperations {
			panic("invalid register")
		}
	}

	return _newRMI(isFloat, operandSize, opCode, reg, rm, immediate)
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) RM Op/En: <opCode> <ModRM:reg (r, w)> <ModRM:r/m (r)>
// (SSE2) A Op/En:     <opCode> <ModRM:reg (r, w)> <ModRM:r/m (r)>
func newRM(
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register,
) modRMSpec {
	return newRMI(isFloat, operandSize, opCode, reg, rm, nil)
}

func _newMI(
	operandSize int,
	opCode []byte,
	opCodeExtension byte,
	rm *architecture.Register,
	immediate []byte,
) modRMSpec {
	spec := modRMSpec{
		requireRexBBit:    (rm.Encoding & 0x08) != 0,
		opCode:            opCode,
		mode:              directModRMMode,
		reg:               opCodeExtension,
		rm:                byte(rm.Encoding & 0x07),
		sibAndOrImmediate: immediate,
	}

	switch operandSize {
	case 1:
		spec.maybeSetRexPrefix(rm.Encoding)
	case 2:
		spec.requireOperandSizePrefix = true
	case 4:
	case 8:
		spec.requireRexWBit = true
	default:
		panic("should never happen")
	}

	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	return spec
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) MI Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <ib|iw|id immediate>
func newMI(
	isUnsigned bool,
	operandSize int,
	opCode []byte,
	opCodeExtension byte,
	rm *architecture.Register,
	immediate interface{}, // either int64 or uint64, matching isUnsigned
) modRMSpec {
	if isUnsigned {
		switch operandSize {
		case 1:
			_ = immediate.(uint8)
		case 2:
			_ = immediate.(uint16)
		case 4:
			_ = immediate.(uint32)
		case 8:
			value := immediate.(uint64)
			// NOTE: immediate are sign extended for 64-bit operand
			if math.MaxInt32 < value {
				panic(fmt.Sprintf(
					"out of bound uint64 sign-extended immediate (%d)",
					value))
			}
		}
	} else {
		switch operandSize {
		case 1:
			_ = immediate.(int8)
		case 2:
			_ = immediate.(int16)
		case 4:
			_ = immediate.(int32)
		case 8:
			value := immediate.(int64)
			if value < math.MinInt32 || math.MaxInt32 < value {
				panic(fmt.Sprintf("out of bound int32 immediate (%d)", value))
			}
		}
	}

	immediateBytes := make([]byte, 8)
	n, err := binary.Encode(immediateBytes, binary.LittleEndian, immediate)
	if err != nil {
		panic(err)
	}
	if n != operandSize {
		panic("should never happen")
	}

	if operandSize == 8 {
		immediateBytes = immediateBytes[:4]
	} else {
		immediateBytes = immediateBytes[:operandSize]
	}

	return _newMI(operandSize, opCode, opCodeExtension, rm, immediateBytes)
}

// Register-direct addressing ModRM instruction of the form:
//
// (shift) MI Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <ib immediate>
func newMI8(
	operandSize int,
	opCode []byte,
	opCodeExtension byte,
	rm *architecture.Register,
	immediate uint8,
) modRMSpec {
	return _newMI(operandSize, opCode, opCodeExtension, rm, []byte{immediate})
}

// Register-direct addressing ModRM instruction of the form:
//
// (general) M Op/En:  <opCode> </digit> <ModRM:r/m (r, w)>
// (general) MC Op/En: <opCode> </digit> <ModRM:r/m (r, w)> <RCX>
func newM(
	operandSize int,
	opCode []byte,
	opCodeExtension byte,
	rm *architecture.Register,
) modRMSpec {
	return _newMI(operandSize, opCode, opCodeExtension, rm, nil)
}

// indirect addressing ModRM instruction of the form:
//
// (general) RM Op/En: <opCode> <ModRM:reg (r, w)>, [<ModRM:r/m (r)>]
// (general) MR Op/En: <opCode> [<ModRM:r/m (r, w)>], <ModRM:reg (r)>
// (SSE2) A Op/En: <opCode> <ModRM:reg (r, w)>, [<ModRM:r/m (r)>]
// (SSE2) B Op/En: <opCode> [<ModRM:r/m (r, w)>], <ModRM:reg (r)>
func newIndirectRM(
	isFloat bool,
	operandSize int,
	opCode []byte,
	reg *architecture.Register,
	rm *architecture.Register, // address
) modRMSpec {
	spec := modRMSpec{
		requireRexRBit: (reg.Encoding & 0x08) != 0,
		requireRexBBit: (rm.Encoding & 0x08) != 0,
		opCode:         opCode,
		mode:           indirectDisp0ModRMMode,
		reg:            byte(reg.Encoding & 0x07),
		rm:             byte(rm.Encoding & 0x07),
	}

	if isFloat {
		if !reg.AllowFloatOperations {
			panic("invalid register")
		}

		switch operandSize {
		case 4:
		case 8:
			spec.requireRexWBit = true
		default:
			panic("should never happen")
		}

		spec.requireOperandSizePrefix = true
	} else {
		if !reg.AllowGeneralOperations {
			panic("invalid register")
		}

		switch operandSize {
		case 1:
			spec.maybeSetRexPrefix(reg.Encoding)
		case 2:
			spec.requireOperandSizePrefix = true
		case 4:
		case 8:
			spec.requireRexWBit = true
		default:
			panic("invalid register")
		}
	}

	if !rm.AllowGeneralOperations {
		panic("invalid register")
	}

	switch spec.rm {
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
		spec.sibAndOrImmediate = []byte{0b00_100_100}
	case 5: // either rbp or r13
		// NOTE: we must use an alternative encoding for rbp/r13 since the default
		// encoding refers to [RIP + disp32] rather than [r/m].

		spec.mode = indirectDisp8ModRMMode // use [<r/m> + disp8] encoding
		spec.sibAndOrImmediate = []byte{0} // immediate
	}

	return spec
}

// Register encoded op code instruction of the form:
//
// (mov) OI Op/En: <opCode + rd (w)> <ib|iw|id|io>
func oiInstruction(
	builder *layout.SegmentBuilder,
	baseOpCode byte,
	register *architecture.Register,
	immediate []byte,
) {
	if !register.AllowGeneralOperations {
		panic("invalid register")
	}

	// +3 for 16-bit prefix, rex prefix, opCode
	instruction := make([]byte, 0, 3+len(immediate))

	requireRex := false
	rex := rexPrefix
	switch len(immediate) {
	case 1:
		// NOTE: rex makes AH/CH/DH/BH inaccessible for 8-bit operand
		requireRex = 4 <= register.Encoding && register.Encoding <= 7
	case 2:
		instruction = append(instruction, operandSizePrefix)
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
