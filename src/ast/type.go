package ast

import (
	"github.com/pattyshack/gt/parseutil"
)

type NumberTypeKind string

const (
	I8  = NumberTypeKind("I8")
	I16 = NumberTypeKind("I16")
	I32 = NumberTypeKind("I32")
	I64 = NumberTypeKind("I64")

	U8  = NumberTypeKind("U8")
	U16 = NumberTypeKind("U16")
	U32 = NumberTypeKind("U32")
	U64 = NumberTypeKind("U64")

	F32 = NumberTypeKind("F32")
	F64 = NumberTypeKind("F64")
)

type NumberType struct {
	isType
	parseutil.StartEndPos

	Kind NumberTypeKind
}

var _ Type = NumberType{}
var _ Validator = NumberType{}

func (numType NumberType) Walk(visitor Visitor) {
	visitor.Enter(numType)
	visitor.Exit(numType)
}

func (numType NumberType) Validate(emitter *parseutil.Emitter) {
	switch numType.Kind {
	case I8, I16, I32, I64, U8, U16, U32, U64, F32, F64: // ok
	default:
		emitter.Emit(numType.Loc(), "unexpected number type (%s)", numType.Kind)
	}
}

type FunctionType struct {
	isType
	parseutil.StartEndPos

	ReturnType     Type
	ParameterTypes []Type
}

var _ Type = FunctionType{}

func (funcType FunctionType) Walk(visitor Visitor) {
	visitor.Enter(funcType)
	funcType.ReturnType.Walk(visitor)
	for _, param := range funcType.ParameterTypes {
		param.Walk(visitor)
	}
	visitor.Exit(funcType)
}
