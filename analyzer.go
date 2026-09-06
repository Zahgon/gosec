package gosec

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"log"

	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/packages"

	"github.com/securego/gosec/v2/analyzers"
	"github.com/securego/gosec/v2/issue"
)

var (
	ErrNoPackageTypeInfo = errors.New("package has no type information")
	ErrNilPackage        = errors.New("nil package provided")
)

const LoadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedTypes |
	packages.NeedTypesSizes |
	packages.NeedTypesInfo |
	packages.NeedSyntax |
	packages.NeedModule |
	packages.NeedEmbedFiles |
	packages.NeedEmbedPatterns

const (
	externalSuppressionJustification = "Globally suppressed."
	aliasOfAllRules                  = "*"
	directivePrefix                  = "//gosec:disable"
)

type ignore struct {
	start        int
	end          int
	suppressions map[string][]issue.SuppressionInfo
}

type ignores map[string][]ignore

func newIgnores() ignores { _ = "STUB: not implemented"; return *new(ignores) }

func (i ignores) parseLine(line string) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (i ignores) add(file string, line string, suppressions map[string]issue.SuppressionInfo) {
	_ = "STUB: not implemented"
	return
}

func (i ignores) get(file string, line string) map[string][]issue.SuppressionInfo {
	_ = "STUB: not implemented"
	return nil
}

type Context struct {
	FileSet      *token.FileSet
	Comments     ast.CommentMap
	Info         *types.Info
	Pkg          *types.Package
	PkgFiles     []*ast.File
	Root         *ast.File
	Imports      *ImportTracker
	Config       Config
	Ignores      ignores
	PassedValues map[string]any
	callCache    map[ast.Node]callInfo
}

func (ctx *Context) GetFileAtNodePos(node ast.Node) *token.File {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) NewIssue(node ast.Node, ruleID, desc string,
	severity, confidence issue.Score,
) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

type Metrics struct {
	NumFiles int `json:"files"`
	NumLines int `json:"lines"`
	NumNosec int `json:"nosec"`
	NumFound int `json:"found"`
}

func (m *Metrics) Merge(other *Metrics) { _ = "STUB: not implemented"; return }

type Analyzer struct {
	ignoreNosec bool
	ruleset     RuleSet

	ruleBuilders      map[string]RuleBuilder
	ruleSuppressed    map[string]bool
	context           *Context
	config            Config
	logger            *log.Logger
	issues            []*issue.Issue
	stats             *Metrics
	errors            map[string][]Error
	tests             bool
	excludeGenerated  bool
	showIgnored       bool
	trackSuppressions bool
	concurrency       int
	analyzerSet       *analyzers.AnalyzerSet
}

func NewAnalyzer(conf Config, tests bool, excludeGenerated bool, trackSuppressions bool, concurrency int, logger *log.Logger) *Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (gosec *Analyzer) SetConfig(conf Config) { _ = "STUB: not implemented"; return }

func (gosec *Analyzer) Config() Config { _ = "STUB: not implemented"; return *new(Config) }

func (gosec *Analyzer) LoadRules(ruleDefinitions map[string]RuleBuilder, ruleSuppressed map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func (gosec *Analyzer) buildPackageRuleset() RuleSet {
	_ = "STUB: not implemented"
	return *new(RuleSet)
}

func (gosec *Analyzer) LoadAnalyzers(analyzerDefinitions map[string]analyzers.AnalyzerDefinition, analyzerSuppressed map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func (gosec *Analyzer) Process(buildTags []string, packagePaths ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (gosec *Analyzer) load(pkgPath string, buildTags []string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gosec *Analyzer) CheckRules(pkg *packages.Package) { _ = "STUB: not implemented"; return }

func (gosec *Analyzer) checkRules(pkg *packages.Package) ([]*issue.Issue, *Metrics, ignores) {
	_ = "STUB: not implemented"
	return nil, nil, *new(ignores)
}

func (gosec *Analyzer) CheckAnalyzers(pkg *packages.Package) { _ = "STUB: not implemented"; return }

func (gosec *Analyzer) checkAnalyzers(pkg *packages.Package, allIgnores ignores) ([]*issue.Issue, *Metrics) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gosec *Analyzer) CheckAnalyzersWithSSA(pkg *packages.Package, ssaResult *buildssa.SSA) {
	_ = "STUB: not implemented"
	return
}

func (gosec *Analyzer) checkAnalyzersWithSSA(pkg *packages.Package, ssaResult *buildssa.SSA, allIgnores ignores) ([]*issue.Issue, *Metrics) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gosec *Analyzer) generatedFiles(pkg *packages.Package) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (gosec *Analyzer) buildSSA(pkg *packages.Package) (*buildssa.SSA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseErrors(pkg *packages.Package) (map[string][]Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gosec *Analyzer) AppendError(file string, err error) { _ = "STUB: not implemented"; return }

func (gosec *Analyzer) appendErrorAt(file string, line, column int, err error) {
	_ = "STUB: not implemented"
	return
}

func findNoSecDirective(group *ast.CommentGroup, noSecDefaultTag, noSecAlternativeTag string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func findNoSecTag(text, tag string) (bool, string) { _ = "STUB: not implemented"; return false, "" }

type astVisitor struct {
	gosec *Analyzer

	ruleset           *RuleSet
	context           *Context
	issues            []*issue.Issue
	stats             *Metrics
	ignoreNosec       bool
	showIgnored       bool
	trackSuppressions bool
}

func (v *astVisitor) activeRuleset() *RuleSet { _ = "STUB: not implemented"; return nil }

func (v *astVisitor) Visit(n ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func (v *astVisitor) updateIgnores() { _ = "STUB: not implemented"; return }

func (v *astVisitor) updateIgnoredRulesForNode(n ast.Node) { _ = "STUB: not implemented"; return }

func (v *astVisitor) ignore(n ast.Node) (map[string]issue.SuppressionInfo, *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *astVisitor) reportInvalidDirective(group *ast.CommentGroup, reason string) {
	_ = "STUB: not implemented"
	return
}

func (gosec *Analyzer) updateIssues(issue *issue.Issue, issues []*issue.Issue, stats *Metrics, allIgnores ignores) []*issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func getSuppressions(ignores ignores, file, line, ruleID string, ruleset RuleSet, analyzerSet *analyzers.AnalyzerSet) ([]issue.SuppressionInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (gosec *Analyzer) Report() ([]*issue.Issue, *Metrics, map[string][]Error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (gosec *Analyzer) Reset() { _ = "STUB: not implemented"; return }
