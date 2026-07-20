package rules

import (
	"go/ast"
	"go/types"
	"regexp"
	"sync"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type secretSerialization struct {
	issue.MetaData
	pattern *regexp.Regexp
	cache   sync.Map
}

type formatSpec struct {
	name            string
	tagKey          string
	marshalerMethod string
	functionSinks   []functionSink
	methodSinks     []methodSink
}

type functionSink struct {
	pkgPath string
	names   []string
}

type methodSink struct {
	pkgPath  string
	typeName string
	method   string
}

type typeAnalysisCacheKey struct {
	typ    types.Type
	tagKey string
}

type sensitiveFieldMatch struct {
	fieldName     string
	serializedKey string
	found         bool
}

var g117Formats = []formatSpec{
	{
		name:            "JSON",
		tagKey:          "json",
		marshalerMethod: "MarshalJSON",
		functionSinks: []functionSink{
			{pkgPath: "encoding/json", names: []string{"Marshal", "MarshalIndent"}},
		},
		methodSinks: []methodSink{
			{pkgPath: "encoding/json", typeName: "Encoder", method: "Encode"},
		},
	},
	{
		name:            "YAML",
		tagKey:          "yaml",
		marshalerMethod: "MarshalYAML",
		functionSinks: []functionSink{
			{pkgPath: "go.yaml.in/yaml/v3", names: []string{"Marshal"}},
			{pkgPath: "gopkg.in/yaml.v3", names: []string{"Marshal"}},
			{pkgPath: "gopkg.in/yaml.v2", names: []string{"Marshal"}},
			{pkgPath: "sigs.k8s.io/yaml", names: []string{"Marshal"}},
		},
		methodSinks: []methodSink{
			{pkgPath: "go.yaml.in/yaml/v3", typeName: "Encoder", method: "Encode"},
			{pkgPath: "gopkg.in/yaml.v3", typeName: "Encoder", method: "Encode"},
			{pkgPath: "gopkg.in/yaml.v2", typeName: "Encoder", method: "Encode"},
		},
	},
	{
		name:            "XML",
		tagKey:          "xml",
		marshalerMethod: "MarshalXML",
		functionSinks: []functionSink{
			{pkgPath: "encoding/xml", names: []string{"Marshal", "MarshalIndent"}},
		},
		methodSinks: []methodSink{
			{pkgPath: "encoding/xml", typeName: "Encoder", method: "Encode"},
		},
	},
	{
		name:   "TOML",
		tagKey: "toml",
		functionSinks: []functionSink{
			{pkgPath: "github.com/pelletier/go-toml", names: []string{"Marshal"}},
			{pkgPath: "github.com/pelletier/go-toml/v2", names: []string{"Marshal"}},
		},
		methodSinks: []methodSink{
			{pkgPath: "github.com/pelletier/go-toml", typeName: "Encoder", method: "Encode"},
			{pkgPath: "github.com/pelletier/go-toml/v2", typeName: "Encoder", method: "Encode"},
			{pkgPath: "github.com/BurntSushi/toml", typeName: "Encoder", method: "Encode"},
		},
	},
}

func (r *secretSerialization) Match(n ast.Node, ctx *gosec.Context) (*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var customMarshalerMethods = map[string]bool{
	"MarshalJSON": true,
	"MarshalYAML": true,
	"MarshalXML":  true,
	"MarshalText": true,
	"MarshalTOML": true,
	"MarshalBSON": true,
}

func isInsideCustomMarshaler(callExpr *ast.CallExpr, ctx *gosec.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func typeImplementsMarshaler(typ types.Type, methodName string) bool {
	_ = "STUB: not implemented"
	return false
}

func elementNamedType(typ types.Type) *types.Named { _ = "STUB: not implemented"; return nil }

func compositeLitFieldIsTransformed(expr ast.Expr, fieldName string) bool {
	_ = "STUB: not implemented"
	return false
}

func isNamedTypeInPackage(typ types.Type, pkgPath, typeName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *secretSerialization) findSerializedArgument(callExpr *ast.CallExpr, ctx *gosec.Context) (ast.Expr, formatSpec, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(formatSpec), false
}

func callMatchesMethodSink(callExpr *ast.CallExpr, ctx *gosec.Context, sink methodSink) bool {
	_ = "STUB: not implemented"
	return false
}

func callMatchesPackageFunction(callExpr *ast.CallExpr, ctx *gosec.Context, pkgPath string, names ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func importAliasMatchesPath(ctx *gosec.Context, alias, pkgPath string) bool {
	_ = "STUB: not implemented"
	return false
}

func importAliasPathContains(ctx *gosec.Context, alias, fragment string) bool {
	_ = "STUB: not implemented"
	return false
}

func packageNameFromPath(path string) string { _ = "STUB: not implemented"; return "" }

func packagePathMatches(actual, expected string) bool { _ = "STUB: not implemented"; return false }

func (r *secretSerialization) findSensitiveFieldForType(typ types.Type, tagKey string) sensitiveFieldMatch {
	_ = "STUB: not implemented"
	return *new(sensitiveFieldMatch)
}

func (r *secretSerialization) findSensitiveFieldForTypeWithVisited(typ types.Type, tagKey string, visited map[types.Type]struct{}) sensitiveFieldMatch {
	_ = "STUB: not implemented"
	return *new(sensitiveFieldMatch)
}

func (r *secretSerialization) findSensitiveSerializedField(st *types.Struct, tagKey string) sensitiveFieldMatch {
	_ = "STUB: not implemented"
	return *new(sensitiveFieldMatch)
}

func isSecretCandidateType(typ types.Type) bool { _ = "STUB: not implemented"; return false }

func serializedNameFromTag(defaultName, tag, tagKey string) (name string, omitted bool) {
	_ = "STUB: not implemented"
	return "", false
}

func NewSecretSerialization(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	_ = "STUB: not implemented"
	return *new(gosec.Rule), nil
}
