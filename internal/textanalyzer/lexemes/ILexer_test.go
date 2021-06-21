package lexemes_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LexerTextCase struct {
	Text    string
	Lexemes []lexemes.Lexeme
}

func (lt LexerTextCase) Execute(t *testing.T) {
	actualLexemes := lexemes.ReadLexemeListS(lt.Text)
	assert.Equal(t, lt.Lexemes, actualLexemes)
}

func TestNullLexemeIsUndefined(t *testing.T) {
	assert.True(t, lexemes.NullLexeme.IsUndefined())
}

func TestNewLexeme(t *testing.T) {
	token := tokens.NewTokenPlus(tokens.TOKEN_WD, 12, "Привёт", false, false)
	lexeme := lexemes.NewLexeme(2, true, *token)
	assert.Equal(t, "привет", lexeme.Value())
	assert.Equal(t, 2, lexeme.StatementPosition())
	assert.True(t, lexeme.IsLastInStatement())
	assert.Equal(t, token.Start(), lexeme.Start())
	assert.Equal(t, token.Finish(), lexeme.Finish())
}

func TestSingleStatement(t *testing.T) {
	LexerTextCase{
		"2 Привет мир #23 и здравствуй!",
		[]lexemes.Lexeme{
			*lexemes.NewLexeme(
				0,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 2, "Привет", false, false),
			),
			*lexemes.NewLexeme(
				1,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 15, "мир", false, false),
			),
			*lexemes.NewLexeme(
				2,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 26, "и", false, false),
			),
			*lexemes.NewLexeme(
				3,
				true,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 29, "здравствуй", false, false),
			),
		},
	}.Execute(t)
}

func TestMultiStatement(t *testing.T) {
	LexerTextCase{
		"2 Привет мир #23 и здравствуй! Ещё предложение. \n И еще.",
		[]lexemes.Lexeme{
			*lexemes.NewLexeme(
				0,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 2, "Привет", false, false),
			),
			*lexemes.NewLexeme(
				1,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 15, "мир", false, false),
			),
			*lexemes.NewLexeme(
				2,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 26, "и", false, false),
			),
			*lexemes.NewLexeme(
				3,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 29, "здравствуй", false, false),
			),
			*lexemes.NewLexeme(
				4,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 51, "Ещё", false, false),
			),
			*lexemes.NewLexeme(
				5,
				true,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 58, "предложение", false, false),
			),
			*lexemes.NewLexeme(
				0,
				false,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 84, "И", false, false),
			),
			*lexemes.NewLexeme(
				1,
				true,
				*tokens.NewTokenPlus(tokens.TOKEN_WD, 87, "еще", false, false),
			),
		},
	}.Execute(t)
}
