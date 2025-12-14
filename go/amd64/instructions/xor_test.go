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

func TestXorInt8(t *testing.T) {
	// xor eax, ecx
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt16(t *testing.T) {
	// xor edx, ebx
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt32(t *testing.T) {
	// xor ebp, esi
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt64(t *testing.T) {
	// xor rdi, r8
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x33, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint8(t *testing.T) {
	// xor r9d, r10d
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x33, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint16(t *testing.T) {
	// xor r11d, r12d
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x33, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint32(t *testing.T) {
	// xor r13d, r14d
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x33, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint64(t *testing.T) {
	// xor r15, rax
	builder := layout.NewSegmentBuilder()
	xor(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x33, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt8Immediate(t *testing.T) {
	// xor bl, 0x12
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(builder, ir.Int8, registers.Rbx, int8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xf3, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt16Immediate(t *testing.T) {
	// xor bx, 0x1234
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(builder, ir.Int16, registers.Rbx, int16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xf3, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt32Immediate(t *testing.T) {
	// xor ebx, 0x12345678
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(
		builder,
		ir.Int32,
		registers.Rbx,
		int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xf3, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorInt64Immediate(t *testing.T) {
	// xor rbx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(
		builder,
		ir.Int64,
		registers.Rbx,
		int64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xf3, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint8Immediate(t *testing.T) {
	// xor sil, 0x12
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(builder, ir.Uint8, registers.Rsi, uint8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x80, 0xf6, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint16Immediate(t *testing.T) {
	// xor si, 0x1234
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(builder, ir.Uint16, registers.Rsi, uint16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xf6, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint32Immediate(t *testing.T) {
	// xor esi, 0x12345678
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(
		builder,
		ir.Uint32,
		registers.Rsi,
		uint32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xf6, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestXorUint64Immediate(t *testing.T) {
	// xor rsi, 0x3456789a
	builder := layout.NewSegmentBuilder()
	xorIntImmediate(
		builder,
		ir.Uint64,
		registers.Rsi,
		uint64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xf6, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectXorIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(int16(0x0a0b))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Int16,
		Operation: &ir.BinaryOperation{
			Kind: ir.Xor,
			Src1: imm,
			Src2: src,
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
		[]byte{0x66, 0x81, 0xf3, 0x0b, 0x0a},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectXorInt(t *testing.T) {
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
		Type: ir.Int8,
		Operation: &ir.BinaryOperation{
			Kind: ir.Xor,
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
			src1Register: registers.Rax,
			src2Register: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectXorUintImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint8(0xff))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Uint8,
		Operation: &ir.BinaryOperation{
			Kind: ir.Xor,
			Src1: imm,
			Src2: src,
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
	expect.Equal(t, []byte{0x80, 0xf0, 0xff}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectXorUint(t *testing.T) {
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
		Type: ir.Uint64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Xor,
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
	expect.Equal(t, []byte{0x4c, 0x33, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
