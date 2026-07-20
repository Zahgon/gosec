package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type ssrf struct {
	callListRule
}

func (r *ssrf) ResolveVar(n *ast.CallExpr, c *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *ssrf) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSSRFCheck(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
