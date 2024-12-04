package amd64

import (
	"github.com/pattyshack/chickadee/platform"
)

var (
	// Unconditional jump has no constraints
	jumpConstraints = platform.NewInstructionConstraints()

	intConditionalJumpConstraints = newConditionalJumpConstraints(
		ArchitectureRegisters.General)
	floatConditionalJumpConstraints = newConditionalJumpConstraints(
		ArchitectureRegisters.Float)

	intAssignOpConstraint = newAssignOpConstraints(
		ArchitectureRegisters.General)
	floatAssignOpConstraint = newAssignOpConstraints(ArchitectureRegisters.Float)

	intUnaryOpConstraints   = newUnaryOpConstraints(ArchitectureRegisters.General)
	floatUnaryOpConstraints = newUnaryOpConstraints(ArchitectureRegisters.Float)

	intBinaryOpConstraints = newBinaryOpConstraints(
		ArchitectureRegisters.General)
	floatBinaryOpConstraints = newBinaryOpConstraints(ArchitectureRegisters.Float)

	// TODO func call / ret constraints
)

func newConditionalJumpConstraints(
	candidates []*platform.Register,
) *platform.InstructionConstraints {
	constraints := platform.NewInstructionConstraints()

	// Conditional jump compare two source registers without clobbering them.
	// There's no destination register.
	constraints.AddRegisterSource(constraints.Select(false, candidates...))
	constraints.AddRegisterSource(constraints.Select(false, candidates...))

	return constraints
}

func newAssignOpConstraints(
	candidates []*platform.Register,
) *platform.InstructionConstraints {
	constraints := platform.NewInstructionConstraints()

	// Copy from source to destination without clobbering the source register.
	constraints.AddRegisterSource(constraints.Select(false, candidates...))
	constraints.SetRegisterDestination(constraints.Select(true, candidates...))

	return constraints
}

func newUnaryOpConstraints(
	candidates []*platform.Register,
) *platform.InstructionConstraints {
	constraints := platform.NewInstructionConstraints()

	// Destination reuses the source register.
	reg := constraints.Select(true, candidates...)
	constraints.AddRegisterSource(reg)
	constraints.SetRegisterDestination(reg)

	return constraints
}

func newBinaryOpConstraints(
	candidates []*platform.Register,
) *platform.InstructionConstraints {
	constraints := platform.NewInstructionConstraints()

	// Destination reuses the first source register, the second source register is
	// not clobbered.
	src1 := constraints.Select(true, candidates...)
	constraints.AddRegisterSource(src1)
	constraints.SetRegisterDestination(src1)
	constraints.AddRegisterSource(constraints.Select(false, candidates...))

	return constraints
}