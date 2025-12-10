package architecture

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/layout"
)

type RegisterConstraint struct {
	Clobbered bool

	AnyGeneral bool
	AnyFloat   bool

	Require *Register
}

// When a source register (constraint) is mapped to a definition chunk, the
// register may or may not be clobbered.
//
// When a source register (constraint) is not mapped to any definition chunk,
// the register is a scratch register and it must be clobbered.
//
// Destination register (constraint) are always clobbered.
//
// Note that for call conventions, caller-saved registers that aren't used
// for passing arguments are represented as scratch registers (callee-saved
// registers that aren't used for passing arguments are specified in
// SourceRegisters).
//
// TODO: create a pseudo definition for call's frame pointer materialization
type RegisterMapping struct {
	*RegisterConstraint

	*ir.DefinitionChunk
}

type StackEntryMapping struct {
	*StackEntry

	*ir.Definition
}

// The selected instruction's input/output data placement constraints.
type InstructionConstraints struct {
	// The number of input registers used by this instruction.  For the return
	// instruction, this includes all callee-saved registers that must be
	// restored before returning.
	RegisterSources []RegisterMapping

	// The number of stack entries used by this instruction (The stack layout is
	// determined by the stack entries' offsets).  For the call instruction,
	// stack sources are used for passing stack arguments.  For the return
	// instruction, stack sources (containing at most one element) is used for
	// returning value on stack.
	StackSources []StackEntryMapping

	// The number of output registers used by this instruction.  All destination
	// register constraints must be clobbered.  NOTE: source and destination
	// register mapping that shares the same register constraint object will map
	// both source and destination to the same register.
	RegisterDestinations []RegisterMapping

	// Only used by call instruction.  This is nil if the destination value is
	// on registers.
	StackDestination *StackEntryMapping
}

// The selected machine architecture instruction(s) used to implement the
// ir.Instruction.  Instruction constraints are determined as part of the
// instruction selection process (Registers are selected afterward based on
// the constraints).
type MachineInstruction interface {
	Instruction() ir.Instruction

	Constraints() InstructionConstraints

	EmitTo(
		builder *layout.SegmentBuilder,
		selectedRegisters map[*RegisterConstraint]*Register)
}
