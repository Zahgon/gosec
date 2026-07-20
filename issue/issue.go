package issue

import (
	"go/ast"
	"go/token"
	"os"

	"github.com/securego/gosec/v2/cwe"
)

type Score int

const (
	Low Score = iota

	Medium

	High
)

const SnippetOffset = 1

func GetCweByRule(id string) *cwe.Weakness { _ = "STUB: not implemented"; return nil }

var ruleToCWE = map[string]string{
	"G101": "798",
	"G102": "200",
	"G103": "242",
	"G104": "703",
	"G106": "322",
	"G107": "88",
	"G108": "200",
	"G109": "190",
	"G110": "409",
	"G111": "22",
	"G112": "400",
	"G113": "444",
	"G707": "93",
	"G708": "94",
	"G709": "502",
	"G114": "676",
	"G115": "190",
	"G116": "838",
	"G117": "499",
	"G118": "400",
	"G119": "200",
	"G120": "400",
	"G121": "346",
	"G122": "367",
	"G123": "295",
	"G124": "614",
	"G201": "89",
	"G202": "89",
	"G203": "79",
	"G204": "78",
	"G301": "276",
	"G302": "276",
	"G303": "377",
	"G304": "22",
	"G305": "22",
	"G306": "276",
	"G307": "276",
	"G401": "328",
	"G402": "295",
	"G403": "310",
	"G404": "338",
	"G405": "327",
	"G406": "328",
	"G407": "1204",
	"G408": "287",
	"G501": "327",
	"G502": "327",
	"G503": "327",
	"G504": "327",
	"G505": "327",
	"G506": "327",
	"G507": "327",
	"G601": "118",
	"G602": "118",
	"G701": "89",
	"G702": "78",
	"G703": "22",
	"G704": "918",
	"G705": "79",
	"G706": "117",
	"G710": "601",
}

type Issue struct {
	Severity     Score             `json:"severity"`
	Confidence   Score             `json:"confidence"`
	Cwe          *cwe.Weakness     `json:"cwe"`
	RuleID       string            `json:"rule_id"`
	What         string            `json:"details"`
	File         string            `json:"file"`
	Code         string            `json:"code"`
	Line         string            `json:"line"`
	Col          string            `json:"column"`
	NoSec        bool              `json:"nosec"`
	Suppressions []SuppressionInfo `json:"suppressions"`
	Autofix      string            `json:"autofix,omitempty"`
}

type SuppressionInfo struct {
	Kind          string `json:"kind"`
	Justification string `json:"justification"`
}

func (i *Issue) FileLocation() string { _ = "STUB: not implemented"; return "" }

type MetaData struct {
	RuleID     string
	Severity   Score
	Confidence Score
	What       string
}

func NewMetaData(id, what string, severity, confidence Score) MetaData {
	_ = "STUB: not implemented"
	return *new(MetaData)
}

func (m MetaData) ID() string { _ = "STUB: not implemented"; return "" }

func (c Score) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c Score) String() string { _ = "STUB: not implemented"; return "" }

func CodeSnippet(file *os.File, start int64, end int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func codeSnippetStartLine(node ast.Node, fobj *token.File) int64 {
	_ = "STUB: not implemented"
	return 0
}

func codeSnippetEndLine(node ast.Node, fobj *token.File) int64 { _ = "STUB: not implemented"; return 0 }

func New(fobj *token.File, node ast.Node, ruleID, desc string, severity, confidence Score) *Issue {
	_ = "STUB: not implemented"
	return nil
}

func (i *Issue) WithSuppressions(suppressions []SuppressionInfo) *Issue {
	_ = "STUB: not implemented"
	return nil
}

func GetLine(fobj *token.File, node ast.Node) string { _ = "STUB: not implemented"; return "" }
