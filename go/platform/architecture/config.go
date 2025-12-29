package architecture

type Config struct {
	Name string

	Registers RegisterSet

	InstructionSet

	CallConventions
}
