package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <general dest> = <general src>
//
// https://www.felixcloutier.com/x86/mov
//
// NOTE: we'll use 32-bit variant when possible.
//
// 8/16/32/64-bit (RM Op/En): 8B /r (32/64-bit variants)
func copyGeneral(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	if src.Encoding == dest.Encoding { // no-op
		return
	}

	if srcSize != 8 {
		srcSize = 4
	}

	rmInstruction(builder, false, srcSize, []byte{0x8B}, dest, src)
}

// [<address>] = <general src>
//
// https://www.felixcloutier.com/x86/mov
//
// 8-bit (MR Op/En):        88 /r
// 16/32/64-bit (MR Op/En): 89 /r
func copyGeneralToMemory(
	builder *layout.SegmentBuilder,
	destAddress *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	opCode := []byte{0x89}
	if srcSize == 1 {
		opCode = []byte{0x88}
	}

	indirectModRMInstruction(builder, false, srcSize, opCode, src, destAddress)
}

// <general dest> = [<address>]
//
// https://www.felixcloutier.com/x86/mov
//
// 8-bit (RM Op/En):        8A /r
// 16/32/64-bit (RM Op/En): 8B /r
func copyMemoryToGeneral(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	srcAddress *architecture.Register,
) {
	opCode := []byte{0x8B}
	if srcSize == 1 {
		opCode = []byte{0x8A}
	}

	indirectModRMInstruction(builder, false, srcSize, opCode, dest, srcAddress)
}

// <float dest> = <float src>
//
// https://www.felixcloutier.com/x86/movss
// https://www.felixcloutier.com/x86/movsd
//
// 32/64-bit (A Op/En): 0F 10 /r
func copyFloat(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	if src.Encoding == dest.Encoding { // no-op
		return
	}

	rmInstruction(builder, true, srcSize, []byte{0x0F, 0x10}, dest, src)
}

// [<address>] = <float src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// 32-bit (B Op/En): 66 0F 7E /r
// 64-bit (B Op/En): 66 REX.W OF 7E /r
func copyFloatToMemory(
	builder *layout.SegmentBuilder,
	destAddress *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	indirectModRMInstruction(
		builder,
		true,
		srcSize,
		[]byte{0x0F, 0x7E},
		src,
		destAddress)
}

// <float dest> = [<address>]
//
// https://www.felixcloutier.com/x86/movd:movq
//
// 32-bit (A Op/En): 66 0F 6E /r
// 64-bit (A Op/En): 66 REX.W OF 6E /r
func copyMemoryToFloat(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	srcAddress *architecture.Register,
) {
	indirectModRMInstruction(
		builder,
		true,
		srcSize,
		[]byte{0x0F, 0x6E},
		dest,
		srcAddress)
}

// <general dest> = <float src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// NOTE: we'll use 32-bit variant when possible.
//
// 8/16/32-bit src (B Op/En): 66 0F 7E /r
// 64-bit src (B Op/En):      66 REX.W 0F 7E /r
func copyFloatToGeneral(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	baseRex := rexPrefix // 32-bit variant
	if srcSize == 8 {
		baseRex |= rexWBit // 64-bit variant
	}

	// NOTE: this uses int16 style (operand size prefixed) MR Op/En (src before
	// dest) encoding.
	modRMInstruction(
		builder,
		false,
		2,
		baseRex,
		[]byte{0x0F, 0x7E},
		directModRMMode,
		src.Encoding,
		dest.Encoding,
		nil) // immediate
}

// <float dest> = <general src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// NOTE: we'll use 32-bit variant when possible
//
// 8/16/32-bit src (A Op/En): 66 0F 6E /r
// 64-bit src (A Op/En):      66 REX.W 0F 6E /r
func copyGeneralToFloat(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	baseRex := rexPrefix // 32-bit variant
	if srcSize == 8 {
		baseRex |= rexWBit // 64-bit variant
	}

	// NOTE: this uses int16 style (operand size prefixed) RM Op/En (dest before
	// src) encoding.
	modRMInstruction(
		builder,
		false,
		2,
		baseRex,
		[]byte{0x0F, 0x6E},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate
}

// <int/float dest> = <int/float immediate>
//
// NOTE: This operates only on general registers
//
// https://www.felixcloutier.com/x86/mov
//
// 8-bit (OI Op/En):  B0 + rb ib
// 16-bit (OI Op/En): B8 + rw iw
// 32-bit (OI Op/En): B8 + rd id
// 64-bit (OI Op/En): B8 + rd io
func setImmediate(
	builder *layout.SegmentBuilder,
	dest *architecture.Register, // general register
	immediate []byte,
) {
	operandSize := len(immediate)

	baseOpCode := byte(0xB8)
	if operandSize == 1 {
		baseOpCode = 0xB0
	}

	oiInstruction(builder, operandSize, baseOpCode, dest, immediate)
}

// <extended int/uint dest> = <int/uint src>
//
// See _extendUnsignedInt and _extendSignedInt documentation.
func extendInt(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	extend := _extendUnsignedInt
	srcSize := 0
	switch size := srcType.(type) {
	case ir.UnsignedIntType:
		srcSize = int(size)
	case ir.SignedIntType:
		extend = _extendSignedInt
		srcSize = int(size)
	default:
		panic("should never happen")
	}

	if srcSize >= destSize {
		panic("should never happen")
	}

	extend(builder, destSize, dest, srcSize, src)
}

// <extended uint dest> = <uint src>
//
// https://www.felixcloutier.com/x86/movzx
// https://www.felixcloutier.com/x86/mov
//
// NOTE: we'll always use 32-bits operand for zero extension.  The upper 32
// bits are automatically zero-ed when a 32-bit operand instruction is used
// (see Intel manual, Volume 1, Section 3.4.1.1 General-Purpose Registers in
// 64-Bit Mode).
//
// uint8 -> uint16/uint32/uint64 (movzx RM Op/En): 0F B6 /r
// uint16 -> uint32/uint64 (movzx RM Op/En):       0F B7 /r
// uint32 -> uint64 (mov RM Op/En):                8B /r
func _extendUnsignedInt(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	var opCode []byte
	switch srcSize {
	case 1:
		opCode = []byte{0x0F, 0xB6}
	case 2:
		opCode = []byte{0x0F, 0xB7}
	case 4:
		// NOTE: even when dest == src, we need to explicitly "mov" to zero the
		// upper bytes.
		opCode = []byte{0x8B}
	default:
		panic("should never happen")
	}

	destSize = 4 // see above NOTE

	rmInstruction(builder, false, destSize, opCode, dest, src)
}

// <extended int dest> = <int src>
//
// https://www.felixcloutier.com/x86/movsx:movsxd
//
// NOTE: we'll extend to 32-bit when possible.
//
// int8 -> int16/int32/int64 (movsx RM Op/En): 0F BE /r
// int16 -> int32/int64 (movsx RM Op/En):      0F BF /r
// int32 -> int64 (movsxd RM Op/En):           63 /r
func _extendSignedInt(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	srcSize int,
	src *architecture.Register,
) {
	var opCode []byte
	switch srcSize {
	case 1:
		opCode = []byte{0x0F, 0xBE}
	case 2:
		opCode = []byte{0x0F, 0xBF}
	case 4:
		opCode = []byte{0x63}
	default:
		panic("should never happen")
	}

	if destSize != 8 {
		destSize = 4
	}

	rmInstruction(builder, false, destSize, opCode, dest, src)
}
