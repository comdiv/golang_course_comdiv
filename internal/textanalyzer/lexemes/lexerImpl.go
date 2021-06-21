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
		lexeme:         NewLexeme(0, false, tokens.UNDEFINED_TOKEN),
		prepared:       NewLexeme(0, false, tokens.UNDEFINED_TOKEN),
	}
}

func (l *lexerImpl) Next() *Lexeme {
	for {
		token := l.tokenizer.Next().Copy()

		if token.IsUnknown() {
			continue
		}

		if token.IsEof() {
			if !l.prepared.IsUndefined() {
				l.lexeme, l.prepared = l.prepared, l.lexeme
				l.prepared.token = token
				l.prepared.stPosition = 0
				l.lexeme.isLast = true
				l.lexeme.cachedValue = ""
				return l.lexeme
			}
			l.lexeme.token = token
			l.lexeme.stPosition = 0
			l.lexeme.isLast = false
			l.lexeme.cachedValue = ""
			return l.lexeme
		}

		if token.IsWord() {
			l.statementIndex++
			if !l.prepared.IsUndefined() {
				l.lexeme, l.prepared = l.prepared, l.lexeme
				l.prepared.token = token
				l.prepared.stPosition = l.statementIndex
				l.prepared.isLast = false
				l.lexeme.cachedValue = ""
				return l.lexeme
			}
			l.prepared.token = token
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
				l.prepared.token = tokens.UNDEFINED_TOKEN
				l.prepared.isLast = false
				l.prepared.stPosition = 0
				l.prepared.cachedValue = ""
				l.lexeme.cachedValue = ""
				return l.lexeme
			}
		}
	}
}
