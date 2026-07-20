package taint

import (
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/securego/gosec/v2/issue"
)

type RuleInfo struct {
	ID          string
	Description string
	Severity    string
	CWE         string
}

func NewGosecAnalyzer(rule *RuleInfo, config *Config) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func makeAnalyzerRunner(rule *RuleInfo, config *Config) func(*analysis.Pass) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil
}

func newIssue(analyzerID string, desc string, fileSet *token.FileSet,
	pos token.Pos, severity, confidence issue.Score,
) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func issueCodeSnippet(fileSet *token.FileSet, pos token.Pos) string {
	_ = "STUB: not implemented"
	return ""
}
