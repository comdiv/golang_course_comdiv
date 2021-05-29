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
	wasDelimiters bool
	wasWS         bool
	si            int
}

var _ ITokenizer = &tokenizerImpl{}

const MAX_WORD_LENGTH = 50

func newTokenizerImpl(in io.Reader) *tokenizerImpl {
	return &tokenizerImpl{
		reader:   bufio.NewReader(in),
		position: -1,
		buf:      make([]byte, 0, MAX_WORD_LENGTH),
	}
}

func (t *tokenizerImpl) Next() Token {
	if t.eof {
		return EofToken(t.position)
	}
	t.buf = t.buf[:0]
	t.si = t.position
	if t.next == 0 {
		t.si = t.si + 1
	}
	t.overlapLength = 0
	t.wasDelimiters = false
	t.wasWS = false

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
				return EofToken(t.position)
			}
			return t.BuildToken()
		}
		if b == ' ' || b == '\n' || b == '\r' || b == '\t' {
			if len(t.buf) > 0 && !t.wasWS {
				t.next = b
				return t.BuildToken()
			}
			t.wasWS = true
			t.add(b)
			continue
		}

		if b == ',' || b == ':' || b == '.' || b == '-' || b == ';' || b == '!' || b == '?' {
			if len(t.buf) > 0 && !t.wasDelimiters {
				t.next = b
				return t.BuildToken()
			}
			t.wasDelimiters = true
			t.add(b)
			continue
		}

		if t.wasDelimiters || t.wasWS {
			t.next = b
			return t.BuildToken()
		}

		t.add(b)

	}
}

func (t *tokenizerImpl) BuildToken() Token {
	if t.overlapLength == 0 {
		var tp TokenType
		switch {
		case t.wasWS:
			tp = TOKEN_WS
		case t.wasDelimiters:
			tp = TOKEN_DM
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
		return NewToken(tp, t.si, string(t.buf))
	}
	return NewLargeToken(t.si, t.position)
}

func (t *tokenizerImpl) add(b byte) {
	if len(t.buf) == MAX_WORD_LENGTH {
		t.overlapLength++
	} else {
		t.buf = append(t.buf, b)
	}
}
