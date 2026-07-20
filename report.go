package gosec

import (
	"github.com/securego/gosec/v2/issue"
)

type ReportInfo struct {
	Errors       map[string][]Error `json:"Golang errors"`
	Issues       []*issue.Issue
	Stats        *Metrics
	GosecVersion string
}

func NewReportInfo(issues []*issue.Issue, metrics *Metrics, errors map[string][]Error) *ReportInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReportInfo) WithVersion(version string) *ReportInfo { _ = "STUB: not implemented"; return nil }
