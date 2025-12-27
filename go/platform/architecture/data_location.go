package architecture

import (
	"github.com/pattyshack/chickadee/ir"
)

// Stack space allocated for a data location.  Unlike data chunk / location,
// the stack entry may be uninitialized / invalid.
type StackEntry struct {
	ir.Type

	// The offset is relative to the top of the current stack frame
	Offset int
}

type ChunkLocation struct {
	*ir.DefinitionChunk

	*Register   // Location of this particular chunk
	*StackEntry // Location of the entire value rather than individual chunks
}

type ValueLocation struct {
	// NOTE: there is at least one copy of each chunk at all times while the
	// value is alive.
	Chunks map[*ir.DefinitionChunk][]ChunkLocation
}
