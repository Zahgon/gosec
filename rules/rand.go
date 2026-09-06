package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type weakRand struct {
	callListRule
}

func NewWeakRandCheck(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
