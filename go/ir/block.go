package ir

import (
	"github.com/pattyshack/gt/parseutil"
)

// A stright-line / basic block
type Block struct {
	// NOTE: we'll only keep track of line position on a per block basis. Split
	// instructions into multiple blocks if fine grain line position tracking is
	// needed.
	//
	// TODO: emit line debug information
	parseutil.StartEndPos

	Label string

	Operations []*Definition

	// When nil, this implicitly jumps to the next block
	ControlFlow ControlFlowInstruction

	// internal

	Function *FunctionDefinition

	Parents []*Block
	// The jump child branch (if exist) is always before the fallthrough child
	// branch (if exist)
	Children []*Block

	Phis map[string]*Phi
}

type Phi struct {
	Dest *Definition

	// Value is usually a variable reference, but could be constant after
	// optimization.
	Srcs map[*Block]Value
}

type Instruction interface {
	isInstruction()

	SetParentBlock(*Block)
}

type instruction struct {
	Block *Block
}

func (*instruction) isInstruction() {}

func (inst *instruction) SetParentBlock(block *Block) {
	inst.Block = block
}

type ControlFlowInstruction interface {
	Instruction
	isControlFlow()
}
type controlFlowInstruction struct {
	instruction
}

func (inst *controlFlowInstruction) isControlFlow() {}
