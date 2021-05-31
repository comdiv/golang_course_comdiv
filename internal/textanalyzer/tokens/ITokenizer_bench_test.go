package tokens_test

import (
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/tokens"
	"testing"
)

func BenchmarkTokenizer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := tokens.CountTokensR(testdata_test.TestDataReader())
		if result < 30000 {
			panic("Что-то он не то считает")
		}
	}
}
