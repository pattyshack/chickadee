package call

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"

	"github.com/pattyshack/chickadee/amd64/registers"
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
)

func TestRegistersPicker(t *testing.T) {
	picker := newRegistersPicker(
		sysVGeneralParameterRegisters,
		sysVFloatParameterRegisters)

	selected := picker.Pick([]bool{false, false})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Rdi,
			registers.Rsi,
		},
		selected)

	selected = picker.Pick([]bool{true, false, true})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Xmm0,
			registers.Rdx,
			registers.Xmm1,
		},
		selected)

	selected = picker.Pick([]bool{false, false, false, false})
	expect.Nil(t, selected)

	selected = picker.Pick([]bool{false, false, true})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Rcx,
			registers.R8,
			registers.Xmm2,
		},
		selected)

	selected = picker.Pick([]bool{false, false, true})
	expect.Nil(t, selected)

	selected = picker.Pick([]bool{true, true, false, true})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Xmm3,
			registers.Xmm4,
			registers.R9,
			registers.Xmm5,
		},
		selected)

	selected = picker.Pick([]bool{true})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Xmm6,
		},
		selected)

	selected = picker.Pick([]bool{true, true})
	expect.Nil(t, selected)

	selected = picker.Pick([]bool{true})
	expect.Equal(
		t,
		[]*architecture.Register{
			registers.Xmm7,
		},
		selected)
}

func TestSysVClassifyBasicTypes(t *testing.T) {
	sysV := sysVLite{}

	for _, intType := range []ir.Type{
		ir.Int8,
		ir.Int16,
		ir.Int32,
		ir.Int64,
		ir.Uint8,
		ir.Uint16,
		ir.Uint32,
		ir.Uint64,
	} {
		request, onStack := sysV.classify(intType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{false}, request)

		addressType := ir.NewAddressType(intType)
		request, onStack = sysV.classify(addressType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{false}, request)

		varArrayAddrType := ir.NewVariableLengthArrayAddressType(intType)
		request, onStack = sysV.classify(varArrayAddrType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{false}, request)
	}

	for _, floatType := range []ir.Type{ir.Float32, ir.Float64} {
		request, onStack := sysV.classify(floatType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{true}, request)

		addressType := ir.NewAddressType(floatType)
		request, onStack = sysV.classify(addressType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{false}, request)

		varArrayAddrType := ir.NewVariableLengthArrayAddressType(floatType)
		request, onStack = sysV.classify(varArrayAddrType)
		expect.False(t, onStack)
		expect.Equal(t, []bool{false}, request)
	}

	funcType := ir.NewFunctionType(
		ir.SysVLiteCallConvention,
		[]ir.Type{},
		ir.Int32)
	request, onStack := sysV.classify(funcType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false}, request)
}

