package main

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	args := app.NewTextAnalyzerArgsF()
	args.Parse()

	if args.Cpuprof() != "" {
		f, err := os.Create(args.Cpuprof())
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	filter := index.NewTermFilter(
		index.TermFilterOptions{
			MinLen:       args.Minlen(),
			IncludeFirst: args.UseFirst(),
			IncludeLast:  args.UseLast(),
			ReverseFreq:  args.Nonfreq(),
		},
	)
	var mode index.ReadMode
	if args.Json() {
		mode = index.MODE_PARALLEL_JSON
	} else {
		mode = index.MODE_PLAIN
	}
	stats,_ := index.CollectFromReader(os.Stdin, index.CollectConfig{Filter:filter, Mode:mode})
	result := stats.Find(args.Size(), filter)
	for _, v := range result {
		fmt.Printf("%v\n", *v)
	}

	if args.Memprof() != "" {
		f, err := os.Create(args.Memprof())
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
