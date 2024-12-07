package analyzer

import (
	"fmt"
	"sync"

	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/chickadee/architecture"
	"github.com/pattyshack/chickadee/ast"
	"github.com/pattyshack/chickadee/platform"
)

type callRetConstraints struct {
	platform platform.Platform

	sync.WaitGroup

	mutex            sync.Mutex
	entryConstraints map[ast.SourceEntry]*architecture.InstructionConstraints
}

func (collector *callRetConstraints) process(entry ast.SourceEntry) {
	defer collector.Done()

	funcDef, ok := entry.(*ast.FunctionDefinition)
	if !ok {
		return
	}

	callSpec := collector.platform.CallSpec(funcDef.CallConvention)
	constraints := callSpec.CallRetConstraints(funcDef.Type().(ast.FunctionType))

	collector.mutex.Lock()
	defer collector.mutex.Unlock()

	collector.entryConstraints[entry] = constraints
}

func (collector *callRetConstraints) Get() map[ast.SourceEntry]*architecture.InstructionConstraints {
	collector.Wait()
	return collector.entryConstraints
}

func CollectSignaturesAndCallRetConstraints(
	entries []ast.SourceEntry,
	targetPlatform platform.Platform,
	emitter *parseutil.Emitter,
) (
	map[string]ast.SourceEntry,
	*callRetConstraints,
) {
	constraints := &callRetConstraints{
		platform:         targetPlatform,
		entryConstraints: map[ast.SourceEntry]*architecture.InstructionConstraints{},
	}

	result := map[string]ast.SourceEntry{}
	for _, source := range entries {
		if source.HasDeclarationSyntaxError() {
			continue
		}

		switch entry := source.(type) {
		case *ast.FunctionDefinition:
			prev, ok := result[entry.Label]
			if ok {
				emitter.Emit(
					entry.Loc(),
					"definition (%s) previously defined at (%s)",
					entry.Label,
					prev.Loc())
				continue
			}

			result[entry.Label] = entry
			constraints.Add(1)
			go constraints.process(entry)
		default:
			panic(fmt.Sprintf("%s: unhandled SourceEntry", source.Loc()))
		}
	}

	return result, constraints
}
