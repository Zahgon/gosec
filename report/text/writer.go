package text

import (
	_ "embed"
	"io"
	"text/template"

	"github.com/gookit/color"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

var (
	errorTheme   = color.New(color.FgLightWhite, color.BgRed)
	warningTheme = color.New(color.FgBlack, color.BgYellow)
	defaultTheme = color.New(color.FgWhite, color.BgBlack)

	//go:embed template.txt
	templateContent string
)

func WriteReport(w io.Writer, data *gosec.ReportInfo, enableColor bool) error {
	_ = "STUB: not implemented"
	return nil
}

func plainTextFuncMap(enableColor bool) template.FuncMap {
	_ = "STUB: not implemented"
	return *new(template.FuncMap)
}

func highlight(t string, s issue.Score, ignored bool) string { _ = "STUB: not implemented"; return "" }

func printCodeSnippet(issue *issue.Issue) string { _ = "STUB: not implemented"; return "" }

func parseLine(line string) (int, int) { _ = "STUB: not implemented"; return 0, 0 }
