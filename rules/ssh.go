package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type sshHostKey struct {
	callListRule
}

func NewSSHHostKey(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
