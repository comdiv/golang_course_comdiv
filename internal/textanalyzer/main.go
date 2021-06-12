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
	var stats *index.TermStatCollection
	if args.Json() {
		stats = index.CollectStatsFromJson(os.Stdin, filter)
	} else {
		stats = index.CollectStats(os.Stdin, filter, 0)
	}
	result := stats.Find(args.Size(), filter)
	for _, v := range result {
		fmt.Printf("%v\n", *v)
	}
}
