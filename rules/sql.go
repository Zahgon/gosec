package rules

import (
	"go/ast"
	"regexp"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type sqlStatement struct {
	issue.MetaData
	gosec.CallList

	patterns []*regexp.Regexp
}

var sqlCallIdents = map[string]map[string]int{
	"*database/sql.Conn": {
		"ExecContext":     1,
		"QueryContext":    1,
		"QueryRowContext": 1,
		"PrepareContext":  1,
	},
	"*database/sql.DB": {
		"Exec":            0,
		"ExecContext":     1,
		"Query":           0,
		"QueryContext":    1,
		"QueryRow":        0,
		"QueryRowContext": 1,
		"Prepare":         0,
		"PrepareContext":  1,
	},
	"*database/sql.Tx": {
		"Exec":            0,
		"ExecContext":     1,
		"Query":           0,
		"QueryContext":    1,
		"QueryRow":        0,
		"QueryRowContext": 1,
		"Prepare":         0,
		"PrepareContext":  1,
	},
}

var (
	sqlRegexp       = regexp.MustCompile("(?i)(SELECT|DELETE|INSERT|UPDATE|INTO|FROM|WHERE)( |\n|\r|\t)")
	sqlFormatRegexp = regexp.MustCompile("%[^bdoxXfFp]")
)

func findQueryArg(call *ast.CallExpr, ctx *gosec.Context) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (s *sqlStatement) MatchPatterns(str string) bool { _ = "STUB: not implemented"; return false }

type sqlStrConcat struct {
	sqlStatement
}

func (s *sqlStrConcat) findInjectionInBranch(ctx *gosec.Context, branch []ast.Expr) *ast.BinaryExpr {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlStrConcat) checkQuery(call *ast.CallExpr, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlStrConcat) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSQLStrConcat(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}

type sqlStrFormat struct {
	gosec.CallList
	sqlStatement
	fmtCalls      gosec.CallList
	noIssue       gosec.CallList
	noIssueQuoted gosec.CallList
}

func (s *sqlStrFormat) checkQuery(call *ast.CallExpr, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlStrFormat) checkFormatting(n ast.Node, ctx *gosec.Context) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlStrFormat) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSQLStrFormat(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
