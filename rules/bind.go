package rules

import (
	"go/ast"
	"regexp"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type bindsToAllNetworkInterfaces struct {
	callListRule
	pattern *regexp.Regexp
}

func (r *bindsToAllNetworkInterfaces) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBindsToAllNetworkInterfaces(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
