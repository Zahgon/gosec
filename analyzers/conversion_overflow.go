package analyzers

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"
)

func newConversionOverflowAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

type conversionPair struct {
	src types.BasicKind
	dst types.BasicKind
}

type overflowState struct {
	*BaseAnalyzerState
	msgCache map[conversionPair]string
}

func newOverflowState(pass *analysis.Pass) *overflowState { _ = "STUB: not implemented"; return nil }

func runConversionOverflow(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *overflowState) isSafeConversion(instr *ssa.Convert, dstInt IntTypeInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func hasOverflow(srcInfo, dstInfo IntTypeInfo) bool { _ = "STUB: not implemented"; return false }

func isSameWidthPlatformConversion(src, dst types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func isPlatformWordType(k types.BasicKind) bool { _ = "STUB: not implemented"; return false }

func (s *overflowState) hasRangeCheck(v ssa.Value, dstInt IntTypeInfo, block *ssa.BasicBlock) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *overflowState) validateRangeLimits(v ssa.Value, res *rangeResult, dstInt IntTypeInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func signedMinForUnsignedSize(size int) int64 { _ = "STUB: not implemented"; return 0 }

func signedMaxForUnsignedSize(size int) int64 { _ = "STUB: not implemented"; return 0 }

func (s *overflowState) isSafeFromPredecessor(v ssa.Value, dstInt IntTypeInfo, pred *ssa.BasicBlock, targetBlock *ssa.BasicBlock) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *overflowState) isSafeIfEdgeResult(v ssa.Value, dstInt IntTypeInfo, result *rangeResult) bool {
	_ = "STUB: not implemented"
	return false
}
