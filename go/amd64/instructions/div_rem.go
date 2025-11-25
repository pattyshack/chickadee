package instructions

import (
	"github.com/pattyshack/chickadee/amd64"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

var (
	cwd = []byte{int16OperandPrefix, 0x99}
	cdq = []byte{0x99}
	cqo = []byte{rexPrefix | rexWBit, 0x99}
)

// (<int/uint quotient RAX>, <int/uint remainder RDX>) =
//
//	<int/uint upper RDX>:<int/uint lower RAX> / <int/uint divisor>
//
// NOTE: RAX and <divisor> (cannot be RDX) are the input registers.
//
// Unsigned int | Signed int:
//
// 8-bit (uses 32-bit div/idiv):
//
//	movzx eax, al                | movsx eax, al
//	movzx <divisor>d, <divisor>b | movsx <divisor>d, <divisor>b
//	xor edx, edx                 | cdq
//	div <divisor>d               | idiv <divisor>d
//
// 16-bit:
//
//	xor edx, edx                 | cwd
//	div <divisor>w               | idiv <divisor>w
//
// 32-bit:
//
//	xor edx, edx                 | cdq
//	div <divisor>d               | idiv <divisor>d
//
// 64-bit:
//
//	xor rdx, rdx                 | cqo
//	div <divisor>q               | idiv <divisor>q
//
// https://www.felixcloutier.com/x86/cwd:cdq:cqo
// https://www.felixcloutier.com/x86/div
// https://www.felixcloutier.com/x86/idiv
//
// cdq/cdq/cqo (ZO Op/En):        99
// div 8/16/32/64-bit (M Op/En):  F7 /6
// idiv 8/16/32/64-bit (M Op/En): F7 /7
func divRemInt(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	divisor *architecture.Register,
) {
	if divisor == amd64.Rdx {
		panic("cannot use rdx as divisor")
	}

	switch size := simpleType.(type) {
	case ir.UnsignedIntType:
		if size == 1 {
			extendInt(builder, 4, amd64.Rax, size, amd64.Rax)
			if divisor != amd64.Rax {
				extendInt(builder, 4, divisor, size, divisor)
			}
			size = 4
		}

		xor(builder, size, amd64.Rdx, amd64.Rdx)

		// div
		mInstruction(builder, int(size), []byte{0xF7}, 6, divisor)

	case ir.SignedIntType:
		if size == 1 {
			extendInt(builder, 4, amd64.Rax, size, amd64.Rax)
			if divisor != amd64.Rax {
				extendInt(builder, 4, divisor, size, divisor)
			}
			size = 4
		}

		switch size {
		case 2:
			builder.AppendBasicData(cwd)
		case 4:
			builder.AppendBasicData(cdq)
		case 8:
			builder.AppendBasicData(cqo)
		default:
			panic("should never happen")
		}

		// idiv
		mInstruction(builder, int(size), []byte{0xF7}, 7, divisor)
	default:
		panic("should never happen")
	}
}

// <float dest> /= <float src>
//
// https://www.felixcloutier.com/x86/divss
// https://www.felixcloutier.com/x86/divsd
//
// 32/64-bit (A Op/En): 0F 5E /r
func divFloat(
	builder *layout.SegmentBuilder,
	simpleType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
) {
	operandSize := int(simpleType.(ir.FloatType))
	rmInstruction(builder, true, operandSize, []byte{0x0F, 0x5E}, dest, src)
}
