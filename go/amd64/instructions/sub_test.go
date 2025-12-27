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

func TestSubInt8(t *testing.T) {
	// sub bl, al
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2a, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub bpl, cl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rbp, registers.Rcx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xe9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub sil, cl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rsi, registers.Rcx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xf1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub dil, cl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rdi, registers.Rcx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xf9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub dl, bpl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rdx, registers.Rbp)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xd5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub dl, sil
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rdx, registers.Rsi)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub dl, dil
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rdx, registers.Rdi)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x2a, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9b, dl
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x2a, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub cl, r12b
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x2a, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10b, r11b
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int8, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x2a, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt16(t *testing.T) {
	// sub bp, si
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9w, dx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x44, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub cx, r12w
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x41, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10w, r11w
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int16, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x45, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt32(t *testing.T) {
	// sub ebp, esi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9d, edx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub ecx, r12d
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10d, r11d
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int32, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt64(t *testing.T) {
	// sub rbp, rsi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r9, rdx
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.R9, registers.Rdx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x2b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub rcx, r12
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.Rcx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x2b, 0xcc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// sub r10, r11
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Int64, registers.R10, registers.R11)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4d, 0x2b, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint8(t *testing.T) {
	// sub bl, al
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint8, registers.Rbx, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2a, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint16(t *testing.T) {
	// sub bp, si
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint16, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint32(t *testing.T) {
	// sub ebp, esi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint64(t *testing.T) {
	// sub rbp, rsi
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Uint64, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubFloat32(t *testing.T) {
	// subss xmm1, xmm7
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm1, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x5c, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm1, xmm12
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm1, registers.Xmm12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x41, 0x0f, 0x5c, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm10, xmm3
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm10, registers.Xmm3)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x5c, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// subss xmm9, xmm14
	builder = layout.NewSegmentBuilder()
	sub(builder, ir.Float32, registers.Xmm9, registers.Xmm14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x45, 0x0f, 0x5c, 0xce},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubFloat64(t *testing.T) {
	// subsd xmm5, xmm6
	builder := layout.NewSegmentBuilder()
	sub(builder, ir.Float64, registers.Xmm5, registers.Xmm6)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x5c, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt8Immediate(t *testing.T) {
	// sub cl, 0x12
	builder := layout.NewSegmentBuilder()
	subIntImmediate(builder, ir.Int8, registers.Rcx, int8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xe9, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt16Immediate(t *testing.T) {
	// sub cx, 0x1234
	builder := layout.NewSegmentBuilder()
	subIntImmediate(builder, ir.Int16, registers.Rcx, int16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xe9, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt32Immediate(t *testing.T) {
	// sub ecx, 0x12345678
	builder := layout.NewSegmentBuilder()
	subIntImmediate(
		builder,
		ir.Int32,
		registers.Rcx,
		int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xe9, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubInt64Immediate(t *testing.T) {
	// sub rcx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	subIntImmediate(
		builder,
		ir.Int64,
		registers.Rcx,
		int64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xe9, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint8Immediate(t *testing.T) {
	// sub r12b, 0x12
	builder := layout.NewSegmentBuilder()
	subIntImmediate(builder, ir.Uint8, registers.R12, uint8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x80, 0xec, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint16Immediate(t *testing.T) {
	// sub r12w, 0x1234
	builder := layout.NewSegmentBuilder()
	subIntImmediate(builder, ir.Uint16, registers.R12, uint16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x81, 0xec, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint32Immediate(t *testing.T) {
	// sub r12d, 0x12345678
	builder := layout.NewSegmentBuilder()
	subIntImmediate(
		builder,
		ir.Uint32,
		registers.R12,
		uint32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x81, 0xec, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSubUint64Immediate(t *testing.T) {
	// sub r12, 0x3456789a
	builder := layout.NewSegmentBuilder()
	subIntImmediate(
		builder,
		ir.Uint64,
		registers.R12,
		uint64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x49, 0x81, 0xec, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubIntRightImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(int32(0x01020304))
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
			Kind: ir.Sub,
			Src1: src,
			Src2: imm,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, srcRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, srcRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xe9, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubIntLeftImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(int32(0x01020304))
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
			Kind: ir.Sub,
			Src1: imm,
			Src2: src,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))
	expect.Equal(t, immChunk, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, srcChunk, constraints.RegisterSources[1].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, clobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		clobberedRegister)

	nonClobberedRegister := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, nonClobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		nonClobberedRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, clobberedRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister:    registers.Rbp,
			nonClobberedRegister: registers.Rsi,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x2b, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubIntIgnoreCheapSrc2(t *testing.T) {
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
			Kind: ir.Sub,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				src2Chunk: struct{}{},
			},
		})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))

	expect.Equal(
		t,
		src1Chunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(
		t,
		src2Chunk,
		constraints.RegisterSources[1].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, clobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		clobberedRegister)

	nonClobberedRegister := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, nonClobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		nonClobberedRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, clobberedRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister:    registers.Rdi,
			nonClobberedRegister: registers.R8,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x2b, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubUintRightImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint32(0x01020304))
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
			Kind: ir.Sub,
			Src1: src,
			Src2: imm,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(binaryMIOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 1, len(constraints.RegisterSources))
	expect.Equal(t, srcChunk, constraints.RegisterSources[0].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	srcRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, srcRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, srcRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister: registers.Rcx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xe9, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubUint(t *testing.T) {
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
			Kind: ir.Sub,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				src2Chunk: struct{}{},
			},
		})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))

	expect.Equal(
		t,
		src1Chunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(
		t,
		src2Chunk,
		constraints.RegisterSources[1].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, clobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		clobberedRegister)

	nonClobberedRegister := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, nonClobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		nonClobberedRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, clobberedRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister:    registers.Rdi,
			nonClobberedRegister: registers.R8,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x2b, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSubFloat(t *testing.T) {
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
		Type: ir.Float64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Sub,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				src2Chunk: struct{}{},
			},
		})

	_, ok := instruction.(binaryRMOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))

	expect.Equal(
		t,
		src1Chunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(
		t,
		src2Chunk,
		constraints.RegisterSources[1].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	clobberedRegister := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, clobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   true,
			Require:    nil,
		},
		clobberedRegister)

	nonClobberedRegister := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, nonClobberedRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: false,
			AnyFloat:   true,
			Require:    nil,
		},
		nonClobberedRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, clobberedRegister == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			clobberedRegister:    registers.Xmm1,
			nonClobberedRegister: registers.Xmm7,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x5c, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
