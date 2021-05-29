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

func ReadTokenListR(reader io.Reader) []Token {
	return ReadTokenList(New(reader))
}

func ReadTokenList(tokenizer ITokenizer) []Token {
	result := make([]Token, 0)
	for {
		token := tokenizer.Next()
		if token.Type() == TOKEN_EOF {
			break
		}
		result = append(result, token)
	}
	return result
}
