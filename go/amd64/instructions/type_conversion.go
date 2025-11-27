package instructions

import (
	"github.com/pattyshack/chickadee/amd64"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <dest-sized int/uint dest> = <src-sized int/uint src>
//
// NOTE: both dest and src are general registers.
func convertIntToInt(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	destSize := destType.Size()
	srcSize := srcType.Size()
	if destSize > srcSize {
		extendInt(builder, destSize, dest, srcType, src)
	} else {
		// NOTE: when dest size is smaller than src size, the src's upper bytes
		// may be copied to dest, but the dest type will ignore those bytes,
		// effectively truncated the value.
		copyGeneral(builder, dest, srcSize, src)
	}
}

// <dest-sized float dest> = <src-sized float src>
//
// NOTE: both dest and src are float registers
//
// https://www.felixcloutier.com/x86/cvtsd2ss
// https://www.felixcloutier.com/x86/cvtss2sd
//
// float32 -> float64 (A Op/En): 0F 5A / r (4 byte operand variant)
// float64 -> float32 (A Op/En): 0F 5A / r (8 byte operand variant)
func convertFloatToFloat(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	destSize := destType.Size()
	srcSize := srcType.Size()
	if srcSize == destSize {
		copyFloat(builder, dest, srcSize, src)
	} else {
		rmInstruction(builder, true, srcSize, []byte{0x0F, 0x5A}, dest, src)
	}
}

// <int/uint dest> = <float src>
//
// NOTE: dest is a general register and src is a float register.  The converted
// int bytes may be larger than the dest size (the dest type will ignore those
// bytes).
//
// NOTE: we'll follow c conversion of truncating the decimals
// (cvttss2si/cvttsd2si) rather than rounding (cvtss2si/cvtsd2si).
//
// https://www.felixcloutier.com/x86/cvttss2si
// https://www.felixcloutier.com/x86/cvttsd2si
//
// 8/16/32-bit dest cvttss2si/cvttsdsi (A Op/En): 0F 2D /r
// 64-bit dest cvttss2si/cvttsdsi (A Op/En):      REX.W 0F 2D /r
func convertFloatToInt(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	if !dest.AllowGeneralOperations {
		panic("invalid register")
	}

	if !src.AllowFloatOperations {
		panic("invalid register")
	}

	operandSize := int(srcType.(ir.FloatType))

	baseRex := rexPrefix
	if destType.Size() > 4 {
		baseRex |= rexWBit // coverts float to int64 instead of int32
	}

	modRMInstruction(
		builder,
		true, // isFloat
		operandSize,
		baseRex,
		[]byte{0x0F, 0x2D},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate
}

// <dest-sized float dest> = <src-sized signed int src>
//
// NOTE: dest is a float register and src is a general register.  8/16-bit src
// is clobbered by signed extension. However, since the lower bits
// are preserved, it's safe for reuse.
//
// https://www.felixcloutier.com/x86/cvtsi2ss
// https://www.felixcloutier.com/x86/cvtsi2sd
//
// 8/16/32-bit src cvtsi2ss/cvtsi2sd (A Op/En): 0F 2A /r
// 64-bit src cvtsi2ss/cvtsi2sd (A Op/En):      REX.W 0F 2A /r
func convertSignedIntToFloat(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	if !dest.AllowFloatOperations {
		panic("invalid register")
	}

	if !src.AllowGeneralOperations {
		panic("invalid register")
	}

	baseRex := rexPrefix
	srcSize := int(srcType.(ir.SignedIntType))
	if srcSize < 4 {
		extendInt(builder, 4, src, srcType, src)
	} else if srcSize > 4 {
		baseRex |= rexWBit
	}

	operandSize := int(destType.(ir.FloatType))

	modRMInstruction(
		builder,
		true, // isFloat,
		operandSize,
		baseRex,
		[]byte{0x0F, 0x2A},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate
}

// <dest-sized float dest> = <uint8/uint16/uint32 src>
//
// NOTE: dest is a float register and src is a general register.  src is
// clobbered by zero extension. However, since the lower bits are preserved,
// it's safe for reuse.
//
// https://www.felixcloutier.com/x86/cvtsi2ss
// https://www.felixcloutier.com/x86/cvtsi2sd
//
// 8/16/32-bit src cvtsi2ss/cvtsi2sd (A Op/En): 0F 2A /r
// 64-bit src cvtsi2ss/cvtsi2sd (A Op/En):      REX.W 0F 2A /r
func convertSmallUintToFloat(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	srcType ir.Type,
	src *architecture.Register,
) {
	if !dest.AllowFloatOperations {
		panic("invalid register")
	}

	if !src.AllowGeneralOperations {
		panic("invalid register")
	}

	baseRex := rexPrefix
	extendedSize := 4
	switch int(srcType.(ir.UnsignedIntType)) {
	case 1, 2: // use 32-bit operand variant
	case 4:
		baseRex |= rexWBit
		extendedSize = 8
	default: // uint64 is handled differently
		panic("should never happen")
	}

	extendInt(builder, extendedSize, src, srcType, src)

	operandSize := int(destType.(ir.FloatType))
	modRMInstruction(
		builder,
		true, // isFloat,
		operandSize,
		baseRex,
		[]byte{0x0F, 0x2A},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate
}

// <dest-sized float dest> = <uint64 src>
//
// NOTE: src and scratch are general registers while dest is a float register.
// Both src and scratch are clobbered.
//
// NOTE: uint64 must be special-cased since cvtsi2s* expects a signed integer.
// When the src is "negative", we'll convert the src using 2*ceil(src/2) to
// work around the sign bit (this matches gcc behavior).
//
//	if int64(src) >= 0 {
//		  cvtsi2s* <src>  // can directly convert to float
//		} else {  // work around sign bit via 2*ceil(src/2)
//
//		  <scratch> = <src>
//		  <scratch> = shr <scratch>, 1 // divide by 2
//		  <src> = and <src>, 1         // the remainder when divided by 2
//		  <src> = or <src>, <scratch>  // round up src
//
//		  <dest> = [cvtsi2s*] <src>
//
//		  <dest> = add <dest> <dest> // dest = 2 * dest
//		}
//
// https://www.felixcloutier.com/x86/cvtsi2ss
// https://www.felixcloutier.com/x86/cvtsi2sd
//
// 64-bit src cvtsi2ss/cvtsi2sd (A Op/En):      REX.W 0F 2A /r
func convertUint64ToFloat(
	builder *layout.SegmentBuilder,
	destType ir.Type,
	dest *architecture.Register,
	src *architecture.Register,
	scratch *architecture.Register,
) {
	if !dest.AllowFloatOperations {
		panic("invalid register")
	}

	if !src.AllowGeneralOperations {
		panic("invalid register")
	}

	if !scratch.AllowGeneralOperations {
		panic("invalid register")
	}

	operandSize := int(destType.(ir.FloatType))
	nonNegative := "non negative"
	end := "end"

	instructions := &layout.SegmentBuilder{}

	// if int64(src) >= 0
	jgeIntImmediate(instructions, nonNegative, ir.Int64, src, []byte{0, 0, 0, 0})

	//
	// negative src branch
	//

	// <scratch> = <src>
	copyGeneral(instructions, scratch, 8, src)

	// <scratch> = <scratch> / 2
	shrImmediate(instructions, ir.Uint64, scratch, []byte{1, 0, 0, 0})

	// <src> = <src> % 2
	andImmediate(instructions, ir.Uint32, src, []byte{1, 0, 0, 0})

	// round up <src>
	or(instructions, ir.Uint64, src, scratch)

	// cvtsi2ss or cvtsi2sd
	modRMInstruction(
		instructions,
		true, // isFloat,
		operandSize,
		rexPrefix|rexWBit,
		[]byte{0x0F, 0x2A},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate

	// <dest> = 2 * <dest>
	add(instructions, destType, dest, dest)

	jump(instructions, end)

	//
	// non-negative src branch
	//

	instructions.AppendData(
		nil,
		layout.Definitions{
			Labels: []*layout.Symbol{
				{
					Kind:   layout.BasicBlockKind,
					Name:   nonNegative,
					Offset: 0,
				},
			},
		},
		layout.Relocations{})

	// cvtsi2ss or cvtsi2sd
	modRMInstruction(
		instructions,
		true, // isFloat,
		operandSize,
		rexPrefix|rexWBit,
		[]byte{0x0F, 0x2A},
		directModRMMode,
		dest.Encoding,
		src.Encoding,
		nil) // immediate

	//
	// end of inlined conversion function
	//

	// Add end label
	instructions.AppendData(
		nil,
		layout.Definitions{
			Labels: []*layout.Symbol{
				{
					Kind:   layout.BasicBlockKind,
					Name:   end,
					Offset: 0,
				},
			},
		},
		layout.Relocations{})

	segment, err := instructions.Finalize(amd64.ArchitectureLayout)
	if err != nil {
		panic(err)
	}

	if len(segment.Relocations.Labels) > 0 ||
		len(segment.Relocations.Symbols) > 0 {

		panic("should never happen")
	}

	segment.Definitions = layout.Definitions{}
	segment.Relocations = layout.Relocations{}

	builder.Append(segment)
}
