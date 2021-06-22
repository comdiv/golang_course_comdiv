package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func HttpMain(args *TextAnalyzerArgs) {
	var wg sync.WaitGroup
	// handler to close from main server for gracefull shutdown
	var pprofserver *http.Server
	if args.PprofHttpMode() == PPROF_SELF {
		wg.Add(1)
		pprofserver = PprofStartStandaloneServer(args.Pprofhttp(), func(){wg.Done()})
	}
	wg.Add(1)
	var mainmux *http.ServeMux
	if args.PprofHttpMode() == PPROF_SELF {
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
		if args.PprofHttpMode() == PPROF_MAIN {
			fmt.Println("Start listen main with pprof on "+strconv.Itoa(args.Http()))
		}else{
			fmt.Println("Start listen main on "+strconv.Itoa(args.Http()))
		}
		fmt.Println(mainserver.ListenAndServe())
	}()
	wg.Wait()
}
