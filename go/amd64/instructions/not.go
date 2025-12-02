package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <int dest> = !<int dest>
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

	mInstruction(builder, operandSize, []byte{0xF7}, 2, dest)
}