func TestSysVClassify8BitArray(t *testing.T) {
	sysV := sysVLite{}

	for _, int8Type := range []ir.Type{ir.Int8, ir.Uint8} {
		arrayType := ir.NewArrayType(int8Type, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		for numElements := 1; numElements < 9; numElements++ {
			arrayType := ir.NewArrayType(int8Type, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{false}, request)
		}

		for numElements := 9; numElements < 17; numElements++ {
			arrayType := ir.NewArrayType(int8Type, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{false, false}, request)
		}

		arrayType = ir.NewArrayType(int8Type, 17)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}
}

func TestSysVClassify16BitArray(t *testing.T) {
	sysV := sysVLite{}

	for _, int16Type := range []ir.Type{ir.Int16, ir.Uint16} {
		arrayType := ir.NewArrayType(int16Type, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		for numElements := 1; numElements < 5; numElements++ {
			arrayType := ir.NewArrayType(int16Type, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{false}, request)
		}

		for numElements := 5; numElements < 9; numElements++ {
			arrayType := ir.NewArrayType(int16Type, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{false, false}, request)
		}

		arrayType = ir.NewArrayType(int16Type, 9)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}
}

func TestSysVClassify32BitArray(t *testing.T) {
	sysV := sysVLite{}

	for _, elementType := range []ir.Type{ir.Int32, ir.Uint32, ir.Float32} {
		_, isFloat := elementType.(*ir.FloatType)

		arrayType := ir.NewArrayType(elementType, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		for numElements := 1; numElements < 3; numElements++ {
			arrayType := ir.NewArrayType(elementType, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{isFloat}, request)
		}

		for numElements := 3; numElements < 5; numElements++ {
			arrayType := ir.NewArrayType(elementType, numElements)
			request, onStack := sysV.classify(arrayType)
			expect.False(t, onStack)
			expect.NotNil(t, request)
			expect.Equal(t, []bool{isFloat, isFloat}, request)
		}

		arrayType = ir.NewArrayType(elementType, 5)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}
}

func TestSysVClassify64BitArray(t *testing.T) {
	sysV := sysVLite{}

	for _, elementType := range []ir.Type{ir.Int64, ir.Uint64, ir.Float64} {
		_, isFloat := elementType.(*ir.FloatType)

		arrayType := ir.NewArrayType(elementType, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		arrayType = ir.NewArrayType(elementType, 1)
		request, onStack = sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{isFloat}, request)

		arrayType = ir.NewArrayType(elementType, 2)
		request, onStack = sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{isFloat, isFloat}, request)

		arrayType = ir.NewArrayType(elementType, 3)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}
}

func TestSysVClassifyMultiChunkElementArray(t *testing.T) {
	sysV := sysVLite{}

	for _, valueType := range []ir.Type{ir.Int64, ir.Float64} {
		_, isFloat := valueType.(*ir.FloatType)

		elementType := ir.NewArrayType(valueType, 2)

		arrayType := ir.NewArrayType(elementType, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		arrayType = ir.NewArrayType(elementType, 1)
		request, onStack = sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{isFloat, isFloat}, request)

		arrayType = ir.NewArrayType(elementType, 2)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}

	for _, valueType := range []ir.Type{ir.Int64, ir.Float64} {
		elementType := ir.NewArrayType(valueType, 3)

		arrayType := ir.NewArrayType(elementType, 0)
		request, onStack := sysV.classify(arrayType)
		expect.False(t, onStack)
		expect.NotNil(t, request)
		expect.Equal(t, []bool{}, request)

		arrayType = ir.NewArrayType(elementType, 1)
		request, onStack = sysV.classify(arrayType)
		expect.True(t, onStack)
		expect.Nil(t, request)
	}
}

func TestSysVClassifyStruct(t *testing.T) {
	sysV := sysVLite{}

	structType := ir.NewStructType([]ir.Field{
		// INTEGER + INTEGER = INTEGER
		{
			Name: "i",
			Type: ir.Int32,
		},
		{
			Name: "j",
			Type: ir.Int32,
		},

		// SSE + SSE = SSE
		{
			Name: "x",
			Type: ir.Float32,
		},
		{
			Name: "y",
			Type: ir.Float32,
		},
	})
	request, onStack := sysV.classify(structType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false, true}, request)

	outerStructType := ir.NewStructType([]ir.Field{
		{
			Name: "s",
			Type: structType,
		},
	})
	request, onStack = sysV.classify(outerStructType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false, true}, request)

	structType = ir.NewStructType([]ir.Field{
		// INTEGER + SSE = INTEGER
		{
			Name: "i",
			Type: ir.Int32,
		},
		{
			Name: "j",
			Type: ir.Float32,
		},
	})
	request, onStack = sysV.classify(structType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false}, request)

	outerStructType = ir.NewStructType([]ir.Field{
		{
			Name: "s",
			Type: structType,
		},
	})
	request, onStack = sysV.classify(outerStructType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false}, request)

	structType = ir.NewStructType([]ir.Field{
		// INTEGER + INTEGER + INTEGER = INTEGER
		{
			Name: "i",
			Type: ir.Int8,
		},
		{
			Name: "j",
			Type: ir.Int16,
		},
		{
			Name: "k",
			Type: ir.Int32,
		},
	})
	request, onStack = sysV.classify(structType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false}, request)

	outerStructType = ir.NewStructType([]ir.Field{
		{
			Name: "s",
			Type: structType,
		},
	})
	request, onStack = sysV.classify(outerStructType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{false}, request)

	outerStructType = ir.NewStructType([]ir.Field{
		{
			Name: "f",
			Type: ir.Float32,
		},
		{
			Name: "s",
			Type: structType,
		},
	})
	request, onStack = sysV.classify(outerStructType)
	expect.False(t, onStack)
	expect.Equal(t, []bool{true, false}, request)

	structType = ir.NewStructType([]ir.Field{
		{
			Name: "i",
			Type: ir.Int64,
		},
		{
			Name: "j",
			Type: ir.Uint64,
		},
		{
			Name: "k",
			Type: ir.Float64,
		},
	})
	request, onStack = sysV.classify(structType)
	expect.True(t, onStack)
	expect.Nil(t, request)

	outerStructType = ir.NewStructType([]ir.Field{
		{
			Name: "s",
			Type: structType,
		},
	})
	request, onStack = sysV.classify(outerStructType)
	expect.True(t, onStack)
	expect.Nil(t, request)
}

