package sonar

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartColumn int `json:"startColumn,omitempty"`
	EtartColumn int `json:"endColumn,omitempty"`
}

type Location struct {
	Message   string     `json:"message"`
	FilePath  string     `json:"filePath"`
	TextRange *TextRange `json:"textRange,omitempty"`
}

type Issue struct {
	RuleID             string      `json:"ruleId"`
	PrimaryLocation    *Location   `json:"primaryLocation"`
	EffortMinutes      int         `json:"effortMinutes"`
	SecondaryLocations []*Location `json:"secondaryLocations,omitempty"`
}

type Impact struct {
	SoftwareQuality string `json:"softwareQuality"`
	Severity        string `json:"severity"`
}

type Rule struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	EngineID           string    `json:"engineId"`
	CleanCodeAttribute string    `json:"cleanCodeAttribute,omitempty"`
	Impacts            []*Impact `json:"impacts"`
}

type Report struct {
	Rules  []*Rule  `json:"rules"`
	Issues []*Issue `json:"issues"`
}
