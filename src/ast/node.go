package ast

import (
	"fmt"
	"strings"

	"github.com/pattyshack/gt/parseutil"
)

type Node interface {
	parseutil.Locatable
	Walk(Visitor)
}

type Visitor interface {
	Enter(Node)
	Exit(Node)
}

type Validator interface {
	Validate(*parseutil.Emitter)
}

type Line interface { // used only by the parser
	IsLine()
}

type SourceEntry interface {
	Node
	Line
	Validator
	isSourceEntry()

	// Internal

	HasDeclarationSyntaxError() bool
	SetHasDeclarationSyntaxError(bool)

	Type() Type
}

type sourceEntry struct {
	hasDeclarationSyntaxError bool
}

func (sourceEntry) IsLine()        {}
func (sourceEntry) isSourceEntry() {}

func (entry *sourceEntry) HasDeclarationSyntaxError() bool {
	return entry.hasDeclarationSyntaxError
}

func (entry *sourceEntry) SetHasDeclarationSyntaxError(val bool) {
	entry.hasDeclarationSyntaxError = val
}

// %-prefixed local variable definition.  Note that the '%' prefix is not part
// of the name and is only used by the parser.
type VariableDefinition struct {
	parseutil.StartEndPos

	Name string // require

	Type Type // optional. Type is check/inferred during type checking

	// Internal (set during ssa construction)
	ParentInstruction Instruction // nil for phi and func parameters
	DefUses           map[*VariableReference]struct{}
}

var _ Node = &VariableDefinition{}
var _ Validator = &VariableDefinition{}

func (def *VariableDefinition) ReplaceReferencesWith(value Value) {
	for ref, _ := range def.DefUses {
		ref.ReplaceWith(value)
	}
}

func (def *VariableDefinition) Walk(visitor Visitor) {
	visitor.Enter(def)
	if def.Type != nil {
		def.Type.Walk(visitor)
	}
	visitor.Exit(def)
}

func (def *VariableDefinition) Validate(emitter *parseutil.Emitter) {
	if def.Name == "" {
		emitter.Emit(def.Loc(), "empty variable definition name")
	}

	if strings.HasPrefix(def.Name, "%") {
		emitter.Emit(def.Loc(), "%%-prefixed name is reserved for internal use")
	}

	if def.Type != nil {
		validateUsableType(def.Type, emitter)
	}
}

func (def *VariableDefinition) AddRef(ref *VariableReference) {
	if def.DefUses == nil {
		def.DefUses = map[*VariableReference]struct{}{}
	}

	def.DefUses[ref] = struct{}{}
	ref.UseDef = def
}

func (def *VariableDefinition) NewRef(
	pos parseutil.StartEndPos,
) *VariableReference {
	ref := &VariableReference{
		StartEndPos: pos,
		Name:        def.Name,
	}
	def.AddRef(ref)
	return ref
}

func (def *VariableDefinition) String() string {
	result := "%" + def.Name
	if def.Type != nil {
		result += " " + def.Type.String()
	}
	return result
}

// Local variable, global label, or immediate
type Value interface {
	Node
	isValue()

	// Internal

	// What this reference refers to.  For now:
	// - local variable reference returns a *VariableDefinition
	// - global label reference returns a string
	// - immediate returns an int / float
	Definition() interface{}

	// NOTE: A copy of newVal, not newVal itself, is used to replace the
	// original value (current object).  The current object is discarded as part
	// of this call.
	ReplaceWith(newVal Value)

	Copy(pos parseutil.StartEndPos) Value
	Discard() // clear graph node references

	SetParentInstruction(Instruction)

	Type() Type

	String() string
}

type value struct {
	// Internal (set during ssa construction)
	ParentInstruction Instruction
}

func (val *value) Discard() {
	val.ParentInstruction = nil
}

func (val *value) SetParentInstruction(ins Instruction) {
	val.ParentInstruction = ins
}

// @-prefixed label for various definitions/declarations.  Note that the '@'
// prefix is not part of the name and is only used by the parser.
type GlobalLabelReference struct {
	value
	parseutil.StartEndPos

	Label string

	// Internal

	Signature SourceEntry
}

var _ Node = &GlobalLabelReference{}
var _ Validator = &GlobalLabelReference{}
var _ Value = &GlobalLabelReference{}

func (GlobalLabelReference) isValue() {}

func (ref *GlobalLabelReference) Definition() interface{} {
	return ref.Label
}

func (ref *GlobalLabelReference) ReplaceWith(newVal Value) {
	newVal = newVal.Copy(ref.StartEnd())
	newVal.SetParentInstruction(ref.ParentInstruction)
	ref.ParentInstruction.replaceSource(ref, newVal)
	ref.Discard()
}

func (ref *GlobalLabelReference) Copy(pos parseutil.StartEndPos) Value {
	copied := *ref
	copied.StartEndPos = pos
	return &copied
}

func (ref *GlobalLabelReference) Walk(visitor Visitor) {
	visitor.Enter(ref)
	visitor.Exit(ref)
}

