package instructions

import (
	"github.com/pattyshack/chickadee/platform/architecture"
)

var InstructionSet = architecture.InstructionSet{
	NotUint: unaryMSelector{
		encodeM: not,
	},
	NotInt: unaryMSelector{
		encodeM: not,
	},

	NegInt: unaryMSelector{
		encodeM: negSignedInt,
	},
	NegFloat: negFloatSelector{
		f32: unaryMSelector{
			encodeM: negFloat32,
		},
	},

	UintToUint: conversionSelector{
		srcIsFloat:       false,
		destIsFloat:      false,
		encodeConversion: convertIntToInt,
	},
	IntToUint: conversionSelector{
		srcIsFloat:       false,
		destIsFloat:      false,
		encodeConversion: convertIntToInt,
	},
	FloatToUint: conversionSelector{
		srcIsFloat:       true,
		destIsFloat:      false,
		encodeConversion: convertFloatToInt,
	},

	UintToInt: conversionSelector{
		srcIsFloat:       false,
		destIsFloat:      false,
		encodeConversion: convertIntToInt,
	},
	IntToInt: conversionSelector{
		srcIsFloat:       false,
		destIsFloat:      false,
		encodeConversion: convertIntToInt,
	},
	FloatToInt: conversionSelector{
		srcIsFloat:       true,
		destIsFloat:      false,
		encodeConversion: convertFloatToInt,
	},

	UintToFloat: uintToFloatSelector{
		smallUint: conversionSelector{
			srcIsFloat:       false,
			destIsFloat:      true,
			encodeConversion: convertSmallUintToFloat,
		},
	},
	IntToFloat: conversionSelector{
		srcIsFloat:       false,
		destIsFloat:      true,
		encodeConversion: convertSignedIntToFloat,
	},
	FloatToFloat: conversionSelector{
		srcIsFloat:       true,
		destIsFloat:      true,
		encodeConversion: convertFloatToFloat,
	},

	// TODO 3-address form add via lea
	AddUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    addIntImmediate,
		encodeRM:    add,
	},
	// TODO 3-address form add via lea
	AddInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    addIntImmediate,
		encodeRM:    add,
	},
	AddFloat: commonBinaryOperationSelector{
		isFloat:     true,
		isSymmetric: true,
		encodeRM:    add,
	},

	// TODO mul int/uint need to handle RMI encoding
	MulUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeRM:    mul,
	},
	// TODO mul int/uint need to handle RMI encoding
	MulInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeRM:    mul,
	},
	MulFloat: commonBinaryOperationSelector{
		isFloat:     true,
		isSymmetric: true,
		encodeRM:    mul,
	},

	DivUint: divRemSelector{
		isRem: false,
	},
	DivInt: divRemSelector{
		isRem: false,
	},
	DivFloat: commonBinaryOperationSelector{
		isFloat:     true,
		isSymmetric: true,
		encodeRM:    divFloat,
	},

	RemUint: divRemSelector{
		isRem: true,
	},
	RemInt: divRemSelector{
		isRem: true,
	},

	SubUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: false,
		encodeMI:    subIntImmediate,
		encodeRM:    sub,
	},
	SubInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: false,
		encodeMI:    subIntImmediate,
		encodeRM:    sub,
	},
	SubFloat: commonBinaryOperationSelector{
		isFloat:     true,
		isSymmetric: false,
		encodeRM:    sub,
	},

	ShlUint: shiftSelector{
		encodeMI8: shlIntImmediate,
		encodeMC:  shl,
	},
	ShlInt: shiftSelector{
		encodeMI8: shlIntImmediate,
		encodeMC:  shl,
	},

	ShrUint: shiftSelector{
		encodeMI8: shrIntImmediate,
		encodeMC:  shr,
	},
	ShrInt: shiftSelector{
		encodeMI8: shrIntImmediate,
		encodeMC:  shr,
	},

	AndUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    andIntImmediate,
		encodeRM:    and,
	},
	AndInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    andIntImmediate,
		encodeRM:    and,
	},

	OrUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    orIntImmediate,
		encodeRM:    or,
	},
	OrInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    orIntImmediate,
		encodeRM:    or,
	},

	XorUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    xorIntImmediate,
		encodeRM:    xor,
	},
	XorInt: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeMI:    xorIntImmediate,
		encodeRM:    xor,
	},
}
