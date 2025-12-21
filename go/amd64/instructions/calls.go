package instructions

import (
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// https://www.felixcloutier.com/x86/syscall
//
// ZO Op/En: 0F 05
func syscall(builder *layout.SegmentBuilder) {
	builder.AppendBasicData([]byte{0x0f, 0x05})
}

// call <absolute address in register>
//
// https://www.felixcloutier.com/x86/call
//
// M Op/En: FF /2 (NOTE: without REX.W)
func callAddress(
	builder *layout.SegmentBuilder,
	address *architecture.Register,
) {
	newM(4, []byte{0xff}, 2, address).encode(builder)
}

// call <rel32>
//
// https://www.felixcloutier.com/x86/call
//
// D Op/En: E8 cd
func callSymbol(
	builder *layout.SegmentBuilder,
	symbol string,
) {
	d32Instruction(builder, []byte{0xe8}, layout.FunctionKind, symbol)
}

// https://www.felixcloutier.com/x86/ret
//
// NOTE: All returns are "near" in 64-bit mode since there's only one code
// segment.
//
// ZO Op/En: C3
func ret(builder *layout.SegmentBuilder) {
	builder.AppendBasicData([]byte{0xc3})
}
