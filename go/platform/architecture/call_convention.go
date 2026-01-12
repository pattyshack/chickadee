package architecture

import (
	"github.com/pattyshack/chickadee/ir"
)

type CallConventionComputer interface {
	Compute(*ir.FunctionType) *CallConvention
}

type CallConventions map[ir.CallConventionKind]CallConventionComputer

func (conventions CallConventions) Compute(
	funcType *ir.FunctionType,
) *CallConvention {
	convention, ok := conventions[funcType.CallConventionKind]
	if !ok {
		panic("unsupported call convention kind: " + funcType.CallConventionKind)
	}

	return convention.Compute(funcType)
}

// The value must either be completely in registers or completely on stack.
// RegisterConstraint entries must be in the CallConvention's Registers set.
// StackEntry must be in CallConvention's CallFrameLayout
type ValueMapping struct {
	Registers []*RegisterConstraint
	*StackEntry
}

type ReturnValue struct {
	// When AddressParameter is nil, the return value is returned directly and
	// ReturnMapping directly maps the value to a location.
	//
	// Otherwise, the return value is returned indirectly.  The return value's
	// address is provided by the caller via AddressParameter, a hidden parameter
	// that is not part of the function signature.  The return value is copied
	// to the provided address, and ReturnMapping returns the provided address
	// at the specified location.  The call convention must allocate space on
	// the stack for the return value; however, the compiler may choose to use a
	// different location to reduce copying.
	AddressParameter *RegisterConstraint

	ScratchSpace *StackEntry

	ReturnMapping ValueMapping
}

// NOTE: variadic functions are not supported.  Variabled length arguments must
// be pass in via an array.
type CallConvention struct {
	// One entry for each register in the architecture's register set, used for
	// specifying which registers are caller-saved (Clobbered) and which
	// registers are callee-saved (not Clobbered).  The constraints must be
	// Require rather than AnyGeneral/AnyFloat.
	Registers map[*Register]*RegisterConstraint

	// From top to bottom.  Callee's frame from the caller's point of view.  May
	// include both arguments and return value.  We'll assume the callee is free
	// to clobber these stack entries (The arguments are caller-saved elsewhere
	// if needed).
	CallFrameLayout []*StackEntry
	CallFrameSize   int

	// A scratch (clobbered) register used for holding the function absolute
	// address when the function call is indirect (i.e., not a global symbol
	// reference).
	FunctionAddress *RegisterConstraint

	// Optional hidden parameter that is not part of the function signature.
	//
	// When specified, the compiler will evict the register's content and set
	// the register with the current (caller) frame's base pointer address
	// before invoking the function.
	//
	// Even though the compiler does not make use of the register, the register
	// must be callee-saved to support stack unwinding.
	BasePointer *RegisterConstraint

	Arguments []ValueMapping

	// NOTE: The return value type may not be identical to the function return
	// type.  If the call convention returns the value indirectly, the return
	// value type is an address type.
	ReturnValue
}

func NewCallConvention(
	set RegisterSet,
	calleeSaved []*Register,
) *CallConvention {
	convention := &CallConvention{
		Registers: map[*Register]*RegisterConstraint{},
	}

	for _, register := range calleeSaved {
		convention.addRegister(register, false)
	}

	for _, list := range [][]*Register{set.General, set.Float} {
		for _, register := range list {
			_, ok := convention.Registers[register]
			if ok {
				continue
			}
			convention.addRegister(register, true)
		}
	}

	return convention
}

func (convention *CallConvention) addRegister(
	register *Register,
	clobbered bool,
) {
	_, ok := convention.Registers[register]
	if ok {
		panic("duplicate register: " + register.Name)
	}

	convention.Registers[register] = &RegisterConstraint{
		Clobbered: clobbered,
		Require:   register,
	}
}

func (convention *CallConvention) registersToConstraints(
	registers []*Register,
) []*RegisterConstraint {
	location := []*RegisterConstraint{}
	for _, register := range registers {
		location = append(location, convention.Registers[register])
	}
	return location
}

func (convention *CallConvention) AddStackEntry(t ir.Type) *StackEntry {
	entry := &StackEntry{
		Type: t,
	}
	convention.CallFrameLayout = append(convention.CallFrameLayout, entry)
	return entry
}

func (convention *CallConvention) FinalizeCallFrameLayout() {
	size := 0
	for _, entry := range convention.CallFrameLayout {
		alignment := ir.Alignment(entry.Type.Size())
		mod := size % alignment
		if mod > 0 {
			size += alignment - mod
		}

		entry.Offset = size
		size += entry.Type.Size()
	}

	convention.CallFrameSize = size
}

func (convention *CallConvention) SetFunctionAddressRegister(
	register *Register,
) {
	convention.FunctionAddress = convention.Registers[register]
}

func (convention *CallConvention) SetBasePointer(register *Register) {
	convention.BasePointer = convention.Registers[register]
}

