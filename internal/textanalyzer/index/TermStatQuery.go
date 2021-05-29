package index

import "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"

// TermStatQuery запрос на термы в коллекции
type TermStatQuery struct {
	MinLength    int
	MaxLength    int
	AllowStarts  bool
	AllowEnds    bool
	Size         int
	MostFrequent bool
}

func NewQuery() TermStatQuery {
	return TermStatQuery{
		MinLength:    0,
		MaxLength:    tokens.MAX_WORD_LENGTH,
		AllowStarts:  true,
		AllowEnds:    true,
		Size:         100,
		MostFrequent: true,
	}
}
