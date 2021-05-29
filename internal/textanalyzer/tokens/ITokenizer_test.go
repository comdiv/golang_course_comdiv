package tokens_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TokenizeTextCase struct {
	Text   string
	Tokens []tokens.Token
}

func (tc TokenizeTextCase) Execute(t *testing.T) {
	actualTokens := tokens.ReadTokenListS(tc.Text)
	assert.Equal(t, tc.Tokens, actualTokens)
}

// просто проверяем что можем создать токенайзер
func TestNew(t *testing.T) {
	tokenizer := tokens.New(strings.NewReader("Такой текст"))
	assert.NotNil(t, tokenizer)
}
func TestNewS(t *testing.T) {
	tokenizer := tokens.NewS("Такой текст")
	assert.NotNil(t, tokenizer)
}

func TestTokenEqual(t *testing.T) {
	assert.Equal(t, tokens.EofToken(100), tokens.EofToken(100))
	assert.NotEqual(t, tokens.EofToken(100), tokens.EofToken(101))
}

func Test_SingleWordText(t *testing.T) {
	tokens := tokens.ReadTokenListS("Привет")
	assert.Len(t, tokens, 1)
	assert.Equal(t, "Привет", tokens[0].Value())
	assert.Equal(t, 0, tokens[0].Start())
	assert.Equal(t, 11, tokens[0].Finish())
}

func Test_SimpleSingleWord(t *testing.T) {
	TokenizeTextCase{
		"Привет",
		[]tokens.Token{
			tokens.NewToken(tokens.TOKEN_WD, 0, "Привет"),
		},
	}.Execute(t)
}

func Test_Two_Words_Splitted_With_Ws(t *testing.T) {
	TokenizeTextCase{
		"Привет   мир!",
		[]tokens.Token{
			tokens.NewToken(tokens.TOKEN_WD, 0, "Привет"),
			tokens.NewToken(tokens.TOKEN_WS, 12, "   "),
			tokens.NewToken(tokens.TOKEN_WD, 15, "мир"),
			tokens.NewToken(tokens.TOKEN_DM, 21, "!"),
		},
	}.Execute(t)
}

func Test_InvalidWords(t *testing.T) {
	TokenizeTextCase{
		"Mixрусскийeng 1.32 jkjkjkjkjdkfsdfsdkfjsdkfjksdfjksdfjsdkfjsdkfjsdkfjsdkfjsdkfjksdfjsdkfjsdkfjsdkfjsdkfjskdfjksdfjkdsfjkds",
		[]tokens.Token{
			tokens.NewToken(tokens.TOKEN_UK, 0, "Mixрусскийeng"),
			tokens.NewToken(tokens.TOKEN_WS, 20, " "),
			tokens.NewToken(tokens.TOKEN_UK, 21, "1"),
			tokens.NewToken(tokens.TOKEN_DM, 22, "."),
			tokens.NewToken(tokens.TOKEN_UK, 23, "32"),
			tokens.NewToken(tokens.TOKEN_WS, 25, " "),
			tokens.NewLargeToken(26, 129),
		},
	}.Execute(t)
}
