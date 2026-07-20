package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type weakCryptoUsage struct {
	callListRule
}

func NewUsesWeakCryptographyHash(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewUsesWeakCryptographyEncryption(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewUsesWeakDeprecatedCryptographyHash(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
