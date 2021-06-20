package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	testdata_test "github.com/comdiv/golang_course_comdiv/internal/textanalyzer/testdata"
	"testing"
)

func Benchmark_NonFilteredCollection(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _ = index.CollectFromReader(testdata_test.TestDataReader(), index.CollectConfig{Mode: index.MODE_PLAIN})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJson(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _ = index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Mode: index.MODE_PLAIN})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJsonParallel_8(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _  = index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Mode: index.MODE_PARALLEL_JSON})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}
func Benchmark_NonFilteredCollectionJsonParallel_4(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _  = index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Mode: index.MODE_PARALLEL_JSON, Workers: 4})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJsonParallel_2(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _  = index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Mode: index.MODE_PARALLEL_JSON, Workers: 2})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredCollectionJsonParallel_16(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		stats, _  = index.CollectFromReader(testdata_test.TestDataJsonReader(), index.CollectConfig{Mode: index.MODE_PARALLEL_JSON, Workers: 16})
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredSearch(b *testing.B) {
	stats, _ := index.CollectFromReader(testdata_test.TestDataReader(), index.CollectConfig{Mode: index.MODE_PLAIN})
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
		stats, _ = index.CollectFromReader(testdata_test.TestDataReader(), index.CollectConfig{Filter: query, Mode: index.MODE_PLAIN})
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
	stats, _ := index.CollectFromReader(testdata_test.TestDataReader(),index.CollectConfig{Filter: query, Mode: index.MODE_PLAIN})
	var result []*index.TermStat

	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}
