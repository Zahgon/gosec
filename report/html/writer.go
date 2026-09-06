package html

import (
	_ "embed"
	"io"

	"github.com/securego/gosec/v2"
)

//go:embed template.html
var templateContent string

func WriteReport(w io.Writer, data *gosec.ReportInfo) error { _ = "STUB: not implemented"; return nil }
