package prefixsearch

import (
	"sort"
	"strings"
)

// PrefixSet 前缀集合
type PrefixSet []string

// NewSortedPrefixSet ...
func NewSortedPrefixSet(prefixes ...string) PrefixSet {
	sort.Strings(prefixes)
	return prefixes
}

// FindLongestPrefix find longest prefix
//  PrefixSet: [a,a/b,a/b/c,x/y/z] x: a/b/c/d returns: a/b/c
//  if NotFound then returns: ""
func (ps PrefixSet) FindLongestPrefix(x string) (prefix string) {
	l := len(ps)
	if l == 0 {
		return ""
	}

	i := sort.SearchStrings(ps, x)

	if i < l && ps[i] == x {
		return x // 完全匹配
	}

	if i > 0 && strings.HasPrefix(x, ps[i-1]) {
		return ps[i-1] // 前缀匹配
	}

	return "" // not found
}
