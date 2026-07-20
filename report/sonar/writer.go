package sonar

import (
	"io"

	"github.com/securego/gosec/v2"
)

func WriteReport(w io.Writer, data *gosec.ReportInfo, rootPaths []string) error {
	_ = "STUB: not implemented"
	return nil
}
