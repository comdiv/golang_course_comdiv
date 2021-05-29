package tokenizer_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokenizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewToken(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.TOKEN_WD, 10, "привет")
	assert.Equal(t, tokenizer.TOKEN_WD, token.Type())
	assert.Equal(t, "привет", token.Value())
	assert.Equal(t, 10, token.Start())
	assert.Equal(t, 21, token.Finish())
	assert.Equal(t, 12, token.Length())
}
