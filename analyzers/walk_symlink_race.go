package analyzers

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const msgWalkSymlinkRace = "Filesystem operation in filepath.Walk/WalkDir callback uses race-prone path; consider root-scoped APIs (e.g. os.Root) to prevent symlink TOCTOU traversal"

func newWalkSymlinkRaceAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func runWalkSymlinkRaceAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type walkSymlinkRaceState struct {
	*BaseAnalyzerState
	issuesByPos map[token.Pos]*issue.Issue
}

func newWalkSymlinkRaceState(pass *analysis.Pass) *walkSymlinkRaceState {
	_ = "STUB: not implemented"
	return nil
}

func (s *walkSymlinkRaceState) resolveFunctions(v ssa.Value) []*ssa.Function {
	_ = "STUB: not implemented"
	return nil
}

func (s *walkSymlinkRaceState) scanCallbackForRaceSinks(fn *ssa.Function, pathParam *ssa.Parameter) {
	_ = "STUB: not implemented"
	return
}

func (s *walkSymlinkRaceState) addIssue(pos token.Pos) { _ = "STUB: not implemented"; return }

func walkCallbackArgIndex(common *ssa.CallCommon) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func filesystemSinkArgIndexes(common *ssa.CallCommon) ([]int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isRootScopedFilesystemCall(callee *ssa.Function) bool { _ = "STUB: not implemented"; return false }

func isOSRootType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func isStringType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func pathDependsOn(value ssa.Value, target ssa.Value, depth int, visited map[ssa.Value]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

func storedValues(ptr ssa.Value) []ssa.Value { _ = "STUB: not implemented"; return nil }
