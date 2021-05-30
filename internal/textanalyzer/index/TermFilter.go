package index

type TermFilter struct {
	minlen       int
	includeFirst bool
	includeLast  bool
	reverseFreq  bool
}

func (t *TermFilter) Minlen() int {
	return t.minlen
}

func (t *TermFilter) IncludeFirst() bool {
	return t.includeFirst
}

func (t *TermFilter) IncludeLast() bool {
	return t.includeLast
}

func (t *TermFilter) ReverseFreq() bool {
	return t.reverseFreq
}

func NewTermFilter(minlen int, includeFirst bool, includeLast bool, reverseFreq bool) *TermFilter {
	return &TermFilter{minlen: minlen, includeFirst: includeFirst, includeLast: includeLast, reverseFreq: reverseFreq}
}

func (filter *TermFilter) Matches(v *TermStat) bool {
	return (filter.IncludeFirst() || v.FirstCount() == 0) && (filter.IncludeLast() || v.LastCount() == 0) && v.Len() >= filter.Minlen()
}
