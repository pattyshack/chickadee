package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

var (
	float64SignMask = uint64(0x8000000000000000)

	float32SignMask = uint32(0x80000000)
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
	if operandSize == 1 {
		opCode = []byte{0xF6}
	}

	newM(operandSize, opCode, 3, dest).encode(builder)
}

// <float32 dest> = -<float32 dest>
//
// (NOTE: dest is GENERAL, not float, register)
//
// gcc computes float negation by xor-ing against a negation mask constant
// written to .rodata, which simply flips the float's sign bit (snippet from
// godbolt):
//
//	negFloat(float):
//	        xorps   xmm0, XMMWORD PTR .LC0[rip]
//	        ret
//	.LC0:
//	        .long   -2147483648
//	        .long   0
//	        .long   0
//	        .long   0
//
// float32's mask (.LC0): -2147483648 = 0x80000000
//
// Instead of loading the mask from memory.  We'll utilize general registers
// to flip the sign bit.  This enables us to encode the mask immediate
// directly into the instructions.
func negFloat32(
	builder *layout.SegmentBuilder,
	dest *architecture.Register, // general, not float, register
) {
	xorIntImmediate(builder, ir.Uint32, dest, float32SignMask)
}

// <float64 dest> = -<float64 src>
//
// (NOTE: dest and src are distinct GENERAL, not float, registers)
//
// gcc computes float negation by xor-ing against a negation mask constant
// written to .rodata, which simply flips the float's sign bit (snippet from
// godbolt):
//
//	negDouble(double):
//	        xorpd   xmm0, XMMWORD PTR .LC1[rip]
//	        ret
//	.LC1:
//	        .long   0
//	        .long   -2147483648
//	        .long   0
//	        .long   0
//
// float64's mask (.LC1): (-2147483648, 0) = 0x8000000000000000
//
// Instead of loading the mask from memory.  We'll utilize general registers
// to flip the sign bit.  This enables us to encode the mask immediate
// directly into the instructions.
func negFloat64(
	builder *layout.SegmentBuilder,
	dest *architecture.Register, // general, not float, register
	src *architecture.Register, // general, not float, register
) {
	if dest.Encoding == src.Encoding {
		panic("registers must be distinct")
	}

	setIntImmediate(builder, dest, float64SignMask)
	xor(builder, ir.Uint64, dest, src)
}
