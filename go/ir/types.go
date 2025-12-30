package ir

const (
	generalRegisterSize = 8
	addressSize         = generalRegisterSize
)

type Type interface {
	isTypeExpression()

	Equals(Type) bool

	Size() int

	Chunks() []*TypeChunk
}

// For internal use only.
//
// Each type is partition into chunks that could fit into 8-byte registers.
type TypeChunk struct {
	Values []TypeChunkValue
}

// For internal use only.
//
// How values are packed within a type chunk
type TypeChunkValue struct {
	// NOTE: memory address offset = 8 * Index + Offset
	Index  int // relative to the beginning of the type's value.
	Offset int // relative to the current chunk

	ValueType Type

	// NOTE: only used for tracking aggregate (struct / array) value type that
	// spans multiple chunks.
	ValueTypeChunk *TypeChunk
}

func simpleTypeChunk(t Type) []*TypeChunk {
	return []*TypeChunk{
		{
			Values: []TypeChunkValue{
				{
					Index:          0,
					Offset:         0,
					ValueType:      t,
					ValueTypeChunk: nil,
				},
			},
		},
	}
}

type UnsignedIntType struct {
	ByteSize int

	// Internal (not part of the type signature)

	chunks []*TypeChunk
}

func (*UnsignedIntType) isTypeExpression() {}

func (this *UnsignedIntType) Equals(other Type) bool {
	if this == other {
		return true
	}

	otherInt, ok := other.(*UnsignedIntType)
	if !ok {
		return false
	}

	return this.ByteSize == otherInt.ByteSize
}

func (t *UnsignedIntType) Size() int {
	return t.ByteSize
}

func (t *UnsignedIntType) Chunks() []*TypeChunk {
	if t.chunks == nil {
		t.chunks = simpleTypeChunk(t)
	}
	return t.chunks
}

var (
	Uint8 = &UnsignedIntType{
		ByteSize: 1,
	}

	Uint16 = &UnsignedIntType{
		ByteSize: 2,
	}

	Uint32 = &UnsignedIntType{
		ByteSize: 4,
	}

	Uint64 = &UnsignedIntType{
		ByteSize: 8,
	}
)

type SignedIntType struct {
	ByteSize int

	// Internal (not part of the type signature)

	chunks []*TypeChunk
}

func (*SignedIntType) isTypeExpression() {}

func (this *SignedIntType) Equals(other Type) bool {
	if this == other {
		return true
	}

	otherInt, ok := other.(*SignedIntType)
	if !ok {
		return false
	}

	return this.ByteSize == otherInt.ByteSize
}

func (t *SignedIntType) Size() int {
	return t.ByteSize
}

func (t *SignedIntType) Chunks() []*TypeChunk {
	if t.chunks == nil {
		t.chunks = simpleTypeChunk(t)
	}
	return t.chunks
}

var (
	Int8 = &SignedIntType{
		ByteSize: 1,
	}

	Int16 = &SignedIntType{
		ByteSize: 2,
	}

	Int32 = &SignedIntType{
		ByteSize: 4,
	}

	Int64 = &SignedIntType{
		ByteSize: 8,
	}
)

type FloatType struct {
	ByteSize int

	// Internal (not part of the type signature)

	chunks []*TypeChunk
}

func (*FloatType) isTypeExpression() {}

func (this *FloatType) Equals(other Type) bool {
	if this == other {
		return true
	}

	otherFloat, ok := other.(*FloatType)
	if !ok {
		return false
	}

	return this.ByteSize == otherFloat.ByteSize
}

func (t *FloatType) Size() int {
	return t.ByteSize
}

func (t *FloatType) Chunks() []*TypeChunk {
	if t.chunks == nil {
		t.chunks = simpleTypeChunk(t)
	}
	return t.chunks
}

var (
	Float32 = &FloatType{
		ByteSize: 4,
	}

	Float64 = &FloatType{
		ByteSize: 8,
	}
)

