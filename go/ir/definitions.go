package ir

type FunctionDefinition struct {
	Name string
	// NOTE: The declaration type must be FunctionType.
	Type Type

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

// Local variable definition
type LocalDefinition struct {
	Name string
	Type Type // optional.  Type is checked/inferred during type checking

	// Internal
	Instruction Instruction
	DefUses     map[*LocalReference]struct{}
}
