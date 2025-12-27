package architecture

type Register struct {
	// NOTE: name must be '%'-prefixed
	Name string

	// Architecture specific instruction encoding. e.g., X.Reg on x64
	Encoding int

	AllowGeneralOperations bool
	AllowFloatOperations   bool

	// Internally assigned index used for ordering / sorting.  Indices are in the
	// ordered given by NewRegisterSet.
	Index int
}

// NOTE: name must be '%'-prefixed
func NewGeneralRegister(name string, encoding int) *Register {
	return &Register{
		Name:                   name,
		Encoding:               encoding,
		AllowGeneralOperations: true,
	}
}

// NOTE: name must be '%'-prefixed
func NewFloatRegister(name string, encoding int) *Register {
	return &Register{
		Name:                 name,
		Encoding:             encoding,
		AllowFloatOperations: true,
	}
}

// NOTE: The register set does not keep track of non-interchangable/specialized
// registers (e.g., instruction pointer, stack pointer, control/status
// registers, debug registers, etc.) since their instruction usage have no
// degrees of freedom.
type RegisterSet struct {
	// All general/float registers are usable for temporary data storage and
	// register spilling.
	Data []*Register

	// The set of registers usable for general operations.
	General []*Register

	// The set of registers usable for float operations.
	Float []*Register
}

func NewRegisterSet(registers ...*Register) RegisterSet {
	set := RegisterSet{
		Data: registers,
	}

	names := map[string]struct{}{}
	for idx, register := range registers {
		if register.Name == "" {
			panic("no register name")
		}

		_, ok := names[register.Name]
		if ok {
			panic("added duplicate register: " + register.Name)
		}
		names[register.Name] = struct{}{}

		register.Index = idx

		if register.AllowGeneralOperations {
			set.General = append(set.General, register)
		}

		if register.AllowFloatOperations {
			set.Float = append(set.Float, register)
		}
	}

	return set
}

func (set RegisterSet) Get(name string) *Register {
	for _, register := range set.Data {
		if register.Name == name {
			return register
		}
	}

	panic("unknown register: " + name)
}
