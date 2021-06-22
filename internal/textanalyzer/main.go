package main

import (
	"context"
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := app.NewTextAnalyzerArgsF()
	args.Parse()
	app.PprofStartCpuIfRequired(args)

	if args.IsHttpMode() {
		var wg sync.WaitGroup
		// handler to close from main server for gracefull shutdown
		var pprofserver *http.Server
		if args.PprofHttpMode() == app.PPROF_SELF {
			wg.Add(1)
			// for pprof using default mux
			pprofserver = &http.Server{Addr: "127.0.0.1:"+strconv.Itoa(args.Pprofhttp())}
			go func() {
				defer wg.Done()
				fmt.Println("Start listen pprof on "+strconv.Itoa(args.Pprofhttp()))
				fmt.Println(pprofserver.ListenAndServe())
			}()
		}
		wg.Add(1)
		var mainmux *http.ServeMux
		if args.PprofHttpMode() == app.PPROF_SELF {
			mainmux = http.NewServeMux()
		}else{
			mainmux = http.DefaultServeMux
		}
		mainserver := &http.Server{Addr: "127.0.0.1:"+strconv.Itoa(args.Http()), Handler: mainmux}
		mainmux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
			if pprofserver != nil {
				pprofserver.Shutdown(context.Background())
			}
			mainserver.Shutdown(context.Background())
		})
		go func() {
			defer wg.Done()
			if args.PprofHttpMode() == app.PPROF_MAIN {
				fmt.Println("Start listen main with pprof on "+strconv.Itoa(args.Http()))
			}else{
				fmt.Println("Start listen main on "+strconv.Itoa(args.Http()))
			}
			fmt.Println(mainserver.ListenAndServe())
		}()
		wg.Wait()
	} else {
		filter := args.GetStatisticsFilter()
		stats := index.CollectFromReader(os.Stdin, args.GetCollectorConfig())
		result := stats.Find(args.Size(), filter)
		for _, v := range result {
			fmt.Printf("%v\n", *v)
		}
	}

	app.PprofWriteMemoryIfRequired(args)
}
