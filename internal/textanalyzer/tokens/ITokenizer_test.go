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

func Test_StripApos(t *testing.T) {
	tokens := tokens.ReadTokenListS("'Приве'т'")
	assert.Len(t, tokens, 1)
	assert.Equal(t, "Привет", tokens[0].Value())
	assert.Equal(t, 0, tokens[0].Start())
	assert.Equal(t, 11, tokens[0].Finish())
}

func Test_SimpleSingleWord(t *testing.T) {
	TokenizeTextCase{
		"Привет",
		[]tokens.Token{
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 0, "Привет", false, false),
		},
	}.Execute(t)
}

func Test_Two_Words_Splitted_With_Ws(t *testing.T) {
	TokenizeTextCase{
		"Привет   мир!",
		[]tokens.Token{
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 0, "Привет", false, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 12, "   ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 15, "мир", false, false),
			*tokens.NewTokenPlus(tokens.TOKEN_DM, 21, "!", true, false),
		},
	}.Execute(t)
}

func Test_InvalidWords(t *testing.T) {
	TokenizeTextCase{
		"Mixрусскийeng 1.32 jkjkjkjkjdkfsdfsdkfjsdkfjksdfjksdfjsdkfjsdkfjsdkfjsdkfjsdkfjksdfjsdkfjsdkfjsdkfjsdkfjskdfjksdfjkdsfjkds",
		[]tokens.Token{
			*tokens.NewTokenPlus(tokens.TOKEN_UK, 0, "Mixрусскийeng", false, true),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 20, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_NB, 21, "1.32", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 25, " ", true, false),
			*tokens.NewLargeToken(26, 129),
		},
	}.Execute(t)
}

func Test_Statement_Ends(t *testing.T) {
	TokenizeTextCase{
		"A b! C d. E f? G 1.1 h\n i j",
		[]tokens.Token{
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 0, "A", true, true),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 1, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 2, "b", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_DM, 3, "!", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 4, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 5, "C", true, true),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 6, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 7, "d", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_ES, 8, ".", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 9, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 10, "E", true, true) ,
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 11, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 12, "f", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_DM, 13, "?", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 14, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 15, "G", true, true),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 16, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_NB, 17, "1.1", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 20, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 21, "h", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_ES, 22, "\n", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 23, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 24, "i", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WS, 25, " ", true, false),
			*tokens.NewTokenPlus(tokens.TOKEN_WD, 26, "j", true, false),
		},
	}.Execute(t)
}
