package allocator

import (
	"fmt"
	"sort"

	"github.com/pattyshack/chickadee/architecture"
	"github.com/pattyshack/chickadee/ast"
	"github.com/pattyshack/chickadee/platform"
)

// All distances are in terms of instruction distance relative to the beginning
// of the block (phis are at distance zero).
type LiveRange struct {
	Start int // first live distance (relative to beginning of current block)
	End   int // last use distance, inclusive.
}

type PreferredAllocation struct {
	// Instruction distance where the variable is required
	Use int

	// Def could be nil since the source ast.Value could be an immediate
	// or a global label.
	Def *ast.VariableDefinition

	// Which chunk of the definition maps to this register
	ChunkIdx int
}

func (pref PreferredAllocation) String() string {
	if pref.Def != nil {
		return fmt.Sprintf(
			"%d %s %d (%s)",
			pref.Use,
			pref.Def.Name,
			pref.ChunkIdx,
			pref.Def.Loc())
	}
	return fmt.Sprintf("%d (nil) %d", pref.Use, pref.ChunkIdx)
}

// The block's execution state at a particular point in time.
type BlockState struct {
	platform.Platform
	*ast.Block

	LiveIn  LiveSet
	LiveOut LiveSet

	NextInstIdx int

	// Note: unused definitions are not included in LiveRanges
	LiveRanges map[*ast.VariableDefinition]LiveRange

	Constraints map[ast.Instruction]*architecture.InstructionConstraints

	// Preferences are always at least one instruction ahead of NextInstIdx.
	Preferences map[*architecture.Register][]PreferredAllocation

	// Where data are located immediately prior to executing the block.
	// Every entry maps one-to-one to the corresponding live in set.
	LocationIn LocationSet

	// Where data are located immediately after the block executed.
	// Every entry maps one-to-one to the corresponding live in set.
	LocationOut LocationSet
}

func (state *BlockState) GenerateConstraints(targetPlatform platform.Platform) {
	constraints := map[ast.Instruction]*architecture.InstructionConstraints{}
	for _, inst := range state.Instructions {
		constraints[inst] = targetPlatform.InstructionConstraints(inst)
	}

	state.Constraints = constraints
}

// Note: A block's preferences cannot be precomputed since the block's
// preferences could changes due to its children's LocationIn (set by a
// different parent).
func (state *BlockState) ComputeLiveRangesAndPreferences(
	blockStates map[*ast.Block]*BlockState,
) {
	liveRanges := map[*ast.VariableDefinition]LiveRange{}
	state.LiveRanges = liveRanges
	state.Preferences = map[*architecture.Register][]PreferredAllocation{}

	defStarts := map[*ast.VariableDefinition]int{}
	for def, _ := range state.LiveIn {
		defStarts[def] = 0
	}

	for idx, inst := range state.Instructions {
		dist := idx + 1

		def := inst.Destination()
		if def != nil {
			defStarts[def] = dist
		}

		constraints := state.Constraints[inst]
		for srcIdx, src := range inst.Sources() {
			var def *ast.VariableDefinition
			ref, ok := src.(*ast.VariableReference)
			if ok {
				def = ref.UseDef
				liveRanges[ref.UseDef] = LiveRange{
					Start: defStarts[def],
					End:   dist,
				}
			}

			// Don't collect preferences from the first instruction since we only
			// care about look ahead preferences.
			if idx == 0 {
				continue
			}

			state.collectConstraintPreferences(
				dist,
				def,
				constraints.Sources[srcIdx])
		}
	}

	// Generate preference from children LocationIn populated by other parents.
	for _, childBlock := range state.Children {
		childState := blockStates[childBlock]
		if childState.LocationIn == nil {
			continue
		}
		state.computePreferencesFromChildLocationIn(childState.LocationIn)
	}

	state.computeLiveRangesAndPreferencesFromLiveOut(blockStates, defStarts)
}

func (state *BlockState) computePreferencesFromChildLocationIn(
	childLocationIn LocationSet,
) {
	sortedDefs := []*ast.VariableDefinition{}
	for def, _ := range childLocationIn {
		sortedDefs = append(sortedDefs, def)
	}

	sort.Slice(
		sortedDefs,
		func(i int, j int) bool { return sortedDefs[i].Name < sortedDefs[j].Name })

	dist := len(state.Instructions) + 1 // +1 for phi
	for _, def := range sortedDefs {
		loc := childLocationIn[def]

		for chunkIdx, reg := range loc.Registers {
			state.maybeAddPreference(reg, dist, def, chunkIdx)
		}
	}
}

func (state *BlockState) computeLiveRangesAndPreferencesFromLiveOut(
	blockStates map[*ast.Block]*BlockState,
	defStarts map[*ast.VariableDefinition]int,
) {
	type usage struct {
		inst ast.Instruction
		dist int
	}

	sortedUsages := []*usage{}
	usages := map[ast.Instruction]*usage{}
	for def, info := range state.LiveOut {
		maxDist := 0
		for inst, dist := range info.NextUse {
			if dist > maxDist {
				maxDist = dist
			}

			_, ok := usages[inst]
			if !ok {
				use := &usage{
					inst: inst,
					dist: dist,
				}
				usages[inst] = use
				sortedUsages = append(sortedUsages, use)
			}
		}

		state.LiveRanges[def] = LiveRange{
			Start: defStarts[def],
			// Note: global live range could be longer, but this is a good enough
			// estimate for this block.
			End: maxDist,
		}
	}

	sort.Slice(
		sortedUsages,
		func(i int, j int) bool {
			return sortedUsages[i].dist < sortedUsages[j].dist
		})

	for _, use := range sortedUsages {
		inst := use.inst
		_, ok := inst.(*ast.Phi)
		if ok { // phi copy has no preferences
			continue
		}

		constraints := blockStates[inst.ParentBlock()].Constraints[inst]
		for srcIdx, src := range inst.Sources() {
			ref, ok := src.(*ast.VariableReference)
			if !ok {
				continue
			}

			def := ref.UseDef
			_, ok = state.LiveOut[def]
			if !ok {
				continue
			}

			state.collectConstraintPreferences(
				use.dist,
				def,
				constraints.Sources[srcIdx])
		}
	}
}

func (state *BlockState) collectConstraintPreferences(
	dist int,
	def *ast.VariableDefinition,
	constraint *architecture.LocationConstraint,
) {
	for chunkIdx, candidate := range constraint.Registers {
		if candidate.Require == nil {
			continue
		}

		state.maybeAddPreference(candidate.Require, dist, def, chunkIdx)
	}
}

func (state *BlockState) maybeAddPreference(
	reg *architecture.Register,
	dist int,
	def *ast.VariableDefinition,
	chunkIdx int,
) {
	list, ok := state.Preferences[reg]
	if ok && list[len(list)-1].Use == dist {
		// confliciting preferences (just keep the first preference)
		return
	}

	state.Preferences[reg] = append(
		list,
		PreferredAllocation{
			Use:      dist,
			Def:      def,
			ChunkIdx: chunkIdx,
		})
}
