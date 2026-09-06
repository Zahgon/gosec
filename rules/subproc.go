package rules

import (
	"go/ast"
	"go/token"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type subprocess struct {
	callListRule
}

func getEnclosingBodyStart(pos token.Pos, ctx *gosec.Context) token.Pos {
	_ = "STUB: not implemented"
	return *new(token.Pos)
}

func (r *subprocess) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *subprocess) isContext(n ast.Node, ctx *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func NewSubproc(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
