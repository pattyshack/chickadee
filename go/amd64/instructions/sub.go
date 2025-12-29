package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint/float dest> -= <int/uint/float src>
//
// https://www.felixcloutier.com/x86/sub (not sign sensitive)
// https://www.felixcloutier.com/x86/subss
// https://www.felixcloutier.com/x86/subsd
//
// int 8-bit (RM Op/En):        2A /r
// int 16/32/64-bit (RM Op/En): 2B /r
// float 32/64-bit (A Op/En):   0F 5C /r
func sub(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	isFloat := false
	opCode := []byte{0x2B}
	switch simpleType.(type) {
	case ir.SignedIntType:
	case *ir.UnsignedIntType:
	case ir.FloatType:
		isFloat = true
		opCode = []byte{0x0F, 0x5C}
	default:
		panic("should never happen")
	}

	operandSize := simpleType.Size()
	if !isFloat && operandSize == 1 {
		opCode = []byte{0x2A}
	}

	newRM(isFloat, operandSize, opCode, dest, src).encode(builder)
}

// <int/uint dest> -= <int/uint immediate>
//
// https://www.felixcloutier.com/x86/sub
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (MI Op/En):     80 /5 ib
// 16-bit (MI Op/En):    81 /5 iw
// 32/64-bit (MI Op/En): 81 /5 id
func subIntImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	immediate interface{}, // int64 or uint64
) {
	isUnsigned := false
	switch simpleType.(type) {
	case ir.SignedIntType:
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

	newMI(isUnsigned, operandSize, opCode, 5, dest, immediate).encode(builder)
}
