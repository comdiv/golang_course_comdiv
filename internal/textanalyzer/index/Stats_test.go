package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectStats(t *testing.T) {
	stats := index.CollectStatsS("Тут несколько Одинаковых термов именно тут и именно Термов!", nil, 0)
	assert.Equal(t, 0, stats.Terms()["ТУТ"].FirstIndex())
	assert.Equal(t, 2, stats.Terms()["ТУТ"].Count())
	assert.Equal(t, 1, stats.Terms()["ТУТ"].FirstCount())
	assert.Equal(t, 0, stats.Terms()["ТУТ"].LastCount())
	assert.Equal(t, 3, stats.Terms()["ТУТ"].Len())

	assert.Equal(t, 3, stats.Terms()["ТЕРМОВ"].FirstIndex())
	assert.Equal(t, 2, stats.Terms()["ТЕРМОВ"].Count())
	assert.Equal(t, 0, stats.Terms()["ТЕРМОВ"].FirstCount())
	assert.Equal(t, 1, stats.Terms()["ТЕРМОВ"].LastCount())
	assert.Equal(t, 6, stats.Terms()["ТЕРМОВ"].Len())

	docorder := stats.DocOrderIndex()
	assert.Equal(t, "ТУТ", docorder[0].Value())
	assert.Equal(t, "НЕСКОЛЬКО", docorder[1].Value())
	assert.Equal(t, "ОДИНАКОВЫХ", docorder[2].Value())
	assert.Equal(t, "ТЕРМОВ", docorder[3].Value())
	assert.Equal(t, "ИМЕННО", docorder[4].Value())
	assert.Equal(t, "И", docorder[5].Value())

	freqorder := stats.FreqOrderIndex()
	assert.Equal(t, "ТУТ", freqorder[0].Value())
	assert.Equal(t, "ТЕРМОВ", freqorder[1].Value())
	assert.Equal(t, "ИМЕННО", freqorder[2].Value())
	assert.Equal(t, "НЕСКОЛЬКО", freqorder[3].Value())
	assert.Equal(t, "ОДИНАКОВЫХ", freqorder[4].Value())
	assert.Equal(t, "И", freqorder[5].Value())

}

func TestTask_10_4_no_start_no_finish(t *testing.T) {
	// испходные условия - слова только из середины фраз и длина не менее 4 символов
	query := index.NewTermFilterArgs(4, false, false, false)
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats := index.CollectStats(testdata_test.TestDataReader(), query, 0)

	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"LIKE", "TOLD", "LOOKED", "MARRY", "WENT", "LOVE", "WANT", "INTO", "TOOK", "CANT",
	}, resultWords)
}

func TestTask_10_4_no_start_no_finish_json_json(t *testing.T) {
	// испходные условия - слова только из середины фраз и длина не менее 4 символов
	query := index.NewTermFilterArgs(4, false, false, false)
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats := index.CollectStatsFromJson(testdata_test.TestDataJsonReader(), query)

	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"LIKE", "TOLD", "LOOKED", "MARRY", "WENT", "LOVE", "WANT", "INTO", "TOOK", "CANT",
	}, resultWords)
}

func TestTermStat_Merge(t *testing.T) {
	t1 := index.NewTermStat("x")
	t1.SetCount(10)
	t1.SetFirstCount(2)
	t1.SetLastCount(3)
	t1.SetFirstPart(4)
	t1.SetFirstIndex(45)

	// этот перекроет индексы (так как глава меньше)
	t2 := index.NewTermStat("x")
	t2.SetCount(5)
	t2.SetFirstCount(1)
	t2.SetLastCount(1)
	t2.SetFirstPart(3)
	t2.SetFirstIndex(134)

	// этот не перекроет индексы так как в более поздней главе
	t3 := index.NewTermStat("x")
	t3.SetCount(6)
	t3.SetFirstCount(2)
	t3.SetLastCount(2)
	t3.SetFirstPart(5)
	t3.SetFirstIndex(14)

	// этот перекроет индекс так как более раннее определение
	t4 := index.NewTermStat("x")
	t4.SetCount(1)
	t4.SetFirstCount(1)
	t4.SetLastCount(1)
	t4.SetFirstPart(3)
	t4.SetFirstIndex(14)

	// этот не перекроет индекс, в той же главе но индекс больше
	t5 := index.NewTermStat("x")
	t5.SetCount(1)
	t5.SetFirstCount(1)
	t5.SetLastCount(1)
	t5.SetFirstPart(3)
	t5.SetFirstIndex(33)

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
