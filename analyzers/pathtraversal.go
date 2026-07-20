package analyzers

import (
	"golang.org/x/tools/go/analysis"

	"github.com/securego/gosec/v2/taint"
)

func PathTraversal() taint.Config { _ = "STUB: not implemented"; return *new(taint.Config) }

func newPathTraversalAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}
