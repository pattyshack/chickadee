package instructions

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	amd64 "github.com/pattyshack/chickadee/amd64/layout"
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

func TestFloat32ToFloat32(t *testing.T) {
	// movss xmm1, xmm2 (no-op conversion)
	builder := layout.NewSegmentBuilder()
	convertFloatToFloat(
		builder,
		ir.Float32,
		registers.Xmm1,
		ir.Float32,
		registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf3, 0x0f, 0x10, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToFloat64(t *testing.T) {
	// cvtss2sd xmm10, xmm5
	builder := layout.NewSegmentBuilder()
	convertFloatToFloat(
		builder,
		ir.Float64,
		registers.Xmm10,
		ir.Float32,
		registers.Xmm5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x5a, 0xd5},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToInt8(t *testing.T) {
	// cvtss2si ebp, xmm7
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int8,
		registers.Rbp,
		ir.Float32,
		registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x0f, 0x2d, 0xef},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToInt16(t *testing.T) {
	// cvtss2si edx, xmm14
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int16,
		registers.Rdx,
		ir.Float32,
		registers.Xmm14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x41, 0x0f, 0x2d, 0xd6},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToInt32(t *testing.T) {
	// cvtss2si r9d, xmm4
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int32,
		registers.R9,
		ir.Float32,
		registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x2d, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToInt64(t *testing.T) {
	// cvtss2si rcx, xmm3
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int64,
		registers.Rcx,
		ir.Float32,
		registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x48, 0x0f, 0x2d, 0xcb},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToUint8(t *testing.T) {
	// cvtss2si ebp, xmm7
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint8,
		registers.Rbp,
		ir.Float32,
		registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x0f, 0x2d, 0xef},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToUint16(t *testing.T) {
	// cvtss2si edx, xmm14
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint16,
		registers.Rdx,
		ir.Float32,
		registers.Xmm14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x41, 0x0f, 0x2d, 0xd6},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToUint32(t *testing.T) {
	// cvtss2si r9d, xmm4
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint32,
		registers.R9,
		ir.Float32,
		registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x44, 0x0f, 0x2d, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat32ToUint64(t *testing.T) {
	// cvtss2si rcx, xmm3
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint64,
		registers.Rcx,
		ir.Float32,
		registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf3, 0x48, 0x0f, 0x2d, 0xcb},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToFloat32(t *testing.T) {
	// cvtsd2ss xmm1, xmm2
	builder := layout.NewSegmentBuilder()
	convertFloatToFloat(
		builder,
		ir.Float32,
		registers.Xmm1,
		ir.Float64,
		registers.Xmm2)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(t, []byte{0xf2, 0x0f, 0x5a, 0xca}, segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToFloat64(t *testing.T) {
	// movsd xmm10, xmm5 (no-op conversion)
	builder := layout.NewSegmentBuilder()
	convertFloatToFloat(
		builder,
		ir.Float64,
		registers.Xmm10,
		ir.Float64,
		registers.Xmm5)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x44, 0x0f, 0x10, 0xd5},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToInt8(t *testing.T) {
	// cvtsd2si ebp, xmm7
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int8,
		registers.Rbp,
		ir.Float64,
		registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x0f, 0x2d, 0xef},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToInt16(t *testing.T) {
	// cvtsd2si edx, xmm14
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int16,
		registers.Rdx,
		ir.Float64,
		registers.Xmm14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x41, 0x0f, 0x2d, 0xd6},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToInt32(t *testing.T) {
	// cvtsd2si r9d, xmm4
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int32,
		registers.R9,
		ir.Float64,
		registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x44, 0x0f, 0x2d, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToInt64(t *testing.T) {
	// cvtsd2si rcx, xmm3
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Int64,
		registers.Rcx,
		ir.Float64,
		registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x48, 0x0f, 0x2d, 0xcb},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToUint8(t *testing.T) {
	// cvtsd2si ebp, xmm7
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint8,
		registers.Rbp,
		ir.Float64,
		registers.Xmm7)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x0f, 0x2d, 0xef},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToUint16(t *testing.T) {
	// cvtsd2si edx, xmm14
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint16,
		registers.Rdx,
		ir.Float64,
		registers.Xmm14)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x41, 0x0f, 0x2d, 0xd6},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToUint32(t *testing.T) {
	// cvtsd2si r9d, xmm4
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint32,
		registers.R9,
		ir.Float64,
		registers.Xmm4)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x44, 0x0f, 0x2d, 0xcc},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}

func TestFloat64ToUint64(t *testing.T) {
	// cvtsd2si rcx, xmm3
	builder := layout.NewSegmentBuilder()
	convertFloatToInt(
		builder,
		ir.Uint64,
		registers.Rcx,
		ir.Float64,
		registers.Xmm3)
	segment, err := builder.Finalize(amd64.ArchitectureLayout)
	expect.Nil(t, err)
	expect.Equal(
		t,
		[]byte{0xf2, 0x48, 0x0f, 0x2d, 0xcb},
		segment.Content.Flatten())
	expect.Equal(t, layout.Definitions{}, segment.Definitions)
	expect.Equal(t, layout.Relocations{}, segment.Relocations)
}
