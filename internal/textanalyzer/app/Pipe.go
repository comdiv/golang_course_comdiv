package app

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"os"
)

func PipeMain(args *TextAnalyzerArgs) {
	filter := args.GetStatisticsFilter()
	stats := index.CollectFromReader(os.Stdin, args.GetCollectorConfig())
	result := stats.Find(args.Size(), filter)
	for _, v := range result {
		fmt.Printf("%v\n", *v)
	}
}
