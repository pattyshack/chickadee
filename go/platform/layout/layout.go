// Assumptions:
// - all relocations are position independent (i.e., use pc relative offset)
//   and can be incrementally linked

package layout

import (
	"fmt"
)

type Content struct {
	Size       int64
	DataChunks [][]byte
}

func (content *Content) Append(chunk []byte) {
	content.Size += int64(len(chunk))
	content.DataChunks = append(content.DataChunks, chunk)
}

func (content *Content) MaybePad(alignment int64, padding []byte) error {
	alignedSize := ((content.Size + alignment - 1) / alignment) * alignment
	numNeeded := alignedSize - content.Size

	if numNeeded == 0 {
		return nil
	}

	if numNeeded%int64(len(padding)) != 0 {
		return fmt.Errorf(
			"cannot pad content (needed: %d padding: %d)",
			numNeeded,
			len(padding))
	}

	chunk := make([]byte, numNeeded)
	if len(padding) != 1 || padding[0] != 0 {
		remaining := chunk
		for len(remaining) > 0 {
			n := copy(remaining, padding)
			remaining = remaining[n:]
		}
	}

	content.Append(chunk)
	return nil
}

func (content *Content) Merge(other Content) {
	content.Size += other.Size
	content.DataChunks = append(content.DataChunks, other.DataChunks...)
}

func (content Content) Flatten() []byte {
	if len(content.DataChunks) == 1 {
		return content.DataChunks[0]
	}

	result := make([]byte, content.Size)
	dest := result
	for _, chunk := range content.DataChunks {
		n := copy(dest, chunk)
		dest = dest[n:]
	}

	return result
}

func (content Content) Peek(startOffset int64) ([]byte, error) {
	for _, chunk := range content.DataChunks {
		chunkSize := int64(len(chunk))
		if startOffset < chunkSize {
			return chunk[startOffset:], nil
		}
		startOffset -= chunkSize
	}

	return nil, fmt.Errorf("snippet out of range")
}

type Segment struct {
	Content
	Definitions
	Relocations
}

func (segment Segment) ShiftAll(offset int64) {
	segment.Definitions.Shift(offset)
	segment.Relocations.Shift(offset)
}

func (segment Segment) Defs() Definitions {
	return segment.Definitions
}

func (segment Segment) Relocs() Relocations {
	return segment.Relocations
}

func (segment *Segment) SetRelocations(unresolved Relocations) {
	segment.Relocations = unresolved
}

type SegmentBuilder struct {
	Size     int64
	Segments []Segment
}

func NewSegmentBuilder() *SegmentBuilder {
	return &SegmentBuilder{}
}

func (builder *SegmentBuilder) Append(segment Segment) {
	segment.ShiftAll(builder.Size)

	builder.Size += segment.Size
	builder.Segments = append(builder.Segments, segment)
}

func (builder *SegmentBuilder) AppendData(
	data []byte,
	defs Definitions,
	relocs Relocations,
) {
	builder.Append(
		Segment{
			Content: Content{
				Size:       int64(len(data)),
				DataChunks: [][]byte{data},
			},
			Definitions: defs,
			Relocations: relocs,
		})
}

func (builder *SegmentBuilder) AppendBasicData(data []byte) {
	builder.AppendData(data, Definitions{}, Relocations{})
}

func (builder *SegmentBuilder) Finalize(
	config ArchitectureConfig,
) (
	Segment,
	error,
) {
	defs, labels, symbols, err := MergeDefinitions(builder.Segments...)
	if err != nil {
		return Segment{}, err
	}
	merged := Segment{
		Definitions: defs,
		Relocations: MergeRelocations(builder.Segments...),
	}

	buffered := Content{}
	for _, segment := range builder.Segments {
		buffered.Merge(segment.Content)
		if buffered.Size > config.MergeContentThreshold {
			merged.Content.Append(buffered.Flatten())
			buffered = Content{}
		}
	}

	if buffered.Size > 0 {
		merged.Content.Append(buffered.Flatten())
	}

	err = Link(&merged, labels, symbols, config.Relocator)
	if err != nil {
		return Segment{}, err
	}

	return merged, nil
}

type BSSSegment struct {
	Size int64
	Definitions
}

func (segment BSSSegment) ShiftAll(offset int64) {
	segment.Definitions.Shift(offset)
}

func (segment BSSSegment) Defs() Definitions {
	return segment.Definitions
}

