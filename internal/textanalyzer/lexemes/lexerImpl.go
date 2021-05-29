package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"strings"
)

type lexerImpl struct {
	tokenizer      tokens.ITokenizer
	statementIndex int
	prepared       Lexeme
}

var _ ILexer = &lexerImpl{}

func newLexerImpl(tokenizer tokens.ITokenizer) *lexerImpl {
	return &lexerImpl{tokenizer: tokenizer, statementIndex: -1}
}

func (l *lexerImpl) Next() Lexeme {
	var result Lexeme
	for {
		token := l.tokenizer.Next()

		if token.IsUnknown() {
			continue
		}

		if token.IsEof() {
			eof := NewLexeme(0, false, token)
			if !l.prepared.IsUndefined() {
				l.prepared.isLast = true
				result, l.prepared = l.prepared, eof
				return result
			}
			return eof
		}

		if token.IsWord() {
			l.statementIndex++
			word := NewLexeme(l.statementIndex, false, token)
			if !l.prepared.IsUndefined() {
				result, l.prepared = l.prepared, word
				return result
			} else {
				l.prepared = word
				continue
			}
		}

		// остались только варианты пробелов и знаков препинаний и если это конец предложения то надо
		// особым образом обработать
		if isStatementDelimiter(token) {
			l.statementIndex = -1
			if !l.prepared.IsUndefined() {
				l.prepared.isLast = true
				result, l.prepared = l.prepared, NullLexeme
				return result
			}
		}

	}
}

func isStatementDelimiter(t tokens.Token) bool {
	return strings.ContainsAny(t.Value(), ".!?\n\r")
}
