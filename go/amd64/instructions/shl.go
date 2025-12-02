package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int/uint dest> <<= <uint8 RCX>
//
// https://www.felixcloutier.com/x86/sal:sar:shl:shr
//
// NOTE: we'll use the 32-bit variant whenever possible since the operand size
// does not change the result bits.
//
// 8/16/32/64-bit (MC Op/En): D3 /4 (32/64-bit operand variants)
func shl(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
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

	if operandSize != 64 {
		operandSize = 32
	}

	mcInstruction(builder, operandSize, []byte{0xD3}, 4, dest)
}

// <int/uint dest> <<= <imm8>
//
// https://www.felixcloutier.com/x86/sal:sar:shl:shr
//
// NOTE: we'll use the 32-bit variant whenever possible since the operand size
// does not change the result bits.
//
// 8/16/32/64-bit (MI8 Op/En): C1 /4 ib (32/64-bit operand variants)
func shlIntImmediate(
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

	if operandSize != 64 {
		operandSize = 32
	}

	mi8Instruction(builder, operandSize, []byte{0xC1}, 4, dest, immediate)
}