func (segment BSSSegment) Relocs() Relocations {
	return Relocations{}
}

func (segment *BSSSegment) Pad(alignment int64) {
	mod := segment.Size % alignment
	if mod > 0 {
		segment.Size += alignment - mod
	}
}

type BSSSegmentBuilder struct {
	Size     int64
	Segments []BSSSegment
}

func (builder *BSSSegmentBuilder) Append(segment BSSSegment) {
	segment.ShiftAll(builder.Size)

	builder.Size += segment.Size
	builder.Segments = append(builder.Segments, segment)
}

func (builder *BSSSegmentBuilder) AppendObject(name string, size int64) {
	builder.Append(
		BSSSegment{
			Size: size,
			Definitions: Definitions{
				Symbols: []*Symbol{
					{
						Kind:    ObjectKind,
						Section: BSSSection,
						Name:    name,
						Size:    size,
					},
				},
			},
		})
}

func (builder *BSSSegmentBuilder) Finalize() (BSSSegment, error) {
	defs, _, _, err := MergeDefinitions(builder.Segments...)
	if err != nil {
		return BSSSegment{}, err
	}

	return BSSSegment{
		Size:        builder.Size,
		Definitions: defs,
	}, nil
}

type ObjectFileBuilder struct {
	Text         SegmentBuilder
	Init         SegmentBuilder
	ReadOnlyData SegmentBuilder
	Data         SegmentBuilder
	BSS          BSSSegmentBuilder
}

func NewObjectFileBuilder() ObjectFileBuilder {
	return ObjectFileBuilder{}
}

func (builder *ObjectFileBuilder) Merge(file ObjectFile) {
	builder.Text.Append(file.Text)
	builder.Init.Append(file.Init)
	builder.ReadOnlyData.Append(file.ReadOnlyData)
	builder.Data.Append(file.Data)
	builder.BSS.Append(file.BSS)
}

func (builder *ObjectFileBuilder) Finalize(
	config Config,
) (
	ObjectFile,
	error,
) {
	file := ObjectFile{}

	text, err := builder.Text.Finalize(config.Architecture)
	if err != nil {
		return ObjectFile{}, err
	}
	file.Text = text

	init, err := builder.Init.Finalize(config.Architecture)
	if err != nil {
		return ObjectFile{}, err
	}
	file.Init = init

	readOnlyData, err := builder.ReadOnlyData.Finalize(config.Architecture)
	if err != nil {
		return ObjectFile{}, err
	}
	file.ReadOnlyData = readOnlyData

	data, err := builder.Data.Finalize(config.Architecture)
	if err != nil {
		return ObjectFile{}, err
	}
	file.Data = data

	bss, err := builder.BSS.Finalize()
	if err != nil {
		return ObjectFile{}, err
	}
	file.BSS = bss

	return file, nil
}

type ObjectFile struct {
	Text         Segment
	Init         Segment
	ReadOnlyData Segment
	Data         Segment
	BSS          BSSSegment
}

