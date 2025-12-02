package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

var (
	// 0x8000000000000000
	float64SignMaskBytes = []byte{0, 0, 0, 0, 0, 0, 0, 0x80}

	// 0x80000000
	float32SignMaskBytes = []byte{0, 0, 0, 0x80}
)

// <int dest> = -<int dest>
//
// https://www.felixcloutier.com/x86/neg
//
// 8-bit (M Op/En):        F6 /3
// 16/32/64-bit (M Op/En): F7 /3
func negSignedInt(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
) {
	operandSize := int(simpleType.(ir.SignedIntType))

	opCode := []byte{0xF7}
	if operandSize == 8 {
		opCode = []byte{0xF6}
	}

	mInstruction(builder, operandSize, opCode, 3, dest)
}

// <float dest> = -<float src>
//
// (NOTE: dest and src are distinct GENERAL, not float, registers)
//
// gcc computes float negation by xor-ing against a negation mask constant
// written to .rodata, which simply flips the float's sign bit (snippet from
// godbolt):
//
//	negFloat(float):
//	        xorps   xmm0, XMMWORD PTR .LC0[rip]
//	        ret
//	negDouble(double):
//	        xorpd   xmm0, XMMWORD PTR .LC1[rip]
//	        ret
//	.LC0:
//	        .long   -2147483648
//	        .long   0
//	        .long   0
//	        .long   0
//	.LC1:
//	        .long   0
//	        .long   -2147483648
//	        .long   0
//	        .long   0
//
// float32's mask (.LC0): -2147483648 = 0x80000000
// float64's mask (.LC1): (-2147483648, 0) = 0x8000000000000000
//
// Instead of loading the mask from memory.  We'll utilize general registers
// to flip the sign bit.  This enables us to encode the mask immediate
// directly into the instructions.
//
// Aside: Since xor supports 32-bit immediate, we can in theory implement
// float32 negation using a single register.  We can't do the same for float64.
func negFloat(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register, // general, not float, register
	src *architecture.Register, // general, not float, register
) {
	if dest.Encoding == src.Encoding {
		panic("registers must be distinct")
	}

	var xorType ir.Type
	var maskBytes []byte
	switch simpleType.(ir.FloatType) {
	case ir.Float32:
		xorType = ir.Uint32
		maskBytes = float32SignMaskBytes
	case ir.Float64:
		xorType = ir.Uint64
		maskBytes = float64SignMaskBytes
	default:
		panic("should never happen")
	}

	setIntImmediate(builder, dest, maskBytes)
	xor(builder, xorType, dest, src)
}
