package ir

type Reference interface {
	isReference()

	Type() Type
}

// Function reference
type FunctionReference struct {
	Name string

	// Internal
	FuncType *FunctionType
}

func (*FunctionReference) isReference() {}

func (ref *FunctionReference) Type() Type {
	return ref.FuncType
}

// Global variable reference
type AddressReference struct {
	Name string

	// Internal
	ValueType Type
}

func (*AddressReference) isReference() {}

func (ref *AddressReference) Type() Type {
	return AddressType{
		ValueType: ref.ValueType,
	}
}

// Global constant reference
type ConstantReference struct {
	Name string

	// Internal
	ConstantType Type
	Content      []byte
}

func (*ConstantReference) isReference() {}

func (ref *ConstantReference) Type() Type {
	return ref.ConstantType
}

// Local variable reference
type LocalReference struct {
	Name string

	// Internal
	UseDef *LocalDefinition
}

func (*LocalReference) isReference() {}

func (ref *LocalReference) Type() Type {
	return ref.UseDef.Type
}

// Local immediate
type Immediate struct {
	ImmediateType Type
	Reference     []byte
}

func (*Immediate) isReference() {}

func (imm *Immediate) Type() Type {
	return imm.ImmediateType
}

type ElementIndex struct {
	IsArrayIndex bool // false indicate field index

	Field string
	Index int
}

// (Similar to llvm's getelementptr, but operates on all data types) This
// represents data chunk / address offset calculation that will be used for
// accessing/modifying the referenced value.
type Value struct {
	Reference Reference

	// Empty list indicate the full value is used.
	SubElement []ElementIndex

	// Only applicable when reference's type is address type.  Access/modify the
	// dereference subelement value rather than the address itself.
	IndirectAccess bool
}
