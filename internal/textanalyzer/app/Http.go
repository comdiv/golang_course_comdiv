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
	mainmux := createMux(args)
	application := &httpApplicationContext{
		args: args,
		mux: mainmux,
		mainserver: &http.Server{Addr: "127.0.0.1:"+strconv.Itoa(args.Http()), Handler: mainmux},
		pprofserver: pprofserver,
	}
	application.setupHandlers()
	go func() {
		defer wg.Done()
		application.start()
	}()
	wg.Wait()
}

func createMux(args *TextAnalyzerArgs) *http.ServeMux {
	var result *http.ServeMux
	if args.PprofHttpMode() == PPROF_SELF {
		result = http.NewServeMux()
	}else{
		result = http.DefaultServeMux
	}
	return result
}

type httpApplicationContext struct {
	args *TextAnalyzerArgs
	mux *http.ServeMux
	mainserver *http.Server
	pprofserver *http.Server
}

func (a *httpApplicationContext) stop() {
	if a.pprofserver != nil {
		a.pprofserver.Shutdown(context.Background())
	}
	a.mainserver.Shutdown(context.Background())
}

func (a *httpApplicationContext) setupHandlers() {
	a.mux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		a.stop()
	})
}

func (a *httpApplicationContext) start(){
	if a.args.PprofHttpMode() == PPROF_MAIN {
		fmt.Println("start listen main with pprof on "+strconv.Itoa(a.args.Http()))
	}else{
		fmt.Println("start listen main on "+strconv.Itoa(a.args.Http()))
	}
	fmt.Println(a.mainserver.ListenAndServe())
}
