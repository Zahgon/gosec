package rules

import (
	"go/ast"
	"regexp"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type secretPattern struct {
	name   string
	regexp *regexp.Regexp
}

type entropyCacheKey string

type secretPatternCacheKey string

var secretsPatterns = [...]secretPattern{
	{
		name:   "RSA private key",
		regexp: regexp.MustCompile(`-----BEGIN RSA PRIVATE KEY-----`),
	},
	{
		name:   "SSH (DSA) private key",
		regexp: regexp.MustCompile(`-----BEGIN DSA PRIVATE KEY-----`),
	},
	{
		name:   "SSH (EC) private key",
		regexp: regexp.MustCompile(`-----BEGIN EC PRIVATE KEY-----`),
	},
	{
		name:   "PGP private key block",
		regexp: regexp.MustCompile(`-----BEGIN PGP PRIVATE KEY BLOCK-----`),
	},
	{
		name:   "Slack Token",
		regexp: regexp.MustCompile(`xox[pborsa]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32}`),
	},
	{
		name:   "AWS API Key",
		regexp: regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	},
	{
		name:   "AWS Temporary Access Key",
		regexp: regexp.MustCompile(`ASIA[0-9A-Z]{16}`),
	},
	{
		name:   "Amazon MWS Auth Token",
		regexp: regexp.MustCompile(`amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
	},
	{
		name:   "AWS AppSync GraphQL Key",
		regexp: regexp.MustCompile(`da2-[a-z0-9]{26}`),
	},
	{
		name:   "GitHub personal access token",
		regexp: regexp.MustCompile(`ghp_[a-zA-Z0-9]{36}`),
	},
	{
		name:   "GitHub fine-grained access token",
		regexp: regexp.MustCompile(`github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59}`),
	},
	{
		name:   "GitHub action temporary token",
		regexp: regexp.MustCompile(`ghs_[a-zA-Z0-9]{36}`),
	},
	{
		name:   "Google API Key",
		regexp: regexp.MustCompile(`AIza[0-9A-Za-z\-_]{35}`),
	},

	{
		name:   "Google Cloud Platform OAuth",
		regexp: regexp.MustCompile(`[0-9]+-[0-9A-Za-z_]{32}\.apps\.googleusercontent\.com`),
	},

	{
		name:   "Google (GCP) Service-account",
		regexp: regexp.MustCompile(`"type": "service_account"`),
	},

	{
		name:   "Google OAuth Access Token",
		regexp: regexp.MustCompile(`ya29\.[0-9A-Za-z\-_]+`),
	},

	{
		name:   "Generic API Key",
		regexp: regexp.MustCompile(`[aA][pP][iI]_?[kK][eE][yY].*[''|"][0-9a-zA-Z]{32,45}[''|"]`),
	},
	{
		name:   "Generic Secret",
		regexp: regexp.MustCompile(`[sS][eE][cC][rR][eE][tT].*[''|"][0-9a-zA-Z]{32,45}[''|"]`),
	},
	{
		name:   "Heroku API Key",
		regexp: regexp.MustCompile(`[hH][eE][rR][oO][kK][uU].*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`),
	},
	{
		name:   "MailChimp API Key",
		regexp: regexp.MustCompile(`[0-9a-f]{32}-us[0-9]{1,2}`),
	},
	{
		name:   "Mailgun API Key",
		regexp: regexp.MustCompile(`key-[0-9a-zA-Z]{32}`),
	},
	{
		name:   "Password in URL",
		regexp: regexp.MustCompile(`[a-zA-Z]{3,10}://[a-zA-Z0-9\.\-\_\+]{1,64}:[a-zA-Z0-9\.\-\_\!\$\%\&\*\+\=\^\(\)]{1,128}@[a-zA-Z0-9\.\-\_]+(:[0-9]+)?(/[^"'\s]*)?(["'\s]|$)`),
	},
	{
		name:   "Slack Webhook",
		regexp: regexp.MustCompile(`https://hooks\.slack\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`),
	},
	{
		name:   "Stripe API Key",
		regexp: regexp.MustCompile(`sk_live_[0-9a-zA-Z]{24}`),
	},
	{
		name:   "Stripe Restricted API Key",
		regexp: regexp.MustCompile(`rk_live_[0-9a-zA-Z]{24}`),
	},
	{
		name:   "Square Access Token",
		regexp: regexp.MustCompile(`sq0atp-[0-9A-Za-z\-_]{22}`),
	},
	{
		name:   "Square OAuth Secret",
		regexp: regexp.MustCompile(`sq0csp-[0-9A-Za-z\-_]{43}`),
	},
	{
		name:   "Telegram Bot API Key",
		regexp: regexp.MustCompile(`[0-9]+:AA[0-9A-Za-z\-_]{33}`),
	},
	{
		name:   "Twilio API Key",
		regexp: regexp.MustCompile(`SK[0-9a-fA-F]{32}`),
	},
	{
		name:   "Twitter Access Token",
		regexp: regexp.MustCompile(`[tT][wW][iI][tT][tT][eE][rR].*[1-9][0-9]+-[0-9a-zA-Z]{40}`),
	},
	{
		name:   "Twitter OAuth",
		regexp: regexp.MustCompile(`[tT][wW][iI][tT][tT][eE][rR].*[''|"][0-9a-zA-Z]{35,44}[''|"]`),
	},

	{
		name:   "GitHub personal access token",
		regexp: regexp.MustCompile(`ghp_[a-zA-Z0-9]{36}`),
	},
	{
		name:   "GitHub fine-grained access token",
		regexp: regexp.MustCompile(`github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59}`),
	},
	{
		name:   "GitHub action temporary token",
		regexp: regexp.MustCompile(`ghs_[a-zA-Z0-9]{36}`),
	},
}

type credentials struct {
	issue.MetaData
	pattern          *regexp.Regexp
	entropyThreshold float64
	perCharThreshold float64
	truncate         int
	ignoreEntropy    bool
	minEntropyLength int
}

func truncate(s string, n int) string { _ = "STUB: not implemented"; return "" }

func (r *credentials) isHighEntropyString(str string) bool { _ = "STUB: not implemented"; return false }

type secretResult struct {
	ok          bool
	patternName string
}

func (r *credentials) isSecretPattern(str string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (r *credentials) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *credentials) matchAssign(assign *ast.AssignStmt, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *credentials) matchValueSpec(valueSpec *ast.ValueSpec, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *credentials) matchEqualityCheck(binaryExpr *ast.BinaryExpr, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *credentials) matchCompositeLit(lit *ast.CompositeLit, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHardcodedCredentials(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
