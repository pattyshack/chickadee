package platform

import (
	"github.com/pattyshack/chickadee/platform/architecture"
	"github.com/pattyshack/chickadee/platform/layout"
	"github.com/pattyshack/chickadee/platform/layout/executable"
)

type OperatingSystem string

const (
	Linux = OperatingSystem("linux")
)

type Config struct {
	OperatingSystem

	Architecture architecture.Config

	Layout           layout.Config
	ExecutableFormat executable.Config
}
