package ir

const (
	generalRegisterSize = 8
	addressSize         = generalRegisterSize
)

type Type interface {
	isTypeExpression()

	Equals(Type) bool

	Alignment() int

	Size() int
}

type UnsignedIntType struct {
	ByteSize int
}

func (*UnsignedIntType) isTypeExpression() {}

func (thisInt *UnsignedIntType) Equals(other Type) bool {
	otherInt, ok := other.(*UnsignedIntType)
	if !ok {
		return false
	}

	return thisInt == otherInt // Same pointer
}

func (t *UnsignedIntType) Alignment() int {
	return t.ByteSize
}

func (t *UnsignedIntType) Size() int {
	return t.ByteSize
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
}

func (*SignedIntType) isTypeExpression() {}

func (thisInt *SignedIntType) Equals(other Type) bool {
	otherInt, ok := other.(*SignedIntType)
	if !ok {
		return false
	}

	return thisInt == otherInt // Same pointer
}

func (t *SignedIntType) Alignment() int {
	return t.ByteSize
}

func (t *SignedIntType) Size() int {
	return t.ByteSize
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

type FloatType int // in bytes

func (FloatType) isTypeExpression() {}

func (thisFloat FloatType) Equals(other Type) bool {
	otherFloat, ok := other.(FloatType)
	if !ok {
		return false
	}

	return thisFloat == otherFloat
}

func (t FloatType) Alignment() int {
	return int(t)
}

func (t FloatType) Size() int {
	return int(t)
}

const (
	Float32 = FloatType(4)
	Float64 = FloatType(8)
)

// NOTE: Unlike c pointer, int8 address type is not the same as int8 array
// address type.  We don't support general pointer arithmetic and only
// struct/array address type is index accessible.
type AddressType struct {
	ValueType Type
}

func (AddressType) isTypeExpression() {}

func (thisAddr AddressType) Equals(other Type) bool {
	otherAddr, ok := other.(AddressType)
	if !ok {
		return false
	}

	return thisAddr.ValueType.Equals(otherAddr.ValueType)
}

func (AddressType) Alignment() int {
	return addressSize
}

func (AddressType) Size() int {
	return addressSize
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

func (*FunctionType) Alignment() int {
	return addressSize
}

func (*FunctionType) Size() int {
	return addressSize
}

type ArrayType struct {
	ElementType Type
	// NOTE: use -1 for unknown length array, non-negative for fixed length array
	NumElements int

	// NOTE: (internal use only) This is not part of the type signature.
	ComputedSize int
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
	otherArray, ok := other.(*ArrayType)
	if !ok {
		return false
	}

	return thisArray.NumElements == otherArray.NumElements &&
		thisArray.ElementType.Equals(otherArray.ElementType)
}

func (t *ArrayType) Alignment() int {
	return generalRegisterSize
}

func (t *ArrayType) Size() int {
	if t.NumElements < 0 {
		panic("unknown number of elements in array")
	}

	if t.ComputedSize == 0 {
		t.ComputedSize = roundUpLastChunk(t.NumElements * t.ElementType.Size())
	}
	return t.ComputedSize
}

type Field struct {
	Name string
	Type Type

	// NOTE: (internal use only) This is not part of the type signature.
	ComputedOffset int // relative to the beginning of the record
}

// NOTE: we do not support packed struct since that complicates data
// location book keeping / code generation (a simple value in a packed struct
// may span multiple data chunks)
type StructType struct {
	Fields []*Field

	// NOTE: (internal use only) This is not part of the type signature.
	ComputedSize int
}

func (*StructType) isTypeExpression() {}

func (thisStruct *StructType) Equals(other Type) bool {
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

func (t *StructType) Alignment() int {
	return generalRegisterSize
}

func (t *StructType) Size() int {
	if t.ComputedSize == 0 {
		t.computeSize()
	}
	return t.ComputedSize
}

func (t *StructType) computeSize() {
	t.ComputedSize = 0
	for _, field := range t.Fields {
		t.addAlignmentPadding(field.Type.Alignment())
		field.ComputedOffset = t.ComputedSize
		t.ComputedSize += field.Type.Size()
	}

	t.ComputedSize = roundUpLastChunk(t.ComputedSize)
}

func (t *StructType) addAlignmentPadding(alignment int) {
	mod := t.ComputedSize % alignment
	if mod > 0 {
		t.ComputedSize += alignment - mod
	}
}

// NOTE: we'll round up the size to the nearest power of 2 that fits within
// a general register since memory operations typically operate in power of
// 2 units.
func roundUpLastChunk(size int) int {
	lastChunkSize := size % generalRegisterSize
	if lastChunkSize > 0 {
		roundUpSize := 1
		for roundUpSize < lastChunkSize {
			roundUpSize <<= 1
		}

		size += (roundUpSize - lastChunkSize)
	}

	return size
}
