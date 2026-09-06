package analyzers

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const msgTLSResumptionVerifyPeerBypass = "tls.Config uses VerifyPeerCertificate while session resumption may remain enabled and VerifyConnection is not set; resumed sessions can bypass custom certificate checks"

func newTLSResumptionVerifyPeerAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

type tlsConfigState struct {
	verifyPeerSet              bool
	verifyPeerPos              token.Pos
	verifyConnectionSet        bool
	sessionTicketsDisabledTrue bool
	clientSessionCacheSet      bool
	getConfigForClientSet      bool
	getConfigForClientPos      token.Pos
	getConfigForClientFns      []*ssa.Function
}

type tlsResumptionState struct {
	*BaseAnalyzerState
	configs     map[ssa.Value]*tlsConfigState
	issuesByPos map[token.Pos]*issue.Issue
}

func newTLSResumptionState(pass *analysis.Pass) *tlsResumptionState {
	_ = "STUB: not implemented"
	return nil
}

func runTLSResumptionVerifyPeerAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *tlsResumptionState) trackTLSConfigFieldStore(store *ssa.Store) {
	_ = "STUB: not implemented"
	return
}

func (s *tlsResumptionState) getOrCreateConfigState(root ssa.Value) *tlsConfigState {
	_ = "STUB: not implemented"
	return nil
}

func (s *tlsResumptionState) resolveFunctions(v ssa.Value) []*ssa.Function {
	_ = "STUB: not implemented"
	return nil
}

func (s *tlsResumptionState) reportDirectTLSConfigs() { _ = "STUB: not implemented"; return }

func (s *tlsResumptionState) reportGetConfigForClientBypassCandidates() {
	_ = "STUB: not implemented"
	return
}

func (s *tlsResumptionState) getConfigForClientReturnsRiskyTLSConfig(fns []*ssa.Function) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *tlsResumptionState) extractTLSConfigsFromValue(v ssa.Value, visited map[ssa.Value]struct{}, depth int) []*tlsConfigState {
	_ = "STUB: not implemented"
	return nil
}

func (s *tlsResumptionState) addIssue(pos token.Pos) { _ = "STUB: not implemented"; return }

func tlsConfigRoot(v ssa.Value, depth int) ssa.Value {
	_ = "STUB: not implemented"
	return *new(ssa.Value)
}

func tlsConfigFieldName(fieldAddr *ssa.FieldAddr) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func isTLSConfigPointerType(t types.Type) bool { _ = "STUB: not implemented"; return false }

func boolConstValue(v ssa.Value) (bool, bool) { _ = "STUB: not implemented"; return false, false }

func isNilValue(v ssa.Value) bool { _ = "STUB: not implemented"; return false }
