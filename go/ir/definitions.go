package ir

type FunctionDefinition struct {
	Name string
	Type *FunctionType

	// 1-to-1 mapping between parameter types and names
	ParameterNames []string

	Blocks []*Block

	// A potential entry (aka main) function.  The function signature must be
	// "func()"; the function can freely clobber all general/float registers.
	//
	// The compiler will generate a entry point stub in .text which sets up the
	// stack, calls the .init section function, calls the entry function, and
	// finally exit.
	IsEntryFunction bool
}

// Global variable/constant definition.
type ObjectDefinition struct {
	Name string
	// NOTE: The declaration type cannot not be FunctionType.
	Type Type

	// Content must either be nil or the same length as the Type's size.  If
	// Content is nil, we'll default to all zero bytes.
	Content []byte
}

// Logical compilation unit that forms a single object file.
type CompilationUnit struct {
	// The function is populated into .text.  Functions can access the function
	// using FunctionReference.
	FunctionDefinitions []*FunctionDefinition

	// When init function name is not empty, populate a function call into .init.
	// The function signature must be "func()" and require no pre-call setup
	// (except basic stack alignment); the function can freely clobber all
	// general/float registers.
	InitFunction string

	// The global constant's content is populated into .rodata.  Functions can
	// access the constant using ConstantReference which directly exposes the
	// constant's value (The compiler convert constant references to immediates
	// during compilation)
	//
	// XXX: maybe skip writing to .rodata entirely?
	ConstantDefinitions []*ObjectDefinition

	// The global variable's content is populated into .data (or .bss).
	// Functions can access the global variable using AddressReference which
	// exposes the object via an address indirection.
	VariableDefinitions []*ObjectDefinition
}

// NOTE: A local definition is both an instruction statement and a definition.
// Grouping the operation with the definition simplifies constant propagation,
// value re-materialization, etc.
type Definition struct {
	instruction

	// Side effect only operations may use empty string name to indicate the
	// value is inaccessible.
	Name string
	Type

	Operation // nil iff IsFunctionParameter is true

	// Internal

	// Used for defining function parameters.
	IsFunctionParameter bool

	Chunks []*DefinitionChunk

	DefUse map[*LocalReference]struct{}
}

// For internal use only.
//
// Each value is partition into chunks that could fit into 8-byte registers.
// Each copy of a chunk is either completely on memory or completely in
// register.
type DefinitionChunk struct {
	*Definition

	// The offset is relative to the beginning of the definition's value.
	Offset int

	// Number of valid bytes in this chunk (e.g., size could be smaller than
	// the register size)
	Size int
}
