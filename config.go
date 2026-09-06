package gosec

import (
	"io"
)

const (
	Globals = "global"

	ExcludeRulesKey = "exclude-rules"
)

type GlobalOption string

const (
	Nosec GlobalOption = "nosec"

	ShowIgnored GlobalOption = "show-ignored"

	Audit GlobalOption = "audit"

	NoSecAlternative GlobalOption = "#nosec"

	ExcludeRules GlobalOption = "exclude"

	IncludeRules GlobalOption = "include"

	SSA GlobalOption = "ssa"

	NoSecRequireRules GlobalOption = "nosec-require-rules"

	NoSecRequireJustification GlobalOption = "nosec-require-justification"
)

func NoSecTag(tag string) string { _ = "STUB: not implemented"; return "" }

type Config map[string]interface{}

func NewConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func (c Config) keyToGlobalOptions(key string) GlobalOption {
	_ = "STUB: not implemented"
	return *new(GlobalOption)
}

func (c Config) convertGlobals() { _ = "STUB: not implemented"; return }

func (c Config) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c Config) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c Config) Get(section string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) Set(section string, value interface{}) { _ = "STUB: not implemented"; return }

func (c Config) GetGlobal(option GlobalOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c Config) SetGlobal(option GlobalOption, value string) { _ = "STUB: not implemented"; return }

func (c Config) IsGlobalEnabled(option GlobalOption) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c Config) GetExcludeRules() ([]PathExcludeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) SetExcludeRules(rules []PathExcludeRule) { _ = "STUB: not implemented"; return }
