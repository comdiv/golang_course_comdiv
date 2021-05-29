package tokenizer

// Token токен - охватывает минимальный примитив из исходного текста
type Token struct {
	// тип токена
	tp   TokenType
	si   int
	ei   int
	data string
}

// TokenType перечисление для типов токена
type TokenType byte

const (
	// TOKEN_UD токен неопределенного типа
	TOKEN_UD TokenType = 0
	// TOKEN_WD токен отдельного слова
	TOKEN_WD TokenType = 1
	// TOKEN_WS токен с пробельными символами
	TOKEN_WS TokenType = 2
	// TOKEN_DM токен с разделителями
	TOKEN_DM TokenType = 4
	// TOKEN_LC токен (без значения) для слитных участок текста, не разбирается
	TOKEN_LC TokenType = 8
)

func NewToken(tp TokenType, si int, data string) *Token {
	return &Token{tp: tp, si: si, ei: si + len(data) - 1, data: data}
}

func NewLargeToken(si int, ei int) *Token {
	return &Token{tp: TOKEN_LC, si: si, ei: ei}
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
