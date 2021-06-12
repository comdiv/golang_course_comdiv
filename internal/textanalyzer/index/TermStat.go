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



func NewLexemeStat(value string) *TermStat {
	return &TermStat{value: value, len: len([]rune(value)), firstIndex: -1, firstPart: -1}
}

func (l *TermStat) Merge(other *TermStat) *TermStat {
	l.count += other.count
	l.firstCount += other.count
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

func (l *TermStat) FullIndex() int64 {
	return int64(l.firstPart) * 10000000000000 + int64(l.firstIndex)
}