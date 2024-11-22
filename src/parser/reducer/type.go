package reducer

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/chickadee/ast"
	"github.com/pattyshack/chickadee/parser/lr"
)

func (Reducer) ToNumberType(
	token *lr.TokenValue,
) (
	ast.Type,
	error,
) {
	return ast.NumberType{
		StartEndPos: token.StartEndPos,
		Kind:        ast.NumberTypeKind(token.Value),
	}, nil
}

func (Reducer) ToFuncType(
	funcKW *lr.TokenValue,
	lparen *lr.TokenValue,
	parameterTypes []ast.Type,
	rparen *lr.TokenValue,
	retType ast.Type,
) (
	ast.Type,
	error,
) {
	return ast.FunctionType{
		StartEndPos:    parseutil.NewStartEndPos(funcKW.Loc(), retType.End()),
		ParameterTypes: parameterTypes,
		ReturnType:     retType,
	}, nil
}
