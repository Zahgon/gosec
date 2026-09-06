package analyzers

import (
	"errors"
	"go/types"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"
)

var errNoFound = errors.New("no found")

type bound int

const (
	lowerUnbounded bound = iota
	upperUnbounded
	unbounded
	upperBounded
	bounded
)

func newSliceBoundsAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

type valOffset struct {
	val    ssa.Value
	offset int
}

type sliceBoundsState struct {
	*BaseAnalyzerState
	trackCache map[trackCacheKey]*trackCacheValue
	valQueue   []valOffset
}

var (
	trackValuePool = sync.Pool{
		New: func() any {
			return &trackCacheValue{
				violations: make([]ssa.Instruction, 0, 4),
				ifs:        make(map[ssa.If]*ssa.BinOp),
			}
		},
	}
	trackMapPool = sync.Pool{
		New: func() any {
			return make(map[trackCacheKey]*trackCacheValue, 32)
		},
	}
)

type trackCacheKey struct {
	node     ssa.Node
	sliceCap int
}

type trackCacheValue struct {
	violations []ssa.Instruction
	ifs        map[ssa.If]*ssa.BinOp
}

func newSliceBoundsState(pass *analysis.Pass) *sliceBoundsState {
	_ = "STUB: not implemented"
	return nil
}

func (s *sliceBoundsState) Release() { _ = "STUB: not implemented"; return }

func (s *sliceBoundsState) acquireTrackCacheValue() *trackCacheValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *sliceBoundsState) releaseTrackCacheValue(res *trackCacheValue) {
	_ = "STUB: not implemented"
	return
}

func (v *trackCacheValue) Reset() { _ = "STUB: not implemented"; return }

func (s *sliceBoundsState) Reset() { _ = "STUB: not implemented"; return }

func runSliceBounds(pass *analysis.Pass) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func extractLenBound(binop *ssa.BinOp) (ssa.Value, int, bool) {
	_ = "STUB: not implemented"
	return *new(ssa.Value), 0, false
}

func extractIndexOffset(indexVal ssa.Value, loopVar ssa.Value) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func decomposeIndex(v ssa.Value) (ssa.Value, int) {
	_ = "STUB: not implemented"
	return *new(ssa.Value), 0
}

func (s *sliceBoundsState) trackSliceBounds(depth int, sliceCap int, slice ssa.Node, violations *[]ssa.Instruction, ifs map[ssa.If]*ssa.BinOp) {
	_ = "STUB: not implemented"
	return
}

func (s *sliceBoundsState) extractIntValueIndexAddr(refinstr *ssa.IndexAddr, sliceCap int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *sliceBoundsState) checkAllSlicesBounds(depth int, sliceCap int, slice *ssa.Slice, violations *[]ssa.Instruction, ifs map[ssa.If]*ssa.BinOp) {
	_ = "STUB: not implemented"
	return
}

func extractSliceIfLenCondition(call *ssa.Call) (*ssa.If, *ssa.BinOp) {
	_ = "STUB: not implemented"
	return nil, nil
}

func invBound(bound bound) bound { _ = "STUB: not implemented"; return *new(bound) }

var errExtractBinOp = errors.New("unable to extract constant from binop")

func extractBinOpBound(binop *ssa.BinOp) (bound, int, error) {
	_ = "STUB: not implemented"
	return *new(bound), 0, nil
}

func isSliceIndexInsideBounds(h int, index int) bool { _ = "STUB: not implemented"; return false }

func extractArrayLen(t types.Type) (int, bool) { _ = "STUB: not implemented"; return 0, false }
