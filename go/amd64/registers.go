package amd64

import (
	"github.com/pattyshack/chickadee/platform/architecture"
)

var (
	rax = architecture.NewGeneralRegister("rax")
	rbx = architecture.NewGeneralRegister("rbx")
	rcx = architecture.NewGeneralRegister("rcx")
	rdx = architecture.NewGeneralRegister("rdx")
	rbp = architecture.NewGeneralRegister("rbp")
	rsi = architecture.NewGeneralRegister("rsi")
	rdi = architecture.NewGeneralRegister("rdi")
	r8  = architecture.NewGeneralRegister("r8")
	r9  = architecture.NewGeneralRegister("r9")
	r10 = architecture.NewGeneralRegister("r10")
	r11 = architecture.NewGeneralRegister("r11")
	r12 = architecture.NewGeneralRegister("r12")
	r13 = architecture.NewGeneralRegister("r13")
	r14 = architecture.NewGeneralRegister("r14")
	r15 = architecture.NewGeneralRegister("r15")

	xmm0  = architecture.NewFloatRegister("xmm0")
	xmm1  = architecture.NewFloatRegister("xmm1")
	xmm2  = architecture.NewFloatRegister("xmm2")
	xmm3  = architecture.NewFloatRegister("xmm3")
	xmm4  = architecture.NewFloatRegister("xmm4")
	xmm5  = architecture.NewFloatRegister("xmm5")
	xmm6  = architecture.NewFloatRegister("xmm6")
	xmm7  = architecture.NewFloatRegister("xmm7")
	xmm8  = architecture.NewFloatRegister("xmm8")
	xmm9  = architecture.NewFloatRegister("xmm9")
	xmm10 = architecture.NewFloatRegister("xmm10")
	xmm11 = architecture.NewFloatRegister("xmm11")
	xmm12 = architecture.NewFloatRegister("xmm12")
	xmm13 = architecture.NewFloatRegister("xmm13")
	xmm14 = architecture.NewFloatRegister("xmm14")
	xmm15 = architecture.NewFloatRegister("xmm15")

	registerSet = architecture.NewRegisterSet(
		rax, rbx, rcx, rdx, rbp, rsi, rdi, r8, r9, r10, r11, r12, r13, r14, r15,
		xmm0, xmm1, xmm2, xmm3, xmm4, xmm5, xmm6, xmm7,
		xmm8, xmm9, xmm10, xmm11, xmm12, xmm13, xmm14, xmm15)
)
