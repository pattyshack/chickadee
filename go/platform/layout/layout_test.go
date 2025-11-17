package layout

import (
	"fmt"
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"
)

var (
	testConfig = Config{
		MergeContentThreshold:    20,
		RegisterAlignment:        5,
		MemoryPageSize:           10,
		ExecutableImageStartPage: 1,
		InstructionPadding:       []byte("!"),
		DataPadding:              []byte("#"),
		InitSymbol:               "initFunc",
		InitEpilogue:             []byte("(init_end)"),
		Relocator:                newTestRelocator(),
	}
)

type testRelocator struct{}

func (r testRelocator) Relocate(
	symbol *Symbol,
	startOffset int64,
	snippet []byte,
) error {
	switch symbol.Kind {
	case BasicBlockKind:
		return r.relocateBasicBlock(symbol, startOffset, snippet)
	case FunctionKind:
		return r.relocateFunction(symbol, startOffset, snippet)
	case ObjectKind:
		return r.relocateObject(symbol, startOffset, snippet)
	default:
		return fmt.Errorf("unsupported symbol kind (%s)", symbol.Kind)
	}
}

func (testRelocator) relocateBasicBlock(
	symbol *Symbol,
	startOffset int64,
	snippet []byte,
) error {
	// +6 (XXXX + "B\n") since the jump is relative to the next instruction
	result := fmt.Sprintf("%d", symbol.Offset-(startOffset+6))
	for len(result) < 4 {
		result += "_"
	}

	if len(result) != 4 {
		panic("should never happen in test")
	}

	copy(snippet, []byte(result))
	return nil
}

func (testRelocator) relocateFunction(
	symbol *Symbol,
	startOffset int64,
	snippet []byte,
) error {
	// +8 (XXXXXX + "F\n") since the jump is relative to the next instruction
	result := fmt.Sprintf("%d", symbol.Offset-(startOffset+8))
	for len(result) < 6 {
		result += "_"
	}

	if len(result) != 6 {
		panic("should never happen in test")
	}

	copy(snippet, []byte(result))
	return nil
}

func (testRelocator) relocateObject(
	symbol *Symbol,
	startOffset int64,
	snippet []byte,
) error {
	// +10 (XXXXXXXX + "O\n") since the jump is relative to the next instruction
	result := fmt.Sprintf("%d", symbol.Offset-(startOffset+10))
	for len(result) < 8 {
		result += "_"
	}

	if len(result) != 8 {
		panic("should never happen in test")
	}

	copy(snippet, []byte(result))
	return nil
}

func newTestRelocator() Relocator {
	return testRelocator{}
}

type LayoutSuite struct{}

func TestLayout(t *testing.T) {
	suite.RunTests(t, &LayoutSuite{})
}

func (LayoutSuite) TestLinkBasicBlocks(t *testing.T) {
	builder := SegmentBuilder{}

	builder.AppendBasicData([]byte("mov eax, 1\n"))

	builder.AppendData(
		[]byte("cmp edi, 1\n"),
		Definitions{
			Labels: []*Symbol{
				{
					Kind: BasicBlockKind,
					Name: ".loop",
				},
			},
		},
		Relocations{})

	builder.AppendData(
		[]byte("jle XXXXB\n"),
		Definitions{},
		Relocations{
			Labels: []*Relocation{
				{
					Name:   ".return",
					Offset: 4, // replaces XXXX
				},
			},
		})

	builder.AppendBasicData([]byte("imul eax, edi\n"))
	builder.AppendBasicData([]byte("sub edi, 1\n"))

	builder.AppendData(
		[]byte("jump XXXXB\n"),
		Definitions{},
		Relocations{
			Labels: []*Relocation{
				{
					Name:   ".loop",
					Offset: 5, // replaces XXXX
				},
			},
		})

	builder.AppendData(
		[]byte("ret\n"),
		Definitions{
			Labels: []*Symbol{
				{
					Kind: BasicBlockKind,
					Name: ".return",
				},
			},
		},
		Relocations{})

	builder.AppendData(
		[]byte("unresolved XXXXB\n"),
		Definitions{},
		Relocations{
			Labels: []*Relocation{
				{
					Name:   ".unresolved",
					Offset: 11,
				},
			},
		})

	segment, err := builder.Finalize(testConfig)
	expect.Nil(t, err)

	expect.Equal(
		t,
		[]*Symbol{
			{
				Kind:   BasicBlockKind,
				Name:   ".loop",
				Offset: 11,
			},
			{
				Kind:   BasicBlockKind,
				Name:   ".return",
				Offset: 68,
			},
		},
		segment.Definitions.Labels)
	expect.Equal(t, 0, len(segment.Definitions.Symbols))

	expect.Equal(
		t,
		Relocations{
			Labels: []*Relocation{
				{
					Name:   ".unresolved",
					Offset: 83,
				},
			},
		},
		segment.Relocations)

	expect.True(t, len(segment.Content.DataChunks) > 1)
	expect.Equal(
		t,
		"mov eax, 1\n"+
			"cmp edi, 1\n"+
			"jle 36__B\n"+ // distance between ret (68) and imul (32)
			"imul eax, edi\n"+
			"sub edi, 1\n"+
			"jump -57_B\n"+ // distance between cmp (11) and ret (68)
			"ret\n"+
			"unresolved XXXXB\n",
		string(segment.Content.Flatten()))
}

