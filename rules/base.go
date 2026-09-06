package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type callListRule struct {
	issue.MetaData
	calls gosec.CallList
}

func newCallListRule(id, what string, severity, confidence issue.Score) callListRule {
	_ = "STUB: not implemented"
	return *new(callListRule)
}

func (r *callListRule) Add(selector, ident string) *callListRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *callListRule) AddAll(selector string, idents ...string) *callListRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *callListRule) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
