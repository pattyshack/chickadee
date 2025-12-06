package instructions

import (
	"encoding/binary"

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
	destSize int,
	dest *architecture.Register,
	src *architecture.Register,
) {
	if src.Encoding == dest.Encoding { // no-op
		return
	}

	if destSize != 8 {
		destSize = 4
	}

	newRM(false, destSize, []byte{0x8B}, dest, src).encode(builder)
}

// [<address>] = <general src>
//
// https://www.felixcloutier.com/x86/mov
//
// 8-bit (MR Op/En):        88 /r
// 16/32/64-bit (MR Op/En): 89 /r
func copyGeneralToMemory(
	builder *layout.SegmentBuilder,
	destSize int,
	destAddress *architecture.Register,
	src *architecture.Register,
) {
	opCode := []byte{0x89}
	if destSize == 1 {
		opCode = []byte{0x88}
	}

	newIndirectRM(false, destSize, opCode, src, destAddress).encode(builder)
}

// <general dest> = [<address>]
//
// https://www.felixcloutier.com/x86/mov
//
// 8-bit (RM Op/En):        8A /r
// 16/32/64-bit (RM Op/En): 8B /r
func copyMemoryToGeneral(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	srcAddress *architecture.Register,
) {
	opCode := []byte{0x8B}
	if destSize == 1 {
		opCode = []byte{0x8A}
	}

	newIndirectRM(false, destSize, opCode, dest, srcAddress).encode(builder)
}

// <float dest> = <float src>
//
// https://www.felixcloutier.com/x86/movss
// https://www.felixcloutier.com/x86/movsd
//
// 32/64-bit (A Op/En): 0F 10 /r
func copyFloat(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	src *architecture.Register,
) {
	if src.Encoding == dest.Encoding { // no-op
		return
	}

	newRM(true, destSize, []byte{0x0F, 0x10}, dest, src).encode(builder)
}

// [<address>] = <float src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// 32-bit (B Op/En): 66 0F 7E /r
// 64-bit (B Op/En): 66 REX.W OF 7E /r
func copyFloatToMemory(
	builder *layout.SegmentBuilder,
	destSize int,
	destAddress *architecture.Register,
	src *architecture.Register,
) {
	newIndirectRM(
		true,
		destSize,
		[]byte{0x0F, 0x7E},
		src,
		destAddress,
	).encode(builder)
}

// <float dest> = [<address>]
//
// https://www.felixcloutier.com/x86/movd:movq
//
// 32-bit (A Op/En): 66 0F 6E /r
// 64-bit (A Op/En): 66 REX.W OF 6E /r
func copyMemoryToFloat(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	srcAddress *architecture.Register,
) {
	newIndirectRM(
		true,
		destSize,
		[]byte{0x0F, 0x6E},
		dest,
		srcAddress,
	).encode(builder)
}

// <general dest> = <float src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// NOTE: we'll use 32-bit variant when possible.
//
// 8/16/32-bit dest (B Op/En): 66 0F 7E /r
// 64-bit dest (B Op/En):      66 REX.W 0F 7E /r
func copyFloatToGeneral(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	src *architecture.Register,
) {
	if !dest.AllowGeneralOperations || !src.AllowFloatOperations {
		panic("invalid registers")
	}

	if destSize < 8 {
		destSize = 4
	}

	// NOTE: this uses int64 style MR Op/En (src before	dest) encoding, plus
	// operand size prefix.
	spec := _newRMI(
		false,
		destSize,
		[]byte{0x0F, 0x7E},
		src,
		dest,
		nil)
	spec.requireOperandSizePrefix = true
	spec.encode(builder)
}

// <float dest> = <general src>
//
// https://www.felixcloutier.com/x86/movd:movq
//
// 32-bit dest (A Op/En): 66 0F 6E /r
// 64-bit dest (A Op/En): 66 REX.W 0F 6E /r
func copyGeneralToFloat(
	builder *layout.SegmentBuilder,
	destSize int,
	dest *architecture.Register,
	src *architecture.Register,
) {
	if !src.AllowGeneralOperations || !dest.AllowFloatOperations {
		panic("invalid registers")
	}

	// NOTE: this uses int64 style RM Op/En (dest before src) encoding, plus
	// operand size prefix.
	spec := _newRMI(
		false,
		destSize,
		[]byte{0x0F, 0x6E},
		dest,
		src,
		nil)
	spec.requireOperandSizePrefix = true
	spec.encode(builder)
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
func setIntImmediate(
	builder *layout.SegmentBuilder,
	dest *architecture.Register, // general register
	immediate interface{}, // int* or uint*
) {
	isZero := false
	switch value := immediate.(type) {
	case int8:
		isZero = value == 0
	case int16:
		isZero = value == 0
	case int32:
		isZero = value == 0
	case int64:
		isZero = value == 0
	case uint8:
		isZero = value == 0
	case uint16:
		isZero = value == 0
	case uint32:
		isZero = value == 0
	case uint64:
		isZero = value == 0
	default:
		panic("should never happen")
	}

	if isZero {
		xor(builder, ir.Int32, dest, dest)
		return
	}

	immediateBytes := make([]byte, 8)
	n, err := binary.Encode(immediateBytes, binary.LittleEndian, immediate)
	if err != nil {
		panic(err)
	}
	immediateBytes = immediateBytes[:n]

	baseOpCode := byte(0xB8)
	if n == 1 {
		baseOpCode = 0xB0
	}

	oiInstruction(builder, baseOpCode, dest, immediateBytes)
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

	spec := newRM(false, destSize, opCode, dest, src)
	if srcSize == 1 {
		spec.maybeSetRexPrefix(src.Encoding)
	}

	spec.encode(builder)
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

	spec := newRM(false, destSize, opCode, dest, src)
	if srcSize == 1 {
		spec.maybeSetRexPrefix(src.Encoding)
	}

	spec.encode(builder)
}
