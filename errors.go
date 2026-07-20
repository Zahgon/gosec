package gosec

type Error struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Err    string `json:"error"`
}

func NewError(line, column int, err string) *Error { _ = "STUB: not implemented"; return nil }

func sortErrors(allErrors map[string][]Error) { _ = "STUB: not implemented"; return }