func (LayoutSuite) TestLinkFunctions(t *testing.T) {
	builder := SegmentBuilder{}

	unresolvedObj := []byte("read  XXXXXXXXO\nret\n") // len = 20
	builder.AppendData(
		unresolvedObj,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind: FunctionKind,
					Name: "UnresolvedObj",
					Size: int64(len(unresolvedObj)),
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "UnknownObj",
					Offset: 6,
				},
			},
		})

	unresolvedFunc := []byte("call XXXXXXF\nret\n") // len = 17
	builder.AppendData(
		unresolvedFunc,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind: FunctionKind,
					Name: "UnresolvedFunc",
					Size: int64(len(unresolvedFunc)),
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "UnknownFunc",
					Offset: 5,
				},
			},
		})

	one := []byte("mov eax, 1\nret\n") // len = 15
	builder.AppendData(
		one,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind: FunctionKind,
					Name: "One",
					Size: int64(len(one)),
				},
			},
		},
		Relocations{})

	bar := []byte("call  XXXXXXF\nret\n") // len = 18
	builder.AppendData(
		bar,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind: FunctionKind,
					Name: "Bar",
					Size: int64(len(bar)),
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "Foo",
					Offset: 6,
				},
			},
		})

	foo := []byte("call XXXXXXF\nret\n") // len = 17
	builder.AppendData(
		foo,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind: FunctionKind,
					Name: "Foo",
					Size: int64(len(foo)),
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "One",
					Offset: 5,
				},
			},
		})

	segment, err := builder.Finalize(testConfig)
	expect.Nil(t, err)

	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:   FunctionKind,
					Name:   "UnresolvedObj",
					Offset: 0,
					Size:   int64(len(unresolvedObj)),
				},
				{
					Kind:   FunctionKind,
					Name:   "UnresolvedFunc",
					Offset: 20,
					Size:   int64(len(unresolvedFunc)),
				},
				{
					Kind:   FunctionKind,
					Name:   "One",
					Offset: 20 + 17,
					Size:   int64(len(one)),
				},
				{
					Kind:   FunctionKind,
					Name:   "Bar",
					Offset: 20 + 17 + 15,
					Size:   int64(len(bar)),
				},
				{
					Kind:   FunctionKind,
					Name:   "Foo",
					Offset: 20 + 17 + 15 + 18,
					Size:   int64(len(foo)),
				},
			},
		},
		segment.Definitions)

	expect.Equal(
		t,
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "UnknownObj",
					Offset: 6,
				},
				{
					Name:   "UnknownFunc",
					Offset: 25,
				},
			},
		},
		segment.Relocations)

	expect.Equal(
		t,
		"read  XXXXXXXXO\nret\n"+ // UnresolvedObj
			"call XXXXXXF\nret\n"+ // UnresolvedFunc
			"mov eax, 1\nret\n"+ // One
			"call  4_____F\nret\n"+ // Bar
			"call -46___F\nret\n", // Foo
		string(segment.Content.Flatten()))
}

