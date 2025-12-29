package instructions

import (
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

var (
	cwd = []byte{operandSizePrefix, 0x99}
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
//	xor edx, edx                 | cqo
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
	if divisor == registers.Rdx {
		panic("cannot use rdx as divisor")
	}

	switch t := simpleType.(type) {
	case *ir.UnsignedIntType:
		if t.ByteSize == 1 {
			extendInt(builder, 4, registers.Rax, t, registers.Rax)
			if divisor != registers.Rax {
				extendInt(builder, 4, divisor, t, divisor)
			}
			t = ir.Uint32
		}

		setImmediate(builder, registers.Rdx, int32(0))

		// div
		newM(t.ByteSize, []byte{0xF7}, 6, divisor).encode(builder)

	case *ir.SignedIntType:
		if t.ByteSize == 1 {
			extendInt(builder, 4, registers.Rax, t, registers.Rax)
			if divisor != registers.Rax {
				extendInt(builder, 4, divisor, t, divisor)
			}
			t = ir.Int32
		}

		switch t.ByteSize {
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
		newM(t.ByteSize, []byte{0xF7}, 7, divisor).encode(builder)
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
	newRM(true, operandSize, []byte{0x0F, 0x5E}, dest, src).encode(builder)
}
