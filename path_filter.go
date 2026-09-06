package gosec

import (
	"regexp"

	"github.com/securego/gosec/v2/issue"
)

type PathExcludeRule struct {
	Path  string   `json:"path"`
	Rules []string `json:"rules"`
}

type compiledPathRule struct {
	pathRegex  *regexp.Regexp
	ruleSet    map[string]bool
	excludeAll bool
	original   PathExcludeRule
}

type PathExclusionFilter struct {
	rules []compiledPathRule
}

func NewPathExclusionFilter(rules []PathExcludeRule) (*PathExclusionFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *PathExclusionFilter) ShouldExclude(filePath, ruleID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *PathExclusionFilter) FilterIssues(issues []*issue.Issue) ([]*issue.Issue, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func ParseCLIExcludeRules(input string) ([]PathExcludeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MergeExcludeRules(configRules, cliRules []PathExcludeRule) []PathExcludeRule {
	_ = "STUB: not implemented"
	return nil
}

func (f *PathExclusionFilter) String() string { _ = "STUB: not implemented"; return "" }
