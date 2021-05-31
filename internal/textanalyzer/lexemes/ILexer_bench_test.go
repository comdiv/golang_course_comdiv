package lexemes_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"testing"
)

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := lexemes.CountLexemesR(testdata_test.TestDataReader())
		if result < 13000 {
			panic("Что-то он не то считает")
		}
	}
}
