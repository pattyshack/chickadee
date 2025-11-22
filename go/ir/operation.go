package ir

// Initialize value of ValueType with zeros.  When AllocatedOnStack is true,
// the value is allocated on stack and the value's address is assigned to
// destination.  When AllocatedOnStack is false, the zero value is assigned
// to the destination.
type InitializeOperation struct {
	instruction

	Dest *LocalDefinition

	AllocateOnStack bool
	ValueType       Type
}

type UnaryOperationKind string

const (
	Copy = UnaryOperationKind("copy")

	Neg = UnaryOperationKind("neg")
)

type UnaryOperation struct {
	instruction

	Kind UnaryOperationKind

	Dest *LocalDefinition
	Src  *Value
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
	instruction

	Kind BinaryOperationKind

	Dest *LocalDefinition
	Src1 *Value
	Src2 *Value
}

type FuncCallKind string

const (
	Call = FuncCallKind("call")
)

type FuncCall struct {
	instruction

	Kind FuncCallKind

	Dest *LocalDefinition // nil if the function doesn't return any value
	Func *Value
	Args []*Value
}
