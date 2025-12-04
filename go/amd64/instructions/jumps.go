package instructions

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

// jmp <label>
//
// https://www.felixcloutier.com/x86/jmp
//
// (D Op/En): E9 cd
func jump(builder *layout.SegmentBuilder, label string) {
	d32Instruction(builder, []byte{0xE9}, layout.BasicBlockKind, label)
}

// Compare two registers of the same class and set eflags.  The srcs are left
// unmodified.  This pairs with the jcc instruction to implement a conditional
// jump.
//
// https://www.felixcloutier.com/x86/cmp
// https://www.felixcloutier.com/x86/comiss
// https://www.felixcloutier.com/x86/comisd
//
// int 8-bit (RM Op/En):        3A /r
// int 16/32/64-bit (RM Op/En): 3B /r
// float32 (comiss A Op/En):    0F 2F /r (encodes as 32-bit int RM Op/En)
// float64 (comisd A Op/En):    0F 2F /r (encodes as 16-bit int RM Op/En)
func compare(
	builder *layout.SegmentBuilder,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	operandSize := 0
	isFloat := false
	switch size := compareType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		operandSize = int(size)
	case ir.FloatType:
		isFloat = true
		operandSize = int(size)
	}

	if !isFloat {
		opCode := []byte{0x3B}
		if operandSize == 1 {
			opCode = []byte{0x3A}
		}

		newRM(false, operandSize, opCode, src1, src2).encode(builder)
		return
	}

	// NOTE: we can't use the newRM function since the float operations encoding
	// behave more like 32/16-bit int RM Op/En instructions.

	if !src1.AllowFloatOperations || src2.AllowFloatOperations {
		panic("invalid register")
	}

	// float32 is encoded like int32 operations whereas float64 is encode like
	// int16 operations.
	if operandSize == 8 { // float64
		operandSize = 2
	}

	_newRMI(
		false,
		operandSize, // uses int encoding convention
		[]byte{0x0F, 0x2F},
		src1,
		src2,
		nil, // immediate
	).encode(builder)
}

// Compare register with int immediate and set eflags.  The src is left
// unmodified.  This pairs with the jcc instruction to implement a conditional
// jump.
//
// https://www.felixcloutier.com/x86/cmp
//
// NOTE: immediate is sign extended for 64-bit operand.  Other operand sizes
// are not sign sensitive.
//
// 8-bit (MI Op/En):     80 /7 ib
// 16-bit (MI Op/En):    81 /7 iw
// 32/64-bit (MI Op/En): 81 /7 id
func compareIntImmediate(
	builder *layout.SegmentBuilder,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	isUnsigned := false
	operandSize := 0
	switch size := compareType.(type) {
	case ir.SignedIntType:
		operandSize = int(size)
	case ir.UnsignedIntType:
		isUnsigned = true
		operandSize = int(size)
	default:
		panic("should never happen")
	}

	opCode := []byte{0x81}
	if operandSize == 1 {
		opCode = []byte{0x80}
	}

	newMI(isUnsigned, operandSize, opCode, 7, src, immediate).encode(builder)
}

// je <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// int/uint/float (JE D Op/En): 0F 84 cd
func je(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	compare(builder, compareType, src1, src2)
	d32Instruction(builder, []byte{0x0F, 0x84}, layout.BasicBlockKind, label)
}

// je <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// int/uint/float (JE D Op/En): 0F 84 cd
func jeIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, []byte{0x0F, 0x84}, layout.BasicBlockKind, label)
}

// jne <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// int/uint/float (JNE D Op/En): 0F 85 cd
func jne(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	compare(builder, compareType, src1, src2)
	d32Instruction(builder, []byte{0x0F, 0x85}, layout.BasicBlockKind, label)
}

// jne <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// int/uint/float (JE D Op/En): 0F 85 cd
func jneIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, []byte{0x0F, 0x85}, layout.BasicBlockKind, label)
}

// jlt <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JB D Op/En): 0F 82 cd
// int (JL D Op/En):        0F 8C cd
func jlt(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8C}
	default:
		opCode = []byte{0x0F, 0x82}
	}

	compare(builder, compareType, src1, src2)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jlt <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JB D Op/En): 0F 82 cd
// int (JL D Op/En):        0F 8C cd
func jltIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8C}
	default:
		opCode = []byte{0x0F, 0x82}
	}

	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jle <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JBE D Op/En): 0F 86 cd
// int (JLE D Op/En):        0F 8E cd
func jle(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8E}
	default:
		opCode = []byte{0x0F, 0x86}
	}

	compare(builder, compareType, src1, src2)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jle <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JB D Op/En): 0F 86 cd
// int (JL D Op/En):        0F 8E cd
func jleIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8E}
	default:
		opCode = []byte{0x0F, 0x86}
	}

	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jgt <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JA D Op/En): 0F 87 cd
// int (JG D Op/En):        0F 8F cd
func jgt(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8F}
	default:
		opCode = []byte{0x0F, 0x87}
	}

	compare(builder, compareType, src1, src2)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jgt <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JB D Op/En): 0F 87 cd
// int (JL D Op/En):        0F 8F cd
func jgtIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8F}
	default:
		opCode = []byte{0x0F, 0x87}
	}

	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jge <label> <int/uint/float src1> <int/uint/float src2>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JAE D Op/En): 0F 83 cd
// int (JGE D Op/En):        0F 8D cd
func jge(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src1 *architecture.Register,
	src2 *architecture.Register,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8D}
	default:
		opCode = []byte{0x0F, 0x83}
	}

	compare(builder, compareType, src1, src2)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}

// jge <label <int/uint src> <int/uint immediate>
//
// https://www.felixcloutier.com/x86/jcc
//
// uint/float (JB D Op/En): 0F 83 cd
// int (JL D Op/En):        0F 8D cd
func jgeIntImmediate(
	builder *layout.SegmentBuilder,
	label string,
	compareType ir.Type,
	src *architecture.Register,
	immediate []byte,
) {
	var opCode []byte
	switch compareType.(type) {
	case ir.SignedIntType:
		opCode = []byte{0x0F, 0x8D}
	default:
		opCode = []byte{0x0F, 0x83}
	}

	compareIntImmediate(builder, compareType, src, immediate)
	d32Instruction(builder, opCode, layout.BasicBlockKind, label)
}
