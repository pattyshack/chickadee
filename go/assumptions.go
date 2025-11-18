// Assumptions:
//
// - We'll only support 64-bit little-enidan architectures. Supporting only
// 64-bit simplifies pointer size, while supporting only little-endianness
// simplifies struct alignment padding, stack layout, etc.  This also
// simplifies compiler code organization.
//
// - Float registers are at least as large as general registers.  When float
// registers are larger than general registers:
//
//	a. We won't attempt to use the float register upper bytes for general data
//	   storage since accessing the upper bytes typically requires clobbering
//	   the lower bytes.
//
//	b. All float registers must be caller saved for all call conventions since
//	   the number of bytes to spill is context dependent.  It's impractical to
//	   spill the full content to memory.
//
// - We can spill to any storage (general/float) register
//
// - A register cannot be partitioned into multiple disjointed registers.  When
// a portion (e.g., AX) of a register is used, the entire register (e.g., RAX)
// is considered occupied.  This simplifies data location accounting.
//
// - Each architecture have exactly one instruction pointer register, one
// stack pointer register.  The instruction pointer and stack pointer are
// always live and hence can't be used for storage and general/float
// operations.  We won't make use of a frame pointer internally, however,
// call conventions will have the option to specify a frame pointer register.
// (XXX: maybe add option to write called address in function prologue)
//
// - We assume user input errors are handled prior to ir translation. The ir
// error checks are only used for guarding against compiler internal
// translation errors and should not be used as part of user input error
// reporting.
//
// - We won't support c style union types since that interacts poorly with type
// systems (maybe allow enum with values?)
package chickadee
