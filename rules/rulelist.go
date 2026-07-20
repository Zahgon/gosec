package rules

import "github.com/securego/gosec/v2"

type RuleDefinition struct {
	ID          string
	Description string
	Create      gosec.RuleBuilder
}

type RuleList struct {
	Rules          map[string]RuleDefinition
	RuleSuppressed map[string]bool
}

func (rl RuleList) RulesInfo() (map[string]gosec.RuleBuilder, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RuleFilter func(string) bool

func NewRuleFilter(action bool, ruleIDs ...string) RuleFilter {
	_ = "STUB: not implemented"
	return *new(RuleFilter)
}

func Generate(trackSuppressions bool, filters ...RuleFilter) RuleList {
	_ = "STUB: not implemented"
	return *new(RuleList)
}