// NOTE: Unlike c pointer, int8 address type is not the same as int8 array
// address type.  We don't support general pointer arithmetic and only
// struct/array address type is index accessible.
type AddressType struct {
	ValueType Type

	// Internal (not part of the type signature)

	chunks []*TypeChunk
}

func (*AddressType) isTypeExpression() {}

func (thisAddr *AddressType) Equals(other Type) bool {
	if thisAddr == other {
		return true
	}

	otherAddr, ok := other.(*AddressType)
	if !ok {
		return false
	}

	return thisAddr.ValueType.Equals(otherAddr.ValueType)
}

func (*AddressType) Size() int {
	return addressSize
}

func (t *AddressType) Chunks() []*TypeChunk {
	if t.chunks == nil {
		t.chunks = simpleTypeChunk(t)
	}
	return t.chunks
}

type CallConventionKind string

const (
	// Simplified System V ABI.
	//
	// On amd64:
	// - This does not support SSEUP, X87, X87UP, COMPLEX_X87 parameter classes
	// as defined in "3.2.3 Parameter Passing", i.e., this does not support
	// legacy x86 values / basic value larger than 64 bits.
	// - This assumes all parameters are "trivial for the purpose of calls",
	// i.e., parameter are shallow copy-able, not C++ object with custom
	// copy constructor.
	SysVLiteCallConvention = CallConventionKind("SysVLite")
)

type FunctionType struct {
	CallConventionKind
	ParameterTypes []Type
	ReturnType     Type // use empty struct for no return value

	// NOTE: (internal use only) This is not part of the type signature.

	chunks []*TypeChunk

	CallConvention interface{} // cached architecture.CallConvention
}

func (*FunctionType) isTypeExpression() {}

func (thisFunction *FunctionType) Equals(other Type) bool {
	otherFunction, ok := other.(*FunctionType)
	if !ok {
		return false
	}

	if thisFunction.CallConventionKind != otherFunction.CallConventionKind ||
		len(thisFunction.ParameterTypes) != len(otherFunction.ParameterTypes) {

		return false
	}

	for idx, parameterType := range thisFunction.ParameterTypes {
		otherParameterType := otherFunction.ParameterTypes[idx]
		if !parameterType.Equals(otherParameterType) {
			return false
		}
	}

	if thisFunction.ReturnType == nil {
		return otherFunction.ReturnType == nil
	}

	return thisFunction.ReturnType.Equals(otherFunction.ReturnType)
}

func (*FunctionType) Size() int {
	return addressSize
}

func (t *FunctionType) Chunks() []*TypeChunk {
	if t.chunks == nil {
		t.chunks = simpleTypeChunk(t)
	}
	return t.chunks
}

type ArrayType struct {
	ElementType Type
	// NOTE: use -1 for unknown length array, non-negative for fixed length array
	NumElements int

	// NOTE: (internal use only) This is not part of the type signature.

	size   int
	chunks []*TypeChunk
}

func NewUnknownLengthArrayType(elementType Type) *ArrayType {
	return &ArrayType{
		ElementType: elementType,
		NumElements: -1,
	}
}

func NewFixedLengthArrayType(elementType Type, numElements int) *ArrayType {
	return &ArrayType{
		ElementType: elementType,
		NumElements: numElements,
	}
}

func (*ArrayType) isTypeExpression() {}

func (thisArray *ArrayType) Equals(other Type) bool {
	if thisArray == other {
		return true
	}

	otherArray, ok := other.(*ArrayType)
	if !ok {
		return false
	}

	return thisArray.NumElements == otherArray.NumElements &&
		thisArray.ElementType.Equals(otherArray.ElementType)
}

func (t *ArrayType) Size() int {
	t.maybeComputeChunks()
	return t.size
}

func (t *ArrayType) Chunks() []*TypeChunk {
	t.maybeComputeChunks()
	return t.chunks
}

