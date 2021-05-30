package index_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"os"
	"testing"
)

func Benchmark_NonFilteredCollection(b *testing.B) {
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		f, e := os.Open("../main_test_text.txt")
		if f == nil || e != nil {
			b.Fatal(f, e)
		}
		stats = index.CollectStats(f, nil)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_NonFilteredSearch(b *testing.B) {

	f, e := os.Open("../main_test_text.txt")
	if f == nil || e != nil {
		b.Fatal(f, e)
	}
	stats := index.CollectStats(f, nil)
	var result []*index.TermStat
	query := index.NewTermFilter(4, false, false, false)
	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}

func Benchmark_PreFilteredCollection(b *testing.B) {
	query := index.NewTermFilter(10, false, false, false)
	var stats *index.TermStatCollection
	for i := 0; i < b.N; i++ {
		f, e := os.Open("../main_test_text.txt")
		if f == nil || e != nil {
			b.Fatal(f, e)
		}
		stats = index.CollectStats(f, query)
		if stats == nil {
			b.Fatal("null stats")
		}
	}
}

func Benchmark_PreFilteredSearch(b *testing.B) {
	query := index.NewTermFilter(10, false, false, false)
	f, e := os.Open("../main_test_text.txt")
	if f == nil || e != nil {
		b.Fatal(f, e)
	}
	stats := index.CollectStats(f, query)
	var result []*index.TermStat

	for i := 0; i < b.N; i++ {
		result = stats.Find(10, query)
		if len(result) != 10 {
			b.Fatal()
		}
	}
}
