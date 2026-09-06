package gosec

import (
	"go/ast"
)

const vendorPath = "vendor/"

type set map[string]bool

type CallList map[string]set

func NewCallList() CallList { _ = "STUB: not implemented"; return *new(CallList) }

func (c CallList) AddAll(selector string, idents ...string) { _ = "STUB: not implemented"; return }

func (c CallList) Add(selector, ident string) { _ = "STUB: not implemented"; return }

func (c CallList) Contains(selector, ident string) bool { _ = "STUB: not implemented"; return false }

func (c CallList) ContainsPointer(selector, indent string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c CallList) ContainsPkgCallExpr(n ast.Node, ctx *Context, stripVendor bool) *ast.CallExpr {
	_ = "STUB: not implemented"
	return nil
}

func (c CallList) ContainsCallExpr(n ast.Node, ctx *Context) *ast.CallExpr {
	_ = "STUB: not implemented"
	return nil
}
