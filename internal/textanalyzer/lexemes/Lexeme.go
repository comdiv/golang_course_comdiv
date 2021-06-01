package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"strings"
)

// Lexeme описатель лексемы в тексте
type Lexeme struct {
	// нормализованное значение лексемы
	value string
	// позиция в предложении
	stPosition int
	// признак последнего слова в предложении
	isLast    bool
	isEof     bool
	isDefined bool
	start     int
	finish    int
}

func (l Lexeme) IsUndefined() bool {
	return !l.isDefined
}

func (l Lexeme) Start() int {
	return l.start
}

func (l Lexeme) Finish() int {
	return l.finish
}

func (l Lexeme) Value() string {
	return l.value
}

func (l Lexeme) Len() int {
	return len([]rune(l.value))
}

func (l Lexeme) StatementPosition() int {
	return l.stPosition
}

func (l Lexeme) IsLastInStatement() bool {
	return l.isLast
}

func (l Lexeme) IsEof() bool {
	return l.isEof
}

var NullLexeme = Lexeme{}

func NewLexeme(pos int, last bool, token *tokens.Token) Lexeme {
	return Lexeme{
		strings.Replace(strings.ToUpper(token.Value()), "Ё", "Е", -1),
		pos,
		last,
		token.IsEof(),
		!token.IsUndefined(),
		token.Start(),
		token.Finish(),
	}
}
