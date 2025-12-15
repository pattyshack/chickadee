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

func TestDivRemUint8(t *testing.T) {
	// movzx eax, al
	// movzx ebp, bpl
	// xor edx, edx
	// div ebp
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint8, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xb6, 0xc0, // movzx
			0x40, 0x0f, 0xb6, 0xed, // movzx
			0x33, 0xd2, // xor
			0xf7, 0xf5, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint16(t *testing.T) {
	// xor edx, edx
	// div di
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint16, registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x66, 0xf7, 0xf7, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint32(t *testing.T) {
	// xor edx, edx
	// div r13d
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint32, registers.R13)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x41, 0xf7, 0xf5, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemUint64(t *testing.T) {
	// xor edx, edx
	// div rcx
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Uint64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor
			0x48, 0xf7, 0xf1, // div
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt8(t *testing.T) {
	// movsx eax, al
	// movsx ebp, bpl
	// cdq
	// idiv ebp
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int8, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xbe, 0xc0, // movsx
			0x40, 0x0f, 0xbe, 0xed, // movsx
			0x99,       // cdq
			0xf7, 0xfd, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt16(t *testing.T) {
	// cwd
	// idiv di
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int16, registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x66, 0x99, // cwd
			0x66, 0xf7, 0xff, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt32(t *testing.T) {
	// cdq
	// idiv r13d
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int32, registers.R13)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x99,             // cdq
			0x41, 0xf7, 0xfd, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivRemInt64(t *testing.T) {
	// cqo
	// idiv rcx
	builder := layout.NewSegmentBuilder()
	divRemInt(builder, ir.Int64, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x48, 0x99, // cqo
			0x48, 0xf7, 0xf9, // idiv
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivFloat32(t *testing.T) {
	// divss xmm3, xmm7
	builder := layout.NewSegmentBuilder()
	divFloat(builder, ir.Float32, registers.Xmm3, registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf3, 0x0f, 0x5e, 0xdf,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivFloat64(t *testing.T) {
	// divsd xmm10, xmm5
	builder := layout.NewSegmentBuilder()
	divFloat(builder, ir.Float64, registers.Xmm10, registers.Xmm5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf2, 0x44, 0x0f, 0x5e, 0xd5,
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivUint(t *testing.T) {
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
			Kind: ir.Div,
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

	_, ok := instruction.(divRemOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 3, len(constraints.RegisterSources))
	expect.Nil(t, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, src1Chunk, constraints.RegisterSources[1].DefinitionChunk)
	expect.Equal(t, src2Chunk, constraints.RegisterSources[2].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	rdx := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, rdx)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rdx,
		},
		rdx)

	rax := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, rax)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rax,
		},
		rax)

	divisor := constraints.RegisterSources[2].RegisterConstraint
	expect.NotNil(t, divisor)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		divisor)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, rax == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			rdx:     registers.Rdx,
			rax:     registers.Rax,
			divisor: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor edx, edx
			0x66, 0xf7, 0xf3, // div bx
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivUintSameSource(t *testing.T) {
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef

	src1 := ir.NewLocalReference("src")
	src1.(*ir.LocalReference).UseDef = srcDef

	src2 := ir.NewLocalReference("src2")
	src2.(*ir.LocalReference).UseDef = srcDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Uint32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Div,
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

	_, ok := instruction.(divRemOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))
	expect.Nil(t, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, srcChunk, constraints.RegisterSources[1].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	rdx := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, rdx)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rdx,
		},
		rdx)

	rax := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, rax)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rax,
		},
		rax)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, rax == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			rdx: registers.Rdx,
			rax: registers.Rax,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor edx, edx
			0xf7, 0xf0, // div eax
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestRemUint(t *testing.T) {
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
			Kind: ir.Rem,
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

	_, ok := instruction.(divRemOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 3, len(constraints.RegisterSources))
	expect.Nil(t, constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(t, src1Chunk, constraints.RegisterSources[1].DefinitionChunk)
	expect.Equal(t, src2Chunk, constraints.RegisterSources[2].DefinitionChunk)

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	expect.Equal(
		t,
		destChunk,
		constraints.RegisterDestinations[0].DefinitionChunk)

	rdx := constraints.RegisterSources[0].RegisterConstraint
	expect.NotNil(t, rdx)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rdx,
		},
		rdx)

	rax := constraints.RegisterSources[1].RegisterConstraint
	expect.NotNil(t, rax)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   false,
			Require:    registers.Rax,
		},
		rax)

	divisor := constraints.RegisterSources[2].RegisterConstraint
	expect.NotNil(t, divisor)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		divisor)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, rdx == destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			rdx:     registers.Rdx,
			rax:     registers.Rax,
			divisor: registers.Rbx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x33, 0xd2, // xor edx, edx
			0x66, 0xf7, 0xf3, // div bx
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivInt(t *testing.T) {
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
		Type: ir.Int32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Div,
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

	_, ok := instruction.(divRemOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 3, len(constraints.RegisterSources))
	expect.Nil(t, constraints.RegisterSources[0].DefinitionChunk)

	rdx := constraints.RegisterSources[0].RegisterConstraint
	rax := constraints.RegisterSources[1].RegisterConstraint
	divisor := constraints.RegisterSources[2].RegisterConstraint

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, rax == destRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			rdx:     registers.Rdx,
			rax:     registers.Rax,
			divisor: registers.Rsi,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x99,       // cdq
			0xf7, 0xfe, // idiv esi
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestRemInt(t *testing.T) {
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
		Type: ir.Int32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Rem,
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

	_, ok := instruction.(divRemOperation)
	expect.True(t, ok)

	constraints := instruction.Constraints()
	expect.Equal(t, 3, len(constraints.RegisterSources))
	expect.Nil(t, constraints.RegisterSources[0].DefinitionChunk)

	rdx := constraints.RegisterSources[0].RegisterConstraint
	rax := constraints.RegisterSources[1].RegisterConstraint
	divisor := constraints.RegisterSources[2].RegisterConstraint

	expect.Equal(t, 1, len(constraints.RegisterDestinations))
	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.True(t, rdx == destRegister)

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			rdx:     registers.Rdx,
			rax:     registers.Rax,
			divisor: registers.Rsi,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x99,       // cdq
			0xf7, 0xfe, // idiv esi
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestDivFloat(t *testing.T) {
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
		Type: ir.Float32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Div,
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
			src1Register: registers.Xmm1,
			src2Register: registers.Xmm2,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x5e, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
