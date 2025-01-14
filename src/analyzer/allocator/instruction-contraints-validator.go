package allocator

import (
	"fmt"

	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/chickadee/analyzer/util"
	"github.com/pattyshack/chickadee/architecture"
	"github.com/pattyshack/chickadee/ast"
	"github.com/pattyshack/chickadee/platform"
)

type InstructionConstraintsValidator struct {
	platform.Platform
	*Allocator
}

func ValidateInstructionConstraints(
	targetPlatform platform.Platform,
	allocator *Allocator,
) util.Pass[ast.SourceEntry] {
	return &InstructionConstraintsValidator{
		Platform:  targetPlatform,
		Allocator: allocator,
	}
}

func (validator *InstructionConstraintsValidator) Process(
	entry ast.SourceEntry,
) {
	funcDef, ok := entry.(*ast.FunctionDefinition)
	if !ok {
		return
	}

	validator.ValidateFunctionDefinition(funcDef)
	for _, block := range validator.BlockStates {
		for _, in := range block.Instructions {
			isFuncCall := false
			constraints := block.Constraints[in]
			switch inst := in.(type) {
			case *ast.FuncCall:
				isFuncCall = true
				validator.ValidateCall(inst, constraints)
			case *ast.Terminal:
				validator.ValidateRet(inst, constraints)
			case *ast.Jump:
				continue
			case ast.Instruction:
				validator.ValidateGenericInstruction(inst, constraints)
			}

			validator.ValidateUniqueRegisters(in.Loc(), constraints, isFuncCall)
		}
	}
}

