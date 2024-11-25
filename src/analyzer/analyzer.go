package analyzer

import (
	"context"

	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/chickadee/ast"
)

func Analyze(sources []ast.SourceEntry, emitter *parseutil.Emitter) {
	collector := NewSignatureCollector(emitter)

	Process(
		sources,
		[][]Pass[[]ast.SourceEntry]{
			{
				ValidateAstSyntax(emitter),
				collector,
			},
		},
		nil)
	if emitter.HasErrors() {
		return
	}

	signatures := collector.Signatures()

	_, abort := context.WithCancel(context.Background())
	// TODO
	//abortCtx, abort := context.WithCancel(context.Background())

	entryEmitters := []*parseutil.Emitter{}
	ParallelProcess(
		sources,
		func(ast.SourceEntry) func(ast.SourceEntry) {
			entryEmitter := &parseutil.Emitter{}
			entryEmitters = append(entryEmitters, entryEmitter)

			passes := [][]Pass[ast.SourceEntry]{
				{InitializeControlFlowGraph(entryEmitter)},
				{ConstructSSA(entryEmitter)},
				{CheckTypes(entryEmitter, signatures)},
			}

			return func(entry ast.SourceEntry) {
				Process(entry, passes, nil)
				if entryEmitter.HasErrors() {
					abort()
				}
			}
		})

	for _, entryEmitter := range entryEmitters {
		emitter.EmitErrors(entryEmitter.Errors()...)
	}
}
