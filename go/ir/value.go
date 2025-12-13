package ir

import (
	"fmt"
)

type Value interface {
	Operation // Copy/assignment operation

	isValue()

	Type() Type

	// For local reference, this returns a real definition of the value.  For
	// other values, this returns a pseudo definition associated with the value.
	Def() *Definition
}

// Global (function/object/constant) definition reference
type GlobalReference struct {
	operation

	Name string

	// Internal

	// REMINDER: deduplicate pseudo definition to reuse definitions / reduce
	// register pressure
	PseudoDefinition *Definition
}

func NewGlobalReference(name string) Value {
	ref := &GlobalReference{
		Name: name,
	}
	return ref
}

func (*GlobalReference) isValue() {}

func (ref *GlobalReference) Type() Type {
	return ref.PseudoDefinition.Type
}

func (ref *GlobalReference) Def() *Definition {
	return ref.PseudoDefinition
}

// Local immediate
type Immediate struct {
	operation

	// int*/uint*/float* for basic types; []byte for array/struct/address types
	Value interface{}

	ImmediateType Type

	// Internal

	// REMINDER: deduplicate pseudo definition to reuse definitions / reduce
	// register pressure
	PseudoDefinition *Definition
}

// int*/uint*/float* immediate
func NewBasicImmediate(value interface{}) Value {
	var t Type
	switch value.(type) {
	case int8:
		t = Int8
	case int16:
		t = Int16
	case int32:
		t = Int32
	case int64:
		t = Int64
	case uint8:
		t = Uint8
	case uint16:
		t = Uint16
	case uint32:
		t = Uint32
	case uint64:
		t = Uint64
	case float32:
		t = Float32
	case float64:
		t = Float64
	default:
		panic(fmt.Sprintf("invalid basic immediate type: %#v", value))
	}

	imm := &Immediate{
		Value:         value,
		ImmediateType: t,
	}

	return imm
}

// array / struct / address immediate
func NewComplexImmediate(immediateType Type, value []byte) Value {
	switch immediateType.(type) {
	case *AddressType:
	case *ArrayType:
	case *StructType:
	default:
		panic(fmt.Sprintf("invalid complex immediate type: %#v", immediateType))
	}

	if immediateType.Size() != len(value) {
		panic(fmt.Sprintf(
			"invalid immediate length (%d != %d)",
			immediateType.Size(),
			len(value)))
	}

	imm := &Immediate{
		Value:         value,
		ImmediateType: immediateType,
	}

	return imm
}

func (*Immediate) isValue() {}

func (imm *Immediate) Type() Type {
	return imm.PseudoDefinition.Type
}

func (imm *Immediate) Def() *Definition {
	return imm.PseudoDefinition
}

// Local variable reference
type LocalReference struct {
	operation

	Name string

	// Internal
	UseDef *Definition
}

func NewLocalReference(name string) Value {
	return &LocalReference{
		Name: name,
	}
}

func (*LocalReference) isValue() {}

func (ref *LocalReference) Type() Type {
	return ref.UseDef.Type
}

func (ref *LocalReference) Def() *Definition {
	return ref.UseDef
}
