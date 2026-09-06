package analyzers

import (
	"golang.org/x/tools/go/analysis"

	"github.com/securego/gosec/v2/taint"
)

type AnalyzerDefinition struct {
	ID          string
	Description string
	Create      AnalyzerBuilder
}

type AnalyzerBuilder func(id string, description string) *analysis.Analyzer

var (
	SQLInjectionRule = taint.RuleInfo{
		ID:          "G701",
		Description: "SQL injection via string concatenation",
		Severity:    "HIGH",
		CWE:         "CWE-89",
	}

	CommandInjectionRule = taint.RuleInfo{
		ID:          "G702",
		Description: "Command injection via user input",
		Severity:    "CRITICAL",
		CWE:         "CWE-78",
	}

	PathTraversalRule = taint.RuleInfo{
		ID:          "G703",
		Description: "Path traversal via user input",
		Severity:    "HIGH",
		CWE:         "CWE-22",
	}

	SSRFRule = taint.RuleInfo{
		ID:          "G704",
		Description: "SSRF via user-controlled URL",
		Severity:    "HIGH",
		CWE:         "CWE-918",
	}

	XSSRule = taint.RuleInfo{
		ID:          "G705",
		Description: "XSS via unescaped user input",
		Severity:    "MEDIUM",
		CWE:         "CWE-79",
	}

	LogInjectionRule = taint.RuleInfo{
		ID:          "G706",
		Description: "Log injection via user input",
		Severity:    "LOW",
		CWE:         "CWE-117",
	}

	SMTPInjectionRule = taint.RuleInfo{
		ID:          "G707",
		Description: "SMTP command/header injection via user input",
		Severity:    "HIGH",
		CWE:         "CWE-93",
	}

	SSTIRule = taint.RuleInfo{
		ID:          "G708",
		Description: "Server-side template injection via text/template",
		Severity:    "CRITICAL",
		CWE:         "CWE-94",
	}

	UnsafeDeserializationRule = taint.RuleInfo{
		ID:          "G709",
		Description: "Unsafe deserialization of untrusted data",
		Severity:    "HIGH",
		CWE:         "CWE-502",
	}

	OpenRedirectRule = taint.RuleInfo{
		ID:          "G710",
		Description: "Open redirect: user-controlled URL flows into http.Redirect",
		Severity:    "MEDIUM",
		CWE:         "CWE-601",
	}

	FormParsingLimitRule = taint.RuleInfo{
		ID:          "G120",
		Description: "Unbounded multipart form parsing can cause memory exhaustion",
		Severity:    "MEDIUM",
		CWE:         "CWE-400",
	}
)

type AnalyzerList struct {
	Analyzers          map[string]AnalyzerDefinition
	AnalyzerSuppressed map[string]bool
}

func (al *AnalyzerList) AnalyzersInfo() (map[string]AnalyzerDefinition, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AnalyzerFilter func(string) bool

func NewAnalyzerFilter(action bool, analyzerIDs ...string) AnalyzerFilter {
	_ = "STUB: not implemented"
	return *new(AnalyzerFilter)
}

var defaultAnalyzers = []AnalyzerDefinition{
	{"G113", "HTTP request smuggling via conflicting headers or bare LF in body parsing", newRequestSmugglingAnalyzer},
	{"G115", "Type conversion which leads to integer overflow", newConversionOverflowAnalyzer},
	{"G118", "Context propagation failure leading to goroutine/resource leaks", newContextPropagationAnalyzer},
	{"G119", "Unsafe redirect policy may propagate sensitive headers", newRedirectHeaderPropagationAnalyzer},
	{"G120", "Unbounded form parsing in HTTP handlers can cause memory exhaustion", newFormParsingLimitAnalyzer},
	{"G121", "Unsafe CrossOriginProtection bypass patterns", newCORSBypassPatternAnalyzer},
	{"G122", "Filesystem TOCTOU race risk in filepath.Walk/WalkDir callbacks", newWalkSymlinkRaceAnalyzer},
	{"G123", "TLS resumption may bypass VerifyPeerCertificate when VerifyConnection is unset", newTLSResumptionVerifyPeerAnalyzer},
	{"G124", "Insecure HTTP cookie configuration missing Secure, HttpOnly, or SameSite attributes", newInsecureCookieAnalyzer},
	{"G602", "Possible slice bounds out of range", newSliceBoundsAnalyzer},
	{"G407", "Use of hardcoded IV/nonce for encryption", newHardCodedNonce},
	{"G408", "Stateful misuse of ssh.PublicKeyCallback leading to auth bypass", newSSHCallbackAnalyzer},
	{"G701", "SQL injection via taint analysis", newSQLInjectionAnalyzer},
	{"G702", "Command injection via taint analysis", newCommandInjectionAnalyzer},
	{"G703", "Path traversal via taint analysis", newPathTraversalAnalyzer},
	{"G704", "SSRF via taint analysis", newSSRFAnalyzer},
	{"G705", "XSS via taint analysis", newXSSAnalyzer},
	{"G706", "Log injection via taint analysis", newLogInjectionAnalyzer},
	{"G707", "SMTP command/header injection via taint analysis", newSMTPInjectionAnalyzer},
	{"G708", "Server-side template injection via taint analysis", newSSTIAnalyzer},
	{"G709", "Unsafe deserialization of untrusted data via taint analysis", newUnsafeDeserializationAnalyzer},
	{"G710", "Open redirect via taint analysis", newOpenRedirectAnalyzer},
}

func Generate(trackSuppressions bool, filters ...AnalyzerFilter) *AnalyzerList {
	_ = "STUB: not implemented"
	return nil
}

func DefaultTaintAnalyzers() []*analysis.Analyzer { _ = "STUB: not implemented"; return nil }
