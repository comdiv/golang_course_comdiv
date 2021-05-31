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
			tokens.NewToken(tokens.TOKEN_ES, 21, "!"),
		},
	}.Execute(t)
}

func Test_InvalidWords(t *testing.T) {
	TokenizeTextCase{
		"Mixрусскийeng 1.32 jkjkjkjkjdkfsdfsdkfjsdkfjksdfjksdfjsdkfjsdkfjsdkfjsdkfjsdkfjksdfjsdkfjsdkfjsdkfjsdkfjskdfjksdfjkdsfjkds",
		[]tokens.Token{
			tokens.NewToken(tokens.TOKEN_UK, 0, "Mixрусскийeng"),
			tokens.NewToken(tokens.TOKEN_WS, 20, " "),
			tokens.NewToken(tokens.TOKEN_NB, 21, "1.32"),
			tokens.NewToken(tokens.TOKEN_WS, 25, " "),
			tokens.NewLargeToken(26, 129),
		},
	}.Execute(t)
}

func Test_Statement_Ends(t *testing.T) {
	TokenizeTextCase{
		"A B! C D. E F? G 1.1 H\n I J",
		[]tokens.Token{
			tokens.NewToken(tokens.TOKEN_WD, 0, "A"),
			tokens.NewToken(tokens.TOKEN_WS, 1, " "),
			tokens.NewToken(tokens.TOKEN_WD, 2, "B"),
			tokens.NewToken(tokens.TOKEN_ES, 3, "!"),
			tokens.NewToken(tokens.TOKEN_WS, 4, " "),
			tokens.NewToken(tokens.TOKEN_WD, 5, "C"),
			tokens.NewToken(tokens.TOKEN_WS, 6, " "),
			tokens.NewToken(tokens.TOKEN_WD, 7, "D"),
			tokens.NewToken(tokens.TOKEN_ES, 8, "."),
			tokens.NewToken(tokens.TOKEN_WS, 9, " "),
			tokens.NewToken(tokens.TOKEN_WD, 10, "E"),
			tokens.NewToken(tokens.TOKEN_WS, 11, " "),
			tokens.NewToken(tokens.TOKEN_WD, 12, "F"),
			tokens.NewToken(tokens.TOKEN_ES, 13, "?"),
			tokens.NewToken(tokens.TOKEN_WS, 14, " "),
			tokens.NewToken(tokens.TOKEN_WD, 15, "G"),
			tokens.NewToken(tokens.TOKEN_WS, 16, " "),
			tokens.NewToken(tokens.TOKEN_NB, 17, "1.1"),
			tokens.NewToken(tokens.TOKEN_WS, 20, " "),
			tokens.NewToken(tokens.TOKEN_WD, 21, "H"),
			tokens.NewToken(tokens.TOKEN_ES, 22, "\n"),
			tokens.NewToken(tokens.TOKEN_WS, 23, " "),
			tokens.NewToken(tokens.TOKEN_WD, 24, "I"),
			tokens.NewToken(tokens.TOKEN_WS, 25, " "),
			tokens.NewToken(tokens.TOKEN_WD, 26, "J"),
		},
	}.Execute(t)
}
