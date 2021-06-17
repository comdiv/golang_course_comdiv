package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"testing"
)

func Benchmark_NonFilteredCollection(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectFromReader(testdata_test.TestDataReader(), nil,0, index.MODE_PLAIN)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJson(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectFromReader(testdata_test.TestDataJsonReader(), nil, 0, index.MODE_JSON)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJsonParallel(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectFromReader(testdata_test.TestDataJsonReader(), nil, 0, index.MODE_PARALLEL_JSON)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredSearch(b *testing.B) {
	stats := index.CollectFromReader(testdata_test.TestDataReader(), nil, 0, index.MODE_PLAIN)
	var result []*index.TermStat
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       4,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}

func Benchmark_PreFilteredCollection(b *testing.B) {
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       10,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats = index.CollectFromReader(testdata_test.TestDataReader(), query, 0, index.MODE_PLAIN)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_PreFilteredSearch(b *testing.B) {
	query := index.NewTermFilter(index.TermFilterOptions{
		MinLen:       10,
		IncludeFirst: false,
		IncludeLast:  false,
		ReverseFreq:  false,
	})
	stats := index.CollectFromReader(testdata_test.TestDataReader(), query, 0, index.MODE_PLAIN)
	var result []*index.TermStat

	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}
