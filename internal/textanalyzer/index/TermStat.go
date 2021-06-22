package index

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
)

type TermStat struct {
	value      string
	firstPart  int
	firstIndex int
	len        int
	count      int
	firstCount int
	lastCount  int
}

type TermStatDto struct {
	Value      string `json:"value"`
	FirstPart  int `json:"first_part"`
	FirstIndex int `json:"first_index"`
	Len        int `json:"len"`
	Count      int `json:"count"`
	FirstCount int `json:"first_count"`
	LastCount  int `json:"last_count"`
}

func (t *TermStat) ToDto() TermStatDto {
	return TermStatDto{
		t.value,
		t.firstPart,
		t.firstIndex,
		t.len,
		t.count,
		t.firstCount,
		t.lastCount,
	}
}

type TermStatConfig struct {
	Value      string
	FirstPart  int
	FirstIndex int
	Count      int
	FirstCount int
	LastCount  int
}

func NewTermStatCustom(config TermStatConfig) *TermStat {
	return &TermStat{
		value:      config.Value,
		firstPart:  config.FirstPart,
		firstIndex: config.FirstIndex,
		len:        len([]rune(config.Value)),
		count:      config.Count,
		firstCount: config.FirstCount,
		lastCount:  config.LastCount,
	}
}

func NewTermStat(value string) *TermStat {
	return &TermStat{value: value, len: len([]rune(value)), firstIndex: -1, firstPart: -1}
}

func (l *TermStat) Merge(other *TermStat) *TermStat {
	l.count += other.count
	l.firstCount += other.firstCount
	l.lastCount += other.lastCount
	if l.firstPart == -1 || l.firstPart > other.firstPart {
		l.firstPart = other.firstPart
		l.firstIndex = other.firstIndex
	} else if l.firstPart == other.firstPart && l.firstIndex > other.firstIndex {
		l.firstIndex = other.firstIndex
	}
	return l
}

func (l *TermStat) Register(lexeme *lexemes.Lexeme, part int, idx int) {
	l.count++
	if lexeme.StatementPosition() == 0 {
		l.firstCount++
	}
	if lexeme.IsLastInStatement() {
		l.lastCount++
	}
	// готовимся к мержингу коллекций - выставляем наилучший первый индекс
	if l.firstPart == -1 || l.firstPart > part {
		l.firstPart = part
		l.firstIndex = idx
	} else if l.firstPart == part && l.firstIndex > idx {
		l.firstIndex = idx
	}
}

func (l *TermStat) Value() string {
	return l.value
}

func (l *TermStat) FirstIndex() int {
	return l.firstIndex
}

func (l *TermStat) Len() int {
	return l.len
}

func (l *TermStat) Count() int {
	return l.count
}

func (l *TermStat) FirstCount() int {
	return l.firstCount
}

func (l *TermStat) LastCount() int {
	return l.lastCount
}
func (l *TermStat) FirstPart() int {
	return l.firstPart
}

// GetSortIndex - формирует синтетический порядковый номер для сквозного упорядочения
// число само по себе не является обозначение позиции
func (l *TermStat) GetSortIndex() int64 {
	return int64(l.firstPart)*10000000000000 + int64(l.firstIndex)
}
