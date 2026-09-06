package rules

import (
	"go/ast"
	"go/types"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type readfile struct {
	callListRule
	pathJoin gosec.CallList
	clean    gosec.CallList

	cleanedVar map[*types.Var]ast.Node

	joinedVar map[*types.Var]ast.Node
}

func (r *readfile) isJoinFunc(n ast.Node, c *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *readfile) isFilepathClean(v *types.Var, _ *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *readfile) trackCleanAssign(assign *ast.AssignStmt, c *gosec.Context) {
	_ = "STUB: not implemented"
	return
}

func (r *readfile) trackJoinAssignStmt(assign *ast.AssignStmt, c *gosec.Context) {
	_ = "STUB: not implemented"
	return
}

func (r *readfile) osRootSuggestion() string { _ = "STUB: not implemented"; return "" }

func (r *readfile) isSafeJoin(call *ast.CallExpr, c *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *readfile) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReadFile(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
