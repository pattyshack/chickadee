package amd64

import (
	"github.com/pattyshack/chickadee/platform/architecture"
)

// Reference: https://wiki.osdev.org/X86-64_Instruction_Encoding#Registers
var (
	rax = architecture.NewGeneralRegister("rax", 0)
	rcx = architecture.NewGeneralRegister("rcx", 1)
	rdx = architecture.NewGeneralRegister("rdx", 2)
	rbx = architecture.NewGeneralRegister("rbx", 3)
	// rsp = 4
	rbp = architecture.NewGeneralRegister("rbp", 5)
	rsi = architecture.NewGeneralRegister("rsi", 6)
	rdi = architecture.NewGeneralRegister("rdi", 7)
	r8  = architecture.NewGeneralRegister("r8", 8)
	r9  = architecture.NewGeneralRegister("r9", 9)
	r10 = architecture.NewGeneralRegister("r10", 10)
	r11 = architecture.NewGeneralRegister("r11", 11)
	r12 = architecture.NewGeneralRegister("r12", 12)
	r13 = architecture.NewGeneralRegister("r13", 13)
	r14 = architecture.NewGeneralRegister("r14", 14)
	r15 = architecture.NewGeneralRegister("r15", 15)

	xmm0  = architecture.NewFloatRegister("xmm0", 0)
	xmm1  = architecture.NewFloatRegister("xmm1", 1)
	xmm2  = architecture.NewFloatRegister("xmm2", 2)
	xmm3  = architecture.NewFloatRegister("xmm3", 3)
	xmm4  = architecture.NewFloatRegister("xmm4", 4)
	xmm5  = architecture.NewFloatRegister("xmm5", 5)
	xmm6  = architecture.NewFloatRegister("xmm6", 6)
	xmm7  = architecture.NewFloatRegister("xmm7", 7)
	xmm8  = architecture.NewFloatRegister("xmm8", 8)
	xmm9  = architecture.NewFloatRegister("xmm9", 9)
	xmm10 = architecture.NewFloatRegister("xmm10", 10)
	xmm11 = architecture.NewFloatRegister("xmm11", 11)
	xmm12 = architecture.NewFloatRegister("xmm12", 12)
	xmm13 = architecture.NewFloatRegister("xmm13", 13)
	xmm14 = architecture.NewFloatRegister("xmm14", 14)
	xmm15 = architecture.NewFloatRegister("xmm15", 15)

	registerSet = architecture.NewRegisterSet(
		rax, rbx, rcx, rdx, rbp, rsi, rdi, r8, r9, r10, r11, r12, r13, r14, r15,
		xmm0, xmm1, xmm2, xmm3, xmm4, xmm5, xmm6, xmm7,
		xmm8, xmm9, xmm10, xmm11, xmm12, xmm13, xmm14, xmm15)
)
