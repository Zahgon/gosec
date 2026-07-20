package gosec

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"sync"
)

var (
	ErrUnexpectedASTNode     = errors.New("unexpected AST node type")
	ErrNoProjectRelativePath = errors.New("no project relative path found")
	ErrNoProjectAbsolutePath = errors.New("no project absolute path found")
)

const envGoModVersion = "GOSECGOVERSION"

func MatchCallByPackage(n ast.Node, c *Context, pkg string, names ...string) (*ast.CallExpr, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func MatchCompLit(n ast.Node, ctx *Context, required string) *ast.CompositeLit {
	_ = "STUB: not implemented"
	return nil
}

func GetInt(n ast.Node) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func GetFloat(n ast.Node) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func GetChar(n ast.Node) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func GetStringRecursive(n ast.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetString(n ast.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetCallObject(n ast.Node, ctx *Context) (*ast.CallExpr, types.Object) {
	_ = "STUB: not implemented"
	return nil, *new(types.Object)
}

type callInfo struct {
	packageName string
	funcName    string
	err         error
}

var callCachePool = sync.Pool{
	New: func() any {
		return make(map[ast.Node]callInfo)
	},
}

func GetCallInfo(n ast.Node, ctx *Context) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func getCallInfo(n ast.Node, ctx *Context) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func GetCallStringArgsValues(n ast.Node, _ *Context) []string {
	_ = "STUB: not implemented"
	return nil
}

func getIdentStringValues(ident *ast.Ident, stringFinder func(ast.Node) (string, error)) []string {
	_ = "STUB: not implemented"
	return nil
}

func GetIdentStringValuesRecursive(ident *ast.Ident) []string {
	_ = "STUB: not implemented"
	return nil
}

func GetIdentStringValues(ident *ast.Ident) []string { _ = "STUB: not implemented"; return nil }

func GetBinaryExprOperands(be *ast.BinaryExpr) []ast.Node { _ = "STUB: not implemented"; return nil }

func GetImportedNames(path string, ctx *Context) (names []string, found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func GetImportPath(name string, ctx *Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetLocation(n ast.Node, ctx *Context) (string, int) { _ = "STUB: not implemented"; return "", 0 }

func Gopath() []string { _ = "STUB: not implemented"; return nil }

func Getenv(key, userDefault string) string { _ = "STUB: not implemented"; return "" }

func GetPkgRelativePath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetPkgAbsPath(pkgPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ConcatString(expr ast.Expr, ctx *Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func FindVarIdentities(n *ast.BinaryExpr, c *Context) ([]*ast.Ident, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func FindModuleRoot(dir string) string { _ = "STUB: not implemented"; return "" }

func PackagePaths(root string, excludes []*regexp.Regexp) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isExcluded(str string, excludes []*regexp.Regexp) bool {
	_ = "STUB: not implemented"
	return false
}

func ExcludedDirsRegExp(excludedDirs []string) []*regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

func RootPath(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }

var (
	goVersionCache struct {
		major, minor, build int
	}
	goVersionOnce sync.Once
)

func GoVersion() (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

type goListOutput struct {
	GoVersion string `json:"GoVersion"`
}

func goModVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseGoVersion(version string) (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func CLIBuildTags(buildTags []string) []string { _ = "STUB: not implemented"; return nil }

func ContainingFile(p interface{ Pos() token.Pos }, ctx *Context) *ast.File {
	_ = "STUB: not implemented"
	return nil
}
