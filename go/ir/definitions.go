package ir

type DeclarationKind string

const (
	// The function is populated into .text.  Functions can access the function
	// using FunctionReference.
	FunctionDeclaration = DeclarationKind("function declaration")

	// The global variable's content is populated into .data (or .bss).
	// Functions can access the global variable using AddressReference which
	// exposes the object via an address indirection.
	VariableDeclaration = DeclarationKind("variable declaration")

	// The global constant's content is populated into .rodata.  Functions can
	// access the constant using ConstantReference which directly exposes the
	// constant's value (The compiler convert constant references to immediates
	// during compilation)
	//
	// XXX: maybe skip writing to .rodata entirely?
	ConstantDeclaration = DeclarationKind("constant declaration")
)

// Global function/variable/constant declaration.
type GlobalDeclaration struct {
	Kind DeclarationKind
	Name string
	Type Type
}

type Definition interface {
	isDefinition()
}

type FunctionDefinition struct {
	// NOTE: The declaration type must be FunctionType.
	GlobalDeclaration

	// 1-to-1 mapping between parameter types and names
	ParameterNames []string

	Blocks []*Block
}

func (*FunctionDefinition) isDefinition() {}

// Global variable/constant definition.
type ObjectDefinition struct {
	// NOTE: The declaration type cannot not be FunctionType.
	GlobalDeclaration

	// Content must either be nil or the same length as the Type's size.  If
	// Content is nil, we'll default to all zero bytes.
	Content []byte
}

func (*ObjectDefinition) isDefinition() {}

// Local variable definition
type LocalDefinition struct {
	Name string
	Type Type

	DefUses map[*LocalReference]struct{}
}

// Logical compilation unit that forms a single object file.
type CompileUnit struct {
	Definitions []Definition
}
