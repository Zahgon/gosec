package analyzers

import "golang.org/x/tools/go/analysis"

type AnalyzerSet struct {
	Analyzers             []*analysis.Analyzer
	AnalyzerSuppressedMap map[string]bool
}

func NewAnalyzerSet() *AnalyzerSet { _ = "STUB: not implemented"; return nil }

func (a *AnalyzerSet) Register(analyzer *analysis.Analyzer, isSuppressed bool) {
	_ = "STUB: not implemented"
	return
}

func (a *AnalyzerSet) IsSuppressed(ruleID string) bool { _ = "STUB: not implemented"; return false }
