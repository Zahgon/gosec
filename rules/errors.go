package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type noErrorCheck struct {
	issue.MetaData
	whitelist gosec.CallList
}

func returnsError(callExpr *ast.CallExpr, ctx *gosec.Context) int {
	_ = "STUB: not implemented"
	return 0
}

func (r *noErrorCheck) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNoErrorCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func toStringSlice(values []interface{}) []string { _ = "STUB: not implemented"; return nil }
