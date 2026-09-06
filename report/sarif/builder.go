package sarif

func NewReport(version string, schema string) *Report { _ = "STUB: not implemented"; return nil }

func (r *Report) WithRuns(runs ...*Run) *Report { _ = "STUB: not implemented"; return nil }

func NewMultiformatMessageString(text string) *MultiformatMessageString {
	_ = "STUB: not implemented"
	return nil
}

func NewRun(tool *Tool) *Run { _ = "STUB: not implemented"; return nil }

func (r *Run) WithTaxonomies(taxonomies ...*ToolComponent) *Run {
	_ = "STUB: not implemented"
	return nil
}

func (r *Run) WithResults(results ...*Result) *Run { _ = "STUB: not implemented"; return nil }

func NewArtifactLocation(uri string) *ArtifactLocation { _ = "STUB: not implemented"; return nil }

func NewRegion(startLine int, endLine int, startColumn int, endColumn int, sourceLanguage string) *Region {
	_ = "STUB: not implemented"
	return nil
}

func (r *Region) WithSnippet(snippet *ArtifactContent) *Region {
	_ = "STUB: not implemented"
	return nil
}

func NewArtifactContent(text string) *ArtifactContent { _ = "STUB: not implemented"; return nil }

func NewTool(driver *ToolComponent) *Tool { _ = "STUB: not implemented"; return nil }

func NewResult(ruleID string, ruleIndex int, level Level, message string, suppressions []*Suppression, autofix string) *Result {
	_ = "STUB: not implemented"
	return nil
}

func NewMessage(text string) *Message { _ = "STUB: not implemented"; return nil }

func (r *Result) WithLocations(locations ...*Location) *Result {
	_ = "STUB: not implemented"
	return nil
}

func NewLocation(physicalLocation *PhysicalLocation) *Location {
	_ = "STUB: not implemented"
	return nil
}

func NewPhysicalLocation(artifactLocation *ArtifactLocation, region *Region) *PhysicalLocation {
	_ = "STUB: not implemented"
	return nil
}

func NewToolComponent(name string, version string, informationURI string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithLanguage(language string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithSemanticVersion(semanticVersion string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithReleaseDateUtc(releaseDateUtc string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithDownloadURI(downloadURI string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithOrganization(organization string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithShortDescription(shortDescription *MultiformatMessageString) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithIsComprehensive(isComprehensive bool) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithMinimumRequiredLocalizedDataSemanticVersion(minimumRequiredLocalizedDataSemanticVersion string) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithTaxa(taxa ...*ReportingDescriptor) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithSupportedTaxonomies(supportedTaxonomies ...*ToolComponentReference) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func (t *ToolComponent) WithRules(rules ...*ReportingDescriptor) *ToolComponent {
	_ = "STUB: not implemented"
	return nil
}

func NewToolComponentReference(name string) *ToolComponentReference {
	_ = "STUB: not implemented"
	return nil
}

func NewSuppression(kind string, justification string) *Suppression {
	_ = "STUB: not implemented"
	return nil
}
