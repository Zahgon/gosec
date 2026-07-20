package testutils

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type MockVisitor struct {
	Context  *gosec.Context
	Callback func(n ast.Node, ctx *gosec.Context) bool
}

func NewMockVisitor() *MockVisitor { _ = "STUB: not implemented"; return nil }

func (v *MockVisitor) Visit(n ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}
