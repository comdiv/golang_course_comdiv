package tokens

import (
	"bufio"
	"io"
	"unicode"
)

type tokenizerImpl struct {
	reader   *bufio.Reader
	position int
	buf      []byte
	eof      bool
	next     byte

	overlapLength int
	wasDM         bool
	wasES         bool
	wasWS         bool
	wasNB         bool
	si            int

	token *Token
}

var _ ITokenizer = &tokenizerImpl{}

const MAX_WORD_LENGTH = 50

func newTokenizerImpl(in io.Reader) *tokenizerImpl {
	return &tokenizerImpl{
		reader:   bufio.NewReader(in),
		position: -1,
		buf:      make([]byte, 0, MAX_WORD_LENGTH),
		token:    &Token{},
	}
}

func (t *tokenizerImpl) Next() *Token {
	if t.eof {
		return t.token.SetEoF(t.position)
	}
	t.buf = t.buf[:0]
	t.si = t.position
	if t.next == 0 {
		t.si = t.si + 1
	}
	t.overlapLength = 0
	t.wasDM = false
	t.wasES = false
	t.wasWS = false
	t.wasNB = false

	for {
		var b byte
		var err error
		if t.next != 0 {
			b = t.next
			t.next = 0
		} else {
			b, err = t.reader.ReadByte()
			t.position++
		}
		if err != nil {
			t.eof = true
			if len(t.buf) == 0 {
				return t.token.SetEoF(t.position)
			}
			return t.BuildToken()
		}

		// игнорируем простые апострофы
		if b == '\'' {
			continue
		}

		if b >= '0' && b <= '9' {
			if len(t.buf) > 0 && !t.wasNB {
				t.next = b
				return t.BuildToken()
			}
			t.wasNB = true
			t.add(b)
			continue
		}

		if (b >= '.' || b == ',') && t.wasNB {
			t.add(b)
			continue
		}

		if b == ' ' || b == '\t' {
			if len(t.buf) > 0 && !t.wasWS {
				t.next = b
				return t.BuildToken()
			}
			t.wasWS = true
			t.add(b)
			continue
		}

		if b == '\n' || b == '\r' || b == '.' /* || b == '!' || b == '?' */ {
			if len(t.buf) > 0 && !t.wasES {
				t.next = b
				return t.BuildToken()
			}
			t.wasES = true
			t.add(b)
			continue
		}

		if b == ',' || b == ':' || b == '-' || b == ';' || b == '!' || b == '?' {
			if len(t.buf) > 0 && !t.wasDM {
				t.next = b
				return t.BuildToken()
			}
			t.wasDM = true
			t.add(b)
			continue
		}

		if t.wasDM || t.wasES || t.wasWS {
			t.next = b
			return t.BuildToken()
		}

		t.add(b)

	}
}

func (t *tokenizerImpl) BuildToken() *Token {
	t.token.si = t.si
	if t.overlapLength == 0 {
		var tp TokenType
		switch {
		case t.wasES:
			tp = TOKEN_ES
		case t.wasWS:
			tp = TOKEN_WS
		case t.wasDM:
			tp = TOKEN_DM
		case t.wasNB:
			tp = TOKEN_NB
		default:
			wasRus := false
			wasLat := false
			wasOther := false
			for _, s := range string(t.buf) {
				switch {
				case unicode.Is(unicode.Latin, s):
					wasLat = true
				case unicode.Is(unicode.Cyrillic, s):
					wasRus = true
				default:
					wasOther = true
				}
				if wasOther {
					break
				}
			}
			tp = TOKEN_WD
			if wasOther || (wasRus && wasLat) {
				tp = TOKEN_UK
			}
		}
		t.token.tp = tp
		t.token.data = t.buf
		t.token.ei = t.si + len(t.token.data) - 1
		return t.token
	}
	t.token.ei = t.position
	t.token.data = t.token.data[:0]
	t.token.tp = TOKEN_LC
	return t.token
}

func (t *tokenizerImpl) add(b byte) {
	if len(t.buf) == MAX_WORD_LENGTH {
		t.overlapLength++
	} else {
		t.buf = append(t.buf, b)
	}
}
