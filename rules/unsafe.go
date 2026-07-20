package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type usingUnsafe struct {
	callListRule
}

func NewUsingUnsafe(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
