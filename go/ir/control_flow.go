package ir

type Jump struct {
	controlFlowInstruction

	Label string
}

type ConditionalJumpKind string

const (
	Jeq = ConditionalJumpKind("Jeq")
)

type ConditionalJump struct {
	controlFlowInstruction

	Kind ConditionalJumpKind

	Label string
	Src1  Value
	Src2  Value
}

type TerminalKind string

const (
	Ret = TerminalKind("ret")
)

type Terminal struct {
	controlFlowInstruction

	Kind TerminalKind

	RetValue *Value // return empty struct for void
}
