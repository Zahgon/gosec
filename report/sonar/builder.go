package sonar

func NewLocation(message string, filePath string, textRange *TextRange) *Location {
	_ = "STUB: not implemented"
	return nil
}

func NewTextRange(startLine int, endLine int) *TextRange { _ = "STUB: not implemented"; return nil }

func NewIssue(ruleID string, primaryLocation *Location, effortMinutes int) *Issue {
	_ = "STUB: not implemented"
	return nil
}

func NewImpact(softwareQuality string, severity string) *Impact {
	_ = "STUB: not implemented"
	return nil
}

func NewRule(id string, name string, description string, engineID string, cleanCodeAttribute string, impacts []*Impact) *Rule {
	_ = "STUB: not implemented"
	return nil
}
