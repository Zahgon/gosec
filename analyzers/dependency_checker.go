package analyzers

import "golang.org/x/tools/go/ssa"

type dependencyKey struct {
	value  ssa.Value
	target ssa.Value
}

type dependencyChecker struct {
	memo     map[dependencyKey]bool
	visiting map[dependencyKey]struct{}
}

func newDependencyChecker() *dependencyChecker { _ = "STUB: not implemented"; return nil }

func (c *dependencyChecker) dependsOn(value ssa.Value, target ssa.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *dependencyChecker) dependsOnDepth(value ssa.Value, target ssa.Value, depth int) bool {
	_ = "STUB: not implemented"
	return false
}
