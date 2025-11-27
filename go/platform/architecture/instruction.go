package architecture

import (
	"github.com/pattyshack/chickadee/platform/layout"
)

type Instruction interface {
	// This assumes that registers are assigned and data are correctly placed.
	Encode(layout.SegmentBuilder) error
}
