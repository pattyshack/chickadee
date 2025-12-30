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
	mulIntImmediate(builder, ir.Int8, registers.R15, registers.Rbx, int8(0x12))
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
		int16(0x1234))
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
		int32(0x12345678))
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
		int64(0x3456789a))
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
	mulIntImmediate(builder, ir.Uint8, registers.Rdx, registers.R10, uint8(0x12))
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
		uint16(0x1234))
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
		uint32(0x12345678))
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
		uint64(0x3456789a))
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
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Int64,
	}
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Int64,
	}
	src2.(*ir.LocalReference).UseDef = src2Def

	dest := &ir.Definition{
		Type: ir.Int64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
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
	expect.Equal(t, []byte{0x4c, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulUint(t *testing.T) {
	src1 := ir.NewLocalReference("src1")
	src1Def := &ir.Definition{
		Name: "src1",
		Type: ir.Uint32,
	}
	src1.(*ir.LocalReference).UseDef = src1Def

	src2 := ir.NewLocalReference("src2")
	src2Def := &ir.Definition{
		Name: "src2",
		Type: ir.Uint32,
	}
	src2.(*ir.LocalReference).UseDef = src2Def

	dest := &ir.Definition{
		Type: ir.Uint32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
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
	expect.Equal(t, []byte{0x44, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Float32,
	}
	src.(*ir.LocalReference).UseDef = srcDef

	imm := ir.NewBasicImmediate(float32(0.5))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Float32,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Float32,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src,
			Src2: imm,
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

// Also verify symmetry
func TestSelectMulIntImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Int64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	for _, immValue := range []int64{
		0x12345678,
		math.MaxInt32,
		math.MinInt32,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immDef := &ir.Definition{
			Name: "imm",
			Type: ir.Int64,
		}
		imm.(*ir.Immediate).PseudoDefinition = immDef

		immBytes := make([]byte, 8)
		_, err := binary.Encode(immBytes, binary.LittleEndian, immValue)
		expect.Nil(t, err)

		for iter := 0; iter < 2; iter++ {
			dest := &ir.Definition{
				Type: ir.Int64,
			}
			destChunk := dest.Chunks()[0]

			if iter == 0 { // right immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Mul,
					Src1: src,
					Src2: imm,
				}
			} else { // left immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Mul,
					Src1: imm,
					Src2: src,
				}
			}

			instruction := architecture.SelectInstruction(
				testConfig,
				dest,
				architecture.SelectorHint{})

			_, ok := instruction.(imulRMIOperation)
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
				append([]byte{0x48, 0x69, 0xc9}, immBytes[:4]...),
				segment.Content.Flatten())
			expect.Equal(t, layout.Definitions{}, segment.Definitions)
			expect.Equal(t, layout.Relocations{}, segment.Relocations)
		}
	}
}

// Also verify symmetry
func TestSelectMulUintImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	for _, immValue := range []uint64{
		0x12345678,
		math.MaxInt32,
		0,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immDef := &ir.Definition{
			Name: "imm",
			Type: ir.Uint64,
		}
		imm.(*ir.Immediate).PseudoDefinition = immDef

		immBytes := make([]byte, 8)
		_, err := binary.Encode(immBytes, binary.LittleEndian, immValue)
		expect.Nil(t, err)

		for iter := 0; iter < 2; iter++ {
			dest := &ir.Definition{
				Type: ir.Uint64,
			}
			destChunk := dest.Chunks()[0]

			if iter == 0 { // right immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Mul,
					Src1: src,
					Src2: imm,
				}
			} else { // left immediate
				dest.Operation = &ir.BinaryOperation{
					Kind: ir.Mul,
					Src1: imm,
					Src2: src,
				}
			}

			instruction := architecture.SelectInstruction(
				testConfig,
				dest,
				architecture.SelectorHint{})

			_, ok := instruction.(imulRMIOperation)
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
				append([]byte{0x48, 0x69, 0xc9}, immBytes[:4]...),
				segment.Content.Flatten())
			expect.Equal(t, layout.Definitions{}, segment.Definitions)
			expect.Equal(t, layout.Relocations{}, segment.Relocations)
		}
	}
}

