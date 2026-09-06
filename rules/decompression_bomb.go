package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type decompressionBombCheck struct {
	issue.MetaData
	readerCalls gosec.CallList
	copyCalls   gosec.CallList
}

func containsReaderCall(node ast.Node, ctx *gosec.Context, list gosec.CallList) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *decompressionBombCheck) Match(node ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDecompressionBombCheck(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
