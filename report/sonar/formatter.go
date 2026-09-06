package sonar

import (
	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

const (
	EffortMinutes = 5

	sonarEngineID           = "gosec"
	sonarSoftwareQuality    = "SECURITY"
	sonarCleanCodeAttribute = "TRUSTWORTHY"
)

func GenerateReport(rootPaths []string, data *gosec.ReportInfo) (*Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFilePath(issue *issue.Issue, rootPaths []string) string {
	_ = "STUB: not implemented"
	return ""
}

func parseTextRange(issue *issue.Issue) (*TextRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getImpactSeverity(s string) string { _ = "STUB: not implemented"; return "" }

func mergeRuleImpacts(existing []*Impact, severity string) []*Impact {
	_ = "STUB: not implemented"
	return nil
}

func compareImpactSeverity(a string, b string) int { _ = "STUB: not implemented"; return 0 }
