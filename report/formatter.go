package report

import (
	"io"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type Format int

const (
	ReportText Format = iota

	ReportJSON

	ReportCSV

	ReportJUnitXML

	ReportSARIF
)

func CreateReport(w io.Writer, format string, enableColor bool, rootPaths []string, data *gosec.ReportInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func filterOutSuppressedIssues(issues []*issue.Issue) []*issue.Issue {
	_ = "STUB: not implemented"
	return nil
}
