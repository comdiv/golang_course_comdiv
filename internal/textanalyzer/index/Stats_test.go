package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectStats(t *testing.T) {
	stats := index.CollectStatsS("Тут несколько Одинаковых термов именно тут и именно Термов!")
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
