package amd64

import (
	"github.com/pattyshack/chickadee/platform/architecture"
)

const (
	RspEncoding = 4
)

// Reference: https://wiki.osdev.org/X86-64_Instruction_Encoding#Registers
var (
	Rax = architecture.NewGeneralRegister("rax", 0)
	Rcx = architecture.NewGeneralRegister("rcx", 1)
	Rdx = architecture.NewGeneralRegister("rdx", 2)
	Rbx = architecture.NewGeneralRegister("rbx", 3)
	// rsp = 4
	Rbp = architecture.NewGeneralRegister("rbp", 5)
	Rsi = architecture.NewGeneralRegister("rsi", 6)
	Rdi = architecture.NewGeneralRegister("rdi", 7)
	R8  = architecture.NewGeneralRegister("r8", 8)
	R9  = architecture.NewGeneralRegister("r9", 9)
	R10 = architecture.NewGeneralRegister("r10", 10)
	R11 = architecture.NewGeneralRegister("r11", 11)
	R12 = architecture.NewGeneralRegister("r12", 12)
	R13 = architecture.NewGeneralRegister("r13", 13)
	R14 = architecture.NewGeneralRegister("r14", 14)
	R15 = architecture.NewGeneralRegister("r15", 15)

	Xmm0  = architecture.NewFloatRegister("xmm0", 0)
	Xmm1  = architecture.NewFloatRegister("xmm1", 1)
	Xmm2  = architecture.NewFloatRegister("xmm2", 2)
	Xmm3  = architecture.NewFloatRegister("xmm3", 3)
	Xmm4  = architecture.NewFloatRegister("xmm4", 4)
	Xmm5  = architecture.NewFloatRegister("xmm5", 5)
	Xmm6  = architecture.NewFloatRegister("xmm6", 6)
	Xmm7  = architecture.NewFloatRegister("xmm7", 7)
	Xmm8  = architecture.NewFloatRegister("xmm8", 8)
	Xmm9  = architecture.NewFloatRegister("xmm9", 9)
	Xmm10 = architecture.NewFloatRegister("xmm10", 10)
	Xmm11 = architecture.NewFloatRegister("xmm11", 11)
	Xmm12 = architecture.NewFloatRegister("xmm12", 12)
	Xmm13 = architecture.NewFloatRegister("xmm13", 13)
	Xmm14 = architecture.NewFloatRegister("xmm14", 14)
	Xmm15 = architecture.NewFloatRegister("xmm15", 15)

	registerSet = architecture.NewRegisterSet(
		Rax, Rbx, Rcx, Rdx, Rbp, Rsi, Rdi, R8, R9, R10, R11, R12, R13, R14, R15,
		Xmm0, Xmm1, Xmm2, Xmm3, Xmm4, Xmm5, Xmm6, Xmm7,
		Xmm8, Xmm9, Xmm10, Xmm11, Xmm12, Xmm13, Xmm14, Xmm15)
)
