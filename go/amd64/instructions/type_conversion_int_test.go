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

func TestInt8ToUint8(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rcx,
		ir.Int8,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToUint16(t *testing.T) {
	// movsx ecx, dil
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Int8,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x0f, 0xbe, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToUint32(t *testing.T) {
	// movsx esi, al
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rsi,
		ir.Int8,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xbe, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToUint64(t *testing.T) {
	// movsx r10, bpl
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.R10,
		ir.Int8,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0f, 0xbe, 0xd5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToUint8(t *testing.T) {
	// mov edx, esi (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rdx,
		ir.Int16,
		registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToUint16(t *testing.T) {
	// mov ebp, r8d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rbp,
		ir.Int16,
		registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToUint32(t *testing.T) {
	// movsx edx, di
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rdx,
		ir.Int16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xbf, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToUint64(t *testing.T) {
	// movsx r9, cx
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.R9,
		ir.Int16,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0f, 0xbf, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToUint8(t *testing.T) {
	// mov ebx, ebp (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rbx,
		ir.Int32,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xdd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToUint16(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Int32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToUint32(t *testing.T) {
	// mov eax, r15d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rax,
		ir.Int32,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xc7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToUint64(t *testing.T) {
	// movsxd r14, r10d
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.R14,
		ir.Int32,
		registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4d, 0x63, 0xf2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToUint8(t *testing.T) {
	// mov esi, r11d (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint8,
		registers.Rsi,
		ir.Int64,
		registers.R11)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xf3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToUint16(t *testing.T) {
	// mov ecx, ebp (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint16,
		registers.Rcx,
		ir.Int64,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xcd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToUint32(t *testing.T) {
	// mov ebx, eax (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint32,
		registers.Rbx,
		ir.Int64,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToUint64(t *testing.T) {
	// mov rdx, rcx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Uint64,
		registers.Rdx,
		ir.Int64,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x8b, 0xd1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToInt8(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rcx,
		ir.Int8,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToInt16(t *testing.T) {
	// movsx ecx, dil
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Int8,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x40, 0x0f, 0xbe, 0xcf}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToInt32(t *testing.T) {
	// movsx esi, al
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rsi,
		ir.Int8,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xbe, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToInt64(t *testing.T) {
	// movsx r10, bpl
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R10,
		ir.Int8,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0f, 0xbe, 0xd5}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToInt8(t *testing.T) {
	// mov edx, esi (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rdx,
		ir.Int16,
		registers.Rsi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd6}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToInt16(t *testing.T) {
	// mov ebp, r8d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rbp,
		ir.Int16,
		registers.R8)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xe8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToInt32(t *testing.T) {
	// movsx edx, di
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rdx,
		ir.Int16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x0f, 0xbf, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToInt64(t *testing.T) {
	// movsx r9, cx
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R9,
		ir.Int16,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4c, 0x0f, 0xbf, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToInt8(t *testing.T) {
	// mov ebx, ebp (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rbx,
		ir.Int32,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xdd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToInt16(t *testing.T) {
	// mov ecx, edx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Int32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToInt32(t *testing.T) {
	// mov eax, r15d (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rax,
		ir.Int32,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xc7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToInt64(t *testing.T) {
	// movsxd r14, r10d
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.R14,
		ir.Int32,
		registers.R10)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x4d, 0x63, 0xf2}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToInt8(t *testing.T) {
	// mov esi, r11d (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int8,
		registers.Rsi,
		ir.Int64,
		registers.R11)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x41, 0x8b, 0xf3}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToInt16(t *testing.T) {
	// mov ecx, ebp (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int16,
		registers.Rcx,
		ir.Int64,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xcd}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToInt32(t *testing.T) {
	// mov ebx, eax (truncate to 32-bit)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int32,
		registers.Rbx,
		ir.Int64,
		registers.Rax)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x8b, 0xd8}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToInt64(t *testing.T) {
	// mov rdx, rcx (no-op type conversion)
	builder := layout.NewSegmentBuilder()
	convertIntToInt(
		builder,
		ir.Int64,
		registers.Rdx,
		ir.Int64,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0x48, 0x8b, 0xd1}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToFloat32(t *testing.T) {
	// movsx ecx, cl
	// cvtsi2ss xmm5, ecx
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float32,
		registers.Xmm5,
		ir.Int8,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xbe, 0xc9, // movsx
			0xf3, 0x0f, 0x2a, 0xe9, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt8ToFloat64(t *testing.T) {
	// movsx ebp, bpl
	// cvtsi2sd xmm2, ebp
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float64,
		registers.Xmm2,
		ir.Int8,
		registers.Rbp)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x40, 0x0f, 0xbe, 0xed, // movsx
			0xf2, 0x0f, 0x2a, 0xd5, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToFloat32(t *testing.T) {
	// movsx edi, di
	// cvtsi2ss xmm0, edi
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float32,
		registers.Xmm0,
		ir.Int16,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x0f, 0xbf, 0xff, // movsx
			0xf3, 0x0f, 0x2a, 0xc7, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt16ToFloat64(t *testing.T) {
	// movsx r15d, r15w
	// cvtsi2sd xmm7, r15d
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float64,
		registers.Xmm7,
		ir.Int16,
		registers.R15)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0x45, 0x0f, 0xbf, 0xff, // movsx
			0xf2, 0x41, 0x0f, 0x2a, 0xff, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToFloat32(t *testing.T) {
	// cvtsi2ss xmm12, ecx
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float32,
		registers.Xmm12,
		ir.Int32,
		registers.Rcx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf3, 0x44, 0x0f, 0x2a, 0xe1, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt32ToFloat64(t *testing.T) {
	// cvtsi2sd xmm7, edx
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float64,
		registers.Xmm7,
		ir.Int32,
		registers.Rdx)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf2, 0x0f, 0x2a, 0xfa, // cvtsi2sd
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToFloat32(t *testing.T) {
	// cvtsi2ss xmm6, rdi
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float32,
		registers.Xmm6,
		ir.Int64,
		registers.Rdi)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf3, 0x48, 0x0f, 0x2a, 0xf7, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestInt64ToFloat64(t *testing.T) {
	// cvtsi2sd xmm4, r11
	builder := layout.NewSegmentBuilder()
	convertSignedIntToFloat(
		builder,
		ir.Float64,
		registers.Xmm4,
		ir.Int64,
		registers.R11)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{
			0xf2, 0x49, 0x0f, 0x2a, 0xe3, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToUint(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int8,
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
	expect.Equal(t, []byte{0x0f, 0xbe, 0xf0}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToUintCheapSource(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int8,
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
			CheapRegisterSources: map[*ir.DefinitionChunk]struct{}{
				srcChunk: struct{}{},
			},
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
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.NotNil(t, destRegister)
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
	expect.Equal(t, []byte{0x0f, 0xbe, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToUintPreferredReuse(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int8,
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
			PreferredRegisterDestination: map[*ir.DefinitionChunk]*ir.DefinitionChunk{
				destChunk: srcChunk,
			},
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
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.NotNil(t, destRegister)
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
	expect.Equal(t, []byte{0x0f, 0xbe, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToUintNoFreeRegisters(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int8,
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
			NumFreeGeneralRegisters: 0,
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
			Clobbered:  true,
			AnyGeneral: true,
			AnyFloat:   false,
			Require:    nil,
		},
		srcRegister)

	destRegister := constraints.RegisterDestinations[0].RegisterConstraint
	expect.NotNil(t, destRegister)
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
	expect.Equal(t, []byte{0x0f, 0xbe, 0xc9}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToInt(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int16,
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
	expect.Equal(t, []byte{0x0f, 0xbf, 0xd7}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestSelectIntToFloat(t *testing.T) {
	src := ir.NewLocalReference("src")
	srcChunk := &ir.DefinitionChunk{}
	srcDef := &ir.Definition{
		Name:   "src",
		Type:   ir.Int32,
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
			0xf3, 0x44, 0x0f, 0x2a, 0xe1, // cvtsi2ss
		},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
