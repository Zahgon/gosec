package goanalysis

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/packages"

	"github.com/securego/gosec/v2/issue"
)

const Doc = `gosec is a static analysis tool that scans Go code for security problems.`

var Analyzer = &analysis.Analyzer{
	Name:     "gosec",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{buildssa.Analyzer},
}

var (
	flagIncludeRules     string
	flagExcludeRules     string
	flagExcludeGenerated bool
	flagMinSeverity      string
	flagMinConfidence    string
)

//nolint:gochecknoinits // Required for go/analysis Analyzer flag registration
func init() {
	Analyzer.Flags.StringVar(&flagIncludeRules, "include", "", "Comma-separated list of rule IDs to include (e.g., G101,G102)")
	Analyzer.Flags.StringVar(&flagExcludeRules, "exclude", "", "Comma-separated list of rule IDs to exclude (e.g., G104)")
	Analyzer.Flags.BoolVar(&flagExcludeGenerated, "exclude-generated", true, "Exclude generated code from analysis")
	Analyzer.Flags.StringVar(&flagMinSeverity, "severity", "low", "Minimum severity: low, medium, or high")
	Analyzer.Flags.StringVar(&flagMinConfidence, "confidence", "low", "Minimum confidence: low, medium, or high")
}

func run(pass *analysis.Pass) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func convertPassToPackage(pass *analysis.Pass) *packages.Package {
	_ = "STUB: not implemented"
	return nil
}

func buildFilters[T any](include, exclude string, newFilter func(bool, ...string) T) []T {
	_ = "STUB: not implemented"
	return nil
}

func parseRuleIDs(s string) []string { _ = "STUB: not implemented"; return nil }

func parseScore(s string) (issue.Score, error) {
	_ = "STUB: not implemented"
	return *new(issue.Score), nil
}

func parsePosition(fset *token.FileSet, iss *issue.Issue) token.Pos {
	_ = "STUB: not implemented"
	return *new(token.Pos)
}
