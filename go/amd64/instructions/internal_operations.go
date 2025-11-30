// non-mov internal operations such as stack push/pop, and relative offset
// address computations.

package instructions

import (
	"encoding/binary"

	"github.com/pattyshack/chickadee/amd64"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// <general dest> = <symbol's address> = <RIP> + <disp32 relocation>
//
// NOTE: All symbols must to be RIP relative to support PIC.
//
// https://www.felixcloutier.com/x86/lea
//
// 64 dest (RM Op/En): REX.W + 8D /r
func computeSymbolAddress(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	symbolName string,
) {
	if !dest.AllowGeneralOperations {
		panic("invalid register")
	}

	// NOTE: We need to use indirectDisp0ModRMMode (00) and set r/m to rbp in
	// order to access [RIP + disp32] computation.  We'll encode the second half
	// of the instruction (displacement) separately.
	modRMInstruction(
		builder,
		false, // isFloat
		8,     // address size
		rexPrefix,
		[]byte{0x8D},
		indirectDisp0ModRMMode,
		dest.Encoding,
		amd64.Rbp.Encoding,
		nil)

	// The displacement bytes, to be relocated.
	builder.AppendData(
		[]byte{0, 0, 0, 0},  // XXX: maybe support addend?
		layout.Definitions{},
		layout.Relocations{
			Symbols: []*layout.Relocation{
				{
					Name: symbolName,
				},
			},
		})
}

// <general dest> = <RSP> + <offset>
//
// NOTE: used for accessing local stack variables.
//
// https://www.felixcloutier.com/x86/lea
//
// 64 dest (RM Op/En): REX.W + 8D /r
func computeStackAddress(
	builder *layout.SegmentBuilder,
	dest *architecture.Register,
	offset int32,
) {
	if !dest.AllowGeneralOperations {
		panic("invalid register")
	}

	if offset < 0 {
		panic("invalid offset")
	}

	// NOTE: RSP can only be accessed via SIB.  We need to use
	// indirectDisp32ModRMMode (10) and set r/m to rsp in order to access
	// [SIB + <disp32>] = [<SIB.base> + <disp32>] computation.

	// 1 sib byte + 4 displacement bytes
	sibAndImmediate := make([]byte, 5)

	// SIB byte = (SIB.scale, SIB.index, SIB.base) where
	//
	// SIB.scale = 00 (factor s = 1); can choose any factor
	// SIB.index = 0.100 (rsp); rsp mode ignores index and scale
	// SIB.base = 0.100 (rsp)
	sibAndImmediate[0] = 0b00_100_100

	_, err := binary.Encode(sibAndImmediate[1:], binary.LittleEndian, offset)
	if err != nil {
		panic(err)
	}

	modRMInstruction(
		builder,
		false, // isFloat
		8,     // address size
		rexPrefix,
		[]byte{0x8D},
		indirectDisp32ModRMMode,
		dest.Encoding,
		amd64.RspEncoding,
		sibAndImmediate)
}

// <RSP> += <int immediate>
//
// https://www.felixcloutier.com/x86/add
//
// NOTE: immediate is sign extended for 64-bit operand.
//
// 64-bit (MI Op/En): 81 /0 id
func _updateStack(
	builder *layout.SegmentBuilder,
	size int32,
) {
	if size == 0 {
		return
	}

	immediate := make([]byte, 4)
	_, err := binary.Encode(immediate, binary.LittleEndian, size)
	if err != nil {
		panic(err)
	}

	modRMInstruction(
		builder,
		false, // isFloat
		8,     // address size
		rexPrefix,
		[]byte{81},
		directModRMMode,
		0, // op code extension
		amd64.RspEncoding,
		immediate)
}

func allocateStackFrame(
	builder *layout.SegmentBuilder,
	size int32,
) {
	if size < 0 {
		panic("invalid stack frame size")
	}

	_updateStack(builder, size)
}

func deallocateStackFrame(
	builder *layout.SegmentBuilder,
	size int32,
) {
	if size < 0 {
		panic("invalid stack frame size")
	}

	_updateStack(builder, -size)
}
