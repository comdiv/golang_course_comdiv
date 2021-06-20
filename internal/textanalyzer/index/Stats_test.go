package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCollectStats(t *testing.T) {
	stats,_ := index.CollectFromString("Тут несколько Одинаковых термов именно тут и именно Термов!", index.CollectConfig{Mode: index.MODE_PLAIN})
	assert.Equal(t, 0, stats.Terms()["тут"].FirstIndex())
	assert.Equal(t, 2, stats.Terms()["тут"].Count())
	assert.Equal(t, 1, stats.Terms()["тут"].FirstCount())
	assert.Equal(t, 0, stats.Terms()["тут"].LastCount())
	assert.Equal(t, 3, stats.Terms()["тут"].Len())

	assert.Equal(t, 3, stats.Terms()["термов"].FirstIndex())
	assert.Equal(t, 2, stats.Terms()["термов"].Count())
	assert.Equal(t, 0, stats.Terms()["термов"].FirstCount())
	assert.Equal(t, 1, stats.Terms()["термов"].LastCount())
	assert.Equal(t, 6, stats.Terms()["термов"].Len())

	docorder := stats.DocOrderIndex()
	assert.Equal(t, "тут", docorder[0].Value())
	assert.Equal(t, "несколько", docorder[1].Value())
	assert.Equal(t, "одинаковых", docorder[2].Value())
	assert.Equal(t, "термов", docorder[3].Value())
	assert.Equal(t, "именно", docorder[4].Value())
	assert.Equal(t, "и", docorder[5].Value())

	freqorder := stats.FreqOrderIndex()
	assert.Equal(t, "тут", freqorder[0].Value())
	assert.Equal(t, "термов", freqorder[1].Value())
	assert.Equal(t, "именно", freqorder[2].Value())
	assert.Equal(t, "несколько", freqorder[3].Value())
	assert.Equal(t, "одинаковых", freqorder[4].Value())
	assert.Equal(t, "и", freqorder[5].Value())

}

func TestTask_10_4_no_start_no_finish(t *testing.T) {
	// испходные условия - слова только из середины фраз и длина не менее 4 символов
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       4,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats,_ := index.CollectFromReader(testdata_test.TestDataReader(), index.CollectConfig{Filter:query, Mode: index.MODE_PLAIN})

	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"like", "told", "looked", "marry", "went", "love", "want", "into", "took", "cant",
	}, resultWords)
}

func TestTask_10_4_no_start_no_finish_json_sync(t *testing.T) {
	// испходные условия - слова только из середины фраз и длина не менее 4 символов
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       4,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats,_ := index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Filter:query, Mode: index.MODE_JSON})

	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"like", "told", "looked", "marry", "went", "love", "want", "into", "took", "cant",
	}, resultWords)
}

func TestJsonAsyncRaceCheck(t *testing.T) {
	// 5 раз по 5 в паралелль загрузов большого файла, в 32 потока
	// нет ассертов только для -race

	for i:=0;i<5;i++{
		var wg sync.WaitGroup
		for j:=0;j<5;j++{
			wg.Add(1)
			go func(){
				defer wg.Done()
				stats, err := index.CollectFromReader(testdata_test.TestDataLargeJsonReader(), index.CollectConfig{Mode: index.MODE_PARALLEL_JSON, Workers: 32})
				if err != nil {
					panic(err)
				}
				if len(stats.Terms()) < 100 {
					panic("Too small result")
				}
			}()
		}
		wg.Wait()
	}

}

func TestTask_10_4_no_start_no_finish_json_async(t *testing.T) {
	// испходные условия - слова только из середины фраз и длина не менее 4 символов
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       4,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats,err := index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Filter:query, Mode: index.MODE_PARALLEL_JSON})
	if err != nil {
		t.Fatal(err)
	}

	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"like", "told", "looked", "marry", "went", "love", "want", "into", "took", "cant",
	}, resultWords)
}

func TestTermStat_Merge(t *testing.T) {
	t1 := index.NewTermStatCustom(index.TermStatConfig{
		Value:      "x",
		Count:      10,
		FirstCount: 2,
		LastCount:  3,
		FirstPart:  4,
		FirstIndex: 45,
	})

	// этот перекроет индексы (так как глава меньше)
	t2 := index.NewTermStatCustom(index.TermStatConfig{
		Value:      "x",
		Count:      5,
		FirstCount: 1,
		LastCount:  1,
		FirstPart:  3,
		FirstIndex: 134,
	})

	// этот не перекроет индексы так как в более поздней главе
	t3 := index.NewTermStatCustom(index.TermStatConfig{
		Value:      "x",
		Count:      6,
		FirstCount: 2,
		LastCount:  2,
		FirstPart:  5,
		FirstIndex: 14,
	})

	// этот перекроет индекс так как более раннее определение
	t4 := index.NewTermStatCustom(index.TermStatConfig{
		Value:      "x",
		Count:      1,
		FirstCount: 1,
		LastCount:  1,
		FirstPart:  3,
		FirstIndex: 14,
	})
	// этот не перекроет индекс, в той же главе но индекс больше
	t5 := index.NewTermStatCustom(index.TermStatConfig{
		Value:      "x",
		Count:      1,
		FirstCount: 1,
		LastCount:  1,
		FirstPart:  3,
		FirstIndex: 33,
	})

	res := index.NewTermStat("x").Merge(t1).Merge(t2).Merge(t3).Merge(t4).Merge(t5)
	assert.Equal(t, 3, res.FirstPart())
	assert.Equal(t, 14, res.FirstIndex())
	assert.Equal(t, 23, res.Count())
	assert.Equal(t, 7, res.FirstCount())
	assert.Equal(t, 8, res.LastCount())

	col1 := index.NewTermStatCollection()
	col1.Terms()["x"] = index.NewTermStat("x").Merge(t1).Merge(t2)
	col2 := index.NewTermStatCollection()
	col2.Terms()["x"] = index.NewTermStat("x").Merge(t3).Merge(t4).Merge(t5)

	col1.Merge(col2)

	assert.Equal(t, res, col1.Terms()["x"])
}
