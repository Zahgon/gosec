package analyzers

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

func newInsecureCookieAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

type cookieState struct {
	allocPos    token.Pos
	secureSet   bool
	httpOnlySet bool
	sameSiteSet bool

	secureTrue   bool
	httpOnlyTrue bool
	sameSiteSafe bool
}

type insecureCookieState struct {
	*BaseAnalyzerState
	cookies     map[ssa.Value]*cookieState
	issuesByPos map[token.Pos]*issue.Issue
}

func newInsecureCookieState(pass *analysis.Pass) *insecureCookieState {
	_ = "STUB: not implemented"
	return nil
}

func runInsecureCookieAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *insecureCookieState) trackCookieFieldStore(store *ssa.Store) {
	_ = "STUB: not implemented"
	return
}

func (s *insecureCookieState) getOrCreateCookieState(root ssa.Value) *cookieState {
	_ = "STUB: not implemented"
	return nil
}

func (s *insecureCookieState) reportInsecureCookies() { _ = "STUB: not implemented"; return }

func (s *insecureCookieState) addIssue(pos token.Pos, msg string) {
	_ = "STUB: not implemented"
	return
}

func isHTTPCookiePointerType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func httpCookieFieldName(fieldAddr *ssa.FieldAddr) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func cookieRoot(v ssa.Value, depth int) ssa.Value {
	_ = "STUB: not implemented"
	return *new(ssa.Value)
}

func intConstValue(c *ssa.Const) (int64, bool) { _ = "STUB: not implemented"; return 0, false }
