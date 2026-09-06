package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/analyzers"
	"github.com/securego/gosec/v2/autofix"
	"github.com/securego/gosec/v2/cmd/vflag"
	"github.com/securego/gosec/v2/issue"
	"github.com/securego/gosec/v2/rules"
)

const (
	usageText = `
gosec - Golang security checker

gosec analyzes Go source code to look for common programming mistakes that
can lead to security problems.

VERSION: %s
GIT TAG: %s
BUILD DATE: %s

USAGE:

	# Check a single package
	$ gosec $GOPATH/src/github.com/example/project

	# Check all packages under the current directory and save results in
	# json format.
	$ gosec -fmt=json -out=results.json ./...

	# Run a specific set of rules (by default all rules will be run):
	$ gosec -include=G101,G203,G401  ./...

	# Run all rules except the provided
	$ gosec -exclude=G101 $GOPATH/src/github.com/example/project/...

	# Exclude specific rules from specific paths
	$ gosec --exclude-rules="cmd/.*:G204,G304" ./...

	# Exclude all rules from scripts directory
	$ gosec --exclude-rules="scripts/.*:*" ./...
`

	aiAPIKeyEnv   = "GOSEC_AI_API_KEY"
	aiProviderEnv = "GOSEC_AI_PROVIDER"
	aiBaseURLEnv  = "GOSEC_AI_BASE_URL"

	exitSuccess = 0
	exitFailure = 1
)

type arrayFlags []string

func (a *arrayFlags) String() string { _ = "STUB: not implemented"; return "" }

func (a *arrayFlags) Set(value string) error { _ = "STUB: not implemented"; return nil }

var (
	flagIgnoreNoSec = flag.Bool("nosec", false, "Ignores #nosec comments when set")

	flagExcludeRules = flag.String("exclude-rules", "",
		`Path-based rule exclusions. Format: "path:rule1,rule2;path2:rule3"
Example: "cmd/.*:G204,G304;test/.*:G101"
Use "*" to exclude all rules for a path: "scripts/.*:*"`)

	flagShowIgnored = flag.Bool("show-ignored", false, "If enabled, ignored issues are printed")

	flagFormat = flag.String("fmt", "text", "Set output format. Valid options are: json, yaml, csv, junit-xml, html, sonarqube, golint, sarif or text")

	flagAlternativeNoSec = flag.String("nosec-tag", "", "Set an alternative string for #nosec. Some examples: #dontanalyze, #falsepositive")

	flagNoSecRequireRules = flag.Bool("nosec-require-rules", false, "Require at least one rule ID (e.g. G401) in every #nosec / //gosec:disable annotation")

	flagNoSecRequireJustification = flag.Bool("nosec-require-justification", false, "Require a `-- justification` in every #nosec / //gosec:disable annotation")

	flagEnableAudit = flag.Bool("enable-audit", false, "Enable audit mode")

	flagOutput = flag.String("out", "", "Set output file for results")

	flagConfig = flag.String("conf", "", "Path to optional config file")

	flagQuiet = flag.Bool("quiet", false, "Only show output when errors are found")

	flagRulesInclude = flag.String("include", "", "Comma separated list of rules IDs to include. (see rule list)")

	flagRulesExclude = vflag.ValidatedFlag{}

	flagExcludeGenerated = flag.Bool("exclude-generated", false, "Exclude generated files")

	flagLogfile = flag.String("log", "", "Log messages to file rather than stderr")

	flagSortIssues = flag.Bool("sort", true, "Sort issues by severity")

	flagBuildTags = flag.String("tags", "", "Comma separated list of build tags")

	flagSeverity = flag.String("severity", "low", "Filter out the issues with a lower severity than the given value. Valid options are: low, medium, high")

	flagConfidence = flag.String("confidence", "low", "Filter out the issues with a lower confidence than the given value. Valid options are: low, medium, high")

	flagConcurrency = flag.Int("concurrency", runtime.NumCPU(), "Concurrency value")

	flagNoFail = flag.Bool("no-fail", false, "Do not fail the scanning, even if issues were found")

	flagScanTests = flag.Bool("tests", false, "Scan tests files")

	flagVersion = flag.Bool("version", false, "Print version and quit with exit code 0")

	flagStdOut = flag.Bool("stdout", false, "Stdout the results as well as write it in the output file")

	flagColor = flag.Bool("color", true, "Prints the text format report with colorization when it goes in the stdout")

	flagRecursive = flag.Bool("r", false, "Appends \"./...\" to the target dir.")

	flagVerbose = flag.String("verbose", "", "Overrides the output format when stdout the results while saving them in the output file.\nValid options are: json, yaml, csv, junit-xml, html, sonarqube, golint, sarif or text")

	flagTrackSuppressions = flag.Bool("track-suppressions", false, "Output suppression information, including its kind and justification")

	flagTerse = flag.Bool("terse", false, "Shows only the results and summary")

	flagAiAPIProvider = flag.String("ai-api-provider", "", autofix.AIProviderFlagHelp)

	flagAiAPIKey = flag.String("ai-api-key", "", "Key to access the AI API")

	flagAiBaseURL = flag.String("ai-base-url", "", "Base URL for AI API (e.g., for OpenAI-compatible services)")

	flagAiSkipSSL = flag.Bool("ai-skip-ssl", false, "Skip SSL certificate verification for AI API")

	flagDirsExclude arrayFlags

	logger *log.Logger
)

func usage() { _ = "STUB: not implemented"; return }

func loadConfig(configFile string) (gosec.Config, error) {
	_ = "STUB: not implemented"
	return *new(gosec.Config), nil
}

func loadRules(include, exclude string) rules.RuleList {
	_ = "STUB: not implemented"
	return *new(rules.RuleList)
}

func loadAnalyzers(include, exclude string) *analyzers.AnalyzerList {
	_ = "STUB: not implemented"
	return nil
}

func getRootPaths(paths []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getPrintedFormat(format string, verbose string) string { _ = "STUB: not implemented"; return "" }

func printReport(format string, color bool, rootPaths []string, reportInfo *gosec.ReportInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func saveReport(filename, format string, rootPaths []string, reportInfo *gosec.ReportInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToScore(value string) (issue.Score, error) {
	_ = "STUB: not implemented"
	return *new(issue.Score), nil
}

func filterIssues(issues []*issue.Issue, severity issue.Score, confidence issue.Score) ([]*issue.Issue, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func computeExitCode(issues []*issue.Issue, errors map[string][]gosec.Error, noFail bool) int {
	_ = "STUB: not implemented"
	return 0
}

func buildPathExclusionFilter(config gosec.Config, cliFlag string) (*gosec.PathExclusionFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }
