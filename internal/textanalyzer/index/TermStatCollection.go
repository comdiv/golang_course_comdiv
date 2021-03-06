package index

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

type TermStatCollection struct {
	terms          map[string]*TermStat
	docOrderIndex  []*TermStat
	freqOrderIndex []*TermStat
	filter         *TermFilter
	errors         []error
}

type ReadMode int

const (
	MODE_PLAIN         = ReadMode(1)
	MODE_JSON          = ReadMode(2)
	MODE_PARALLEL_JSON = ReadMode(4)
)

func (c *TermStatCollection) Terms() map[string]*TermStat {
	return c.terms
}

func (c *TermStatCollection) Merge(other *TermStatCollection) *TermStatCollection {
	// если произошла какая-то ошибка в рамках переданной коллекции, то термы не сводим, только ошибки
	if other.Errors() != nil {
		c.AddErrors(other.Errors())
	} else {
		for i := range other.Terms() {
			my, ok := c.Terms()[i]
			if !ok {
				my = NewTermStat(i)
				c.Terms()[i] = my
			}
			my.Merge(other.Terms()[i])
		}
	}
	// сбросить частотный фильтр
	c.freqOrderIndex = nil
	return c
}

func (c *TermStatCollection) IsIndexReady() bool {
	return c.freqOrderIndex != nil
}

func (c *TermStatCollection) RebuildFrequencyIndex() {

	if len(c.docOrderIndex) == 0 { // it will be not sorted after merge
		c.freqOrderIndex = make([]*TermStat, len(c.terms))
		idx := 0
		for i := range c.terms {
			c.freqOrderIndex[idx] = c.terms[i]
			idx++
		}
	} else {
		c.freqOrderIndex = make([]*TermStat, 0, len(c.terms))
		c.freqOrderIndex = append(c.freqOrderIndex, c.docOrderIndex...)
	}
	sort.SliceStable(c.freqOrderIndex, func(i, j int) bool {
		return c.freqOrderIndex[i].Count() > c.freqOrderIndex[j].Count()
	})
}

func (c *TermStatCollection) DocOrderIndex() []*TermStat {
	return c.docOrderIndex
}

func (c *TermStatCollection) FreqOrderIndex() []*TermStat {
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
	s = NewTermStat(lexeme.Value())
	s.Register(lexeme, part, idx)
	c.docOrderIndex = append(c.docOrderIndex, s)
	c.terms[lexeme.Value()] = s
}

type CollectConfig struct {
	Filter  *TermFilter
	Part    int
	Mode    ReadMode
	Workers int
}

func CollectFromString(text string, config CollectConfig) *TermStatCollection {
	return CollectFromReader(strings.NewReader(text), config)
}

func CollectFromReader(reader io.Reader, config CollectConfig) *TermStatCollection {
	switch config.Mode {
	case MODE_PLAIN:
		return collectStats(reader, config)
	case MODE_JSON:
		return collectStatsFromJson(reader, config)
	case MODE_PARALLEL_JSON:
		return collectStatsFromJsonAsync(reader, config)
	default:
		return collectStats(reader, config)
	}
}

type JsonTextPart struct {
	Number int
	Text   string
	// при работе с горутинами и каналами надо как-то
	// более корректно действовать с ошибками
	Error error
}

func NewErrorPart(err error) *JsonTextPart {
	return &JsonTextPart{Error: err}
}

func readJsonChan(reader io.Reader) <-chan *JsonTextPart {
	result := make(chan *JsonTextPart)
	go func() {
		defer close(result)
		decoder := json.NewDecoder(reader)
		// open brace
		_, err := decoder.Token()
		if err != nil {
			result <- NewErrorPart(err)
			return
		}
		for {
			var value JsonTextPart
			err := decoder.Decode(&value)
			if nil != err {
				// check end of array
				if err.Error() != "expected comma after array element" {
					fmt.Printf("Error: %v\n", err)
					result <- NewErrorPart(err)
				}
				break
			}
			result <- &value
		}
	}()
	return result
}

func processPart(block *JsonTextPart, config *CollectConfig, out chan<- *TermStatCollection) {
	subResult := NewTermStatCollectionF(config.Filter)
	lexer := lexemes.NewS(block.Text)
	idx := -1
	for {
		idx++
		lexeme := lexer.Next()
		if lexeme.IsEof() {
			break
		}
		subResult.Add(lexeme, block.Number, idx)
	}
	out <- subResult
}

func collectStatsFromJsonAsync(
	reader io.Reader,
	config CollectConfig,
) *TermStatCollection {
	workers := config.Workers
	if workers <= 0 {
		workers = 8 // default value
	}
	result := NewTermStatCollectionF(config.Filter)
	subCollections := make(chan *TermStatCollection, 20)
	mergewg := new(sync.WaitGroup)
	mergewg.Add(1)
	go func() {
		defer mergewg.Done()
		for c := range subCollections {
			result.Merge(c)
		}
	}()
	wg := new(sync.WaitGroup)
	input := readJsonChan(reader)
	routinePoolController := make(chan struct{}, workers)
	defer close(routinePoolController)
	for e := range input {
		if e.Error != nil {
			result.AddError(e.Error)
			continue
		}
		routinePoolController <- struct{}{}
		wg.Add(1)
		go func(jtp *JsonTextPart) {
			defer wg.Done()
			defer func() {
				<-routinePoolController
			}()
			processPart(jtp, &config, subCollections)
		}(e)
	}
	wg.Wait()
	close(subCollections)
	mergewg.Wait()
	result.RebuildFrequencyIndex()
	return result
}

func collectStatsFromJson(reader io.Reader, config CollectConfig) *TermStatCollection {
	result := NewTermStatCollectionF(config.Filter)
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
	if err != nil {
		panic(fmt.Errorf("general error in json %v", err))
	}
	result.RebuildFrequencyIndex()
	return result
}

func collectStats(reader io.Reader, config CollectConfig) *TermStatCollection {
	stats := NewTermStatCollectionF(config.Filter)
	lexer := lexemes.NewR(reader)
	idx := -1
	for {
		idx++
		lexeme := lexer.Next()
		if lexeme.IsEof() {
			break
		}
		stats.Add(lexeme, config.Part, idx)
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
			if v == nil {
				// им по идее не откуда браться!!!
				fmt.Println("was nil in collection!!!")
			}
			if filter.MatchesStats(v) {
				result = append(result, v)
				if len(result) == size {
					break
				}
			}
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].GetSortIndex() < result[j].GetSortIndex() })

	return result
}

func (c *TermStatCollection) Errors() []error {
	return c.errors
}

func (c *TermStatCollection) AddError(e error) {
	if c.errors == nil {
		c.errors = make([]error, 0)
	}
	c.errors = append(c.errors, e)
}

func (c *TermStatCollection) AddErrors(e []error) {
	if c.errors == nil {
		c.errors = make([]error, 0)
	}
	c.errors = append(c.errors, e...)
}
