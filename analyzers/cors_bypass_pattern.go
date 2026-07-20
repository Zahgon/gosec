package analyzers

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const (
	msgOverbroadBypassPattern = "Overbroad AddInsecureBypassPattern disables cross-origin protections for too many paths"
	msgRequestBypassPattern   = "AddInsecureBypassPattern argument derived from request data can allow bypass of cross-origin protections"
)

func newCORSBypassPatternAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func runCORSBypassPatternAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func addG121Issue(issues map[token.Pos]*issue.Issue, pass *analysis.Pass, pos token.Pos, what string, severity issue.Score, confidence issue.Score) {
	_ = "STUB: not implemented"
	return
}

func findHTTPRequestParam(fn *ssa.Function) *ssa.Parameter { _ = "STUB: not implemented"; return nil }

func isAddInsecureBypassPatternCall(call *ssa.CallCommon) bool {
	_ = "STUB: not implemented"
	return false
}

func isCrossOriginProtectionType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func extractStringValue(v ssa.Value, depth int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func isOverbroadBypassPattern(pattern string) bool { _ = "STUB: not implemented"; return false }
