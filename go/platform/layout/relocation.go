package layout

import (
	"fmt"
)

type RelocatableInfo interface {
	ShiftAll(offset int64)

	Defs() Definitions
	Relocs() Relocations
}

type Relocatable interface {
	RelocatableInfo

	// Return a sub-slice of the content starting at startOffset. The relocator
	// will directly modify the sub-slice (Note that the sub-slice's length could
	// be longer than what's needed for the relocation)
	Peek(startOffset int64) ([]byte, error)

	SetRelocations(unresolved Relocations)
}

type Relocator interface {
	// Modify the snippet with the relocated address.
	Relocate(symbol *Symbol, startOffset int64, snippet []byte) error
}

type Section string

const (
	UnknownSection       = Section("")
	TextSection          = Section(".text")
	InitSection          = Section(".init")
	ReadOnlyDataSection  = Section(".rodata")
	ReadWriteDataSection = Section(".data")
	BSSSection           = Section(".bss")
)

type SymbolKind string

const (
	BasicBlockKind = SymbolKind("basic block")
	FunctionKind   = SymbolKind("function")
	ObjectKind     = SymbolKind("object")
)

type Symbol struct {
	Kind SymbolKind

	Section // Not set for labels.

	// NOTE: For compatibility with standards such as elf, symbol name must be
	// unique regardless of symbol kind.
	Name string

	// Relative to the start of the content segment.
	Offset int64

	// Symbol's content value range [Offset, Offset + Size).  Not set for label.
	Size int64
}

type Definitions struct {
	Labels []*Symbol // BasicBlockKind

	Symbols []*Symbol // FunctionKind or ObjectKind
}

func (defs *Definitions) Shift(offset int64) {
	for _, entry := range defs.Labels {
		entry.Offset += offset
	}

	for _, entry := range defs.Symbols {
		entry.Offset += offset
	}
}

// NOTE: this assumes definitions have already been shifted
func MergeDefinitions[T RelocatableInfo](
	relocatables ...T,
) (
	Definitions,
	map[string]*Symbol,
	map[string]*Symbol,
	error,
) {
	numLabels := 0
	numSymbols := 0
	for _, relocatable := range relocatables {
		defs := relocatable.Defs()
		numLabels += len(defs.Labels)
		numSymbols += len(defs.Symbols)
	}

	merged := Definitions{}
	if numLabels > 0 {
		merged.Labels = make([]*Symbol, 0, numLabels)
	}
	if numSymbols > 0 {
		merged.Symbols = make([]*Symbol, 0, numSymbols)
	}

	labels := make(map[string]*Symbol, numLabels)
	symbols := make(map[string]*Symbol, numSymbols)

	for _, relocatable := range relocatables {
		defs := relocatable.Defs()

		for _, label := range defs.Labels {
			_, ok := labels[label.Name]
			if ok {
				return Definitions{}, nil, nil, fmt.Errorf(
					"found duplicate label (%s)",
					label.Name)
			}

			if label.Kind != BasicBlockKind {
				return Definitions{}, nil, nil, fmt.Errorf(
					"invalid label %s. unsupported kind (%s)",
					label.Name,
					label.Kind)
			}

			labels[label.Name] = label
			merged.Labels = append(merged.Labels, label)
		}

		for _, symbol := range defs.Symbols {
			_, ok := symbols[symbol.Name]
			if ok {
				return Definitions{}, nil, nil, fmt.Errorf(
					"found duplicate symbol (%s)",
					symbol.Name)
			}

			if symbol.Kind != FunctionKind && symbol.Kind != ObjectKind {
				return Definitions{}, nil, nil, fmt.Errorf(
					"invalid symbol %s. unsupported kind (%s)",
					symbol.Name,
					symbol.Kind)
			}

			symbols[symbol.Name] = symbol
			merged.Symbols = append(merged.Symbols, symbol)
		}
	}

	return merged, labels, symbols, nil
}

type Relocation struct {
	Name   string
	Offset int64 // Offset relative to the start of the content segment.
}

type Relocations struct {
	Labels  []*Relocation
	Symbols []*Relocation
}

func (relocs *Relocations) Shift(offset int64) {
	for _, entry := range relocs.Labels {
		entry.Offset += offset
	}

	for _, entry := range relocs.Symbols {
		entry.Offset += offset
	}
}

// NOTE: this assumes relocations have already been shifted
func MergeRelocations[T RelocatableInfo](relocatables ...T) Relocations {
	numLabels := 0
	numSymbols := 0
	for _, relocatable := range relocatables {
		defs := relocatable.Relocs()
		numLabels += len(defs.Labels)
		numSymbols += len(defs.Symbols)
	}

	merged := Relocations{}
	if numLabels > 0 {
		merged.Labels = make([]*Relocation, 0, numLabels)
	}
	if numSymbols > 0 {
		merged.Symbols = make([]*Relocation, 0, numSymbols)
	}

	for _, relocatable := range relocatables {
		defs := relocatable.Relocs()
		merged.Labels = append(merged.Labels, defs.Labels...)
		merged.Symbols = append(merged.Symbols, defs.Symbols...)
	}

	return merged
}

func link(
	content Relocatable,
	symbols map[string]*Symbol,
	relocations []*Relocation,
	relocator Relocator,
) (
	[]*Relocation,
	error,
) {
	if len(relocations) == 0 {
		return nil, nil
	}

	if len(symbols) == 0 {
		return relocations, nil
	}

	unresolved := make([]*Relocation, 0, len(relocations))
	for _, reloc := range relocations {
		symbol, ok := symbols[reloc.Name]
		if ok {
			snippet, err := content.Peek(reloc.Offset)
			if err != nil {
				return nil, err
			}
			relocator.Relocate(symbol, reloc.Offset, snippet)
		} else {
			unresolved = append(unresolved, reloc)
		}
	}

	if len(unresolved) == 0 {
		return nil, nil
	}
	return unresolved, nil
}

func Link(
	content Relocatable,
	labels map[string]*Symbol,
	symbols map[string]*Symbol,
	relocator Relocator,
) error {
	relocations := content.Relocs()

	unbinded := Relocations{}

	unresolved, err := link(content, labels, relocations.Labels, relocator)
	if err != nil {
		return err
	}
	unbinded.Labels = unresolved

	unresolved, err = link(content, symbols, relocations.Symbols, relocator)
	if err != nil {
		return err
	}
	unbinded.Symbols = unresolved

	content.SetRelocations(unbinded)
	return nil
}
