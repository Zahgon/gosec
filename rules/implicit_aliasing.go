package rules

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type implicitAliasing struct {
	issue.MetaData
	aliases         map[*types.Var]struct{}
	rightBrace      token.Pos
	acceptableAlias []*ast.UnaryExpr
}

func containsUnary(exprs []*ast.UnaryExpr, expr *ast.UnaryExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func getIdentExpr(expr ast.Expr) (*ast.Ident, bool) { _ = "STUB: not implemented"; return nil, false }

func doGetIdentExpr(expr ast.Expr, hasSelector bool) (*ast.Ident, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *implicitAliasing) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewImplicitAliasing(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
