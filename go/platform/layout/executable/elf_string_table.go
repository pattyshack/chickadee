package executable

import (
	"fmt"
	"io"

	"github.com/pattyshack/chickadee/platform/layout"
)

type ElfStringTable struct {
	Size    uint32
	Indices map[string]uint32
	Entries []string
}

func NewElfStringTable() ElfStringTable {
	return ElfStringTable{
		Size: 1, // the table always have a leading '\0' byte
		Indices: map[string]uint32{
			"": 0, // we'll use index 0 as a pseudo empty-string entry
		},
	}
}

func NewElfStringTableFromSymbols(
	symbols []*layout.Symbol,
) ElfStringTable {
	table := ElfStringTable{
		Size:    uint32(1),
		Indices: make(map[string]uint32, len(symbols)+1),
		Entries: make([]string, 0, len(symbols)),
	}

	table.Indices[""] = 0
	for _, symbol := range symbols {
		table.Indices[symbol.Name] = table.Size
		table.Size += uint32(len(symbol.Name)) + 1 // +1 for trailing '\0'
		table.Entries = append(table.Entries, symbol.Name)
	}

	return table
}

func (table *ElfStringTable) MaybeInsert(entry string) (uint32, bool) {
	idx, ok := table.Indices[entry]
	if ok {
		return idx, false
	}

	idx = table.Size
	table.Indices[entry] = idx
	table.Size += uint32(len(entry)) + 1 // +1 for trailing '\0'
	table.Entries = append(table.Entries, entry)

	return idx, true
}

func (table *ElfStringTable) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write([]byte{0})
	numWritten := int64(n)
	if err != nil {
		return 0, fmt.Errorf("failed to write string table: %w", err)
	}

	for _, entry := range table.Entries {
		n, err := writer.Write([]byte(entry))
		numWritten += int64(n)
		if err != nil {
			return numWritten, fmt.Errorf("failed to write string table: %w", err)
		}

		n, err = writer.Write([]byte{0})
		numWritten += int64(n)
		if err != nil {
			return 0, fmt.Errorf("failed to write string table: %w", err)
		}
	}

	if numWritten != int64(table.Size) {
		panic("should never happen")
	}

	return numWritten, nil
}
