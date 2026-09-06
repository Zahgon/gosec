package analyzers

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const (
	msgUnsafeRedirectHeaderCopy = "Unsafe redirect policy may propagate sensitive headers across origins"
	msgSensitiveRedirectHeader  = "Sensitive headers should not be re-added in redirect policy callbacks"
)

var sensitiveRedirectHeaders = map[string]struct{}{
	"authorization":       {},
	"proxy-authorization": {},
	"cookie":              {},
}

func newRedirectHeaderPropagationAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func runRedirectHeaderPropagationAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func collectAnalyzerFunctions(srcFuncs []*ssa.Function) []*ssa.Function {
	_ = "STUB: not implemented"
	return nil
}

func addRedirectIssue(issues map[token.Pos]*issue.Issue, pass *analysis.Pass, pos token.Pos, what string, severity issue.Score, confidence issue.Score) {
	_ = "STUB: not implemented"
	return
}

func findRedirectLikeParams(fn *ssa.Function) (*ssa.Parameter, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isRequestSliceType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func isRequestHeaderStore(store *ssa.Store, reqParam *ssa.Parameter) bool {
	_ = "STUB: not implemented"
	return false
}

func isRequestHeaderValue(val ssa.Value, reqParam *ssa.Parameter) bool {
	_ = "STUB: not implemented"
	return false
}

func isHeaderMutationCall(call *ssa.Call) bool { _ = "STUB: not implemented"; return false }

func isHTTPHeaderType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func extractStringConst(v ssa.Value) string { _ = "STUB: not implemented"; return "" }

func valueDependsOn(value ssa.Value, target ssa.Value, depth int) bool {
	_ = "STUB: not implemented"
	return false
}
