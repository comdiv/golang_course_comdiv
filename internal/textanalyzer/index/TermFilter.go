package index

import "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"

type TermFilter struct {
	minlen       int
	includeFirst bool
	includeLast  bool
	reverseFreq  bool
}
type TermFilterDto struct {
	Minlen       int  `json:"minlen"`
	IncludeFirst bool `json:"include_first"`
	IncludeLast  bool `json:"include_last"`
	ReverseFreq  bool `json:"reverse_freq"`
}

func (t *TermFilter) ToDto() TermFilterDto {
	return TermFilterDto{
		t.minlen,
		t.includeFirst,
		t.includeLast,
		t.reverseFreq,
	}
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

type TermFilterOptions struct {
	MinLen       int
	IncludeFirst bool
	IncludeLast  bool
	ReverseFreq  bool
}

func NewTermFilter(opts TermFilterOptions) *TermFilter {
	return &TermFilter{
		minlen:       opts.MinLen,
		includeFirst: opts.IncludeFirst,
		includeLast:  opts.IncludeLast,
		reverseFreq:  opts.ReverseFreq,
	}
}

// Deprecated: NewTermFilterArgs
func NewTermFilterArgs(minlen int, includeFirst bool, includeLast bool, reverseFreq bool) *TermFilter {
	return &TermFilter{minlen: minlen, includeFirst: includeFirst, includeLast: includeLast, reverseFreq: reverseFreq}
}

func (filter *TermFilter) MatchesStats(v *TermStat) bool {
	return (filter.IncludeFirst() || v.FirstCount() == 0) && (filter.IncludeLast() || v.LastCount() == 0) && v.Len() >= filter.Minlen()
}

func (filter *TermFilter) MatchesLexeme(v *lexemes.Lexeme) bool {
	return v.Len() >= filter.Minlen()
}
