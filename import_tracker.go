package gosec

import (
	"go/ast"
	"go/types"
	"regexp"
)

var versioningPackagePattern = regexp.MustCompile(`v[0-9]+$`)

type ImportTracker struct {
	Imported map[string][]string
}

func NewImportTracker() *ImportTracker { _ = "STUB: not implemented"; return nil }

func (t *ImportTracker) TrackFile(file *ast.File) { _ = "STUB: not implemented"; return }

func (t *ImportTracker) TrackPackages(pkgs ...*types.Package) { _ = "STUB: not implemented"; return }

func (t *ImportTracker) TrackImport(imported *ast.ImportSpec) { _ = "STUB: not implemented"; return }

func importName(importPath string) string { _ = "STUB: not implemented"; return "" }
