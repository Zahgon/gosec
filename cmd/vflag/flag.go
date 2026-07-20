package vflag

type ValidatedFlag struct {
	Value string
}

func (f *ValidatedFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *ValidatedFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }
