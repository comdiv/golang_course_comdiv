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
	isLast bool
	// исходный токен
	token tokens.Token
}

func (l Lexeme) Value() string {
	return l.value
}

func (l Lexeme) StatementPosition() int {
	return l.stPosition
}

func (l Lexeme) IsLastInStatement() bool {
	return l.isLast
}

func (l Lexeme) Token() tokens.Token {
	return l.token
}

func (l Lexeme) IsEof() bool {
	return l.token.Type() == tokens.TOKEN_EOF
}
func (l Lexeme) IsUndefined() bool {
	return l.token.IsUndefined()
}

var NullLexeme = Lexeme{}

func NewLexeme(pos int, last bool, token tokens.Token) Lexeme {
	return Lexeme{
		strings.Replace(strings.ToUpper(token.Value()), "Ё", "Е", -1),
		pos,
		last,
		token,
	}
}