func TestSysVBasicRegistersCall(t *testing.T) {
	structType := ir.NewStructType([]ir.Field{
		{
			Name: "i",
			Type: ir.Int64,
		},
		{
			Name: "f",
			Type: ir.Float32,
		},
	})

	intArrayType := ir.NewArrayType(ir.Int32, 3)
	floatArrayType := ir.NewArrayType(ir.Float64, 2)

	functionType := ir.NewFunctionType(
		ir.SysVLiteCallConvention,
		[]ir.Type{
			ir.Int16,
			ir.Float32,
			structType,
			intArrayType,
			floatArrayType,
		},
		structType)

	sysV := sysVLite{}

	convention := sysV.Compute(functionType)

	calleeSaved := map[string]bool{
		"%rbx": true,
		"%rbp": true,
		"%r12": true,
		"%r13": true,
		"%r14": true,
		"%r15": true,
	}

	expect.Equal(t, len(registers.Registers.Data), len(convention.Registers))
	for _, register := range registers.Registers.Data {
		expect.Equal(
			t,
			!calleeSaved[register.Name],
			convention.Registers[register].Clobbered)
	}

	expect.Nil(t, convention.CallFrameLayout)
	expect.Equal(t, 0, convention.CallFrameSize)

	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered: true,
			Require:   registers.R11,
		},
		convention.FunctionAddress)

	expect.Equal(
		t,
		&architecture.RegisterConstraint{
			Clobbered: false,
			Require:   registers.Rbp,
		},
		convention.BasePointer)

	expect.Equal(t, 5, len(convention.Arguments))

	// ir.Int16
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Rdi,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[0])

	// ir.Float32
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Xmm0,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[1])

	// structType
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Rsi,
					Clobbered: true,
				},
				{
					Require:   registers.Xmm1,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[2])

	// intArrayType
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Rdx,
					Clobbered: true,
				},
				{
					Require:   registers.Rcx,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[3])

	// floatArrayType
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Xmm2,
					Clobbered: true,
				},
				{
					Require:   registers.Xmm3,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[4])

	// return structType
	expect.Equal(
		t,
		architecture.ReturnValue{
			ReturnMapping: architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   registers.Rax,
						Clobbered: true,
					},
					{
						Require:   registers.Xmm0,
						Clobbered: true,
					},
				},
			},
		},
		convention.ReturnValue)
}

func TestSysVGeneralRegistersCall(t *testing.T) {
	functionType := ir.NewFunctionType(
		ir.SysVLiteCallConvention,
		[]ir.Type{
			ir.Int64,
			ir.Int64,
			ir.Int64,
			ir.Int64,
			ir.Int64,
			ir.Int64,
		},
		ir.NewStructType([]ir.Field{
			{
				Name: "i",
				Type: ir.Int64,
			},
			{
				Name: "f",
				Type: ir.Int64,
			},
		}))

	sysV := sysVLite{}

	convention := sysV.Compute(functionType)

	expect.Nil(t, convention.CallFrameLayout)
	expect.Equal(t, 0, convention.CallFrameSize)

	expect.Equal(t, 6, len(convention.Arguments))
	for idx, register := range []*architecture.Register{
		registers.Rdi,
		registers.Rsi,
		registers.Rdx,
		registers.Rcx,
		registers.R8,
		registers.R9,
	} {
		expect.Equal(
			t,
			architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   register,
						Clobbered: true,
					},
				},
			},
			convention.Arguments[idx])
	}

	expect.Equal(
		t,
		architecture.ReturnValue{
			ReturnMapping: architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   registers.Rax,
						Clobbered: true,
					},
					{
						Require:   registers.Rdx,
						Clobbered: true,
					},
				},
			},
		},
		convention.ReturnValue)
}

func TestSysVFloatRegistersCall(t *testing.T) {
	functionType := ir.NewFunctionType(
		ir.SysVLiteCallConvention,
		[]ir.Type{
			ir.Float64,
			ir.Float64,
			ir.Float64,
			ir.Float64,
			ir.Float64,
			ir.Float64,
			ir.Float64,
			ir.Float64,
		},
		ir.NewStructType([]ir.Field{
			{
				Name: "i",
				Type: ir.Float64,
			},
			{
				Name: "f",
				Type: ir.Float64,
			},
		}))

	sysV := sysVLite{}

	convention := sysV.Compute(functionType)

	expect.Nil(t, convention.CallFrameLayout)
	expect.Equal(t, 0, convention.CallFrameSize)

	expect.Equal(t, 8, len(convention.Arguments))
	for idx, register := range []*architecture.Register{
		registers.Xmm0,
		registers.Xmm1,
		registers.Xmm2,
		registers.Xmm3,
		registers.Xmm4,
		registers.Xmm5,
		registers.Xmm6,
		registers.Xmm7,
	} {
		expect.Equal(
			t,
			architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   register,
						Clobbered: true,
					},
				},
			},
			convention.Arguments[idx])
	}

	expect.Equal(
		t,
		architecture.ReturnValue{
			ReturnMapping: architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   registers.Xmm0,
						Clobbered: true,
					},
					{
						Require:   registers.Xmm1,
						Clobbered: true,
					},
				},
			},
		},
		convention.ReturnValue)
}

