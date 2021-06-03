package tokens

// Token токен - охватывает минимальный примитив из исходного текста
type Token struct {
	// тип токена
	tp   TokenType
	si   int
	ei   int
	data []byte
}

func (t *Token) Copy() Token {
	d := make([]byte, len(t.data))
	copy(d, t.data)
	return Token{
		tp:   t.tp,
		si:   t.si,
		ei:   t.ei,
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

func NewToken(tp TokenType, si int, data string) *Token {
	return &Token{tp: tp, si: si, ei: si + len(data) - 1, data: []byte(data)}
}

func NewLargeToken(si int, ei int) *Token {
	return &Token{tp: TOKEN_LC, si: si, ei: ei, data: make([]byte, 0)}
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
