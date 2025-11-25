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
// int 8-bit (RM Op/En):        02 /r
// int 16/32/64-bit (RM Op/En): 03 /r
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
	operandSize := 0
	opCode := []byte{0x03}
	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
	case ir.FloatType:
		isFloat = true
		operandSize = int(size)
		opCode = []byte{0x0F, 0x58}
	default:
		panic("should never happen")
	}

	if !isFloat && operandSize == 1 {
		opCode = []byte{0x02}
	}

	rmInstruction(builder, isFloat, operandSize, opCode, dest, src)
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

	opCode := []byte{0x81}
	if operandSize == 1 {
		opCode = []byte{0x80}
	}

	miInstruction(builder, operandSize, opCode, 0, dest, immediate)
}
