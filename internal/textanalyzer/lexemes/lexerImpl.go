package lexemes

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
)

type lexerImpl struct {
	tokenizer      tokens.ITokenizer
	statementIndex int
	lexeme         *Lexeme
	prepared       *Lexeme
}

var _ ILexer = &lexerImpl{}

func newLexerImpl(tokenizer tokens.ITokenizer) *lexerImpl {
	return &lexerImpl{
		tokenizer:      tokenizer,
		statementIndex: -1,
		lexeme:         NewLexeme(0, false, nil),
		prepared:       NewLexeme(0, false, nil),
	}
}

func (l *lexerImpl) Next() *Lexeme {
	for {
		token := l.tokenizer.Next()

		if token.IsUnknown() {
			continue
		}

		if token.IsEof() {
			if !l.prepared.IsUndefined() {
				l.lexeme, l.prepared = l.prepared, l.lexeme
				l.prepared.token = token
				l.prepared.stPosition = 0
				l.lexeme.isLast = true
				return l.lexeme
			}
			l.lexeme.token = token
			l.lexeme.stPosition = 0
			l.lexeme.isLast = false
			return l.lexeme
		}

		if token.IsWord() {
			l.statementIndex++
			newt := token.Copy()
			if !l.prepared.IsUndefined() {
				l.lexeme, l.prepared = l.prepared, l.lexeme
				l.prepared.token = &newt
				l.prepared.stPosition = l.statementIndex
				l.prepared.isLast = false
				return l.lexeme
			}
			l.prepared.token = &newt
			l.prepared.stPosition = l.statementIndex
			l.prepared.isLast = false
			continue
		}

		// остались только варианты пробелов и знаков препинаний и если это конец предложения то надо
		// особым образом обработать
		if token.IsEoS() {
			l.statementIndex = -1
			if !l.prepared.IsUndefined() {
				l.prepared.isLast = true
				l.lexeme, l.prepared = l.prepared, l.lexeme
				l.prepared.token = nil
				l.prepared.isLast = false
				l.prepared.stPosition = 0
				return l.lexeme
			}
		}
	}
}
