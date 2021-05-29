package tokenizer

// Token токен - охватывает минимальный примитив из исходного текста
type Token struct {
	// тип токена
	tp TokenType
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
