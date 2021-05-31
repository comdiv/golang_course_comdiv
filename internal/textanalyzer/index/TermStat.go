package index

import "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"

type TermStat struct {
	term       string
	fstIndex   int
	len        int
	count      int
	firstCount int
	lastCount  int
}

func NewLexemeStat(value string) *TermStat {
	return &TermStat{term: value, len: len([]rune(value)), fstIndex: -1}
}

func (l *TermStat) Register(lexeme lexemes.Lexeme, idx int) {
	l.count++
	if lexeme.StatementPosition() == 0 {
		l.firstCount++
	}
	if lexeme.IsLastInStatement() {
		l.lastCount++
	}
	if l.fstIndex == -1 {
		l.fstIndex = idx
	}
}

func (l *TermStat) Value() string {
	return l.term
}

func (l *TermStat) FirstIndex() int {
	return l.fstIndex
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
