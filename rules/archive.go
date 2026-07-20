package rules

import (
	"go/ast"
	"go/types"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type archive struct {
	callListRule
	argTypes []string
}

func getArchiveBaseType(expr ast.Expr, ctx *gosec.Context, file *ast.File) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

func (a *archive) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewArchive(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