func (ref *GlobalLabelReference) Validate(emitter *parseutil.Emitter) {
	if ref.Label == "" {
		emitter.Emit(ref.Loc(), "empty global label name")
	}
}

func (ref *GlobalLabelReference) Type() Type {
	if ref.Signature == nil { // failed named binding
		return NewErrorType(ref.StartEndPos)
	}
	return ref.Signature.Type()
}

func (ref *GlobalLabelReference) String() string {
	return "@" + ref.Label
}

// %-prefixed local variable reference.  Note that the '%' prefix is not part
// of the name and is only used by the parser.
type VariableReference struct {
	value
	parseutil.StartEndPos

	Name string // require

	// Internal (set during ssa construction)
	UseDef *VariableDefinition
}

var _ Node = &VariableReference{}
var _ Validator = &VariableReference{}
var _ Value = &VariableReference{}

func (VariableReference) isValue() {}

func (ref *VariableReference) Definition() interface{} {
	return ref.UseDef
}

func (ref *VariableReference) ReplaceWith(newVal Value) {
	newVal = newVal.Copy(ref.StartEnd())
	newVal.SetParentInstruction(ref.ParentInstruction)
	ref.ParentInstruction.replaceSource(ref, newVal)
	ref.Discard()
}

func (ref *VariableReference) Copy(pos parseutil.StartEndPos) Value {
	return ref.UseDef.NewRef(pos)
}

func (ref *VariableReference) Discard() {
	delete(ref.UseDef.DefUses, ref)
	ref.UseDef = nil
	ref.ParentInstruction = nil
}

func (ref *VariableReference) Walk(visitor Visitor) {
	visitor.Enter(ref)
	visitor.Exit(ref)
}

func (ref *VariableReference) Validate(emitter *parseutil.Emitter) {
	if ref.Name == "" {
		emitter.Emit(ref.Loc(), "empty variable reference name")
	}

	if strings.HasPrefix(ref.Name, "%") {
		emitter.Emit(ref.Loc(), "%%-prefixed name is reserved for internal use")
	}
}

func (ref *VariableReference) Type() Type {
	if ref.UseDef == nil { // failed named binding
		return NewErrorType(ref.StartEndPos)
	}
	return ref.UseDef.Type
}

func (ref *VariableReference) String() string {
	return "%" + ref.Name
}

type IntImmediate struct {
	value
	parseutil.StartEndPos

	Value      uint64
	IsNegative bool

	// Internal
	BindedType Type // set by type checker
}

var _ Value = &IntImmediate{}

func NewIntImmediate(
	pos parseutil.StartEndPos,
	value uint64,
	isNegative bool,
) *IntImmediate {
	return &IntImmediate{
		StartEndPos: pos,
		Value:       value,
		IsNegative:  isNegative,
	}
}

func (IntImmediate) isValue() {}

func (imm *IntImmediate) Definition() interface{} {
	return imm.Value
}

func (imm *IntImmediate) ReplaceWith(newVal Value) {
	newVal = newVal.Copy(imm.StartEnd())
	newVal.SetParentInstruction(imm.ParentInstruction)
	imm.ParentInstruction.replaceSource(imm, newVal)
	imm.Discard()
}

func (imm *IntImmediate) Copy(pos parseutil.StartEndPos) Value {
	copied := *imm
	copied.StartEndPos = pos
	return &copied
}

func (imm *IntImmediate) Walk(visitor Visitor) {
	visitor.Enter(imm)
	visitor.Exit(imm)
}

func (imm *IntImmediate) Type() Type {
	if imm.BindedType != nil {
		return imm.BindedType
	} else if imm.IsNegative {
		return NewNegativeIntLiteralType(imm.StartEndPos)
	}
	return NewPositiveIntLiteralType(imm.StartEndPos)
}

func (imm *IntImmediate) String() string {
	sign := ""
	if imm.IsNegative {
		sign = "-"
	}
	return fmt.Sprintf("%s%d", sign, imm.Value)
}

type FloatImmediate struct {
	value
	parseutil.StartEndPos

	Value float64

	// Internal
	BindedType Type // set by type checker
}

var _ Value = &FloatImmediate{}

func (FloatImmediate) isValue() {}

func (imm *FloatImmediate) Definition() interface{} {
	return imm.Value
}

func (imm *FloatImmediate) ReplaceWith(newVal Value) {
	newVal = newVal.Copy(imm.StartEnd())
	newVal.SetParentInstruction(imm.ParentInstruction)
	imm.ParentInstruction.replaceSource(imm, newVal)
	imm.Discard()
}

func (imm *FloatImmediate) Copy(pos parseutil.StartEndPos) Value {
	copied := *imm
	copied.StartEndPos = pos
	return &copied
}

func (imm *FloatImmediate) Walk(visitor Visitor) {
	visitor.Enter(imm)
	visitor.Exit(imm)
}

func (imm *FloatImmediate) Type() Type {
	if imm.BindedType != nil {
		return imm.BindedType
	}
	return NewFloatLiteralType(imm.StartEndPos)
}

func (imm *FloatImmediate) String() string {
	return fmt.Sprintf("%g", imm.Value)
}
