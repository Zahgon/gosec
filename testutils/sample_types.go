package testutils

import "github.com/securego/gosec/v2"

type CodeSample struct {
	Code   []string
	Errors int
	Config gosec.Config
}
