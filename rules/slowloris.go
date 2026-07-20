package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type slowloris struct {
	issue.MetaData
}

func containsReadHeaderTimeout(node *ast.CompositeLit) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *slowloris) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSlowloris(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
