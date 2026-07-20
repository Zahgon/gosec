package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type filePermissions struct {
	issue.MetaData
	mode  int64
	pkgs  []string
	calls []string
}

func getConfiguredMode(conf map[string]interface{}, configKey string, defaultMode int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func modeIsSubset(subset int64, superset int64) bool { _ = "STUB: not implemented"; return false }

func (r *filePermissions) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isOsPerm(n ast.Node) bool { _ = "STUB: not implemented"; return false }

func NewWritePerms(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewFilePerms(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

func NewMkdirPerms(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

type osCreatePermissions struct {
	issue.MetaData
	mode  int64
	pkgs  []string
	calls []string
}

const defaultOsCreateMode = 0o666

func (r *osCreatePermissions) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewOsCreatePerms(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
