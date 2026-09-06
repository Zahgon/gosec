package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type blocklistedImport struct {
	issue.MetaData
	Blocklisted map[string]string
}

func unquote(original string) string { _ = "STUB: not implemented"; return "" }

func (r *blocklistedImport) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBlocklistedImports(id string, _ gosec.Config, blocklist map[string]string) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportMD5(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportDES(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportRC4(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportCGI(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportSHA1(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportMD4(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewBlocklistedImportRIPEMD160(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
