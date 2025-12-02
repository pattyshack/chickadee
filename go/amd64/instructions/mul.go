package instructions

import (
	"fmt"

	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint/float dest> *= <int/uint/float src>
//
// https://www.felixcloutier.com/x86/imul (not sign sensitive)
// https://www.felixcloutier.com/x86/mulss
// https://www.felixcloutier.com/x86/mulsd
//
// NOTE: there is no imul 8-bit operand variant.  However it is safe to use
// larger operand variant since the low bits will be correct. We'll use the
// imul 32-bit variant whenever possible.
//
// int 8/16/32-bit (RM Op/En): 0F AF /r
// float 32/64-bit (A Op/En):  0F 59 /r
func mul(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	isFloat := false
	operandSize := 0
	opCode := []byte{0x0F, 0xAF}

	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
	case ir.FloatType:
		isFloat = true
		operandSize = int(size)
		opCode = []byte{0x0F, 0x59}
	default:
		panic("should never happen")
	}

	if !isFloat && operandSize != 8 {
		operandSize = 4
	}

	rmInstruction(builder, isFloat, operandSize, opCode, dest, src)
}

// <int/uint dest> = <int/uint src> * <immediate>
//
// https://www.felixcloutier.com/x86/imul
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (RMI Op/En):     6B /r ib (32-bit operand variant)
// 16-bit (RMI Op/En):    69 /r iw (16-bit operand variant)
// 32/64-bit (RMI Op/En): 69 /r id (32/64-bit operand variants)
func mulIntImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
	immediate []byte,
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

	opCode := []byte{0x69}
	expectedLength := operandSize
	switch operandSize {
	case 1:
		opCode = []byte{0x6B}
		operandSize = 4
	case 2, 4:
	case 8:
		expectedLength = 4
	default:
		panic("should never happen")
	}

	if len(immediate) != expectedLength {
		panic(fmt.Sprintf(
			"incorrect immediate length (%d != %d)",
			len(immediate),
			expectedLength))
	}

	rmiInstruction(builder, operandSize, opCode, dest, src, immediate)
}