func (LayoutSuite) TestLinkGlobalObjects(t *testing.T) {
	builder := SegmentBuilder{}

	builder.AppendData(
		[]byte("read XXXXXXXXO\n"),
		Definitions{},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "object1",
					Offset: 5,
				},
			},
		})
	builder.AppendData(
		[]byte("read  XXXXXXXXO\n"),
		Definitions{},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "unknown",
					Offset: 6,
				},
			},
		})
	builder.AppendData(
		[]byte("read   XXXXXXXXO\n"),
		Definitions{},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "object2",
					Offset: 7,
				},
			},
		})
	builder.AppendData(
		[]byte("aaaa"),
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:   ObjectKind,
					Name:   "object1",
					Offset: 0,
					Size:   4,
				},
			},
		},
		Relocations{})
	builder.AppendData(
		[]byte("bbbb"),
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:   ObjectKind,
					Name:   "object2",
					Offset: 0,
					Size:   4,
				},
			},
		},
		Relocations{})

	segment, err := builder.Finalize(testConfig)
	expect.Nil(t, err)

	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:   ObjectKind,
					Name:   "object1",
					Offset: 48,
					Size:   4,
				},
				{
					Kind:   ObjectKind,
					Name:   "object2",
					Offset: 52,
					Size:   4,
				},
			},
		},
		segment.Definitions)

	expect.Equal(
		t,
		Relocations{
			Labels: nil,
			Symbols: []*Relocation{
				{
					Name:   "unknown",
					Offset: 21,
				},
			},
		},
		segment.Relocations)

	expect.Equal(
		t,
		"read 33______O\n"+ // 48 - 15
			"read  XXXXXXXXO\n"+
			"read   4_______O\n"+ // 52 - 48
			"aaaa"+ // object1 48
			"bbbb", // object2 52
		string(segment.Content.Flatten()))
}

