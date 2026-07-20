package analyzers

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"

	"github.com/securego/gosec/v2/issue"
)

const defaultSSHCallbackIssueDescription = "Stateful misuse of ssh.PublicKeyCallback leading to auth bypass"

func newSSHCallbackAnalyzer(id string, description string) *analysis.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

type callbackInfo struct {
	makeClosure *ssa.MakeClosure
	closure     *ssa.Function
	storeInstr  ssa.Instruction
}

func runSSHCallbackAnalysis(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type sshCallbackState struct {
	*BaseAnalyzerState
	ssaFuncs []*ssa.Function
}

func newSSHCallbackState(pass *analysis.Pass, funcs []*ssa.Function) *sshCallbackState {
	_ = "STUB: not implemented"
	return nil
}

func (s *sshCallbackState) findCallbackAssignments() []callbackInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *sshCallbackState) analyzeCallback(cb callbackInfo) *issue.Issue {
	_ = "STUB: not implemented"
	return nil
}

func (s *sshCallbackState) hasWritesToCapturedVars(closure *ssa.Function, mkClosure *ssa.MakeClosure) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *sshCallbackState) isWriteToCapturedVar(instr ssa.Instruction, freeVarSet map[*ssa.FreeVar]ssa.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *sshCallbackState) isStoreToCapturedVar(store *ssa.Store, freeVarSet map[*ssa.FreeVar]ssa.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *sshCallbackState) isValueFromCapturedVar(val ssa.Value, freeVarSet map[*ssa.FreeVar]ssa.Value, depth int) bool {
	_ = "STUB: not implemented"
	return false
}
