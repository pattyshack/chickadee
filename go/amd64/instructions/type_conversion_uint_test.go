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

func TestUint8ToUint8(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rcx,
		ir.Uint8,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToUint16(t *testing.T) {
	// movzx ecx, dil
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Uint8,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x0f, 0xb6, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToUint32(t *testing.T) {
	// movzx esi, al
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rsi,
		ir.Uint8,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb6, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToUint64(t *testing.T) {
	// movzx edx, cl
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.Rdx,
		ir.Uint8,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb6, 0xd1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToUint8(t *testing.T) {
	// mov edx, esi (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rdx,
		ir.Uint16,
		registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToUint16(t *testing.T) {
	// mov ebp, r8d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rbp,
		ir.Uint16,
		registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToUint32(t *testing.T) {
	// movzx edx, di
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rdx,
		ir.Uint16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb7, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToUint64(t *testing.T) {
	// movzx r9d, cx
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.R9,
		ir.Uint16,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x0f, 0xb7, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToUint8(t *testing.T) {
	// mov ebx, ebp (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rbx,
		ir.Uint32,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xdd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToUint16(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Uint32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToUint32(t *testing.T) {
	// mov eax, r15d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rax,
		ir.Uint32,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xc7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToUint64(t *testing.T) {
	// mov ebx, ecx (NOTE: upper bytes are set to zero)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.Rbx,
		ir.Uint32,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r10d, r10d (NOTE: This is not no-op; upper bytes are set to zero)
	builder = layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.R10,
		ir.Uint32,
		registers.R10)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x8b, 0xd2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToUint8(t *testing.T) {
	// mov esi, r11d (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rsi,
		ir.Uint64,
		registers.R11)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xf3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToUint16(t *testing.T) {
	// mov ecx, ebp (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Uint64,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xcd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToUint32(t *testing.T) {
	// mov ebx, eax (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rbx,
		ir.Uint64,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToUint64(t *testing.T) {
	// mov rdx, rcx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.Rdx,
		ir.Uint64,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x8b, 0xd1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToInt8(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rcx,
		ir.Uint8,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToInt16(t *testing.T) {
	// movzx ecx, dil
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Uint8,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x0f, 0xb6, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToInt32(t *testing.T) {
	// movzx esi, al
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rsi,
		ir.Uint8,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb6, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToInt64(t *testing.T) {
	// movzx r10d, bpl
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R10,
		ir.Uint8,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x0f, 0xb6, 0xd5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToInt8(t *testing.T) {
	// mov edx, esi (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rdx,
		ir.Uint16,
		registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToInt16(t *testing.T) {
	// mov ebp, r8d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rbp,
		ir.Uint16,
		registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToInt32(t *testing.T) {
	// movzx edx, di
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rdx,
		ir.Uint16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb7, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToInt64(t *testing.T) {
	// movzx r9d, cx
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R9,
		ir.Uint16,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x44, 0x0f, 0xb7, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToInt8(t *testing.T) {
	// mov ebx, ebp (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rbx,
		ir.Uint32,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xdd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToInt16(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Uint32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToInt32(t *testing.T) {
	// mov eax, r15d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rax,
		ir.Uint32,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xc7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToInt64(t *testing.T) {
	// mov ebx, ecx (NOTE: upper bytes are set to zero)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.Rbx,
		ir.Uint32,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)

	// mov r10d, r10d (NOTE: This is not no-op; upper bytes are set to zero)
	builder = layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R10,
		ir.Uint32,
		registers.R10)
	segment, err = builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x45, 0x8b, 0xd2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToInt8(t *testing.T) {
	// mov esi, r11d (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rsi,
		ir.Uint64,
		registers.R11)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xf3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToInt16(t *testing.T) {
	// mov ecx, ebp (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Uint64,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xcd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToInt32(t *testing.T) {
	// mov ebx, eax (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rbx,
		ir.Uint64,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToInt64(t *testing.T) {
	// mov rdx, rcx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.Rdx,
		ir.Uint64,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x8b, 0xd1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToFloat32(t *testing.T) {
	// movzx ecx, cl
	// cvtsi2ss xmm5, ecx
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float32,
		registers.Xmm5,
		ir.Uint8,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xb6, 0xc9, // movzx
			0xf3, 0x0f, 0x2a, 0xe9, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint8ToFloat64(t *testing.T) {
	// movzx ebp, bpl
	// cvtsi2sd xmm2, ebp
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float64,
		registers.Xmm2,
		ir.Uint8,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x0f, 0xb6, 0xed, // movzx
			0xf2, 0x0f, 0x2a, 0xd5, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToFloat32(t *testing.T) {
	// movzx edi, di
	// cvtsi2ss xmm0, edi
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float32,
		registers.Xmm0,
		ir.Uint16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xb7, 0xff, // movzx
			0xf3, 0x0f, 0x2a, 0xc7, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint16ToFloat64(t *testing.T) {
	// movzx r15d, r15w
	// cvtsi2sd xmm7, r15d
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float64,
		registers.Xmm7,
		ir.Uint16,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x45, 0x0f, 0xb7, 0xff, // movzx
			0xf2, 0x41, 0x0f, 0x2a, 0xff, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToFloat32(t *testing.T) {
	// mov ecx, ecx (NOTE: this sets upper bytes to zero)
	// cvtsi2ss xmm12, rcx
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float32,
		registers.Xmm12,
		ir.Uint32,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x8b, 0xc9, // mov
			0xf3, 0x4c, 0x0f, 0x2a, 0xe1, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint32ToFloat64(t *testing.T) {
	// mov ecx, ecx (NOTE: this sets upper bytes to zero)
	// cvtsi2sd xmm7, rdx
	builder := layout.NewSegmentBuilder()
	convertSmallUintToFloat(
		builder,
		ir.Float64,
		registers.Xmm7,
		ir.Uint32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x8b, 0xd2, // mov
			0xf2, 0x48, 0x0f, 0x2a, 0xfa, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToFloat32(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	convertUint64ToFloat(
		builder,
		ir.Float32,
		registers.Xmm7,
		registers.Rcx,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			// comparison basic block

			// 0: cmp rcx, 0
			0x48, 0x81, 0xf9, 0x00, 0x00, 0x00, 0x00,
			// 7: jge <nonNegative offset> (= 43 - 13 = 30 = 0x1e)
			0x0f, 0x8d, 0x1e, 0x00, 0x00, 0x00,

			// negative branch basic block

			// 13: mov rdx, rcx
			0x48, 0x8b, 0xd1,
			// 16: shr rdx, 1
			0x48, 0xc1, 0xea, 0x01,
			// 20: and ecx, 1
			0x81, 0xe1, 0x01, 0x00, 0x00, 0x00,
			// 26: or rcx, rdx
			0x48, 0x0b, 0xca,
			// 29: cvtsi2ss xmm7, rcx
			0xf3, 0x48, 0x0f, 0x2a, 0xf9,
			// 34: addss xmm7, xmm7
			0xf3, 0x0f, 0x58, 0xff,
			// 38: jmp <end offset> (= 48 - 43 = 5)
			0xe9, 0x05, 0x00, 0x00, 0x00,

			// non-negative branch basic block

			// 43: cvtsi2ss xmm7, rcx
			0xf3, 0x48, 0x0f, 0x2a, 0xf9,

			// 48: end basic block
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestUint64ToFloat64(t *testing.T) {
	builder := layout.NewSegmentBuilder()
	convertUint64ToFloat(
		builder,
		ir.Float64,
		registers.Xmm4,
		registers.Rbx,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			// comparison basic block

			// 0: cmp rbx, 0
			0x48, 0x81, 0xfb, 0x00, 0x00, 0x00, 0x00,
			// 7: jge <nonNegative offset> (= 43 - 13 = 30 = 0x1e)
			0x0f, 0x8d, 0x1e, 0x00, 0x00, 0x00,

			// negative branch basic block

			// 13: mov rdi, rbx
			0x48, 0x8b, 0xfb,
			// 16: shr rdi, 1
			0x48, 0xc1, 0xef, 0x01,
			// 20: and ebx, 1
			0x81, 0xe3, 0x01, 0x00, 0x00, 0x00,
			// 26: or rbx, rdi
			0x48, 0x0b, 0xdf,
			// 29: cvtsi2sd xmm7, rbx
			0xf2, 0x48, 0x0f, 0x2a, 0xe3,
			// 34: addsd xmm4, xmm4
			0xf2, 0x0f, 0x58, 0xe4,
			// 38: jmp <end offset> (= 48 - 43 = 5)
			0xe9, 0x05, 0x00, 0x00, 0x00,

			// non-negative branch basic block

			// 43: cvtsi2sd xmm4, rbx
			0xf2, 0x48, 0x0f, 0x2a, 0xe3,

			// 48: end basic block
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectUintToUint(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Uint8,
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Uint32,
		Operation: &ir.UnaryOperation{
			Kind: ir.ToUint32,
			Src:  src,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{
			NumFreeGeneralRegisters: 1,
		})

	_, ok := instruction.(conversionOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 1, len(constraints.RegisterSources))

	expect.Equal(
		t,
		srcChunk,
		constraints.RegisterSources[0].DefinitionChunk)

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
			srcRegister:  registers.Rax,
			destRegister: registers.Rsi,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb6, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectUintToInt(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Uint16,
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Int32,
		Operation: &ir.UnaryOperation{
			Kind: ir.ToInt32,
			Src:  src,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{
			NumFreeGeneralRegisters: 1,
		})

	_, ok := instruction.(conversionOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 1, len(constraints.RegisterSources))

	expect.Equal(
		t,
		srcChunk,
		constraints.RegisterSources[0].DefinitionChunk)

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
			srcRegister:  registers.Rdi,
			destRegister: registers.Rdx,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xb7, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectSmallUintToFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Uint32,
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Float32,
		Operation: &ir.UnaryOperation{
			Kind: ir.ToFloat32,
			Src:  src,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(conversionOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 1, len(constraints.RegisterSources))

	expect.Equal(
		t,
		srcChunk,
		constraints.RegisterSources[0].DefinitionChunk)

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
			AnyGeneral: false,
			AnyFloat:   true,
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
			destRegister: registers.Xmm12,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x8b, 0xc9, // mov(zx)
			0xf3, 0x4c, 0x0f, 0x2a, 0xe1, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectUint64ToFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Uint64,
		Chunks: []*ir.DefinitionChunk{srcChunk},
	}
	srcChunk.Definition = srcDef
	src.(*ir.LocalReference).UseDef = srcDef

	destChunk := &ir.DefinitionChunk{}
	dest := &ir.Definition{
		Type: ir.Float32,
		Operation: &ir.UnaryOperation{
			Kind: ir.ToFloat32,
			Src:  src,
		},
		Chunks: []*ir.DefinitionChunk{destChunk},
	}
	destChunk.Definition = dest

	instruction := architecture.SelectInstruction(
		InstructionSet,
		dest,
		architecture.SelectorHint{})

	_, ok := instruction.(uint64ToFloatOperation)
	expect.True(t, ok)

	// Validate constraints

	constraints := instruction.Constraints()
	expect.Nil(t, constraints.StackSources)
	expect.Nil(t, constraints.StackDestination)

	expect.Equal(t, 2, len(constraints.RegisterSources))

	expect.Equal(
		t,
		srcChunk,
		constraints.RegisterSources[0].DefinitionChunk)
	expect.Nil(t, constraints.RegisterSources[1].DefinitionChunk)

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

	scratchRegister := constraints.RegisterSources[1].RegisterConstraint
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
	expect.NotNil(t, destRegister)
	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered:  true,
			AnyGeneral: false,
			AnyFloat:   true,
			Require:    nil,
		},
		destRegister)
	expect.True(t, srcRegister != scratchRegister)
	expect.True(t, srcRegister != destRegister)
	expect.True(t, scratchRegister != destRegister)

	// Validate encoding

	builder := layout.NewSegmentBuilder()
	instruction.EmitTo(
		builder,
		map[*architecture.RegisterConstraint]*architecture.Register{
			srcRegister:     registers.Rcx,
			scratchRegister: registers.Rdx,
			destRegister:    registers.Xmm7,
		})
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			// comparison basic block

			// 0: cmp rcx, 0
			0x48, 0x81, 0xf9, 0x00, 0x00, 0x00, 0x00,
			// 7: jge <nonNegative offset> (= 43 - 13 = 30 = 0x1e)
			0x0f, 0x8d, 0x1e, 0x00, 0x00, 0x00,

			// negative branch basic block

			// 13: mov rdx, rcx
			0x48, 0x8b, 0xd1,
			// 16: shr rdx, 1
			0x48, 0xc1, 0xea, 0x01,
			// 20: and ecx, 1
			0x81, 0xe1, 0x01, 0x00, 0x00, 0x00,
			// 26: or rcx, rdx
			0x48, 0x0b, 0xca,
			// 29: cvtsi2ss xmm7, rcx
			0xf3, 0x48, 0x0f, 0x2a, 0xf9,
			// 34: addss xmm7, xmm7
			0xf3, 0x0f, 0x58, 0xff,
			// 38: jmp <end offset> (= 48 - 43 = 5)
			0xe9, 0x05, 0x00, 0x00, 0x00,

			// non-negative branch basic block

			// 43: cvtsi2ss xmm7, rcx
			0xf3, 0x48, 0x0f, 0x2a, 0xf9,

			// 48: end basic block
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