// NOTE: start symbol is not part of Config since each module may have its
// own start symbol (a module may be both a library and a binary).
func (file ObjectFile) ToExecutableImage(
	config Config,
	startSymbol string,
) (
	ExecutableImage,
	error,
) {
	pageSize := config.Architecture.MemoryPageSize
	alignment := config.Architecture.RegisterAlignment
	if pageSize%alignment != 0 {
		return ExecutableImage{}, fmt.Errorf(
			"memory page size (%d) not multiples of section alignment (%d)",
			pageSize,
			alignment)
	}

	if file.Text.Size == 0 {
		return ExecutableImage{}, fmt.Errorf("empty .text segment")
	}

	err := file.Text.MaybePad(alignment, config.InstructionPadding)
	if err != nil {
		return ExecutableImage{}, err
	}

	file.Init.Append(config.InitEpilogue)
	file.Init.Definitions.Symbols = append(
		file.Init.Definitions.Symbols,
		&Symbol{
			Kind:    FunctionKind,
			Section: InitSection,
			Name:    config.InitSymbol,
			Offset:  0,
			Size:    file.Init.Size,
		})

	err = file.Init.MaybePad(alignment, config.InstructionPadding)
	if err != nil {
		return ExecutableImage{}, err
	}

	err = file.ReadOnlyData.MaybePad(alignment, config.DataPadding)
	if err != nil {
		return ExecutableImage{}, err
	}

	err = file.Data.MaybePad(alignment, config.DataPadding)
	if err != nil {
		return ExecutableImage{}, err
	}

	file.BSS.Pad(config.Architecture.RegisterAlignment)

	image := ExecutableImage{
		MemoryPageSize:    pageSize,
		RegisterAlignment: alignment,
		Text:              file.Text.Content,
		Init:              file.Init.Content,
		ReadOnlyData:      file.ReadOnlyData.Content,
		Data:              file.Data.Content,
		BSSSize:           file.BSS.Size,
	}

	offset := config.ExecutableImageStartPage * pageSize

	image.ExecutableSegmentStart = offset
	file.Text.ShiftAll(offset)
	offset += image.Text.Size

	file.Init.ShiftAll(offset)
	offset += image.Init.Size

	offset = ((offset + (pageSize - 1)) / pageSize) * pageSize
	image.ReadOnlySegmentStart = offset
	file.ReadOnlyData.ShiftAll(offset)
	offset += image.ReadOnlyData.Size

	offset = ((offset + (pageSize - 1)) / pageSize) * pageSize
	image.ReadWriteSegmentStart = offset
	file.Data.ShiftAll(offset)
	offset += image.Data.Size

	file.BSS.ShiftAll(offset)

	segments := []RelocatableInfo{
		file.Text,
		file.Init,
		file.ReadOnlyData,
		file.Data,
		file.BSS,
	}

	defs, labels, symbols, err := MergeDefinitions(segments...)
	if err != nil {
		return ExecutableImage{}, err
	}

	if len(labels) > 0 {
		return ExecutableImage{}, fmt.Errorf("unexpected labels in object file")
	}

	image.Definitions = defs
	image.Relocations = MergeRelocations(segments...)

	if len(image.Relocations.Labels) > 0 {
		return ExecutableImage{}, fmt.Errorf(
			"unexpected label relocations in object file")
	}

	start, ok := symbols[startSymbol]
	if !ok || start.Kind != FunctionKind {
		return ExecutableImage{}, fmt.Errorf(
			"start function symbol (%s) not found",
			startSymbol)
	}
	image.EntryPoint = start.Offset

	err = Link(&image, labels, symbols, config.Architecture.Relocator)
	if err != nil {
		return ExecutableImage{}, err
	}

	if len(image.Relocations.Symbols) > 0 {
		return ExecutableImage{}, fmt.Errorf(
			"unresolved symbol relocations in object file")
	}

	return image, nil
}

type ExecutableImage struct {
	MemoryPageSize    int64
	RegisterAlignment int64

	EntryPoint int64

	ExecutableSegmentStart int64
	Text                   Content
	Init                   Content

	ReadOnlySegmentStart int64
	ReadOnlyData         Content

	ReadWriteSegmentStart int64
	Data                  Content
	BSSSize               int64

	Definitions
	Relocations
}

func (image ExecutableImage) ShiftAll(offset int64) {
	image.Definitions.Shift(offset)
	image.Relocations.Shift(offset)
}

func (image ExecutableImage) Defs() Definitions {
	return image.Definitions
}

func (image ExecutableImage) Relocs() Relocations {
	return image.Relocations
}

func (image ExecutableImage) Peek(startOffset int64) ([]byte, error) {
	if startOffset < image.ExecutableSegmentStart {
		return nil, fmt.Errorf("snippet out of range")
	}

	start := image.ExecutableSegmentStart
	end := start + image.Text.Size
	if startOffset < end {
		return image.Text.Peek(startOffset - start)
	}

	start = end
	end += image.Init.Size
	if startOffset < end {
		return image.Init.Peek(startOffset - start)
	}

	if startOffset < image.ReadOnlySegmentStart {
		return nil, fmt.Errorf("snippet out of range")
	}

	start = image.ReadOnlySegmentStart
	end = start + image.ReadOnlyData.Size
	if startOffset < end {
		return image.ReadOnlyData.Peek(startOffset - start)
	}

	if startOffset < image.ReadWriteSegmentStart {
		return nil, fmt.Errorf("snippet out of range")
	}

	start = image.ReadWriteSegmentStart
	end = start + image.Data.Size
	if startOffset < end {
		return image.Data.Peek(startOffset - start)
	}

	return nil, fmt.Errorf("snippet out of range")
}

func (image *ExecutableImage) SetRelocations(unresolved Relocations) {
	image.Relocations = unresolved
}