func (LayoutSuite) TestObjectFileAndExecutableImage(t *testing.T) {
	builder := NewObjectFileBuilder()

	textContent := "text XXXXXXXXO; XXXXXXF; " // len = 25, start = 10
	builder.Text.AppendData(
		[]byte(textContent),
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    FunctionKind,
					Section: TextSection,
					Name:    "textFunc",
					Offset:  2, // -> 10 + 2 = 12
					Size:    5,
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "roData",
					Offset: 5,
				},
				{
					Name:   "initFunc",
					Offset: 16,
				},
			},
		})

	initContent := "init XXXXXXXXO; " // len = 16, start = 10 + 25 = 35
	builder.Init.AppendData(
		[]byte(initContent),
		Definitions{},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "rwData",
					Offset: 5,
				},
			},
		})

	rodataContent := "rodata XXXXXXF;  " // len = 17, start = 60
	builder.ReadOnlyData.AppendData(
		[]byte(rodataContent),
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    ObjectKind,
					Section: ReadOnlyDataSection,
					Name:    "roData",
					Offset:  0, // -> 60
					Size:    7,
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "textFunc",
					Offset: 7,
				},
			},
		})

	rwdataContent := "rwdata XXXXXXXXO; " // len = 18, start = 80
	builder.Data.AppendData(
		[]byte(rwdataContent),
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    ObjectKind,
					Section: ReadWriteDataSection,
					Name:    "rwData",
					Offset:  1, // -> 80 + 1 = 81
					Size:    8,
				},
			},
		},
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "bssData",
					Offset: 7,
				},
			},
		})

	builder.BSS.AppendObject("bssData", 9)

	//
	// Verify intermediate object file
	//

	file, err := builder.Finalize(testConfig)
	expect.Nil(t, err)

	expect.Equal(t, textContent, string(file.Text.Content.Flatten()))
	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    FunctionKind,
					Section: TextSection,
					Name:    "textFunc",
					Offset:  2, // -> 10 + 2 = 12
					Size:    5,
				},
			},
		},
		file.Text.Definitions)
	expect.Equal(
		t,
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "roData",
					Offset: 5,
				},
				{
					Name:   "initFunc",
					Offset: 16,
				},
			},
		},
		file.Text.Relocations)

	expect.Equal(t, initContent, string(file.Init.Content.Flatten()))
	expect.Equal(t, Definitions{}, file.Init.Definitions)
	expect.Equal(
		t,
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "rwData",
					Offset: 5,
				},
			},
		},
		file.Init.Relocations)

	expect.Equal(t, rodataContent, string(file.ReadOnlyData.Content.Flatten()))
	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    ObjectKind,
					Section: ReadOnlyDataSection,
					Name:    "roData",
					Offset:  0,
					Size:    7,
				},
			},
		},
		file.ReadOnlyData.Definitions)
	expect.Equal(
		t,
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "textFunc",
					Offset: 7,
				},
			},
		},
		file.ReadOnlyData.Relocations)

	expect.Equal(t, rwdataContent, string(file.Data.Content.Flatten()))
	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    ObjectKind,
					Section: ReadWriteDataSection,
					Name:    "rwData",
					Offset:  1, // -> 80 + 1 = 81
					Size:    8,
				},
			},
		},
		file.Data.Definitions)
	expect.Equal(
		t,
		Relocations{
			Symbols: []*Relocation{
				{
					Name:   "bssData",
					Offset: 7,
				},
			},
		},
		file.Data.Relocations)

	expect.Equal(t, 9, file.BSS.Size)
	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    ObjectKind,
					Section: BSSSection,
					Name:    "bssData",
					Offset:  0,
					Size:    9,
				},
			},
		},
		file.BSS.Definitions)

	//
	// Verify executable image
	//

	image, err := file.ToExecutableImage(testConfig, "textFunc")
	expect.Nil(t, err)

	expect.Equal(t, 12, image.EntryPoint)
	expect.Equal(t, 10, image.ExecutableSegmentStart)
	expect.Equal(t, 70, image.ReadOnlySegmentStart)
	expect.Equal(t, 90, image.ReadWriteSegmentStart)

	expect.Equal(
		t,
		// roData delta = (70 + 0) - (10 + 5) - 10 = 45
		// initFunc delta = (10 + 25 + 0) - (10 + 16) - 8 = 1
		"text 45______O; 1_____F; ",
		string(image.Text.Flatten()))

	expect.Equal(
		t,
		// rwData delta = (90 + 1) - (10 + 25 + 5) - 10 = 41
		"init 41______O; (init_end)!!!!",
		string(image.Init.Flatten()))

	expect.Equal(
		t,
		// textFunc delta = (10 + 2) - (70 + 7) - 8 = -73
		"rodata -73___F;  ###",
		string(image.ReadOnlyData.Flatten()))

	expect.Equal(
		t,
		// bssData delta = (90 + 18 + 2) - (90 + 7) - 10 = 3
		"rwdata 3_______O; ##",
		string(image.Data.Flatten()))

	expect.Equal(
		t,
		Definitions{
			Symbols: []*Symbol{
				{
					Kind:    FunctionKind,
					Section: TextSection,
					Name:    "textFunc",
					Offset:  10 + 2,
					Size:    5,
				},
				{
					Kind:    FunctionKind,
					Section: InitSection,
					Name:    "initFunc",
					Offset:  10 + 25 + 0,
					Size:    26, // original init content (16) + epilogue (10)
				},
				{
					Kind:    ObjectKind,
					Section: ReadOnlyDataSection,
					Name:    "roData",
					Offset:  70 + 0,
					Size:    7,
				},
				{
					Kind:    ObjectKind,
					Section: ReadWriteDataSection,
					Name:    "rwData",
					Offset:  90 + 1,
					Size:    8,
				},
				{
					Kind:    ObjectKind,
					Section: BSSSection,
					Name:    "bssData",
					Offset:  90 + 20 + 0,
					Size:    9,
				},
			},
		},
		image.Definitions)

	expect.Equal(t, Relocations{}, image.Relocations)
}
