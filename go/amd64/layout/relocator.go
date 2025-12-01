package layout

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pattyshack/chickadee/platform/layout"
)

type Rel32Relocator struct{}

func NewRelocator() layout.Relocator {
	return Rel32Relocator{}
}

func (Rel32Relocator) Relocate(
	symbol *layout.Symbol,
	startOffset int64,
	snippet []byte,
) error {
	if len(snippet) < 4 {
		return fmt.Errorf("invalid rel32 relocation. not enough bytes in snippet")
	}

	displacement := int32(0)
	n, err := binary.Decode(snippet, binary.LittleEndian, &displacement)
	if err != nil || n != 4 {
		panic("should never happen")
	}

	delta := symbol.Offset - (startOffset + 4) + int64(displacement)
	if delta < math.MinInt32 || math.MaxInt32 < delta {
		return fmt.Errorf("invalid rel32 relocation. delta overflow (%d)", delta)
	}

	n, err = binary.Encode(snippet, binary.LittleEndian, int32(delta))
	if err != nil || n != 4 {
		panic("should never happen")
	}

	return nil
}
