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

func TestMulInt8(t *testing.T) {
	// imul eax, ecx
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xaf, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt16(t *testing.T) {
	// imul edx, ebx
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xaf, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt32(t *testing.T) {
	// imul ebp, esi
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xaf, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt64(t *testing.T) {
	// imul rdi, r8
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint8(t *testing.T) {
	// imul r9d, r10d
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint16(t *testing.T) {
	// imul r11d, r12d
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint32(t *testing.T) {
	// imul r13d, r14d
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x0f, 0xaf, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint64(t *testing.T) {
	// imul r15, rax
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulFloat32(t *testing.T) {
	// mulss xmm0, xmm2
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Float32, registers.Xmm0, registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x59, 0xc2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulFloat64(t *testing.T) {
	// mulsd xmm1, xmm3
	builder := layout.NewSegmentBuilder()
	mul(builder, ir.Float64, registers.Xmm1, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x59, 0xcb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt8Immediate(t *testing.T) {
	// mul r15d, ebx, 0x12 (NOTE: 32-bit imm8 variant)
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(builder, ir.Int8, registers.R15, registers.Rbx, []byte{0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x6b, 0xfb, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt16Immediate(t *testing.T) {
	// mul r15w, bx, 0x1234
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Int16,
		registers.R15,
		registers.Rbx,
		[]byte{0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x44, 0x69, 0xfb, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt32Immediate(t *testing.T) {
	// mul r15d, ebx, 0x12345678
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Int32,
		registers.R15,
		registers.Rbx,
		[]byte{0x78, 0x56, 0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x44, 0x69, 0xfb, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulInt64Immediate(t *testing.T) {
	// mul r15, rbx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Int64,
		registers.R15,
		registers.Rbx,
		[]byte{0x9a, 0x78, 0x56, 0x34})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x4c, 0x69, 0xfb, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint8Immediate(t *testing.T) {
	// mul edx, r10d, 0x12 (NOTE: 32-bit imm8 variant)
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(builder, ir.Uint8, registers.Rdx, registers.R10, []byte{0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x6b, 0xd2, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint16Immediate(t *testing.T) {
	// mul dx, r10w, 0x1234
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Uint16,
		registers.Rdx,
		registers.R10,
		[]byte{0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x69, 0xd2, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint32Immediate(t *testing.T) {
	// mul edx, r10d, 0x12345678
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Uint32,
		registers.Rdx,
		registers.R10,
		[]byte{0x78, 0x56, 0x34, 0x12})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x69, 0xd2, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestMulUint64Immediate(t *testing.T) {
	// mul rdx, r10, 0x3456789a
	builder := layout.NewSegmentBuilder()
	mulIntImmediate(
		builder,
		ir.Uint64,
		registers.Rdx,
		registers.R10,
		[]byte{0x9a, 0x78, 0x56, 0x34})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x49, 0x69, 0xd2, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulInt(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Chunk := &ir.DefinitionChunk{}
	src1Def := &ir.Definition{
		Name:   "src1",
		Chunks: []*ir.DefinitionChunk{src1Chunk},
	}
	src1Chunk.Definition = src1Def
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Chunk := &ir.DefinitionChunk{}
	src2Def := &ir.Definition{
		Name:   "src2",
		Chunks: []*ir.DefinitionChunk{src2Chunk},
	}
	src2Chunk.Definition = src2Def
	src2.(*ir.LocalReference).UseDef = src2Def

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Int64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
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
	expect.Equal(t, []byte{0x4c, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulUint(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Chunk := &ir.DefinitionChunk{}
	src1Def := &ir.Definition{
		Name:   "src1",
		Chunks: []*ir.DefinitionChunk{src1Chunk},
	}
	src1Chunk.Definition = src1Def
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Chunk := &ir.DefinitionChunk{}
	src2Def := &ir.Definition{
		Name:   "src2",
		Chunks: []*ir.DefinitionChunk{src2Chunk},
	}
	src2Chunk.Definition = src2Def
	src2.(*ir.LocalReference).UseDef = src2Def

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Uint32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
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
	expect.Equal(t, []byte{0x44, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(float32(0.5))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Float32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src,
			Src2: imm,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 2, len(constraints.RegisterSources))
	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	immRegister := constraints.RegisterSources[1].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Xmm0,
			immRegister: registers.Xmm2,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x59, 0xc2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
