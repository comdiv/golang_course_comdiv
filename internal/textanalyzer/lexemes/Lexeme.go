package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"strings"
)

// Lexeme описатель лексемы в тексте
type Lexeme struct {
	// позиция в предложении
	stPosition int
	// признак последнего слова в предложении
	isLast bool
	token  *tokens.Token
}

func (l *Lexeme) IsUndefined() bool {
	if l.token != nil {
		return l.token.IsUndefined()
	}
	return true
}

func (l *Lexeme) Copy() Lexeme {
	newtoken := l.token.Copy()
	return Lexeme{
		l.stPosition,
		l.isLast,
		&newtoken,
	}
}

func (l *Lexeme) Start() int {
	return l.token.Start()
}

func (l *Lexeme) Finish() int {
	return l.token.Finish()
}

func (l *Lexeme) Value() string {
	return strings.Replace(strings.ToUpper(l.token.Value()), "Ё", "Е", -1)
}

func (l *Lexeme) Len() int {
	return len([]rune(l.Value()))
}

func (l *Lexeme) StatementPosition() int {
	return l.stPosition
}

func (l *Lexeme) IsLastInStatement() bool {
	return l.isLast
}

func (l *Lexeme) IsEof() bool {
	return l.token.IsEof()
}

var NullLexeme = Lexeme{}

func NewLexeme(pos int, last bool, token *tokens.Token) *Lexeme {
	return &Lexeme{
		pos,
		last,
		token,
	}
}

func (l *Lexeme) Apply(pos int, last bool, token *tokens.Token) *Lexeme {
	l.stPosition = pos
	l.isLast = last
	l.token = token
	return l
}
