package instructions

import (
	"github.com/pattyshack/chickadee/platform/architecture"
)

var InstructionSet = architecture.InstructionSet{
	Jump: jumpSelector{},

	JeqUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jeIntImmediate,
		encodeLeftImmediate:  jeIntImmediate,
		encode:               je,
	},
	JeqInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jeIntImmediate,
		encodeLeftImmediate:  jeIntImmediate,
		encode:               je,
	},
	JeqFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  je,
	},

	JneUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jneIntImmediate,
		encodeLeftImmediate:  jneIntImmediate,
		encode:               jne,
	},
	JneInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jneIntImmediate,
		encodeLeftImmediate:  jneIntImmediate,
		encode:               jne,
	},
	JneFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  jne,
	},

	JltUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jltIntImmediate,
		encodeLeftImmediate:  jgtIntImmediate, // (i < s) == (s > i)
		encode:               jlt,
	},
	JltInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jltIntImmediate,
		encodeLeftImmediate:  jgtIntImmediate, // (i < s) == (s > i)
		encode:               jlt,
	},
	JltFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  jlt,
	},

	JleUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jleIntImmediate,
		encodeLeftImmediate:  jgeIntImmediate, // (i <= s) == (s >= i)
		encode:               jle,
	},
	JleInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jleIntImmediate,
		encodeLeftImmediate:  jgeIntImmediate, // (i <= s) == (s >= i)
		encode:               jle,
	},
	JleFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  jle,
	},

	JgtUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jgtIntImmediate,
		encodeLeftImmediate:  jltIntImmediate, // (i > s) == (s < i)
		encode:               jgt,
	},
	JgtInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jgtIntImmediate,
		encodeLeftImmediate:  jltIntImmediate, // (i > s) == (s < i)
		encode:               jgt,
	},
	JgtFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  jgt,
	},

	JgeUint: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jgeIntImmediate,
		encodeLeftImmediate:  jleIntImmediate, // (i >= s) == (s <= i)
		encode:               jge,
	},
	JgeInt: conditionalJumpSelector{
		isFloat:              false,
		encodeRightImmediate: jgeIntImmediate,
		encodeLeftImmediate:  jleIntImmediate, // (i >= s) == (s <= i)
		encode:               jge,
	},
	JgeFloat: conditionalJumpSelector{
		isFloat: true,
		encode:  jge,
	},

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

	MulUint: commonBinaryOperationSelector{
		isFloat:     false,
		isSymmetric: true,
		encodeRM:    mul,
	},
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