// Fallback to RM encoding when immediate don't fit
func TestSelectMulIntOversizedImmediate(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Int64,
	}
	src.(*ir.LocalReference).UseDef = srcDef

	for _, immValue := range []int64{
		math.MaxInt32 + 1,
		math.MinInt32 - 1,
	} {
		imm := ir.NewBasicImmediate(immValue)
		immDef := &ir.Definition{
			Name: "imm",
			Type: ir.Int64,
		}
		imm.(*ir.Immediate).PseudoDefinition = immDef

		dest := &ir.Definition{
			Type: ir.Int64,
			Operation: &ir.BinaryOperation{
				Kind: ir.Mul,
				Src1: src,
				Src2: imm,
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
		clobberedRegister := constraints.RegisterSources[0].RegisterConstraint
		nonClobberedRegister := constraints.RegisterSources[1].RegisterConstraint

		builder := layout.NewSegmentBuilder()
		instruction.EmitTo(
			builder,
			map[*architecture.RegisterConstraint]*architecture.Register{
				clobberedRegister:    registers.Rdi,
				nonClobberedRegister: registers.R8,
			})
		segment, err := builder.Finalize(amd64.ArchitectureLayout)
		expect.Nil(t, err)
		expect.Equal(t, []byte{0x49, 0x0f, 0xaf, 0xf8}, segment.Content.Flatten())
		expect.Equal(t, layout.Definitions{}, segment.Definitions)
		expect.Equal(t, layout.Relocations{}, segment.Relocations)
	}
}

func TestSelectMulUintImmediateHasFreeRegisters(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(uint64(0x01020304))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Uint64,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Uint64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src,
			Src2: imm,
		},
	}
	destChunk := dest.Chunks()[0]

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			NumFreeGeneralRegisters: 1,
		})

	_, ok := instruction.(imulRMIOperation)
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
			Clobbered:  false,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.NotNil(t, destRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		destRegister)

	expect.True(t, srcRegister != destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister:  registers.Rcx,
			destRegister: registers.Rdx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x69, 0xd1, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulUintImmediateCheapSoruce(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(uint64(0x01020304))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Uint64,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Uint64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src,
			Src2: imm,
		},
	}
	destChunk := dest.Chunks()[0]

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			NumFreeGeneralRegisters: 1,
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				srcChunk: struct{}{},
			},
		})

	_, ok := instruction.(imulRMIOperation)
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
		[]byte{0x48, 0x69, 0xc9, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectMulUintImmediatePreferredReuse(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcDef := &ir.Definition{
		Name: "src",
		Type: ir.Uint64,
	}
	src.(*ir.LocalReference).UseDef = srcDef
	srcChunk := srcDef.Chunks()[0]

	imm := ir.NewBasicImmediate(uint64(0x01020304))
	immDef := &ir.Definition{
		Name: "imm",
		Type: ir.Uint64,
	}
	imm.(*ir.Immediate).PseudoDefinition = immDef

	dest := &ir.Definition{
		Type: ir.Uint64,
		Operation: &ir.BinaryOperation{
			Kind: ir.Mul,
			Src1: src,
			Src2: imm,
		},
	}
	destChunk := dest.Chunks()[0]

	instruction := architecture.SelectInstruction(
		testConfig,
		dest,
		architecture.SelectorHint{
			NumFreeGeneralRegisters: 1,
			PreferredRegisterDestination: map[*ir.DefinitionChunk]*ir.DefinitionChunk{
				destChunk: srcChunk,
			},
		})

	_, ok := instruction.(imulRMIOperation)
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
		[]byte{0x48, 0x69, 0xc9, 0x04, 0x03, 0x02, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
