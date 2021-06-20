package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
)

// Lexeme описатель лексемы в тексте
type Lexeme struct {
	// позиция в предложении
	stPosition int
	// признак последнего слова в предложении
	isLast bool
	token  tokens.Token
	cachedValue string
}

func (l *Lexeme)  prepareValue() {
	l.cachedValue = l.token.NormalValue()
}

func (l *Lexeme) IsUndefined() bool {
	return l.token.IsUndefined()
}

func (l *Lexeme) Copy() Lexeme {
	return Lexeme{
		l.stPosition,
		l.isLast,
		l.token,
		"",
	}
}

func (l *Lexeme) Start() int {
	return l.token.Start()
}

func (l *Lexeme) Finish() int {
	return l.token.Finish()
}

func (l *Lexeme) Value() string {
	if len(l.cachedValue)==0 {
		l.prepareValue()
	}
	return l.cachedValue
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

func NewLexeme(pos int, last bool, token tokens.Token) *Lexeme {
	return &Lexeme{
		pos,
		last,
		token,
		"",
	}
}

func (l *Lexeme) Apply(pos int, last bool, token tokens.Token) *Lexeme {
	l.stPosition = pos
	l.isLast = last
	l.token = token
	l.cachedValue = ""
	return l
}
