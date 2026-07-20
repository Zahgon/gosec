package analyzers

import (
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const defaultIssueDescription = "Use of hardcoded IV/nonce for encryption"

var tracked = map[string][]int{
	"(crypto/cipher.AEAD).Seal":     {4, 1},
	"crypto/cipher.NewCBCEncrypter": {2, 1},
	"crypto/cipher.NewCFBEncrypter": {2, 1},
	"crypto/cipher.NewCTREncrypter": {2, 1},
	"crypto/cipher.NewCTR":          {2, 1},
	"crypto/cipher.NewOFB":          {2, 1},
	"crypto/cipher.NewCFB":          {2, 1},
	"crypto/cipher.NewCBC":          {2, 1},
}

var dynamicFuncs = map[string]bool{
	"crypto/rand.Read": true,
	"io.ReadFull":      true,
}

var dynamicPkgs = map[string]bool{
	"crypto/rand": true,
	"io":          true,
}

var cipherPkgPrefixes = []string{
	"crypto/cipher",
	"crypto/aes",
}

const (
	statusVisiting = 1 << 0
	statusHard     = 1 << 1
	statusDyn      = 1 << 2
)

func newHardCodedNonce(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func runHardCodedNonce(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type analysisState struct {
	*BaseAnalyzerState
	ssaFuncs   []*ssa.Function
	usageCache map[ssa.Value]uint8
	callerMap  map[string][]*ssa.Call
}

var (
	usageCachePool = sync.Pool{
		New: func() any {
			return make(map[ssa.Value]uint8, 64)
		},
	}
	callerMapPool = sync.Pool{
		New: func() any {
			return make(map[string][]*ssa.Call, 32)
		},
	}
)

type ssaValueAndInstr struct {
	val   ssa.Value
	instr ssa.Instruction
}

func newAnalysisState(pass *analysis.Pass, funcs []*ssa.Function) *analysisState {
	_ = "STUB: not implemented"
	return nil
}

func (s *analysisState) Release() { _ = "STUB: not implemented"; return }

func isAEADOpenCall(c *ssa.Call) bool { _ = "STUB: not implemented"; return false }

func (s *analysisState) getInitialArgs(tracked map[string][]int) []ssaValueAndInstr {
	_ = "STUB: not implemented"
	return nil
}

func (s *analysisState) raiseIssue(val ssa.Value, issueDescription string, fromInstr ssa.Instruction) ([]*issue.Issue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *analysisState) isHardcoded(val ssa.Value) bool { _ = "STUB: not implemented"; return false }

func (s *analysisState) isFuncReturnsHardcoded(fn *ssa.Function) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *analysisState) analyzeUsage(val ssa.Value) uint8 { _ = "STUB: not implemented"; return 0 }

func (s *analysisState) analyzeReferrer(ref ssa.Instruction, val ssa.Value) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func (s *analysisState) allTaintedEventsCovered(val ssa.Value, usage ssa.Instruction) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *analysisState) collectTaintedEvents(val ssa.Value, usage ssa.Instruction, actions *[]RangeAction) {
	_ = "STUB: not implemented"
	return
}

func (s *analysisState) collectCoveredRanges(val ssa.Value, usage ssa.Instruction, actions *[]RangeAction) {
	_ = "STUB: not implemented"
	return
}

func (s *analysisState) isFullDynamicRead(ref ssa.Instruction, val ssa.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *analysisState) resolveAbsoluteRange(val ssa.Value) (ByteRange, bool) {
	_ = "STUB: not implemented"
	return *new(ByteRange), false
}
