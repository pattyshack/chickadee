package ir

type Value interface {
	Operation // Copy/assignment operation

	isValue()

	Type() Type
}

// Function definition reference
type GlobalFunctionReference struct {
	operation

	Name string

	// Internal
	FuncType *FunctionType
}

func (*GlobalFunctionReference) isValue() {}

func (ref *GlobalFunctionReference) Type() Type {
	return ref.FuncType
}

// Global object reference
type GlobalObjectReference struct {
	operation

	Name string

	// Internal
	ValueType Type
}

func (*GlobalObjectReference) isValue() {}

func (ref *GlobalObjectReference) Type() Type {
	return AddressType{
		ValueType: ref.ValueType,
	}
}

// Global constant reference
type GlobalConstantReference struct {
	operation

	Name string

	// Internal
	ConstantType Type
	Content      []byte
}

func (*GlobalConstantReference) isValue() {}

func (ref *GlobalConstantReference) Type() Type {
	return ref.ConstantType
}

// Local variable reference
type LocalReference struct {
	operation

	Name string

	// Internal
	UseDef *Definition
}

func (*LocalReference) isValue() {}

func (ref *LocalReference) Type() Type {
	return ref.UseDef.Type
}

// Local immediate
type Immediate struct {
	operation

	ImmediateType Type
	Bytes         []byte
}

func (*Immediate) isValue() {}

func (imm *Immediate) Type() Type {
	return imm.ImmediateType
}

/*
// TODO: Rethink this. maybe be better to decompose this into operations?
//
// (Similar to llvm's getelementptr, but operates on all data types) This
// represents data chunk / address offset calculation that will be used for
// accessing/modifying the referenced value.
type SubValue struct {
	Reference Value

	// Index into either array or struct field.  Empty list indicate the full
	// value is used.
	SubElementIndices []int

	// Only applicable when reference's type is address type.  Access/modify the
	// dereference subelement value rather than the address itself.
	IndirectAccess bool
}

func (*SubValue) isValue() {}

func (imm *SubValue) Type() Type {
	panic("TODO")
}
*/
