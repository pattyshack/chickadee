package allocator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pattyshack/chickadee/analyzer/util"
	"github.com/pattyshack/chickadee/ast"
)

type AllocatorDebugger struct {
	*Allocator

	debugLiveness      bool
	debugLiveRanges    bool
	debugPreferences   bool
	debugDataLocations bool
	debugStackFrame    bool
	debugOperations    bool
}

func Debug(allocator *Allocator) util.Pass[ast.SourceEntry] {
	return &AllocatorDebugger{
		Allocator: allocator,
		//debugLiveness: true,
		//debugLiveRanges: true,
		//debugPreferences: true,
		debugDataLocations: true,
		debugStackFrame:    true,
		debugOperations:    true,
	}
}

func (debugger *AllocatorDebugger) printLiveness(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	printf("Liveness:\n")
	printf("  # of callee saved registers: %d\n", len(funcDef.PseudoParameters))

	for idx, block := range funcDef.Blocks {
		blockState := debugger.BlockStates[block]

		printf("  Block %d (%s):\n", idx, block.Label)

		printf("    LiveIn:\n")
		calleeSavedCount := 0
		for def, _ := range blockState.LiveIn {
			if strings.HasPrefix(def.Name, "%") {
				calleeSavedCount++
				continue
			}
			printf("      %s (%s)\n", def.Name, def.Loc())
		}

		printf("    LiveOut:\n")
		calleeSavedCount = 0
		for def, _ := range blockState.LiveOut {
			if strings.HasPrefix(def.Name, "%") {
				calleeSavedCount++
				continue
			}
			printf("      %s (%s)\n", def.Name, def.Loc())
		}
	}
}

func (debugger *AllocatorDebugger) printLiveRanges(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	printf("Live Ranges:\n")
	for idx, block := range funcDef.Blocks {
		blockState := debugger.BlockStates[block]

		printf("  Block %d (%s):\n", idx, block.Label)
		for def, liveRange := range blockState.LiveRanges {
			printf(
				"    %s: [%d %d] NextUses: %v (%s)\n",
				def.Name,
				liveRange.Start,
				liveRange.End,
				liveRange.NextUses,
				def.Loc())
		}
	}
}

func (debugger *AllocatorDebugger) printPreferences(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	printf("Preferences:\n")
	for idx, block := range funcDef.Blocks {
		blockState := debugger.BlockStates[block]

		printf("  Block %d (%s):\n", idx, block.Label)
		for reg, list := range blockState.Preferences {
			printf("    Register %s:\n", reg.Name)
			for _, entry := range list {
				printf("      %s\n", entry)
			}
		}
	}
}

func (debugger *AllocatorDebugger) printDataLocations(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	printf("Data Locations:\n")
	for idx, block := range funcDef.Blocks {
		blockState := debugger.BlockStates[block]

		printf("  Block %d (%s):\n", idx, block.Label)
		printf("    LocationIn:\n")
		for _, loc := range blockState.LocationIn {
			printf("      %s\n", loc)
		}

		printf("    ValueLocations Out:\n")
		for def, locs := range blockState.ValueLocations.Values {
			printf("      Definition: %s\n", def)
			for _, loc := range locs {
				printf("        %s\n", loc)
			}
		}
	}
}

func (debugger *AllocatorDebugger) printStackFrame(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	printf("Stack Frame (Size = %d):\n", debugger.StackFrame.TotalFrameSize)
	printf("  Layout (bottom to top):\n")
	for _, entry := range debugger.StackFrame.Layout {
		printf("    %s\n", entry)
	}
}

func (debugger *AllocatorDebugger) printOperations(
	funcDef *ast.FunctionDefinition,
	printf func(string, ...interface{}),
) {
	for idx, block := range funcDef.Blocks {
		blockState := debugger.BlockStates[block]

		printf("  Block %d (%s):\n", idx, block.Label)
		for _, op := range blockState.Operations {
			printf("    %s\n", op.Kind)
			if op.Destination != nil {
				printf("      Destination: %v\n", op.Destination)
			}
			if len(op.Sources) > 0 {
				printf("      Sources:\n")
				for _, src := range op.Sources {
					printf("        %v\n", src)
				}
			}
			if op.Instruction != nil {
				printf(
					"      Instruction: %s // %s\n",
					op.Instruction,
					op.Instruction.Loc())
			}
			if op.Value != nil {
				printf("      Value: %s\n", op.Value)
			}
			if op.StackFrame != nil {
				printf("      StackFrame: Size: %d\n", op.StackFrame.TotalFrameSize)
			}
			if op.SrcRegister != nil {
				printf("      SrcRegister: %s\n", op.SrcRegister.Name)
			}
			if op.DestRegister != nil {
				printf("      DestRegister: %s\n", op.DestRegister.Name)
			}
		}
	}
}

func (debugger *AllocatorDebugger) Process(entry ast.SourceEntry) {
	funcDef, ok := entry.(*ast.FunctionDefinition)
	if !ok {
		return
	}

	buffer := &bytes.Buffer{}
	printf := func(template string, args ...interface{}) {
		fmt.Fprintf(buffer, template, args...)
	}

	printf("Definition: %s\n", funcDef.Label)

	if debugger.debugLiveness {
		printf("------------------------------------------\n")
		debugger.printLiveness(funcDef, printf)
	}

	if debugger.debugLiveRanges {
		printf("------------------------------------------\n")
		debugger.printLiveRanges(funcDef, printf)
	}

	if debugger.debugPreferences {
		printf("------------------------------------------\n")
		debugger.printPreferences(funcDef, printf)
	}

	if debugger.debugDataLocations {
		printf("------------------------------------------\n")
		debugger.printDataLocations(funcDef, printf)
	}

	if debugger.debugStackFrame {
		printf("------------------------------------------\n")
		debugger.printStackFrame(funcDef, printf)
	}

	if debugger.debugOperations {
		printf("------------------------------------------\n")
		debugger.printOperations(funcDef, printf)
	}

	printf("==========================================\n")

	fmt.Println(buffer.String())
}
