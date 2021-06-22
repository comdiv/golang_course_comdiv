package main

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	_ "net/http/pprof"
)

func main() {
	args := app.NewTextAnalyzerArgsF()
	app.PprofStartCpuIfRequired(args)
	defer app.PprofWriteMemoryIfRequired(args)
	if args.IsHttpMode() {
		app.HttpMain(args)
	} else {
		app.PipeMain(args)
	}
}
