package tokens_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"os"
	"testing"
)

func BenchmarkTokenizer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("../main_test_text.txt")
		if err != nil {
			panic(err)
		}
		result := tokens.CountTokensR(f)
		if result < 30000 {
			panic("Что-то он не то считает")
		}
	}
}
