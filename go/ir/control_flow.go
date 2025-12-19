package ir

type Jump struct {
	controlFlowInstruction

	Label string
}

type ConditionalJumpKind string

const (
	Jeq = ConditionalJumpKind("Jeq")
	Jne = ConditionalJumpKind("Jne")
	Jlt = ConditionalJumpKind("Jlt")
	Jle = ConditionalJumpKind("Jle")
	Jgt = ConditionalJumpKind("Jgt")
	Jge = ConditionalJumpKind("Jge")
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
