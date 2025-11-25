package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint dest> >>= <uint8 RCX>
//
// https://www.felixcloutier.com/x86/sal:sar:shl:shr
//
// Signed int arithmetic right shift (sar):
// 8-bit (MC Op/En):        D2 /7
// 16/32/64-bit (MC Op/En): D3 /7
//
// Unsigned int logical right shift (shr):
// 8-bit (MC Op/En):        D2 /5
// 16/32/64-bit (MC Op/En): D3 /5
func shr(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
) {
	operandSize := 0
	opCodeExt := 7
	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
		opCodeExt = 5
	}

	opCode := byte(0xD3)
	if operandSize == 8 {
		opCode = 0xD2
	}

	mcInstruction(builder, operandSize, []byte{opCode}, opCodeExt, dest)
}

// <int/uint dest> >>= <imm8>
//
// https://www.felixcloutier.com/x86/sal:sar:shl:shr
//
// Signed int arithmetic right shift (sar):
// 8-bit (MI8 Op/En):        C0 /7 ib
// 16/32/64-bit (MI8 Op/En): C1 /7 ib
//
// Unsigned int arithmetic right shift (sar):
// 8-bit (MI8 Op/En):        C0 /5 ib
// 16/32/64-bit (MI8 Op/En): C1 /5 ib
func shrImmediate(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	immediate []byte,
) {
	operandSize := 0
	opCodeExt := 7
	switch size := simpleType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
		opCodeExt = 5
	}

	opCode := byte(0xC1)
	if operandSize == 8 {
		opCode = 0xC0
	}

	mi8Instruction(
		builder,
		operandSize,
		[]byte{opCode},
		opCodeExt,
		dest,
		immediate)
}
