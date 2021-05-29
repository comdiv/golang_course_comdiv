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
	si := t.position
	if t.next == 0 {
		si = si + 1
	}
	overlapLength := 0
	wasDelimiters := false
	wasWS := false

	buildToken := func() Token {
		if overlapLength == 0 {
			var tp TokenType
			switch {
			case wasWS:
				tp = TOKEN_WS
			case wasDelimiters:
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
					tp = TOKEN_UD
				}
			}
			return NewToken(tp, si, string(t.buf))
		}
		return NewLargeToken(si, t.position)
	}

	add := func(b byte) {
		if len(t.buf) == MAX_WORD_LENGTH {
			overlapLength++
		} else {
			t.buf = append(t.buf, b)
		}
	}

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
			return buildToken()
		}
		if b == ' ' || b == '\n' || b == '\r' || b == '\t' {
			if len(t.buf) > 0 && !wasWS {
				t.next = b
				return buildToken()
			}
			wasWS = true
			add(b)
			continue
		}

		if b == ',' || b == ':' || b == '.' || b == '-' || b == ';' || b == '!' || b == '?' {
			if len(t.buf) > 0 && !wasDelimiters {
				t.next = b
				return buildToken()
			}
			wasDelimiters = true
			add(b)
			continue
		}

		if wasDelimiters || wasWS {
			t.next = b
			return buildToken()
		}

		add(b)

	}
}
