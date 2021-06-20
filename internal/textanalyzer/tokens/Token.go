package tokens

import (
	"strings"
)

// Token токен - охватывает минимальный примитив из исходного текста
type Token struct {
	// тип токена
	tp   TokenType
	si   int
	ei   int
	hasUpper bool
	isAscii bool
	data []byte
}

func (t *Token) HasUpper() bool {
	return t.hasUpper
}

func (t *Token) IsAscii() bool {
	return t.isAscii
}

func (t *Token) Copy() Token {
	d := make([]byte, len(t.data))
	copy(d, t.data)
	return Token{
		tp:   t.tp,
		si:   t.si,
		ei:   t.ei,
		isAscii: t.isAscii,
		hasUpper: t.hasUpper,
		data: d,
	}
}

// TokenType перечисление для типов токена
type TokenType byte

const (
	// TOKEN_UD токен не определен
	TOKEN_UD TokenType = 0
	// TOKEN_UK токен неопределенного типа строки
	TOKEN_UK TokenType = 1
	// TOKEN_WD токен отдельного слова
	TOKEN_WD TokenType = 2
	// TOKEN_ES токен с пробельными символами
	TOKEN_ES TokenType = 4
	// TOKEN_WS токен с пробельными символами внутри предложения
	TOKEN_WS TokenType = 8
	// TOKEN_DM токен с разделителями внутри предложения помимо пробелов
	TOKEN_DM TokenType = 16
	// TOKEN_LC токен (без значения) для слитных участок текста, не разбирается
	TOKEN_LC TokenType = 32
	// TOKEN_NB нечто численное - требуется для корректной обработки точек и запятых
	TOKEN_NB TokenType = 64
	// TOKEN_EOF признак конца файла
	TOKEN_EOF TokenType = 128
)

func EofToken(si int) *Token {
	return &Token{tp: TOKEN_EOF, si: si, ei: si}
}

func NewTokenPlus(tp TokenType, si int, data string, ascii bool, hasUpper bool) *Token {
	return &Token{tp: tp, si: si, ei: si + len(data) - 1, data: []byte(data), isAscii: ascii, hasUpper: hasUpper}
}

func NewToken(tp TokenType, si int, data string) *Token {
	return &Token{tp: tp, si: si, ei: si + len(data) - 1, data: []byte(data), isAscii: true, hasUpper: false}
}

func NewLargeToken(si int, ei int) *Token {
	return &Token{tp: TOKEN_LC, si: si, ei: ei, data: make([]byte, 0), isAscii: false, hasUpper: false}
}

func (t *Token) Type() TokenType {
	return t.tp
}

func (t *Token) Start() int {
	return t.si
}

func (t *Token) Finish() int {
	return t.ei
}

func (t *Token) Length() int {
	if t.Type() == TOKEN_LC {
		return t.Finish() - t.Start() + 1
	} else {
		return len(t.data)
	}
}

func (t *Token) Value() string {
	return string(t.data)
}

func (t *Token) NormalValue() string {
	result := t.Value()
	if !t.isAscii || t.hasUpper {
		result = strings.ToLower(result)
		if strings.Index(result,"ё") != -1 {
			result = strings.Replace(result, "ё", "е", -1)
		}
	}
	return result
}

func (t *Token) Data() []byte {
	return t.data
}

func (t *Token) IsEof() bool {
	return t.Type() == TOKEN_EOF
}
func (t *Token) IsWord() bool {
	return t.Type() == TOKEN_WD
}
func (t *Token) IsNumber() bool {
	return t.Type() == TOKEN_NB
}

func (t *Token) IsEoS() bool {
	return t.Type() == TOKEN_ES
}

func (t *Token) IsUnknown() bool {
	return t.Type() == TOKEN_UK
}

func (t *Token) IsUndefined() bool {
	return t.Type() == TOKEN_UD
}

func (t *Token) SetEoF(pos int) *Token {
	t.si = pos
	t.ei = pos
	t.tp = TOKEN_EOF
	t.data = t.data[:0]
	return t
}
