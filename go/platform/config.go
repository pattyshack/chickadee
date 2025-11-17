package platform

import (
	"github.com/pattyshack/chickadee/platform/layout"
	"github.com/pattyshack/chickadee/platform/layout/executable"
)

type Architecture string
type OperatingSystem string

const (
	Amd64 = Architecture("amd64")

	Linux = OperatingSystem("linux")
)

type Config struct {
	Architecture
	OperatingSystem

	Layout           layout.Config
	ExecutableFormat executable.Config
}
