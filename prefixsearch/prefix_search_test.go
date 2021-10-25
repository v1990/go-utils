package prefixsearch

import "testing"

func TestPrefixSet_FindLongestPrefix(t *testing.T) {
	tests := []struct {
		name      string
		PredixSet PrefixSet
		query     string
		want      string
	}{
		{name: "t1", PredixSet: NewSortedPrefixSet("a", "a/b", "b/c"), query: "a/b", want: "a/b"},
		{name: "t2", PredixSet: NewSortedPrefixSet("a", "a/b", "b/c"), query: "a", want: "a"},
		{name: "t3", PredixSet: NewSortedPrefixSet("a", "a/b", "b/c"), query: "a/b/c", want: "a/b"},
		{name: "t4", PredixSet: NewSortedPrefixSet("a", "a/b", "b/c"), query: "zzz", want: ""},
		{name: "t5", PredixSet: NewSortedPrefixSet("a"), query: "a/b", want: "a"},
		{name: "t6", PredixSet: NewSortedPrefixSet("a"), query: "a", want: "a"},
		{name: "t7", PredixSet: NewSortedPrefixSet("a"), query: "b", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPrefix := tt.PredixSet.FindLongestPrefix(tt.query); gotPrefix != tt.want {
				t.Errorf("FindLongestPrefix: PrefixSet= %v;query= %v;got= %v;want %v", tt.PredixSet, tt.query, gotPrefix, tt.want)
			}
		})
	}
}
