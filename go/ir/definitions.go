package ir

const (
	// All internal definition names are prefixed by "%"
	PreviousFramePointer = "%previous-frame-pointer"
	CurrentFramePointer  = "%current-frame-pointer"
	ReturnAddress        = "%return-address"
	ReturnValue          = "%return-value"
)

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

	// Internal

	// NOTE: The following definitions span the entire function life time.
	// Function parameters are not tracked at function level since their
	// lifespans could be shorter than the function.

	CalleeSavedRegisters []*Definition

	ReturnValue *Definition

	ReturnAddress *Definition

	// NOTE: PreviousFramePointer is always defined even if the call convention
	// does not provide a value, in which case, the value is zero.
	PreviousFramePointer *Definition
	CurrentFramePointer  *Definition
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

// NOTE: A Definition acts as an instruction statement, a definition for local
// value, or a pseudo definition for non-local values.
//
// Grouping the operation with the definition simplifies constant propagation,
// value re-materialization, etc.
type Definition struct {
	instruction

	// Side effect only operations and pseudo definitions may use empty string
	// name to indicate the value is inaccessible.
	Name string
	Type

	Operation // nil iff this is a function parameter pseudo definition

	// Internal

	// Used for defining function parameters, callee-saved registers,
	// return value, and non-local reference values.
	IsPseudoDefinition bool

	chunks []*DefinitionChunk

	DefUse map[*LocalReference]struct{}
}

func (def *Definition) Chunks() []*DefinitionChunk {
	if def.chunks != nil {
		return def.chunks
	}

	typeChunks := def.Type.Chunks()
	chunks := make([]*DefinitionChunk, len(typeChunks))
	for idx, typeChunk := range typeChunks {
		chunks[idx] = &DefinitionChunk{
			Definition: def,
			TypeChunk:  typeChunk,
		}
	}

	def.chunks = chunks
	return def.chunks
}

// For internal use only.
//
// Each value is partition into chunks that could fit into 8-byte registers.
// Each copy of a chunk is either completely on memory or completely in
// register.
type DefinitionChunk struct {
	*Definition

	*TypeChunk
}
