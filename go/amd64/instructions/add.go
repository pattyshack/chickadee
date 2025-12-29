package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint/float dest> += <int/uint/float src>
//
// https://www.felixcloutier.com/x86/add (not sign sensitive)
// https://www.felixcloutier.com/x86/addss
// https://www.felixcloutier.com/x86/addsd
//
// NOTE: we'll use 32-bit int variant when possible.  This is safe since
// we do not make use of the carry/overflow status flags.
//
// int 8/16/32/64-bit (RM Op/En): 03 /r (32/64 variants)
// float 32/64-bit (A Op/En):   0F 58 /r
//
// TODO: test 3-address code form via LEA (0x8d)
func add(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	isFloat := false
	opCode := []byte{0x03}
	switch simpleType.(type) {
	case ir.SignedIntType:
	case *ir.UnsignedIntType:
	case ir.FloatType:
		isFloat = true
		opCode = []byte{0x0F, 0x58}
	default:
		panic("should never happen")
	}

	operandSize := simpleType.Size()
	if !isFloat && operandSize < 4 {
		operandSize = 4
	}

	newRM(isFloat, operandSize, opCode, dest, src).encode(builder)
}

// <int/uint dest> += <int/uint immediate>
//
// https://www.felixcloutier.com/x86/add
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (MI Op/En):     80 /0 ib
// 16-bit (MI Op/En):    81 /0 iw
// 32/64-bit (MI Op/En): 81 /0 id
//
// TODO: test 3-address code form via LEA (0x8d)
func addIntImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	immediate interface{}, // either int64 or uint64
) {
	isUnsigned := false
	switch simpleType.(type) {
	case ir.SignedIntType:
	case *ir.UnsignedIntType:
		isUnsigned = true
	default:
		panic("should never happen")
	}

	opCode := []byte{0x81}
	operandSize := simpleType.Size()
	if operandSize == 1 {
		opCode = []byte{0x80}
	}

	newMI(isUnsigned, operandSize, opCode, 0, dest, immediate).encode(builder)
}
