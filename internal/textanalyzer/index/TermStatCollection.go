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
}

func (c *TermStatCollection) Terms() map[string]*TermStat {
	return c.terms
}

func (c *TermStatCollection) RebuildIndexes() {
	docorder := make([]*TermStat, 0, len(c.terms))

	for _, v := range c.terms {
		docorder = append(docorder, v)
	}
	sort.SliceStable(docorder, func(i, j int) bool {
		return docorder[i].FirstIndex() < docorder[j].FirstIndex()
	})
	freqorder := make([]*TermStat, 0, len(c.terms))
	freqorder = append(freqorder, docorder...)
	sort.SliceStable(freqorder, func(i, j int) bool {
		return freqorder[i].Count() > freqorder[j].Count()
	})
	c.docOrderIndex = docorder
	c.freqOrderIndex = freqorder
}

func (c *TermStatCollection) DocOrderIndex() []*TermStat {
	if c.docOrderIndex == nil {
		c.RebuildIndexes()
	}
	return c.docOrderIndex
}

func (c *TermStatCollection) FreqOrderIndex() []*TermStat {
	if c.freqOrderIndex == nil {
		c.RebuildIndexes()
	}
	return c.freqOrderIndex
}

func NewTermStatCollection() *TermStatCollection {
	return &TermStatCollection{
		terms: make(map[string]*TermStat, 1024),
	}
}

func (c *TermStatCollection) Add(lexeme lexemes.Lexeme, idx int) {
	c.docOrderIndex = nil
	c.freqOrderIndex = nil
	s, ok := c.terms[lexeme.Value()]
	if ok {
		s.Register(lexeme, idx)
		return
	}
	s = NewLexemeStat(lexeme.Value())
	s.Register(lexeme, idx)
	c.terms[lexeme.Value()] = s
}
func CollectStatsS(text string) *TermStatCollection {
	return CollectStats(strings.NewReader(text))
}
func CollectStats(reader io.Reader) *TermStatCollection {
	stats := NewTermStatCollection()
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
	stats.RebuildIndexes()
	return stats
}

func (c *TermStatCollection) Find(size, minlen int, includeFirst bool, includelLast bool, nonfreq bool) []*TermStat {
	result := make([]*TermStat, 0, size)

	freqs := c.FreqOrderIndex()
	if nonfreq {
		for i := len(freqs) - 1; i >= 0; i-- {
			v := freqs[i]
			if (includeFirst || v.FirstCount() == 0) && (includelLast || v.LastCount() == 0) && v.Len() >= minlen {
				result = append(result, v)
				if len(result) == 10 {
					break
				}
			}
		}
	} else {
		for _, v := range freqs {
			if (includeFirst || v.FirstCount() == 0) && (includelLast || v.LastCount() == 0) && v.Len() >= minlen {
				result = append(result, v)
				if len(result) == 10 {
					break
				}
			}
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].FirstIndex() < result[j].FirstIndex() })

	return result
}