// A required register may only appear once in the source list (including
// pseudo sources).  The same required register may reappear once in the
// destintation to indicate register reuse.
func (validator *InstructionConstraintsValidator) ValidateUniqueRegisters(
	pos parseutil.Location,
	constraints *architecture.InstructionConstraints,
	isFuncDefOrCall bool,
) {
	if !isFuncDefOrCall && len(constraints.PseudoSources) > 0 {
		panic(fmt.Sprintf("invalid: %s", pos))
	}

	numRegistersNeeded := 0
	uniqueSrcCandidates := map[*architecture.RegisterCandidate]struct{}{}
	uniqueSrcRegisters := map[*architecture.Register]struct{}{}
	for _, loc := range constraints.Sources {
		for _, reg := range loc.Registers {
			numRegistersNeeded++

			_, ok := uniqueSrcCandidates[reg]
			if ok {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
			uniqueSrcCandidates[reg] = struct{}{}

			if reg.Require == nil {
				continue
			}

			_, ok = uniqueSrcRegisters[reg.Require]
			if ok {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
			uniqueSrcRegisters[reg.Require] = struct{}{}
		}
	}

	if constraints.Destination != nil {
		uniqueDestCandidates := map[*architecture.RegisterCandidate]struct{}{}
		uniqueDestRegisters := map[*architecture.Register]struct{}{}
		for _, reg := range constraints.Destination.Registers {
			_, ok := uniqueDestCandidates[reg]
			if ok {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
			uniqueDestCandidates[reg] = struct{}{}

			if reg.Require == nil {
				_, ok := uniqueSrcCandidates[reg]
				if !ok { // a new register not used by any source
					numRegistersNeeded++
				}

				continue
			}

			_, ok = uniqueDestRegisters[reg.Require]
			if ok {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
			uniqueDestRegisters[reg.Require] = struct{}{}

			_, ok = uniqueSrcRegisters[reg.Require]
			if !ok {
				numRegistersNeeded++
			}
		}
	}

	// We'll assume all instructions must leave at least one register unused.
	// The allocator will make use of the unused register for data shuffling.
	// This mostly impact call convention since most real architecture
	// instruction use at most 3 registers.
	totalRegisters := len(validator.Platform.ArchitectureRegisters().Data)
	if numRegistersNeeded > totalRegisters-1 {
		panic(fmt.Sprintf("invalid: %s", pos))
	}

	// call convention must fully specify all registers' clobber behavior
	if isFuncDefOrCall && len(constraints.RequiredRegisters) != totalRegisters {
		panic(fmt.Sprintf("invalid: %s", pos))
	}
}

func (validator *InstructionConstraintsValidator) ValidateFunctionDefinition(
	funcDef *ast.FunctionDefinition,
) {
	constraints := validator.CallConvention(funcDef.FuncType).CallConstraints

	// The fist constraint is call's function location, which is not applicable
	// to the function definition.
	if len(constraints.Sources) != len(funcDef.Parameters)+1 {
		panic(fmt.Sprintf("invalid: %s", funcDef.Loc()))
	}

	for idx, constraint := range constraints.Sources[1:] {
		validator.ValidateLocation(
			funcDef.Parameters[idx].Loc(),
			constraint,
			true,
			true)
	}

	if len(funcDef.PseudoParameters) != len(constraints.PseudoSources) {
		panic(fmt.Sprintf("invalid: %s", funcDef.Loc()))
	}

	// callee-saved registers can't be on stack
	for idx, constraint := range constraints.PseudoSources {
		validator.ValidateLocation(
			funcDef.PseudoParameters[idx].Loc(),
			constraint,
			true,
			false)
	}

	if constraints.Destination == nil {
		panic(fmt.Sprintf("invalid: %s", funcDef.Loc()))
	}

	validator.ValidateLocation(
		funcDef.ReturnType.Loc(),
		constraints.Destination,
		true,
		true)

	validator.ValidateUniqueRegisters(funcDef.Loc(), constraints, true)
}

func (validator *InstructionConstraintsValidator) ValidateGenericInstruction(
	inst ast.Instruction,
	constraints *architecture.InstructionConstraints,
) {
	sources := inst.Sources()
	if len(sources) != len(constraints.Sources) {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	for idx, src := range sources {
		validator.ValidateLocation(
			src.Loc(),
			constraints.Sources[idx],
			false,
			false)
	}

	dest := inst.Destination()
	if dest == nil {
		if constraints.Destination != nil {
			panic(fmt.Sprintf("invalid: %s", inst.Loc()))
		}
	} else {
		if constraints.Destination == nil {
			panic(fmt.Sprintf("invalid: %s", dest.Loc()))
		}
		validator.ValidateLocation(
			dest.Loc(),
			constraints.Destination,
			false,
			false)
	}
}

func (validator *InstructionConstraintsValidator) ValidateCall(
	inst *ast.FuncCall,
	constraints *architecture.InstructionConstraints,
) {
	srcs := inst.Sources()
	if len(constraints.Sources) != len(srcs) {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	for idx, src := range srcs {
		// Function location (first source) must be in register
		canBeOnStack := idx != 0
		validator.ValidateLocation(
			src.Loc(),
			constraints.Sources[idx],
			true,
			canBeOnStack)
	}

	if constraints.Destination == nil {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	dest := inst.Destination()
	if dest == nil {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	validator.ValidateLocation(
		dest.Loc(),
		constraints.Destination,
		true,
		true)
}

func (validator *InstructionConstraintsValidator) ValidateRet(
	inst *ast.Terminal,
	constraints *architecture.InstructionConstraints,
) {
	if constraints.Destination != nil {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	if len(inst.CalleeSavedSources)+1 != len(constraints.Sources) {
		panic(fmt.Sprintf("invalid: %s", inst.Loc()))
	}

	// the return value could be on stack
	validator.ValidateLocation(
		inst.RetVal.Loc(),
		constraints.Sources[0],
		true,
		true)

	// callee-saved sources cannot be on stack
	for idx, constraint := range constraints.Sources[1:] {
		validator.ValidateLocation(
			inst.CalleeSavedSources[idx].Loc(),
			constraint,
			true,
			false)
	}
}

func (validator *InstructionConstraintsValidator) ValidateLocation(
	pos parseutil.Location,
	constraint *architecture.LocationConstraint,
	requireConcreteLocation bool,
	canBeOnStack bool,
) {
	if constraint.AnyLocation {
		if constraint.RequireOnStack {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
		if len(constraint.Registers) > 0 {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
		if requireConcreteLocation {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
	} else if constraint.RequireOnStack {
		if len(constraint.Registers) > 0 {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
		if constraint.Size == 0 {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
		if !canBeOnStack {
			panic(fmt.Sprintf("invalid: %s", pos))
		}
	} else { // register candidates
		clobbered := false
		if len(constraint.Registers) > 0 {
			clobbered = constraint.Registers[0].Clobbered
		}
		for _, reg := range constraint.Registers {
			if reg.Clobbered != clobbered {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
			if reg.AnyGeneral || reg.AnyFloat {
				if reg.Require != nil {
					panic(fmt.Sprintf("invalid: %s", pos))
				}
				if requireConcreteLocation {
					panic(fmt.Sprintf("invalid: %s", pos))
				}
			} else if reg.Require == nil {
				panic(fmt.Sprintf("invalid: %s", pos))
			}
		}
	}
}
