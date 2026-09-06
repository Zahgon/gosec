package testutils

import (
	"go/build"

	"golang.org/x/tools/go/packages"

	"github.com/securego/gosec/v2"
)

type buildObj struct {
	pkg    *build.Package
	config *packages.Config
	pkgs   []*packages.Package
}

type TestPackage struct {
	Path   string
	Files  map[string]string
	onDisk bool
	build  *buildObj
}

type Option func(conf *packages.Config)

func WithBuildTags(tags []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewTestPackage() *TestPackage { _ = "STUB: not implemented"; return nil }

func (p *TestPackage) AddFile(filename, content string) { _ = "STUB: not implemented"; return }

func (p *TestPackage) write() error { _ = "STUB: not implemented"; return nil }

func (p *TestPackage) Build(opts ...Option) error { _ = "STUB: not implemented"; return nil }

func (p *TestPackage) CreateContext(filename string, opts ...Option) *gosec.Context {
	_ = "STUB: not implemented"
	return nil
}

func (p *TestPackage) Close() { _ = "STUB: not implemented"; return }

func (p *TestPackage) Pkgs() []*packages.Package { _ = "STUB: not implemented"; return nil }

func (p *TestPackage) PrintErrors() int { _ = "STUB: not implemented"; return 0 }
