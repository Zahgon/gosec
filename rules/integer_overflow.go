package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type integerOverflowCheck struct {
	callListRule
}

func (i *integerOverflowCheck) Match(node ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIntegerOverflowCheck(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
