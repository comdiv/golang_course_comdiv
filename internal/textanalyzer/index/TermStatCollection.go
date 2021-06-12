package index

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	"io"
	"io/ioutil"
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

func (c *TermStatCollection) Merge(other *TermStatCollection) *TermStatCollection {
	for _, v := range other.Terms() {
		my, ok := c.Terms()[v.Value()]
		if !ok {
			my = NewLexemeStat(v.value)
			c.Terms()[v.Value()] = my
		}
		my.Merge(v)
	}
	return c
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

func (c *TermStatCollection) Add(lexeme *lexemes.Lexeme, part int, idx int) {
	if !c.filter.MatchesLexeme(lexeme) {
		return
	}
	// сбрасываем состояние индекса частот
	c.freqOrderIndex = nil
	value := lexeme.Value()
	s, ok := c.terms[value]
	if ok {
		s.Register(lexeme, part, idx)
		return
	}
	s = NewLexemeStat(lexeme.Value())
	s.Register(lexeme, part, idx)
	c.docOrderIndex = append(c.docOrderIndex, s)
	c.terms[lexeme.Value()] = s
}
func CollectStatsFromJsonS(text string, filter *TermFilter) *TermStatCollection {
	return CollectStatsFromJson(strings.NewReader(text), filter)
}

func CollectStatsFromJson(reader io.Reader, filter *TermFilter) *TermStatCollection {
	result := NewTermStatCollectionF(filter)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Errorf("cannot read bytes from reader %v", err))
	}
	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			panic(fmt.Errorf("error in json %v", err))
		}
		part, err := jsonparser.GetInt(value, "number")
		if err != nil {
			panic(fmt.Errorf("error get number from json %v", err))
		}
		text, err := jsonparser.GetString(value, "text")
		if err != nil {
			panic(fmt.Errorf("error get text from json %v", err))
		}
		lexer := lexemes.NewS(text)
		idx := -1
		for {
			idx++
			lexeme := lexer.Next()
			if lexeme.IsEof() {
				break
			}
			result.Add(lexeme, int(part), idx)
		}
	})
	if err!=nil {
		panic(fmt.Errorf("general error in json %v", err))
	}
	result.RebuildFrequencyIndex()
	return result
}

func CollectStatsS(text string, filter *TermFilter, part int) *TermStatCollection {
	return CollectStats(strings.NewReader(text), filter, part)
}

func CollectStats(reader io.Reader, filter *TermFilter , part int) *TermStatCollection {
	stats := NewTermStatCollectionF(filter)
	lexer := lexemes.NewR(reader)
	idx := -1
	for {
		idx++
		lexeme := lexer.Next()
		if lexeme.IsEof() {
			break
		}
		stats.Add(lexeme, part, idx)
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

	sort.Slice(result, func(i, j int) bool { return result[i].FullIndex() < result[j].FullIndex() })

	return result
}
