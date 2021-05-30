package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCollectStats(t *testing.T) {
	stats := index.CollectStatsS("Тут несколько Одинаковых термов именно тут и именно Термов!", nil)
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
	query := index.NewTermFilter(4, false, false, false)
	f, e := os.Open("../main_test_text.txt")
	if f == nil || e != nil {
		t.Fatal(f, e)
	}
	// собираем статистику, используя наш запрос и при построении для оптимизации (учитываться будет только длина)
	stats := index.CollectStats(f, query)
	// берем топ 10 самых частых слов длиной 4+ в порядке docOrder
	result := stats.Find(10, query)

	resultWords := make([]string, 0, len(result))

	for _, s := range result {
		resultWords = append(resultWords, s.Value())
	}

	assert.Equal(t, []string{
		"YOUNG",
		"LOOKED",
		"WENT",
		"UNTIL",
		"INTO",
		"WANTED",
		"HEARD",
		"LITTLE",
		"HAVE",
		"THAN",
	}, resultWords)
}
