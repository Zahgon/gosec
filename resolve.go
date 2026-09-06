package gosec

import (
	"go/ast"
	"go/types"
)

func resolveIdent(n *ast.Ident, c *Context) bool { _ = "STUB: not implemented"; return false }

func resolveValueSpec(n *ast.ValueSpec, c *Context) bool { _ = "STUB: not implemented"; return false }

func resolveAssign(n *ast.AssignStmt, c *Context) bool { _ = "STUB: not implemented"; return false }

func resolveCompLit(n *ast.CompositeLit, c *Context) bool { _ = "STUB: not implemented"; return false }

func resolveBinExpr(n *ast.BinaryExpr, c *Context) bool { _ = "STUB: not implemented"; return false }

func resolveCallExpr(node *ast.CallExpr, c *Context) bool { _ = "STUB: not implemented"; return false }

func builderStringReceiver(node *ast.CallExpr, c *Context) (types.Object, bool) {
	_ = "STUB: not implemented"
	return *new(types.Object), false
}

func isStringBuilderType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func builderWritesAreConst(obj types.Object, strCall *ast.CallExpr, c *Context) bool {
	_ = "STUB: not implemented"
	return false
}

func isEmptyCompositeLit(e ast.Expr) bool { _ = "STUB: not implemented"; return false }

func TryResolve(n ast.Node, c *Context) bool { _ = "STUB: not implemented"; return false }
