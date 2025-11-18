package architecture

import (
	"github.com/pattyshack/chickadee/ir"
)

// Stack space allocated for a data location.  Unlike data chunk / location,
// the stack entry may be uninitialized / invalid.
type StackEntry struct {
	Name string
	ir.Type

	// The offset is relative to the top of the stack
	Offset int
}

// NOTE: data value is partitioned into chunks that could fit into general
// 8-byte registers.  Each copy of a chunk is either completely on memory or
// completely in register.
type DataChunk struct {
	ParentGroup *DataChunkReplicaGroup

	*Register

	// When true, the value is in memory and the register holds the address to
	// the value.  When false and the register is not nil, the register directly
	// holds the data chunk.
	IsIndirect bool

	*StackEntry // address is relative to the stack pointer

	// The offset is relative to the beginning of the type's value.
	Offset int

	// Number of valid bytes in this chunk (e.g., size could be smaller than
	// the register size)
	Size int
}

// Copies of the same data chunk.  The struct indirection simplifies list
// modification.
type DataChunkReplicaGroup struct {
	ParentLocation *DataLocation

	Copies []DataChunk
}

type DataLocation struct {
	Name string
	ir.Type

	// NOTE: there is at least one copy of each chunk at all times while the
	// value is alive.
	Chunks []*DataChunkReplicaGroup
}
