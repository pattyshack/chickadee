package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
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
	// TODO
}

func TestUint64ToFloat64(t *testing.T) {
	// TODO
}
