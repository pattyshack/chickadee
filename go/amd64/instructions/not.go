package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <uint/int dest> = !<uint/int dest>
//
// https://www.felixcloutier.com/x86/not
//
// NOTE: we'll use the 32-bit variant whenever possible since the operand size
// does not change the result bits.
//
// 8/16/32/64-bit (M Op/En): F7 /2 (32/64 bit variants)
func not(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
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

	newM(operandSize, []byte{0xF7}, 2, dest).encode(builder)
}
