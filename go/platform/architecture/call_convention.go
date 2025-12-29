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

// NOTE: variadic functions are not supported.  Variabled length arguments must
// be pass in via an array.
type CallConvention struct {
	// One entry for each register in the architecture's register set, used for
	// specifying which registers are caller-saved (Clobbered) and which
	// registers are callee-saved (not Clobbered).  The constraints must be
	// Require rather than AnyGeneral/AnyFloat.
	Registers map[string]*RegisterConstraint

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

	// Optional callee-saved (not clobbered) register / hidden parameter.
	//
	// When specified, the compiler will evict the register's content and set
	// the register with the current (caller) frame's base pointer address
	// before invoking the function.
	//
	// Even though the compiler does not make use the register, the register
	// must be callee-saved to support stack unwinding.
	BasePointer *RegisterConstraint

	Arguments []ValueMapping

	ReturnValue ValueMapping
}

func NewCallConvention(
	functionAddress *Register,
	basePointer *Register, // could be nil
) *CallConvention {
	call := &CallConvention{
		Registers: map[string]*RegisterConstraint{},
	}

	call.FunctionAddress = call.CallerSaved(functionAddress)

	if basePointer != nil {
		call.BasePointer = call.CalleeSaved(basePointer)
	}

	return call
}

func (call *CallConvention) maybeAddRegister(
	register *Register,
	clobbered bool,
) *RegisterConstraint {
	constraint, ok := call.Registers[register.Name]
	if !ok {
		constraint = &RegisterConstraint{
			Clobbered: clobbered,
			Require:   register,
		}
		call.Registers[register.Name] = constraint
	}

	if clobbered != constraint.Clobbered {
		panic("should never happen")
	}

	return constraint
}

func (call *CallConvention) CallerSaved(
	register *Register,
) *RegisterConstraint {
	return call.maybeAddRegister(register, true)
}

func (call *CallConvention) CalleeSaved(
	register *Register,
) *RegisterConstraint {
	return call.maybeAddRegister(register, false)
}

func (call *CallConvention) AddStackEntry(t ir.Type) *StackEntry {
	entry := &StackEntry{
		Type: t,
	}
	call.CallFrameLayout = append(call.CallFrameLayout, entry)
	return entry
}

func (call *CallConvention) FinalizeCallFrameLayout() {
	offset := 0
	for _, entry := range call.CallFrameLayout {
		entry.Offset = offset
		offset += entry.Type.Size()
	}

	call.CallFrameSize = offset
}

func (call *CallConvention) AddStackArgument(location *StackEntry) {
	call.Arguments = append(
		call.Arguments,
		ValueMapping{
			StackEntry: location,
		})
}

func (call *CallConvention) AddRegisterArgument(
	location []*RegisterConstraint,
) {
	call.Arguments = append(
		call.Arguments,
		ValueMapping{
			Registers: location,
		})
}

func (call *CallConvention) SetStackReturnValue(location *StackEntry) {
	call.ReturnValue = ValueMapping{
		StackEntry: location,
	}
}

func (call *CallConvention) SetRegisterReturnValue(
	location []*RegisterConstraint,
) {
	call.ReturnValue = ValueMapping{
		Registers: location,
	}
}
