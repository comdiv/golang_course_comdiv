package app_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

var testConfig = app.NonFlagsAnalyzerConfig{
	Minlen:   4,
	UseFirst: false,
	UseLast:  false,
}

func TestNewIndexService(t *testing.T) {
	app.NewIndexService(app.NewTextAnalyzerArgsNF(testConfig))
}

func TestIndexingService_Index(t *testing.T) {
	indexingService := app.NewIndexService(app.NewTextAnalyzerArgsNF(testConfig))
	indexingService.Index(1, "this is some terms of some mention")
}

func TestIndexingService_Find(t *testing.T) {
	indexingService := app.NewIndexService(app.NewTextAnalyzerArgsNF(testConfig))
	indexingService.Index(1, "this is some terms of some mention")
	top := indexingService.Find(1, index.NewTermFilter(index.TermFilterOptions{}))
	assert.Equal(t, "some", top[0].Value())
}

//noinline
func useNoNil(i interface{}) {
	if i == nil {
		panic("nil!")
	}
}

func TestForRaceDetection(t *testing.T) {
	indexingService := app.NewIndexService(app.NewTextAnalyzerArgsNF(testConfig))

	// тут мы просто рандомно пишем, читаем и ресетим сервис
	// сходимость данных тут не проверяется - проверяем именно на оперделение race между 100 условными клиентами
	// каждый из которых делает 5000 рандомных команд
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		seed := int64(1234213532523*(i+1))
		randomizer := rand.New(rand.NewSource(seed))
		wg.Add(1)
		go func(_i int) {
			defer wg.Done()
			for r := 0; r < 1000; r++ {
				next := randomizer.Int()
				switch {
				case next%5 == 0:
					useNoNil(indexingService.Find(10, index.NewTermFilter(index.TermFilterOptions{})))
				case next%7 == 0:
					indexingService.Reset()
				default:
					indexingService.Index(_i+r, "it is a text for indexing from some part of text")
				}
			}
		}(i)
	}
	wg.Wait()
}
