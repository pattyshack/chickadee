package ir

type Operation interface {
	isOperation()
}

type operation struct{}

func (operation) isOperation() {}

// Initialize value of ValueType with zeros.  When AllocatedOnStack is true,
// the value is allocated on stack and the value's address is assigned to
// destination.  When AllocatedOnStack is false, the zero value is assigned
// to the destination.
type InitializeOperation struct {
	operation

	AllocateOnStack bool
	ValueType       Type
}

type UnaryOperationKind string

const (
	Neg = UnaryOperationKind("neg")
	Not = UnaryOperationKind("not")

	ToInt8  = UnaryOperationKind("toInt8")
	ToInt16 = UnaryOperationKind("toInt16")
	ToInt32 = UnaryOperationKind("toInt32")
	ToInt64 = UnaryOperationKind("toInt64")

	ToUint8  = UnaryOperationKind("toUint8")
	ToUint16 = UnaryOperationKind("toUint16")
	ToUint32 = UnaryOperationKind("toUint32")
	ToUint64 = UnaryOperationKind("toUint64")

	ToFloat32 = UnaryOperationKind("toFloat32")
	ToFloat64 = UnaryOperationKind("toFloat64")
)

type UnaryOperation struct {
	operation

	Kind UnaryOperationKind
	Src  Value
}

type BinaryOperationKind string

/*
TODO rethink get / set element

// (Similar to llvm's getelementptr, but operates on all data types) This
// represents data chunk / address offset calculation that will be used for
// accessing/modifying the referenced value.
type SubValue struct {
	Reference Value

	// Index into either array or struct field.  Empty list indicate the full
	// value is used.
	SubElementIndices []int

	// Only applicable when reference's type is address type.  Access/modify
	// the dereference subelement value rather than the address itself.
	IndirectAccess bool
}

// Create a new value out of the first source, modified by the second source.
// The first definition is destroyed as part of the process (the same name
// could be reuse for a new definition).
SetElement = BinaryOperationKind("set element")
*/

const (
	Add = BinaryOperationKind("add")
	Mul = BinaryOperationKind("mul")

	Sub = BinaryOperationKind("sub")
	Div = BinaryOperationKind("div")
	Rem = BinaryOperationKind("rem")

	Shl = BinaryOperationKind("shl")
	Shr = BinaryOperationKind("shr")

	And = BinaryOperationKind("and")
	Or  = BinaryOperationKind("or")
	Xor = BinaryOperationKind("xor")
)

type BinaryOperation struct {
	operation

	Kind BinaryOperationKind

	Src1 Value
	Src2 Value
}

type FunctionCallKind string

const (
	Call = FunctionCallKind("call")
)

// NOTE: function always return a value.  Use empty struct for void
type FunctionCall struct {
	operation

	Kind      FunctionCallKind
	Function  Value
	Arguments []Value
}
