package index

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	"io"
	"sort"
	"strings"
)

type TermStatCollection struct {
	terms          map[string]*TermStat
	docOrderIndex  []*TermStat
	freqOrderIndex []*TermStat
	filter         *TermFilter ``
}

func (c *TermStatCollection) Terms() map[string]*TermStat {
	return c.terms
}

func (c *TermStatCollection) RebuildFrequencyIndex() {
	c.freqOrderIndex = make([]*TermStat, 0, len(c.terms))
	c.freqOrderIndex = append(c.freqOrderIndex, c.docOrderIndex...)
	sort.SliceStable(c.freqOrderIndex, func(i, j int) bool {
		return c.freqOrderIndex[i].Count() > c.freqOrderIndex[j].Count()
	})
}

func (c *TermStatCollection) DocOrderIndex() []*TermStat {
	return c.docOrderIndex
}

func (c *TermStatCollection) FreqOrderIndex() []*TermStat {
	if c.freqOrderIndex == nil {
		c.RebuildFrequencyIndex()
	}
	return c.freqOrderIndex
}
func NewTermStatCollection() *TermStatCollection {
	return NewTermStatCollectionF(nil)
}
func NewTermStatCollectionF(filter *TermFilter) *TermStatCollection {
	if filter == nil {
		filter = NewTermFilterArgs(1, true, true, false)
	}
	return &TermStatCollection{
		terms:  make(map[string]*TermStat, 1024),
		filter: filter,
	}
}

func (c *TermStatCollection) Add(lexeme *lexemes.Lexeme, idx int) {
	if !c.filter.MatchesLexeme(lexeme) {
		return
	}
	// сбрасываем состояние индекса частот
	c.freqOrderIndex = nil
	value := lexeme.Value()
	s, ok := c.terms[value]
	if ok {
		s.Register(lexeme, idx)
		return
	}
	s = NewLexemeStat(lexeme.Value())
	s.Register(lexeme, idx)
	c.docOrderIndex = append(c.docOrderIndex, s)
	c.terms[lexeme.Value()] = s
}
func CollectStatsS(text string, filter *TermFilter) *TermStatCollection {
	return CollectStats(strings.NewReader(text), filter)
}
func CollectStats(reader io.Reader, filter *TermFilter) *TermStatCollection {
	stats := NewTermStatCollectionF(filter)
	lexer := lexemes.NewR(reader)
	idx := -1
	for {
		idx++
		lexeme := lexer.Next()
		if lexeme.IsEof() {
			break
		}
		stats.Add(lexeme, idx)
	}
	stats.RebuildFrequencyIndex()
	return stats
}

func (c *TermStatCollection) Get(word string) *TermStat {
	return c.terms[word]
}
func (c *TermStatCollection) Find(size int, filter *TermFilter) []*TermStat {
	result := make([]*TermStat, 0, size)
	freqs := c.FreqOrderIndex()
	if filter == nil {
		filter = c.filter
	}

	if filter.reverseFreq != c.filter.reverseFreq {
		for i := len(freqs) - 1; i >= 0; i-- {
			v := freqs[i]
			if filter.MatchesStats(v) {
				result = append(result, v)
				if len(result) == size {
					break
				}
			}
		}
	} else {
		for i := 0; i < len(freqs); i++ {
			v := freqs[i]
			if filter.MatchesStats(v) {
				result = append(result, v)
				if len(result) == size {
					break
				}
			}
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].FirstIndex() < result[j].FirstIndex() })

	return result
}
