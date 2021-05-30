package main

import (
	"flag"
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"os"
)

func main() {
	var size = flag.Int("size", 10, "Collect top SIZE frequent words")
	var minlen = flag.Int("minlen", 4, "Min length for word in symmbols")
	var useFirst = flag.Bool("first", false, "Include first words of statements")
	var useLast = flag.Bool("last", false, "Include last words of statements")
	var nonfreq = flag.Bool("nonfreq", false, "Less frequent, not more frequent")
	flag.Parse()
	filter := index.NewTermFilter(*minlen, *useFirst, *useLast, *nonfreq)
	stats := index.CollectStats(os.Stdin, filter)
	result := stats.Find(*size, nil)
	for _, v := range result {
		fmt.Printf("%v\n", *v)
	}
}
