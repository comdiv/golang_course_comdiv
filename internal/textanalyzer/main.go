package main

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"os"
)

func main() {
	args := app.NewTextAnalyzerArgsF()
	args.Parse()
	filter := index.NewTermFilter(
		index.TermFilterOptions{
			MinLen:       args.Minlen(),
			IncludeFirst: args.UseFirst(),
			IncludeLast:  args.UseLast(),
			ReverseFreq:  args.Nonfreq(),
		},
	)
	stats := index.CollectStats(os.Stdin, filter)
	result := stats.Find(args.Size(), filter)
	for _, v := range result {
		fmt.Printf("%v\n", *v)
	}
}
