package cwe

type Weakness struct {
	ID          string
	Name        string
	Description string
}

func (w *Weakness) SprintURL() string { _ = "STUB: not implemented"; return "" }

func (w *Weakness) SprintID() string { _ = "STUB: not implemented"; return "" }

func (w *Weakness) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
