package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint dest> ^= <int/uint src>
//
// https://www.felixcloutier.com/x86/xor
//
// NOTE: we'll use the 32-bit variant whenever possible since the operand size
// does not change the result bits.
//
// 8/16/32/64-bit (RM Op/En): 33 /r (32/64 bit variants)
func xor(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	switch simpleType.(type) {
	case *ir.SignedIntType:
	case *ir.UnsignedIntType:
	default:
		panic("should never happen")
	}

	operandSize := simpleType.Size()
	if operandSize != 8 {
		operandSize = 4
	}

	newRM(false, operandSize, []byte{0x33}, dest, src).encode(builder)
}

// <int/uint dest> ^= <int/uint immediate>
//
// https://www.felixcloutier.com/x86/xor
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (MI Op/En):     80 /6 ib
// 16-bit (MI Op/En):    81 /6 iw
// 32/64-bit (MI Op/En): 81 /6 id
func xorIntImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	immediate interface{}, // int64 or uint64
) {
	isUnsigned := false
	switch simpleType.(type) {
	case *ir.SignedIntType:
	case *ir.UnsignedIntType:
		isUnsigned = true
	default:
		panic("should never happen")
	}

	operandSize := simpleType.Size()
	opCode := []byte{0x81}
	if operandSize == 1 {
		opCode = []byte{0x80}
	}

	newMI(isUnsigned, operandSize, opCode, 6, dest, immediate).encode(builder)
}
