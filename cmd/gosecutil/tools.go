package main

import (
	"flag"
	"go/ast"
	"go/token"
	"go/types"
	"os"
)

type (
	command   func(args ...string)
	utilities struct {
		commands map[string]command
		call     []string
	}
)

func newUtils() *utilities { _ = "STUB: not implemented"; return nil }

func (u *utilities) String() string { _ = "STUB: not implemented"; return "" }

func (u *utilities) Set(opt string) error { _ = "STUB: not implemented"; return nil }

func (u *utilities) run(args ...string) { _ = "STUB: not implemented"; return }

func shouldSkip(path string) bool { _ = "STUB: not implemented"; return false }

func dumpAst(files ...string) { _ = "STUB: not implemented"; return }

type context struct {
	fileset  *token.FileSet
	comments ast.CommentMap
	info     *types.Info
	pkg      *types.Package
	config   *types.Config
	root     *ast.File
}

func createContext(filename string) *context { _ = "STUB: not implemented"; return nil }

func printObject(obj types.Object) { _ = "STUB: not implemented"; return }

func checkContext(ctx *context, file string) bool { _ = "STUB: not implemented"; return false }

func dumpCallObj(files ...string) { _ = "STUB: not implemented"; return }

func dumpUses(files ...string) { _ = "STUB: not implemented"; return }

func dumpTypes(files ...string) { _ = "STUB: not implemented"; return }

func dumpDefs(files ...string) { _ = "STUB: not implemented"; return }

func dumpComments(files ...string) { _ = "STUB: not implemented"; return }

func dumpImports(files ...string) { _ = "STUB: not implemented"; return }

func main() {
	tools := newUtils()
	flag.Var(tools, "tool", "Utils to assist with rule development")
	flag.Parse()

	if len(tools.call) > 0 {
		tools.run(flag.Args()...)
		os.Exit(0)
	}
}
