package lexemes_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/lexemes"
	"os"
	"testing"
)

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("../main_test_text.txt")
		if err != nil {
			panic(err)
		}
		result := lexemes.CountLexemesR(f)
		if result < 13000 {
			panic("Что-то он не то считает")
		}
	}
}
