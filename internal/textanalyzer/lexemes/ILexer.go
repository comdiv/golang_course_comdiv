package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"io"
)

type ILexer interface {
	Next() Lexeme
}

func NewS(text string) ILexer {
	return New(tokens.NewS(text))
}

func NewR(reader io.Reader) ILexer {
	return New(tokens.New(reader))
}

func New(tokenizer tokens.ITokenizer) ILexer {
	return newLexerImpl(tokenizer)
}

func ReadLexemeListS(text string) []Lexeme {
	return ReadLexemeList(NewS(text))
}

func ReadLexemeList(lexer ILexer) []Lexeme {
	result := make([]Lexeme, 0)
	for {
		lexeme := lexer.Next()
		if lexeme.IsEof() {
			break
		}
		result = append(result, lexeme)
	}
	return result
}

func CountLexemesR(reader io.Reader) int {
	return CountLexemes(NewR(reader))
}

func CountLexemes(lexer ILexer) int {
	var r int
	for {
		lexeme := lexer.Next()
		r++
		if lexeme.IsEof() {
			break
		}
	}
	return r
}
