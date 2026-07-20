package sarif

import (
	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/cwe"
	"github.com/securego/gosec/v2/issue"
)

func GenerateReport(rootPaths []string, data *gosec.ReportInfo) (*Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addRuleInOrder(rules []*ReportingDescriptor, rule *ReportingDescriptor) ([]*ReportingDescriptor, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func parseSarifRule(i *issue.Issue) *ReportingDescriptor { _ = "STUB: not implemented"; return nil }

func buildSarifReportingDescriptorRelationship(weakness *cwe.Weakness) *ReportingDescriptorRelationship {
	_ = "STUB: not implemented"
	return nil
}

func buildCWETaxonomy(taxa []*ReportingDescriptor) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func parseSarifTaxon(weakness *cwe.Weakness) *ReportingDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func parseSemanticVersion(version string) string { _ = "STUB: not implemented"; return "" }

func buildSarifDriver(rules []*ReportingDescriptor, gosecVersion string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func uuid3(value string) string { _ = "STUB: not implemented"; return "" }

func parseSarifLocation(i *issue.Issue, rootPaths []string) (*Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSarifArtifactLocation(i *issue.Issue, rootPaths []string) *ArtifactLocation {
	_ = "STUB: not implemented"
	return nil
}

func parseSarifRegion(i *issue.Issue) (*Region, error) { _ = "STUB: not implemented"; return nil, nil }

func getSarifLevel(s string) Level { _ = "STUB: not implemented"; return *new(Level) }

func buildSarifSuppressions(suppressions []issue.SuppressionInfo) []*Suppression {
	_ = "STUB: not implemented"
	return nil
}
