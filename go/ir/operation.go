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
	Copy = UnaryOperationKind("copy")

	Neg = UnaryOperationKind("neg")
)

type UnaryOperation struct {
	operation

	Kind UnaryOperationKind
	Src  Value
}

type BinaryOperationKind string

const (
	// Create a new value out of the first source, modified by the second source.
	// The first definition is destroyed as part of the process (the same name
	// could be reuse for a new definition).
	SetElement = BinaryOperationKind("set element")

	Add = BinaryOperationKind("add")
)

type BinaryOperation struct {
	Operation

	Kind BinaryOperationKind

	Src1 Value
	Src2 Value
}

type FuncCallKind string

const (
	Call = FuncCallKind("call")
)

// NOTE: function always return a value.  Use empty struct for void
type FuncCall struct {
	operation

	Kind FuncCallKind
	Func Value
	Args []Value
}
