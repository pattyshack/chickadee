package call

import (
	"github.com/pattyshack/chickadee/ir"
	"github.com/pattyshack/chickadee/platform/architecture"
)

var (
	Conventions = architecture.CallConventions{
		ir.SysVLiteCallConvention: sysVLite{},
	}
)
