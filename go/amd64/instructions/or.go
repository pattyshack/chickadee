package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint dest> |= <int/uint src>
//
// https://www.felixcloutier.com/x86/or
//
// NOTE: we'll use the 32-bit variant whenever possible since the operand size
// does not change the result bits.
//
// 8/16/32/64-bit (RM Op/En): 0B /r (32/64 bit variants)
func or(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	operandSize := 0
	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
	default:
		panic("should never happen")
	}

	if operandSize != 8 {
		operandSize = 4
	}

	newRM(false, operandSize, []byte{0x0B}, dest, src).encode(builder)
}

// <int/uint dest> |= <int/uint immediate>
//
// https://www.felixcloutier.com/x86/or
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (MI Op/En):     80 /1 ib
// 16-bit (MI Op/En):    81 /1 iw
// 32/64-bit (MI Op/En): 81 /1 id
func orIntImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	immediate []byte,
) {
	isUnsigned := true
	operandSize := 0
	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		isUnsigned = false
		operandSize = int(size)
	default:
		panic("should never happen")
	}

	opCode := []byte{0x81}
	if operandSize == 1 {
		opCode = []byte{0x80}
	}

	newMI(isUnsigned, operandSize, opCode, 1, dest, immediate).encode(builder)
}
