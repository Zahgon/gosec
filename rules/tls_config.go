package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

func NewModernTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewIntermediateTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewOldTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
