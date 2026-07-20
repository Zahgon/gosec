package rules

import (
	"go/ast"
	"regexp"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type badTempFile struct {
	callListRule
	args        *regexp.Regexp
	argCalls    gosec.CallList
	nestedCalls gosec.CallList
}

func (t *badTempFile) findTempDirArgs(n ast.Node, c *gosec.Context, suspect ast.Node) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (t *badTempFile) Match(n ast.Node, c *gosec.Context) (gi *issue.Issue, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBadTempFile(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
