package tokens

import (
	"io"
	"strings"
)

// ITokenizer - интерфейс получения токена
type ITokenizer interface {
	Next() Token
}

func New(reader io.Reader) ITokenizer {
	return newTokenizerImpl(reader)
}
func NewS(text string) ITokenizer {
	return newTokenizerImpl(strings.NewReader(text))
}

func ReadTokenListS(text string) []Token {
	return ReadTokenList(New(strings.NewReader(text)))
}

func ReadTokenList(tokenizer ITokenizer) []Token {
	result := make([]Token, 0)
	for {
		token := tokenizer.Next()
		if token.IsEof() {
			break
		}
		result = append(result, token)
	}
	return result
}

func CountTokensR(reader io.Reader) int {
	return CountTokens(New(reader))
}

func CountTokens(tokenizer ITokenizer) int {
	var r int
	for {
		token := tokenizer.Next()
		r++
		if token.IsEof() {
			break
		}
	}
	return r
}
