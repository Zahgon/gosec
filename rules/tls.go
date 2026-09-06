//go:generate tlsconfig

package rules

import (
	"crypto/tls"
	"go/ast"
	"go/types"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type insecureConfigTLS struct {
	issue.MetaData
	MinVersion       int64
	MaxVersion       int64
	requiredType     string
	goodCiphers      []string
	actualMinVersion int64
	actualMaxVersion int64
	minVersionSet    bool
	maxVersionSet    bool
}

var tlsVersionMap = map[string]int64{
	"VersionTLS10": tls.VersionTLS10,
	"VersionTLS11": tls.VersionTLS11,
	"VersionTLS12": tls.VersionTLS12,
	"VersionTLS13": tls.VersionTLS13,
}

func (t *insecureConfigTLS) mapVersion(version string) int64 { _ = "STUB: not implemented"; return 0 }

func (t *insecureConfigTLS) processTLSCipherSuites(n ast.Node, c *gosec.Context) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (t *insecureConfigTLS) resolveTLSVersion(expr ast.Expr, c *gosec.Context) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (t *insecureConfigTLS) resolveBoolConst(expr ast.Expr, c *gosec.Context) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (t *insecureConfigTLS) processTLSConfVal(key ast.Expr, value ast.Expr, c *gosec.Context) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (t *insecureConfigTLS) processTLSConf(n ast.Node, c *gosec.Context) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (t *insecureConfigTLS) findDefinition(obj types.Object, c *gosec.Context) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (t *insecureConfigTLS) isSafeDefault() bool { _ = "STUB: not implemented"; return false }

func (t *insecureConfigTLS) checkVersion(n ast.Node, c *gosec.Context) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (t *insecureConfigTLS) resetVersion() { _ = "STUB: not implemented"; return }

func (t *insecureConfigTLS) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
