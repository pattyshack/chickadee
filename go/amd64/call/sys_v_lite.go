package call

import (
	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
)

var (
	sysVCalleeSavedRegisters = []*architecture.Register{
		registers.Rbx,
		registers.Rbp,
		registers.R12,
		registers.R13,
		registers.R14,
		registers.R15,
	}

	sysVGeneralParameterRegisters = []*architecture.Register{
		registers.Rdi,
		registers.Rsi,
		registers.Rdx,
		registers.Rcx,
		registers.R8,
		registers.R9,
	}

	sysVFloatParameterRegisters = []*architecture.Register{
		registers.Xmm0,
		registers.Xmm1,
		registers.Xmm2,
		registers.Xmm3,
		registers.Xmm4,
		registers.Xmm5,
		registers.Xmm6,
		registers.Xmm7,
	}

	sysVGeneralReturnRegisters = []*architecture.Register{
		registers.Rax,
		registers.Rdx,
	}

	sysVFloatReturnRegisters = []*architecture.Register{
		registers.Xmm0,
		registers.Xmm1,
	}
)

type registersPicker struct {
	Generals []*architecture.Register
	Floats   []*architecture.Register
}

func newRegistersPicker(
	generals []*architecture.Register,
	floats []*architecture.Register,
) *registersPicker {
	return &registersPicker{
		Generals: generals,
		Floats:   floats,
	}
}

func (picker *registersPicker) Pick(
	request []bool, // true for float, false for general
) []*architecture.Register {
	numGeneral := 0
	numFloat := 0
	for _, requestFloat := range request {
		if requestFloat {
			numFloat += 1
		} else {
			numGeneral += 1
		}
	}

	if len(picker.Generals) < numGeneral || len(picker.Floats) < numFloat {
		return nil
	}

	registers := []*architecture.Register{}
	for _, requestFloat := range request {
		if requestFloat {
			registers = append(registers, picker.Floats[0])
			picker.Floats = picker.Floats[1:]
		} else {
			registers = append(registers, picker.Generals[0])
			picker.Generals = picker.Generals[1:]
		}
	}

	return registers
}

// This does not support SSEUP, X87, X87UP, COMPLEX_X87 parameter classes as
// defined in "3.2.3 Parameter Passing"
//
// This also does not support bool.  In particular, this won't automatically
// clear the non-least significant bits.
//
// Basic types classification:
//   - int*, uint*, and pointers are INTEGER
//   - float* are SSE
//
// Aggregate (struct / array) types classification:
//   - all aggregate types are "trivial for the purpose of calls", i.e.,
//     parameters are shallow copy-able
//   - all aggregate types larger than two eightbytes are in MEMORY since SSEUP
//     is not supported (no need to check each eightbyte individually)
//   - if all TypeChunkValue(s) in a eightbyte TypeChunk is SSE, than the
//     eightbyte is SSE.  Otherwise, the eightbyte is INTEGER.
//
// Parameters allocation:
//   - if the eightbyte class is INTEGER, the next available register of the
//     sequence %rdi, %rsi, %rdx, %rcx, %r8, and %r9 is used
//   - if the eightbyte class is SSE, the next available register in
//     %xmm0, ..., %xmm7 is used
//   - %r11 is used for function address (neither callee-saved nor caller-saved)
//   - %rbp is used for base pointer
//   - if there are no registers available for any eightbyte of an argument,
//     the whole argument is passed on stack
//   - if the class is MEMORY, the argument is pass on stack and respects
//     argument's type alignment
//   - stack arguments are pushed on the stack in reversed (right to left)
//     order
//   - if return value class is MEMORY, the caller provide space for the return
//     value and passes the address of this storage in %rdi as if it were the
//     first argument to the function.
//
// Return value allocation:
//   - if return value class is MEMORY, on return, %rax will contain the
//     address that has been passed in by the caller in %rdi
//   - if class is INTEGER, the next available register of the sequence
//     %rax, %rdx is used
//   - if class is SSE, the next available register of the sequence
//     %xmm0, %xmm1 is used
type sysVLite struct {
}

func (sysVLite) classify(valueType ir.Type) ([]bool, bool) {
	chunks := valueType.Chunks()
	if len(chunks) > 2 { // MEMORY class
		return nil, true
	}

	result := []bool{}
	for _, chunk := range chunks {
		// NOTE: when an entire chunk is occupied by an aggregate type value, we
		// need to classify the inner most aggregate type value.
		for len(chunk.Values) == 1 && chunk.Values[0].ValueTypeChunk != nil {
			chunk = chunk.Values[0].ValueTypeChunk
		}

		onFloat := true
		for _, entry := range chunk.Values {
			_, ok := entry.ValueType.(*ir.FloatType)
			if !ok {
				onFloat = false
				break
			}
		}

		result = append(result, onFloat)
	}

	return result, false
}

func (sysV sysVLite) Compute(
	function *ir.FunctionType,
) *architecture.CallConvention {
	convention := architecture.NewCallConvention(
		registers.Registers,
		sysVCalleeSavedRegisters)

	convention.SetFunctionAddressRegister(registers.R11)
	convention.SetBasePointer(registers.Rbp)

	parameterPicker := newRegistersPicker(
		sysVGeneralParameterRegisters,
		sysVFloatParameterRegisters)

	request, returnInMemory := sysV.classify(function.ReturnType)
	if returnInMemory {
		// Reserve %rdi for return value's address hidden parameter (%rax is
		// used to return the same address)
		_ = parameterPicker.Pick([]bool{false})
	} else {
		returnValuePicker := newRegistersPicker(
			sysVGeneralReturnRegisters,
			sysVFloatReturnRegisters)
		registers := returnValuePicker.Pick(request)
		if registers == nil {
			panic("should never happen")
		}

		convention.SetDirectRegisterReturnValue(registers)
	}

	for _, parameter := range function.ParameterTypes {
		request, onStack := sysV.classify(parameter)
		if !onStack {
			registers := parameterPicker.Pick(request)
			if registers == nil { // ran out of registers
				onStack = true
			} else {
				convention.AddRegisterArgument(registers)
			}
		}

		if onStack {
			convention.AddStackArgument(convention.AddStackEntry(parameter))
		}
	}

	if returnInMemory {
		convention.SetIndirectReturnValue(
			registers.Rdi,
			registers.Rax,
			convention.AddStackEntry(function.ReturnType))
	}

	convention.FinalizeCallFrameLayout()
	return convention
}
