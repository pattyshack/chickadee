package instructions

import (
	"encoding/binary"
	"math"
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestAddInt8(t *testing.T) {
	// add eax, ecx
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Int8, registers.Rax, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x03, 0xc1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt16(t *testing.T) {
	// add edx, ebx
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Int16, registers.Rdx, registers.Rbx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x03, 0xd3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt32(t *testing.T) {
	// add ebp, esi
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Int32, registers.Rbp, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x03, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt64(t *testing.T) {
	// add rdi, r8
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Int64, registers.Rdi, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint8(t *testing.T) {
	// add r9d, r10d
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Uint8, registers.R9, registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x03, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint16(t *testing.T) {
	// add r11d, r12d
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Uint16, registers.R11, registers.R12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x03, 0xdc}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint32(t *testing.T) {
	// add r13d, r14d
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Uint32, registers.R13, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x03, 0xee}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint64(t *testing.T) {
	// add r15, rax
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Uint64, registers.R15, registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddFloat32(t *testing.T) {
	// addss xmm0, xmm2
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Float32, registers.Xmm0, registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x58, 0xc2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddFloat64(t *testing.T) {
	// addsd xmm1, xmm3
	builder := layout.NewSegmentBuilder()
	add(builder, ir.Float64, registers.Xmm1, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x58, 0xcb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt8Immediate(t *testing.T) {
	// add al, 0x12
	builder := layout.NewSegmentBuilder()
	addIntImmediate(builder, ir.Int8, registers.Rax, int8(0x12))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x80, 0xc0, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt16Immediate(t *testing.T) {
	// add bx, 0x1234
	builder := layout.NewSegmentBuilder()
	addIntImmediate(builder, ir.Int16, registers.Rbx, int16(0x1234))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xc3, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt32Immediate(t *testing.T) {
	// add ecx, 0x12345678
	builder := layout.NewSegmentBuilder()
	addIntImmediate(
		builder,
		ir.Int32,
		registers.Rcx,
		int32(0x12345678))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x81, 0xc1, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddInt64Immediate(t *testing.T) {
	// add rdx, 0x3456789a
	builder := layout.NewSegmentBuilder()
	addIntImmediate(
		builder,
		ir.Int64,
		registers.Rdx,
		int64(0x3456789a))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x81, 0xc2, 0x9a, 0x78, 0x56, 0x34},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint8Immediate(t *testing.T) {
	// add dil, 0xab
	builder := layout.NewSegmentBuilder()
	addIntImmediate(builder, ir.Uint8, registers.Rdi, uint8(0xab))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x80, 0xc7, 0xab}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint16Immediate(t *testing.T) {
	// add bp, 0xabcd
	builder := layout.NewSegmentBuilder()
	addIntImmediate(builder, ir.Uint16, registers.Rbp, uint16(0xabcd))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x81, 0xc5, 0xcd, 0xab},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint32Immediate(t *testing.T) {
	// add r9d, 0xabcdef01
	builder := layout.NewSegmentBuilder()
	addIntImmediate(
		builder,
		ir.Uint32,
		registers.R9,
		uint32(0xabcdef01))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x81, 0xc1, 0x01, 0xef, 0xcd, 0xab},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestAddUint64Immediate(t *testing.T) {
	// add r14, 0x7cdef012
	builder := layout.NewSegmentBuilder()
	addIntImmediate(
		builder,
		ir.Uint64,
		registers.R14,
		uint64(0x7cdef012))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x49, 0x81, 0xc6, 0x12, 0xf0, 0xde, 0x7c},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	for _, immValue := range []int64{
		0x12345678,
		math.MaxInt32,
		math.MinInt32,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immChunk := &ir.DefinitionChunk{}
		immDef := &ir.Definition{
			Name:   "imm",
			Chunks: []*ir.DefinitionChunk{immChunk},
		}
		immChunk.Definition = immDef
		imm.(*ir.Immediate).PseudoDefinition = immDef

		immBytes := make([]byte, 8)
		_, err := binary.Encode(immBytes, binary.LittleEndian, immValue)
		expect.Nil(t, err)

		for iter := 0; iter < 2; iter++ {
			destChunk := &ir.DefinitionChunk{}
			dest := &ir.Definition{
				Type:   ir.Int64,
				Chunks: []*ir.DefinitionChunk{destChunk},
			}
			destChunk.Definition = dest

			if iter == 0 { // right immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: src,
					Src2: imm,
				}
			} else { // left immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: imm,
					Src2: src,
				}
			}

			instruction := architecture.SelectInstruction(
				InstructionSet,
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
				append([]byte{0x48, 0x81, 0xc1}, immBytes[:4]...),
				segment.Content.Flatten())
			expect.Equal(t, layout.Definitions{}, segment.Definitions)
			expect.Equal(t, layout.Relocations{}, segment.Relocations)
		}
	}
}

func TestSelectAddUIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	for _, immValue := range []uint64{
		0x12345678,
		math.MaxInt32,
		0,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immChunk := &ir.DefinitionChunk{}
		immDef := &ir.Definition{
			Name:   "imm",
			Chunks: []*ir.DefinitionChunk{immChunk},
		}
		immChunk.Definition = immDef
		imm.(*ir.Immediate).PseudoDefinition = immDef

		immBytes := make([]byte, 8)
		_, err := binary.Encode(immBytes, binary.LittleEndian, immValue)
		expect.Nil(t, err)

		for iter := 0; iter < 2; iter++ {
			destChunk := &ir.DefinitionChunk{}
			dest := &ir.Definition{
				Type:   ir.Uint64,
				Chunks: []*ir.DefinitionChunk{destChunk},
			}
			destChunk.Definition = dest

			if iter == 0 { // right immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: src,
					Src2: imm,
				}
			} else { // left immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: imm,
					Src2: src,
				}
			}

			instruction := architecture.SelectInstruction(
				InstructionSet,
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
				append([]byte{0x48, 0x81, 0xc1}, immBytes[:4]...),
				segment.Content.Flatten())
			expect.Equal(t, layout.Definitions{}, segment.Definitions)
			expect.Equal(t, layout.Relocations{}, segment.Relocations)
		}
	}
}

// Fallback to RM encoding when immediate don't fit
func TestSelectAddIntOversizedImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	for _, immValue := range []int64{
		math.MaxInt32 + 1,
		math.MinInt32 - 1,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immChunk := &ir.DefinitionChunk{}
		immDef := &ir.Definition{
			Name:   "imm",
			Chunks: []*ir.DefinitionChunk{immChunk},
		}
		immChunk.Definition = immDef
		imm.(*ir.Immediate).PseudoDefinition = immDef

		for iter := 0; iter < 2; iter++ {
			destChunk := &ir.DefinitionChunk{}
			dest := &ir.Definition{
				Type:   ir.Int64,
				Chunks: []*ir.DefinitionChunk{destChunk},
			}
			destChunk.Definition = dest

			if iter == 0 { // right immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: src,
					Src2: imm,
				}
			} else { // left immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Add,
					Src1: imm,
					Src2: src,
				}
			}

			instruction := architecture.SelectInstruction(
				InstructionSet,
				dest,
				architecture.SelectorHint{})

			_, ok := instruction.(binaryRMOperation)
			expect.True(t, ok)

			// Validate constraints

			constraints := instruction.Constraints()
			expect.Nil(t, constraints.StackSources)
			expect.Nil(t, constraints.StackDestination)

			expect.Equal(t, 2, len(constraints.RegisterSources))

			if iter == 0 { // right immediate
				expect.Equal(
					t,
					srcChunk,
					constraints.RegisterSources[0].DefinitionChunk)
				expect.Equal(
					t,
					immChunk,
					constraints.RegisterSources[1].DefinitionChunk)
			} else { // left immediate
				expect.Equal(
					t,
					immChunk,
					constraints.RegisterSources[0].DefinitionChunk)
				expect.Equal(
					t,
					srcChunk,
					constraints.RegisterSources[1].DefinitionChunk)
			}

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
			expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
			expect.Equal(t, layout.Definitions{}, segment.Definitions)
			expect.Equal(t, layout.Relocations{}, segment.Relocations)
		}
	}
}

// Fallback to RM encoding when immediate don't fit
func TestSelectAddUintOversizedImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(uint64(math.MaxInt32 + 1))
	immChunk := &ir.DefinitionChunk{}
	immDef := &ir.Definition{
		Name:   "imm",
		Chunks: []*ir.DefinitionChunk{immChunk},
	}
	immChunk.Definition = immDef
	imm.(*ir.Immediate).PseudoDefinition = immDef

	for iter := 0; iter < 2; iter++ {
		destChunk := &ir.DefinitionChunk{}
		dest := &ir.Definition{
			Type:   ir.Uint64,
			Chunks: []*ir.DefinitionChunk{destChunk},
		}
		destChunk.Definition = dest

		if iter == 0 { // right immediate
			dest.Operation = &ir.BinaryOperation{
				Kind: ir.Add,
				Src1: src,
				Src2: imm,
			}
		} else { // left immediate
			dest.Operation = &ir.BinaryOperation{
				Kind: ir.Add,
				Src1: imm,
				Src2: src,
			}
		}

		instruction := architecture.SelectInstruction(
			InstructionSet,
			dest,
			architecture.SelectorHint{})

		_, ok := instruction.(binaryRMOperation)
		expect.True(t, ok)

		// Validate constraints

		constraints := instruction.Constraints()
		expect.Nil(t, constraints.StackSources)
		expect.Nil(t, constraints.StackDestination)

		expect.Equal(t, 2, len(constraints.RegisterSources))

		if iter == 0 { // right immediate
			expect.Equal(
				t,
				srcChunk,
				constraints.RegisterSources[0].DefinitionChunk)
			expect.Equal(
				t,
				immChunk,
				constraints.RegisterSources[1].DefinitionChunk)
		} else { // left immediate
			expect.Equal(
				t,
				immChunk,
				constraints.RegisterSources[0].DefinitionChunk)
			expect.Equal(
				t,
				srcChunk,
				constraints.RegisterSources[1].DefinitionChunk)
		}

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
		expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
		expect.Equal(t, layout.Definitions{}, segment.Definitions)
		expect.Equal(t, layout.Relocations{}, segment.Relocations)
	}
}

func TestSelectAddIntNoHint(t *testing.T) {
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
			Kind: ir.Add,
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
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddIntCheapSrc1(t *testing.T) {
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
			Kind: ir.Add,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				src1Chunk: struct{}{},
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
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddIntCheapSrc2(t *testing.T) {
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
			Kind: ir.Add,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
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
		src2Chunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(
		t,
		src1Chunk,
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
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddIntPreferredSrc1(t *testing.T) {
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
			Kind: ir.Add,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{
			PreferredRegisterDestination: map[*ir.DefinitionChunk]*ir.DefinitionChunk{
				destChunk: src1Chunk,
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
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddIntPreferredSrc2(t *testing.T) {
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
			Kind: ir.Add,
			Src1: src1,
			Src2: src2,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{
			PreferredRegisterDestination: map[*ir.DefinitionChunk]*ir.DefinitionChunk{
				destChunk: src2Chunk,
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
		src2Chunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Equal(
		t,
		src1Chunk,
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
	expect.Equal(t, []byte{0x49, 0x03, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectAddFloatNoHint(t *testing.T) {
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
			Kind: ir.Add,
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
			nonClobberedRegister: registers.Xmm3,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x58, 0xcb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
