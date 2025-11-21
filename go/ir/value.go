package ir

type Value interface {
	isValue()

	Type() Type
}

// Function reference
type FunctionReference struct {
	Name string

	// Internal
	FuncType *FunctionType
}

func (*FunctionReference) isValue() {}

func (ref *FunctionReference) Type() Type {
	return ref.FuncType
}

// Global variable reference
type AddressReference struct {
	Name string

	// Internal
	ValueType Type
}

func (*AddressReference) isValue() {}

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

func (*ConstantReference) isValue() {}

func (ref *ConstantReference) Type() Type {
	return ref.ConstantType
}

// Local variable reference
type LocalReference struct {
	Name string

	// Internal
	UseDef *LocalDefinition
}

func (*LocalReference) isValue() {}

func (ref *LocalReference) Type() Type {
	return ref.UseDef.Type
}

// Local immediate
type Immediate struct {
	ImmediateType Type
	Value         []byte
}

func (*Immediate) isValue() {}

func (imm *Immediate) Type() Type {
	return imm.ImmediateType
}
