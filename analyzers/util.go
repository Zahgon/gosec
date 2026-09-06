package analyzers

import (
	"go/token"
	"go/types"
	"math"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/internal/ssautil"
	"github.com/securego/gosec/v2/issue"
)

const MaxDepth = 20

const (
	minInt64  = int64(math.MinInt64)
	maxUint64 = uint64(math.MaxUint64)
	maxInt64  = uint64(math.MaxInt64)
)

type SSAAnalyzerResult = ssautil.SSAAnalyzerResult

type BaseAnalyzerState struct {
	Pass         *analysis.Pass
	Analyzer     *RangeAnalyzer
	Visited      map[ssa.Value]bool
	FuncMap      map[*ssa.Function]bool
	BlockMap     map[*ssa.BasicBlock]bool
	ClosureCache map[ssa.Value]bool
	Depth        int
}

var (
	ErrNoSSAResult    = ssautil.ErrNoSSAResult
	ErrInvalidSSAType = ssautil.ErrInvalidSSAType
)

var (
	visitedPool = sync.Pool{
		New: func() any {
			return make(map[ssa.Value]bool, 64)
		},
	}
	funcMapPool = sync.Pool{
		New: func() any {
			return make(map[*ssa.Function]bool, 32)
		},
	}
	closureCachePool = sync.Pool{
		New: func() any {
			return make(map[ssa.Value]bool, 32)
		},
	}
	blockMapPool = sync.Pool{
		New: func() any {
			return make(map[*ssa.BasicBlock]bool, 32)
		},
	}
)

func NewBaseState(pass *analysis.Pass) *BaseAnalyzerState { _ = "STUB: not implemented"; return nil }

func (s *BaseAnalyzerState) Reset() { _ = "STUB: not implemented"; return }

func (s *BaseAnalyzerState) Release() { _ = "STUB: not implemented"; return }

func (s *BaseAnalyzerState) ResolveFuncs(val ssa.Value, funcs *[]*ssa.Function) {
	_ = "STUB: not implemented"
	return
}

type IntTypeInfo struct {
	Signed bool
	Size   int
	Min    int64
	Max    uint64
}

func isSliceInsideBounds(l, h int, cl, ch int) bool { _ = "STUB: not implemented"; return false }

func isThreeIndexSliceInsideBounds(l, h, maxIdx int, oldCap int) bool {
	_ = "STUB: not implemented"
	return false
}

func BuildDefaultAnalyzers() []*analysis.Analyzer { _ = "STUB: not implemented"; return nil }

func newIssue(analyzerID string, desc string, fileSet *token.FileSet,
	pos token.Pos, severity, confidence issue.Score,
) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func issueCodeSnippet(fileSet *token.FileSet, pos token.Pos) string {
	_ = "STUB: not implemented"
	return ""
}

func GetIntTypeInfo(t types.Type) (IntTypeInfo, error) {
	_ = "STUB: not implemented"
	return *new(IntTypeInfo), nil
}

func GetConstantInt64(v ssa.Value) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func GetConstantUint64(v ssa.Value) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func GetSliceBounds(s *ssa.Slice) (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func GetSliceRange(s *ssa.Slice) (int64, int64) { _ = "STUB: not implemented"; return 0, 0 }

func ComputeSliceNewCap(l, h, maxIdx, oldCap int) int { _ = "STUB: not implemented"; return 0 }

func IsFullSlice(sl *ssa.Slice, bufferLen int64) bool { _ = "STUB: not implemented"; return false }

func IsSubSlice(sub, super *ssa.Slice) bool { _ = "STUB: not implemented"; return false }

func GetBufferLen(val ssa.Value) int64 { _ = "STUB: not implemented"; return 0 }

func BuildCallerMap(funcs []*ssa.Function, callerMap map[string][]*ssa.Call) {
	_ = "STUB: not implemented"
	return
}

func toUint64(i int64) uint64 { _ = "STUB: not implemented"; return 0 }

func toInt64(u uint64) int64 { _ = "STUB: not implemented"; return 0 }

func GetDominators(block *ssa.BasicBlock) []*ssa.BasicBlock { _ = "STUB: not implemented"; return nil }

func IsConstantInTypeRange(constVal *ssa.Const, dstInt IntTypeInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func ExplicitValsInRange(pos []uint, neg []int, dstInt IntTypeInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func TraverseSSA(funcs []*ssa.Function, visitor func(block *ssa.BasicBlock, instr ssa.Instruction)) {
	_ = "STUB: not implemented"
	return
}

type operationInfo struct {
	op      string
	extra   ssa.Value
	flipped bool
}

func minBounds(aVal uint64, aSet bool, bVal uint64, bSet bool, isSrcUnsigned bool) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func maxBounds(aVal uint64, aSet bool, bVal uint64, bSet bool, isSrcUnsigned bool) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func isUint(v ssa.Value) bool { _ = "STUB: not implemented"; return false }

func getRealValueFromOperation(v ssa.Value) (ssa.Value, operationInfo) {
	_ = "STUB: not implemented"
	return *new(ssa.Value), *new(operationInfo)
}

func isEquivalent(a, b ssa.Value) bool { _ = "STUB: not implemented"; return false }

func isSameOrRelated(a, b ssa.Value) bool { _ = "STUB: not implemented"; return false }
