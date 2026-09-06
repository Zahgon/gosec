package gosec

import (
	"go/ast"
	"reflect"

	"github.com/securego/gosec/v2/issue"
)

type Rule interface {
	ID() string
	Match(ast.Node, *Context) (*issue.Issue, error)
}

type RuleBuilder func(id string, c Config) (Rule, []ast.Node)

type RuleSet struct {
	Rules             map[reflect.Type][]Rule
	RuleSuppressedMap map[string]bool
}

func NewRuleSet() RuleSet { _ = "STUB: not implemented"; return *new(RuleSet) }

func (r RuleSet) Register(rule Rule, isSuppressed bool, nodes ...ast.Node) {
	_ = "STUB: not implemented"
	return
}

func (r RuleSet) RegisteredFor(n ast.Node) []Rule { _ = "STUB: not implemented"; return nil }

func (r RuleSet) IsRuleSuppressed(ruleID string) bool { _ = "STUB: not implemented"; return false }