func TestSysVCallWithStack(t *testing.T) {
	array2Type := ir.NewArrayType(ir.Int64, 2)
	array3Type := ir.NewArrayType(ir.Int64, 3)
	array4Type := ir.NewArrayType(ir.Uint64, 4)
	array5Type := ir.NewArrayType(ir.Float64, 5)
	arrayReturnType := ir.NewArrayType(ir.Int32, 5)

	functionType := ir.NewFunctionType(
		ir.SysVLiteCallConvention,
		[]ir.Type{
			array3Type, // stack entry 0
			array2Type, // rsi, rdx  (NOTE: rdi used by return value)
			array4Type, // stack entry 1
			array2Type, // rcx, r8
			array5Type, // stack entry 2
			ir.Int64,   // r9
			ir.Int8,    // stack entry 3
			ir.Int16,   // stack entry 4
		},
		arrayReturnType) // stack entry 5

	sysV := sysVLite{}

	convention := sysV.Compute(functionType)

	expect.Equal(t, 6, len(convention.CallFrameLayout))
	expect.Equal(t, 8, len(convention.Arguments))

	// array3Type stack entry 0
	expect.Equal(
		t,
		architecture.ValueMapping{
			StackEntry: convention.CallFrameLayout[0],
		},
		convention.Arguments[0])

	// array2Type rsi, rdx
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Rsi,
					Clobbered: true,
				},
				{
					Require:   registers.Rdx,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[1])

	// array4Type stack entry 1
	expect.Equal(
		t,
		architecture.ValueMapping{
			StackEntry: convention.CallFrameLayout[1],
		},
		convention.Arguments[2])

	// array2Type rcx, r8
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.Rcx,
					Clobbered: true,
				},
				{
					Require:   registers.R8,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[3])

	// array5Type stack entry 2
	expect.Equal(
		t,
		architecture.ValueMapping{
			StackEntry: convention.CallFrameLayout[2],
		},
		convention.Arguments[4])

	// ir.Int64 r9
	expect.Equal(
		t,
		architecture.ValueMapping{
			Registers: []*architecture.RegisterConstraint{
				{
					Require:   registers.R9,
					Clobbered: true,
				},
			},
		},
		convention.Arguments[5])

	// ir.Int8 stack entry 3
	expect.Equal(
		t,
		architecture.ValueMapping{
			StackEntry: convention.CallFrameLayout[3],
		},
		convention.Arguments[6])

	// ir.Int32 stack entry 4
	expect.Equal(
		t,
		architecture.ValueMapping{
			StackEntry: convention.CallFrameLayout[4],
		},
		convention.Arguments[7])

	expect.Equal(
		t,
		architecture.ReturnValue{
			AddressParameter: &architecture.RegisterConstraint{
				Require:   registers.Rdi,
				Clobbered: true,
			},
			ScratchSpace: convention.CallFrameLayout[5],
			ReturnMapping: architecture.ValueMapping{
				Registers: []*architecture.RegisterConstraint{
					{
						Require:   registers.Rax,
						Clobbered: true,
					},
				},
			},
		},
		convention.ReturnValue)

	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   array3Type,
			Offset: 0,
		},
		convention.CallFrameLayout[0])
	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   array4Type,
			Offset: 24,
		},
		convention.CallFrameLayout[1])
	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   array5Type,
			Offset: 24 + 32,
		},
		convention.CallFrameLayout[2])
	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   ir.Int8,
			Offset: 24 + 32 + 40,
		},
		convention.CallFrameLayout[3])
	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   ir.Int16,
			Offset: 24 + 32 + 40 + 1 + 1, // +1 for alignment
		},
		convention.CallFrameLayout[4])
	expect.Equal(
		t,
		&architecture.StackEntry{
			Type:   arrayReturnType,
			Offset: 24 + 32 + 40 + 1 + 1 + 2 + 4, // +4 for alignment
		},
		convention.CallFrameLayout[5])

	expect.Equal(
		t,
		24+32+40+1+1+2+4+24,
		convention.CallFrameSize)
}
