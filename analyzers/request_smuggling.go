package analyzers

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const (
	msgConflictingHeaders = "Setting both Transfer-Encoding and Content-Length headers may enable request smuggling attacks"
)

func newRequestSmugglingAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func runRequestSmugglingAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type requestSmugglingState struct {
	*BaseAnalyzerState
	ssaFuncs []*ssa.Function

	headerOps map[ssa.Value]*headerTracker
}

type headerTracker struct {
	hasTransferEncoding bool
	hasContentLength    bool
	tePos               token.Pos
	clPos               token.Pos
}

func newRequestSmugglingState(pass *analysis.Pass, funcs []*ssa.Function) *requestSmugglingState {
	_ = "STUB: not implemented"
	return nil
}

func (s *requestSmugglingState) Release() { _ = "STUB: not implemented"; return }

func (s *requestSmugglingState) trackHeaderOperation(instr ssa.Instruction) {
	_ = "STUB: not implemented"
	return
}

func (s *requestSmugglingState) isHTTPHeaderSet(call *ssa.Call) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *requestSmugglingState) extractStringConstant(val ssa.Value) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *requestSmugglingState) findResponseWriter(headerSetCall *ssa.Call) ssa.Value {
	_ = "STUB: not implemented"
	return *new(ssa.Value)
}

func (s *requestSmugglingState) isHeaderMethodCall(call *ssa.Call) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *requestSmugglingState) detectHeaderConflicts() []*issue.Issue {
	_ = "STUB: not implemented"
	return nil
}
