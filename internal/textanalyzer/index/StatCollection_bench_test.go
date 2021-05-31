package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"testing"
)

func Benchmark_NonFilteredCollection(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectStats(testdata_test.TestDataReader(), nil)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredSearch(b *testing.B) {
	stats := index.CollectStats(testdata_test.TestDataReader(), nil)
	var result []*index.TermStat
	query := index.NewTermFilterArgs(4, false, false, false)
	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}

func Benchmark_PreFilteredCollection(b *testing.B) {
	query := index.NewTermFilterArgs(10, false, false, false)
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectStats(testdata_test.TestDataReader(), query)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_PreFilteredSearch(b *testing.B) {
	query := index.NewTermFilterArgs(10, false, false, false)
	stats := index.CollectStats(testdata_test.TestDataReader(), query)
	var result []*index.TermStat

	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}
