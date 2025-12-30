package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestAndInt8(t *testing.T) {
	// and eax, ecx
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x23, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt16(t *testing.T) {
	// and edx, ebx
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x23, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt32(t *testing.T) {
	// and ebp, esi
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x23, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt64(t *testing.T) {
	// and rdi, r8
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x23, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint8(t *testing.T) {
	// and r9d, r10d
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x23, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint16(t *testing.T) {
	// and r11d, r12d
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x23, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint32(t *testing.T) {
	// and r13d, r14d
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x23, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint64(t *testing.T) {
	// and r15, rax
	builder := layout.NewSegmentBuilder()
	and(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x23, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt8Immediate(t *testing.T) {
	// and al, 0x12
	builder := layout.NewSegmentBuilder()
	andIntImmediate(builder, ir.Int8, registers.Rax, int8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xe0, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt16Immediate(t *testing.T) {
	// and bx, 0x1234
	builder := layout.NewSegmentBuilder()
	andIntImmediate(builder, ir.Int16, registers.Rbx, int16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xe3, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt32Immediate(t *testing.T) {
	// and ecx, 0x12345678
	builder := layout.NewSegmentBuilder()
	andIntImmediate(
		builder,
		ir.Int32,
		registers.Rcx,
		int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xe1, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndInt64Immediate(t *testing.T) {
	// and rdx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	andIntImmediate(
		builder,
		ir.Int64,
		registers.Rdx,
		int64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xe2, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint8Immediate(t *testing.T) {
	// and r13b, 0xab
	builder := layout.NewSegmentBuilder()
	andIntImmediate(builder, ir.Uint8, registers.R13, uint8(0xab))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x80, 0xe5, 0xab}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint16Immediate(t *testing.T) {
	// and r13w, 0xabcd
	builder := layout.NewSegmentBuilder()
	andIntImmediate(builder, ir.Uint16, registers.R13, uint16(0xabcd))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x81, 0xe5, 0xcd, 0xab},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint32Immediate(t *testing.T) {
	// and r13d, 0xabcdef01
	builder := layout.NewSegmentBuilder()
	andIntImmediate(
		builder,
		ir.Uint32,
		registers.R13,
		uint32(0xabcdef01))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x81, 0xe5, 0x01, 0xef, 0xcd, 0xab},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAndUint64Immediate(t *testing.T) {
	// and rbp, 0x7cdef012
	builder := layout.NewSegmentBuilder()
	andIntImmediate(
		builder,
		ir.Uint64,
		registers.Rbp,
		uint64(0x7cdef012))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xe5, 0x12, 0xf0, 0xde, 0x7c},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAndIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Int16,
	}
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(int16(0x0a0b))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Int16,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Int16,
		Operation: &ir.BinaryOperation{
			Kind: ir.And,
			Src1: imm,
			Src2: src,
		},
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 1, len(constraints.RegisterSources))
	srcRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xe3, 0x0b, 0x0a},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAndInt(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Int8,
	}
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Int8,
	}
	src2.(*ir.LocalReference).UseDef = src2Def

	dest := &ir.Definition{
		Type: ir.Int8,
		Operation: &ir.BinaryOperation{
			Kind: ir.And,
			Src1: src1,
			Src2: src2,
		},
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 2, len(constraints.RegisterSources))
	src1Register := constraints.RegisterSources[0].RegisterConstraint
	src2Register := constraints.RegisterSources[1].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			src1Register: registers.Rax,
			src2Register: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x23, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAndUintImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint8,
	}
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint8(0xff))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Uint8,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Uint8,
		Operation: &ir.BinaryOperation{
			Kind: ir.And,
			Src1: imm,
			Src2: src,
		},
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 1, len(constraints.RegisterSources))
	srcRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rax,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xe0, 0xff}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAndUint(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Uint64,
	}
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Uint64,
	}
	src2.(*ir.LocalReference).UseDef = src2Def

	dest := &ir.Definition{
		Type: ir.Uint64,
		Operation: &ir.BinaryOperation{
			Kind: ir.And,
			Src1: src1,
			Src2: src2,
		},
	}

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 2, len(constraints.RegisterSources))
	src1Register := constraints.RegisterSources[0].RegisterConstraint
	src2Register := constraints.RegisterSources[1].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			src1Register: registers.R15,
			src2Register: registers.Rax,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x23, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
