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

func TestShrInt8(t *testing.T) {
	// sar r14b, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int8, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd2, 0xfe}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt16(t *testing.T) {
	// sar si, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int16, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xfe}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt32(t *testing.T) {
	// sar ebx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int32, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd3, 0xfb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt64(t *testing.T) {
	// sar rcx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Int64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xd3, 0xf9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint8(t *testing.T) {
	// shr al, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint8, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xd2, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint16(t *testing.T) {
	// shr dx, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint16, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xea}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint32(t *testing.T) {
	// shr r10d, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint32, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xd3, 0xea}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint64(t *testing.T) {
	// shr r8, cl
	builder := layout.NewSegmentBuilder()
	shr(builder, ir.Uint64, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xd3, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt8Immediate(t *testing.T) {
	// sar r14b, 1
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int8, registers.R14, uint8(1))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc0, 0xfe, 1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt16Immediate(t *testing.T) {
	// sar si, 2
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int16, registers.Rsi, uint8(2))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xc1, 0xfe, 2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt32Immediate(t *testing.T) {
	// sar ebx, 3
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int32, registers.Rbx, uint8(3))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xfb, 3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrInt64Immediate(t *testing.T) {
	// sar rcx, 4
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Int64, registers.Rcx, uint8(4))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0xc1, 0xf9, 4}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint8Immediate(t *testing.T) {
	// shr al, 5
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint8, registers.Rax, uint8(5))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc0, 0xe8, 5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint16Immediate(t *testing.T) {
	// shr dx, 6
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint16, registers.Rdx, uint8(6))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xc1, 0xea, 6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint32Immediate(t *testing.T) {
	// shr r10d, 7
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint32, registers.R10, uint8(7))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0xc1, 0xea, 7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestShrUint64Immediate(t *testing.T) {
	// shr r8, 8
	builder := layout.NewSegmentBuilder()
	shrIntImmediate(builder, ir.Uint64, registers.R8, uint8(8))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0xc1, 0xe8, 8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectShrIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint8(4))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Int32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Shr,
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

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Equal(t, 1, len(constraints.RegisterSources))
	srcRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xf9, 4}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectShrInt(t *testing.T) {
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
		Type: ir.Int16,
		Operation: &ir.BinaryOperation{
			Kind: ir.Shr,
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

	_, ok := instruction.(binaryMOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 2, len(constraints.RegisterSources))

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xfb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectShrUintImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint8(4))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Uint32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Shr,
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

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Equal(t, 1, len(constraints.RegisterSources))
	srcRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xc1, 0xe9, 4}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectShrUint(t *testing.T) {
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
		Type: ir.Uint16,
		Operation: &ir.BinaryOperation{
			Kind: ir.Shr,
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

	_, ok := instruction.(binaryMOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 2, len(constraints.RegisterSources))

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xd3, 0xeb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
