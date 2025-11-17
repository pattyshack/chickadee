package ir

const (
	generalRegisterSize = 8
	pointerSize         = generalRegisterSize
)

type Type interface {
	isTypeExpression()

	Equals(Type) bool

	Alignment() int

	Size() int
}

type UnsignedIntType int // in bytes

func (UnsignedIntType) isTypeExpression() {}

func (thisInt UnsignedIntType) Equals(other Type) bool {
	otherInt, ok := other.(UnsignedIntType)
	if !ok {
		return false
	}

	return thisInt == otherInt
}

func (t UnsignedIntType) Alignment() int {
	return int(t)
}

func (t UnsignedIntType) Size() int {
	return int(t)
}

const (
	Uint8  = UnsignedIntType(1)
	Uint16 = UnsignedIntType(2)
	Uint32 = UnsignedIntType(4)
	Uint64 = UnsignedIntType(8)
)

type SignedIntType int // in bytes

func (SignedIntType) isTypeExpression() {}

func (thisInt SignedIntType) Equals(other Type) bool {
	otherInt, ok := other.(SignedIntType)
	if !ok {
		return false
	}

	return thisInt == otherInt
}

func (t SignedIntType) Alignment() int {
	return int(t)
}

func (t SignedIntType) Size() int {
	return int(t)
}

const (
	Int8  = SignedIntType(1)
	Int16 = SignedIntType(2)
	Int32 = SignedIntType(4)
	Int64 = SignedIntType(8)
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

type PointerType struct {
	ValueType Type
}

func (PointerType) isTypeExpression() {}

func (thisPtr PointerType) Equals(other Type) bool {
	otherPtr, ok := other.(PointerType)
	if !ok {
		return false
	}

	return thisPtr.ValueType.Equals(otherPtr.ValueType)
}

func (PointerType) Alignment() int {
	return pointerSize
}

func (PointerType) Size() int {
	return pointerSize
}

type CallConventionKind string

type FunctionType struct {
	CallConventionKind
	ParameterTypes []Type
	ReturnType     Type // nil for no return value

	// TODO compute register constraints / stack layout
	// NOTE: (internal use only) This is not part of the type signature.
	//
	// When the call stack layout is used by the caller, the layout offsets are
	// relative to the stack pointer.  When the call stack layout is used by the
	// callee, the layout offsets are relative to the return address (bottom of
	// the stack frame).
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
	return pointerSize
}

func (*FunctionType) Size() int {
	return pointerSize
}

type ArrayType struct {
	NumElements int
	ElementType Type
}

func (ArrayType) isTypeExpression() {}

func (thisArray ArrayType) Equals(other Type) bool {
	otherArray, ok := other.(ArrayType)
	if !ok {
		return false
	}

	return thisArray.NumElements == otherArray.NumElements &&
		thisArray.ElementType.Equals(otherArray.ElementType)
}

func (t ArrayType) Alignment() int {
	return t.ElementType.Alignment()
}

func (t ArrayType) Size() int {
	return t.NumElements * t.ElementType.Size()
}

type Field struct {
	Name string
	Type Type

	// NOTE: (internal use only) This is not part of the type signature.
	ComputedOffset int // relative to the beginning of the record
}

type StructType struct {
	Fields []*Field

	// When true, fields are not alignment padded.  When false, fields are
	// alignment padded, following c struct convention.
	IsPacked bool

	// NOTE: (internal use only) This is not part of the type signature.
	ComputedAlignment int
	ComputedSize      int
}

func (*StructType) isTypeExpression() {}

func (thisStruct *StructType) Equals(other Type) bool {
	otherStruct, ok := other.(*StructType)
	if !ok {
		return false
	}

	if thisStruct.IsPacked != otherStruct.IsPacked ||
		len(thisStruct.Fields) != len(otherStruct.Fields) {
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
	if t.ComputedAlignment == 0 {
		t.computeSizeAndAlignment()
	}
	return t.ComputedAlignment
}

func (t *StructType) Size() int {
	if t.ComputedSize == 0 {
		t.computeSizeAndAlignment()
	}
	return t.ComputedSize
}

func (t *StructType) computeSizeAndAlignment() {
	t.ComputedAlignment = 0
	t.ComputedSize = 0
	for _, field := range t.Fields {
		alignment := field.Type.Alignment()
		if alignment > t.ComputedAlignment {
			t.ComputedAlignment = alignment
		}

		if !t.IsPacked {
			t.addAlignmentPadding(alignment)
		}

		field.ComputedOffset = t.ComputedSize
		t.ComputedSize += field.Type.Size()
	}

	if !t.IsPacked {
		t.addAlignmentPadding(t.ComputedAlignment)

		// NOTE: we'll round up the size to the nearest power of 2 that fits within
		// a general register since memory operations typically operate in power of
		// 2 units.
		lastChunkSize := t.ComputedSize % generalRegisterSize
		if lastChunkSize > 0 {
			roundUpSize := 1
			for roundUpSize < lastChunkSize {
				roundUpSize <<= 1
			}

			t.ComputedSize += (roundUpSize - lastChunkSize)
		}
	}
}

func (t *StructType) addAlignmentPadding(alignment int) {
	mod := t.ComputedSize % alignment
	if mod > 0 {
		t.ComputedSize += alignment - mod
	}
}
