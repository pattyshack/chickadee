package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestCopyGeneral8(t *testing.T) {
	// mov edi, ecx
	builder := layout.NewSegmentBuilder()
	copyGeneral(builder, 1, registers.Rdi, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x8b, 0xf9},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneral16(t *testing.T) {
	// mov edx, r8d
	builder := layout.NewSegmentBuilder()
	copyGeneral(builder, 2, registers.Rdx, registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8b, 0xd0},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneral32(t *testing.T) {
	// mov r13d, ebp
	builder := layout.NewSegmentBuilder()
	copyGeneral(builder, 4, registers.R13, registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x44, 0x8b, 0xed},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneral64(t *testing.T) {
	// mov rdx, rsi
	builder := layout.NewSegmentBuilder()
	copyGeneral(builder, 8, registers.Rdx, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x8b, 0xd6},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToMemory8(t *testing.T) {
	// mov [rbx], r14b
	builder := layout.NewSegmentBuilder()
	copyGeneralToMemory(builder, 1, registers.Rbx, registers.R14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x44, 0x88, 0x33},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToMemory16(t *testing.T) {
	// mov [rbp+0], dx (disp8 encoding)
	builder := layout.NewSegmentBuilder()
	copyGeneralToMemory(builder, 2, registers.Rbp, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x89, 0x55, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToMemory32(t *testing.T) {
	// mov [r12], ecx (SIB encoding)
	builder := layout.NewSegmentBuilder()
	copyGeneralToMemory(builder, 4, registers.R12, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x89, 0x0c, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToMemory64(t *testing.T) {
	// mov [rdx], rsi
	builder := layout.NewSegmentBuilder()
	copyGeneralToMemory(builder, 8, registers.Rdx, registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x89, 0x32},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToGeneral8(t *testing.T) {
	// mov dl, [rcx]
	builder := layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rdx, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x8a, 0x11},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov al, [rsi]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rax, registers.Rsi)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x8a, 0x06},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r8b, [rbx]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.R8, registers.Rbx)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x44, 0x8a, 0x03},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov bl, [r9]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rbx, registers.R9)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8a, 0x19},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov dl, [r12] (SIB encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rdx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8a, 0x14, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov dil, [rbp+0] (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rdi, registers.Rbp)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x40, 0x8a, 0x7d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov cl, [r13+0] (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 1, registers.Rcx, registers.R13)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8a, 0x4d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToGeneral16(t *testing.T) {
	// mov di, [rcx]
	builder := layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 2, registers.Rdi, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x8b, 0x39},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov si, [r14]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 2, registers.Rsi, registers.R14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x8b, 0x36},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov dx, [r12] (SIB encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 2, registers.Rdx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x8b, 0x14, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov cx, [r13+0] (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 2, registers.Rcx, registers.R13)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x8b, 0x4d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToGeneral32(t *testing.T) {
	// mov edi, [rcx]
	builder := layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 4, registers.Rdi, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x8b, 0x39},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov esi, [r14]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 4, registers.Rsi, registers.R14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8b, 0x36},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov edx, [r12] (SIB encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 4, registers.Rdx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8b, 0x14, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov ecx, [r13+0] (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 4, registers.Rcx, registers.R13)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0x8b, 0x4d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToGeneral64(t *testing.T) {
	// mov rdi, [rcx]
	builder := layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 8, registers.Rdi, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0x8b, 0x39},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov rsi, [r14]
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 8, registers.Rsi, registers.R14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x49, 0x8b, 0x36},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov rdx, [r12] (SIB encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 8, registers.Rdx, registers.R12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x49, 0x8b, 0x14, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r10, [r13+0] (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyMemoryToGeneral(builder, 8, registers.R10, registers.R13)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x4d, 0x8b, 0x55, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloat32(t *testing.T) {
	// movss xmm6, xmm1
	builder := layout.NewSegmentBuilder()
	copyFloat(builder, 4, registers.Xmm6, registers.Xmm1)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x0f, 0x10, 0xf1},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movss xmm6, xmm12
	builder = layout.NewSegmentBuilder()
	copyFloat(builder, 4, registers.Xmm6, registers.Xmm12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x41, 0x0f, 0x10, 0xf4},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movss xmm9, xmm1
	builder = layout.NewSegmentBuilder()
	copyFloat(builder, 4, registers.Xmm9, registers.Xmm1)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x10, 0xc9},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloat64(t *testing.T) {
	// movsd xmm6, xmm1
	builder := layout.NewSegmentBuilder()
	copyFloat(builder, 8, registers.Xmm6, registers.Xmm1)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x0f, 0x10, 0xf1},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movsd xmm6, xmm12
	builder = layout.NewSegmentBuilder()
	copyFloat(builder, 8, registers.Xmm6, registers.Xmm12)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x41, 0x0f, 0x10, 0xf4},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movsd xmm9, xmm1
	builder = layout.NewSegmentBuilder()
	copyFloat(builder, 8, registers.Xmm9, registers.Xmm1)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x44, 0x0f, 0x10, 0xc9},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToMemory32(t *testing.T) {
	// movd [rdx], xmm10
	builder := layout.NewSegmentBuilder()
	copyFloatToMemory(builder, 4, registers.Rdx, registers.Xmm10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x44, 0x0f, 0x7e, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToMemory64(t *testing.T) {
	// movq [rcx], xmm0 (standard encoding)
	builder := layout.NewSegmentBuilder()
	copyFloatToMemory(builder, 8, registers.Rcx, registers.Xmm0)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x48, 0x0f, 0x7e, 0x01},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movq [r12], xmm14 (SIB encoding)
	builder = layout.NewSegmentBuilder()
	copyFloatToMemory(builder, 8, registers.R12, registers.Xmm14)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x4d, 0x0f, 0x7e, 0x34, 0x24},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movq [rbp+0], xmm7 (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyFloatToMemory(builder, 8, registers.Rbp, registers.Xmm7)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x48, 0x0f, 0x7e, 0x7d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movq [r13+0], xmm9 (disp8 encoding)
	builder = layout.NewSegmentBuilder()
	copyFloatToMemory(builder, 8, registers.R13, registers.Xmm9)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x4d, 0x0f, 0x7e, 0x4d, 0x00},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToFloat32(t *testing.T) {
	// movd xmm3, [rdi]
	builder := layout.NewSegmentBuilder()
	copyMemoryToFloat(builder, 4, registers.Xmm3, registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x0f, 0x6e, 0x1f},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyMemoryToFloat64(t *testing.T) {
	// movq xmm6, [rcx]
	builder := layout.NewSegmentBuilder()
	copyMemoryToFloat(builder, 8, registers.Xmm6, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x48, 0x0f, 0x6e, 0x31},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToGeneral8(t *testing.T) {
	// movd ebp, xmm2
	builder := layout.NewSegmentBuilder()
	copyFloatToGeneral(builder, 1, registers.Rbp, registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x0f, 0x7e, 0xd5},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToGeneral16(t *testing.T) {
	// movd edx, xmm12
	builder := layout.NewSegmentBuilder()
	copyFloatToGeneral(builder, 2, registers.Rdx, registers.Xmm12)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x44, 0x0f, 0x7e, 0xe2},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToGeneral32(t *testing.T) {
	// movd edi, xmm5
	builder := layout.NewSegmentBuilder()
	copyFloatToGeneral(builder, 4, registers.Rdi, registers.Xmm5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x0f, 0x7e, 0xef}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyFloatToGeneral64(t *testing.T) {
	// movq rbx, xmm3
	builder := layout.NewSegmentBuilder()
	copyFloatToGeneral(builder, 8, registers.Rbx, registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x48, 0x0f, 0x7e, 0xdb},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToFloat32(t *testing.T) {
	// movd xmm7, ecx
	builder := layout.NewSegmentBuilder()
	copyGeneralToFloat(builder, 4, registers.Xmm7, registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0x0f, 0x6e, 0xf9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movd xmm3, r9d
	builder = layout.NewSegmentBuilder()
	copyGeneralToFloat(builder, 4, registers.Xmm3, registers.R9)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x41, 0x0f, 0x6e, 0xd9},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestCopyGeneralToFloat64(t *testing.T) {
	// movq xmm2, rdx
	builder := layout.NewSegmentBuilder()
	copyGeneralToFloat(builder, 8, registers.Xmm2, registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x48, 0x0f, 0x6e, 0xd2},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// movq xmm10, rdi
	builder = layout.NewSegmentBuilder()
	copyGeneralToFloat(builder, 8, registers.Xmm10, registers.Rdi)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x66, 0x4c, 0x0f, 0x6e, 0xd7},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateInt8(t *testing.T) {
	// xor ebp, ebp
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rbp, int8(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xed}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov bpl, -1
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rbp, int8(-1))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0xb5, 0xff}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateInt16(t *testing.T) {
	// xor edx, edx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rdx, int16(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xd2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov dx, 0x1234
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rdx, int16(0x1234))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xba, 0x34, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateInt32(t *testing.T) {
	// xor r15d, r15d
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.R15, int32(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x33, 0xff}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r15d, 0x12345678
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.R15, int32(0x12345678))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0xbf, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateInt64(t *testing.T) {
	// xor ecx, ecx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, int64(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov rcx, 0x1234567890abcdef
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, int64(0x1234567890abcdef))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0xb9, 0xef, 0xcd, 0xab, 0x90, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateUint8(t *testing.T) {
	// xor ebx, ebx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rbx, uint8(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xdb}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov bl, 5
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rbx, uint8(5))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xb3, 0x05}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateUint16(t *testing.T) {
	// xor edx, edx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rdx, uint16(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xd2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov dx, 0x1234
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rdx, uint16(0x1234))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x66, 0xba, 0x34, 0x12}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateUint32(t *testing.T) {
	// xor r15d, r15d
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.R15, uint32(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x33, 0xff}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r15d, 0x12345678
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.R15, uint32(0x12345678))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x41, 0xbf, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateUint64(t *testing.T) {
	// xor ecx, ecx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, uint64(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov rcx, 0x1234567890abcdef
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, uint64(0x1234567890abcdef))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0xb9, 0xef, 0xcd, 0xab, 0x90, 0x78, 0x56, 0x34, 0x12},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSetImmediateFloat32(t *testing.T) {
	// xor ecx, ecx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, float32(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov ecx, 3.1415 (= 0x40490e56)
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, float32(3.1415))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xb9, 0x56, 0x0e, 0x49, 0x40},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

// (= 0x400921cac083126f)
func TestSetImmediateFloat64(t *testing.T) {
	// xor ecx, ecx
	builder := layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, float64(0))
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x33, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov rcx, 3.1415 (= 0x400921cac083126f)
	builder = layout.NewSegmentBuilder()
	setImmediate(builder, registers.Rcx, float64(3.1415))
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0x48, 0xb9, 0x6f, 0x12, 0x83, 0xc0, 0xca, 0x21, 0x09, 0x40},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