func (t *ArrayType) maybeComputeChunks() {
	if t.chunks != nil {
		return
	}

	if t.NumElements < 0 {
		panic("unknown number of elements in array")
	}

	elementSize := t.ElementType.Size()
	dataSize := t.NumElements * elementSize
	numChunks := (dataSize + generalRegisterSize - 1) / generalRegisterSize

	t.size = numChunks * generalRegisterSize
	if t.size == 0 {
		t.chunks = []*TypeChunk{}
		return
	}

	chunks := make([]*TypeChunk, 0, numChunks)
	if elementSize <= generalRegisterSize {
		currentChunk := &TypeChunk{}
		currentSize := 0
		for i := 0; i < t.NumElements; i++ {
			if currentSize+elementSize > generalRegisterSize {
				chunks = append(chunks, currentChunk)
				currentChunk = &TypeChunk{}
				currentSize = 0
			}

			currentChunk.Values = append(
				currentChunk.Values,
				TypeChunkValue{
					Index:     len(chunks),
					Offset:    currentSize,
					ValueType: t.ElementType,
				})
			currentSize += elementSize
		}

		chunks = append(chunks, currentChunk)
	} else {
		// NOTE: Aggregate type sizes are always chunk aligned
		elementChunks := t.ElementType.Chunks()
		for i := 0; i < t.NumElements; i++ {
			for _, elementChunk := range elementChunks {
				chunks = append(
					chunks,
					&TypeChunk{
						Values: []TypeChunkValue{
							{
								Index:          len(chunks),
								Offset:         0,
								ValueType:      t.ElementType,
								ValueTypeChunk: elementChunk,
							},
						},
					})
			}
		}
	}

	if len(chunks) != numChunks {
		panic("should never happen")
	}
	t.chunks = chunks
}

type Field struct {
	Name string
	Type Type
}

// NOTE: we do not support packed struct since that complicates data
// location book keeping / code generation (a simple value in a packed struct
// may span multiple data chunks)
type StructType struct {
	Fields []*Field

	// NOTE: (internal use only) This is not part of the type signature.

	size   int
	chunks []*TypeChunk
}

func (*StructType) isTypeExpression() {}

func (thisStruct *StructType) Equals(other Type) bool {
	if thisStruct == other {
		return true
	}

	otherStruct, ok := other.(*StructType)
	if !ok {
		return false
	}

	if len(thisStruct.Fields) != len(otherStruct.Fields) {
		return false
	}

	for idx, field := range thisStruct.Fields {
		otherField := otherStruct.Fields[idx]
		if field.Name != otherField.Name || !field.Type.Equals(otherField.Type) {
			return false
		}
	}

	return true
}

func (t *StructType) Size() int {
	t.maybeComputeChunks()
	return len(t.chunks) * generalRegisterSize
}

func (t *StructType) Chunks() []*TypeChunk {
	t.maybeComputeChunks()
	return t.chunks
}

func (t *StructType) maybeComputeChunks() {
	if t.chunks != nil {
		return
	}

	chunks := []*TypeChunk{}
	currentChunk := &TypeChunk{}
	currentSize := 0
	for _, field := range t.Fields {
		fieldSize := field.Type.Size()
		fieldAlignment := alignment(fieldSize)

		mod := currentSize % fieldAlignment
		if mod > 0 {
			currentSize += fieldAlignment - mod
		}

		if currentSize+fieldSize > generalRegisterSize {
			chunks = append(chunks, currentChunk)
			currentChunk = &TypeChunk{}
			currentSize = 0
		}

		currentChunk.Values = append(
			currentChunk.Values,
			TypeChunkValue{
				Index:     len(chunks),
				Offset:    currentSize,
				ValueType: field.Type,
			})
		currentSize += fieldSize
	}

	chunks = append(chunks, currentChunk)

	t.size = len(chunks) * generalRegisterSize
	t.chunks = chunks
}

func alignment(typeSize int) int {
	switch typeSize {
	case 0, 1, 2, 4, 8:
		return typeSize
	}

	if typeSize%8 != 0 {
		panic("should never happen")
	}

	return 8
}
