package taint

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

const maxTaintDepth = 50

const maxCallerEdges = 32

func isContextType(t types.Type) bool { _ = "STUB: not implemented"; return false }

type Source struct {
	Package string

	Name string

	Pointer bool

	IsFunc bool
}

type Sink struct {
	Package string

	Receiver string

	Method string

	Pointer bool

	CheckArgs []int

	ArgTypeGuards map[int]string
}

func resolveOriginalType(v ssa.Value) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

func guardsSatisfied(args []ssa.Value, sink Sink, prog *ssa.Program) bool {
	_ = "STUB: not implemented"
	return false
}

func lookupNamedType(typePath string, prog *ssa.Program) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

type Sanitizer struct {
	Package string

	Receiver string

	Method string

	Pointer bool
}

type Result struct {
	Source Source

	Sink Sink

	SinkPos token.Pos

	Path []*ssa.Function
}

type Config struct {
	Sources []Source

	Sinks []Sink

	Sanitizers []Sanitizer
}

type paramKey struct {
	fn       *ssa.Function
	paramIdx int
}

type Analyzer struct {
	config          *Config
	sources         map[string]Source
	funcSrcs        map[string]Source
	sinks           map[string]Sink
	sanitizers      map[string]struct{}
	callGraph       *callgraph.Graph
	prog            *ssa.Program
	paramTaintCache map[paramKey]bool
}

func (a *Analyzer) SetCallGraph(cg *callgraph.Graph) { _ = "STUB: not implemented"; return }

func New(config *Config) *Analyzer { _ = "STUB: not implemented"; return nil }

func formatSourceKey(src Source) string { _ = "STUB: not implemented"; return "" }

func formatSinkKey(sink Sink) string { _ = "STUB: not implemented"; return "" }

func formatSanitizerKey(san Sanitizer) string { _ = "STUB: not implemented"; return "" }

func (a *Analyzer) Analyze(prog *ssa.Program, srcFuncs []*ssa.Function) []Result {
	_ = "STUB: not implemented"
	return nil
}

func (a *Analyzer) analyzeFunctionSinks(fn *ssa.Function) []Result {
	_ = "STUB: not implemented"
	return nil
}

func (a *Analyzer) isSinkCall(call *ssa.Call) (Sink, bool) {
	_ = "STUB: not implemented"
	return *new(Sink), false
}

func (a *Analyzer) isSanitizerCall(call *ssa.Call) bool { _ = "STUB: not implemented"; return false }

func (a *Analyzer) isTainted(v ssa.Value, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isSourceType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func mayHaveExternalCallers(fn *ssa.Function) bool { _ = "STUB: not implemented"; return false }

func (a *Analyzer) isSourceFuncCall(call *ssa.Call) bool { _ = "STUB: not implemented"; return false }

func (a *Analyzer) isParameterTainted(param *ssa.Parameter, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFreeVarTainted(fv *ssa.FreeVar, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldAccessTainted(fa *ssa.FieldAddr, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldTaintedOnValue(v ssa.Value, fieldIdx int, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldOfAllocTainted(alloc *ssa.Alloc, fieldIdx int, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldAccessOnPointerTainted(unop *ssa.UnOp, fieldIdx int, fn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldTaintedViaCall(call *ssa.Call, fieldIdx int, callee *ssa.Function, callerFn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isFieldOfAllocTaintedInCallee(alloc *ssa.Alloc, fieldIdx int, callee *ssa.Function, call *ssa.Call, callerFn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) isCalleValueTainted(v ssa.Value, callee *ssa.Function, call *ssa.Call, callerFn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) doTaintedArgsFlowToReturn(call *ssa.Call, callee *ssa.Function, callerFn *ssa.Function, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Analyzer) valueReachableFromParams(v ssa.Value, taintedParams map[*ssa.Parameter]bool, visited map[ssa.Value]bool, depth int) bool {
	_ = "STUB: not implemented"
	return false
}

func traceToAlloc(v ssa.Value) *ssa.Alloc { _ = "STUB: not implemented"; return nil }

func traceToAllocImpl(v ssa.Value, seen map[ssa.Value]bool) *ssa.Alloc {
	_ = "STUB: not implemented"
	return nil
}

func (a *Analyzer) buildPath(fn *ssa.Function) []*ssa.Function {
	_ = "STUB: not implemented"
	return nil
}
