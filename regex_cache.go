package gosec

import "regexp"

type regexCacheKey struct {
	Re  *regexp.Regexp
	Str string
}

func RegexMatchWithCache(re *regexp.Regexp, s string) bool { _ = "STUB: not implemented"; return false }