func (convention *CallConvention) AddStackArgument(location *StackEntry) {
	convention.Arguments = append(
		convention.Arguments,
		ValueMapping{
			StackEntry: location,
		})
}

func (convention *CallConvention) AddRegisterArgument(
	registers []*Register,
) {
	convention.Arguments = append(
		convention.Arguments,
		ValueMapping{
			Registers: convention.registersToConstraints(registers),
		})
}

func (convention *CallConvention) SetDirectStackReturnValue(
	location *StackEntry,
) {
	convention.ReturnValue = ReturnValue{
		ReturnMapping: ValueMapping{
			StackEntry: location,
		},
	}
}

func (convention *CallConvention) SetDirectRegisterReturnValue(
	registers []*Register,
) {
	convention.ReturnValue = ReturnValue{
		ReturnMapping: ValueMapping{
			Registers: convention.registersToConstraints(registers),
		},
	}
}

func (convention *CallConvention) SetIndirectReturnValue(
	addressParameter *Register,
	addressReturnValue *Register,
	scratchSpace *StackEntry,
) {
	convention.ReturnValue = ReturnValue{
		AddressParameter: convention.Registers[addressParameter],
		ScratchSpace:     scratchSpace,
		ReturnMapping: ValueMapping{
			Registers: convention.registersToConstraints(
				[]*Register{addressReturnValue}),
		},
	}
}

func (convention *CallConvention) CallConstraints(
	config Config,
	instruction *ir.Definition,
	call *ir.FunctionCall,
) InstructionConstraints {
	constraints := InstructionConstraints{}
	used := map[*RegisterConstraint]struct{}{}

	_, isIndirect := call.Function.(*ir.LocalReference)
	if isIndirect {
		constraints.RegisterSources = append(
			constraints.RegisterSources,
			RegisterMapping{
				RegisterConstraint: convention.FunctionAddress,
				DefinitionChunk:    call.Function.Def().Chunks()[0],
			})
		used[convention.FunctionAddress] = struct{}{}
	}

	if convention.BasePointer != nil {
		constraints.RegisterSources = append(
			constraints.RegisterSources,
			RegisterMapping{
				RegisterConstraint: convention.BasePointer,
				DefinitionChunk: instruction.Block.Function.
					CurrentFramePointer.Chunks()[0],
			})
		used[convention.BasePointer] = struct{}{}
	}

	// Indirect return value
	if convention.ReturnValue.AddressParameter != nil {
		// XXX: For now, we will unconditionally return the value via the
		// allocated scratch space.
		//
		// As an optimization, we can use the destination's stack entry location
		// directly if we can ensure that source and destination do not interfere
		// with each other.
		constraints.RegisterSources = append(
			constraints.RegisterSources,
			RegisterMapping{
				RegisterConstraint: convention.ReturnValue.AddressParameter,
				TempStackLocation:  convention.ReturnValue.ScratchSpace,
			})
		used[convention.ReturnValue.AddressParameter] = struct{}{}
	}

	for argIdx, argument := range call.Arguments {
		mapping := convention.Arguments[argIdx]
		if mapping.StackEntry != nil {
			constraints.StackSources = append(
				constraints.StackSources,
				StackEntryMapping{
					StackEntry: mapping.StackEntry,
					Definition: argument.Def(),
				})
		} else {
			for chunkIdx, chunk := range argument.Def().Chunks() {
				register := mapping.Registers[chunkIdx]
				constraints.RegisterSources = append(
					constraints.RegisterSources,
					RegisterMapping{
						RegisterConstraint: register,
						DefinitionChunk:    chunk,
					})
				used[register] = struct{}{}
			}
		}
	}

	// Evict unused caller-saved registers
	for _, register := range config.Registers.Data {
		registerConstraint := convention.Registers[register]
		if !registerConstraint.Clobbered {
			continue
		}

		_, ok := used[registerConstraint]
		if ok {
			continue
		}

		constraints.RegisterSources = append(
			constraints.RegisterSources,
			RegisterMapping{
				RegisterConstraint: registerConstraint,
			})
	}

	if convention.ReturnValue.AddressParameter == nil { // Direct return value
		if convention.ReturnValue.ReturnMapping.StackEntry != nil {
			constraints.StackDestination = &StackEntryMapping{
				StackEntry: convention.ReturnValue.ReturnMapping.StackEntry,
				Definition: instruction,
			}
		} else {
			for chunkIdx, chunk := range instruction.Chunks() {
				constraints.RegisterDestinations = append(
					constraints.RegisterDestinations,
					RegisterMapping{
						RegisterConstraint: convention.ReturnValue.ReturnMapping.
							Registers[chunkIdx],
						DefinitionChunk: chunk,
					})
			}
		}
	} else { // Indirect return value
		// NOTE: The value returned by register is the scratch space's address.
		// We'll simply ignore it.
		constraints.StackDestination = &StackEntryMapping{
			StackEntry: convention.ReturnValue.ScratchSpace,
			Definition: instruction,
		}
	}

	return constraints
}
