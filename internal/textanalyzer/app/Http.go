package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func HttpMain(args *TextAnalyzerArgs) {
	var wg sync.WaitGroup
	// handler to close from main server for gracefull shutdown
	var pprofserver *http.Server
	if args.PprofHttpMode() == PPROF_SELF {
		wg.Add(1)
		pprofserver = PprofStartStandaloneServer(args.Pprofhttp(), func() { wg.Done() })
	}
	wg.Add(1)
	mainmux := createMux(args)
	application := &httpApplicationContext{
		args:            args,
		mux:             mainmux,
		mainserver:      &http.Server{Addr: "127.0.0.1:" + strconv.Itoa(args.Http()), Handler: mainmux},
		pprofserver:     pprofserver,
		indexingService: NewIndexService(args),
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
	} else {
		result = http.DefaultServeMux
	}
	return result
}

type httpApplicationContext struct {
	args            *TextAnalyzerArgs
	mux             *http.ServeMux
	mainserver      *http.Server
	pprofserver     *http.Server
	indexingService *IndexingService
}

func (a *httpApplicationContext) stop() {
	if a.pprofserver != nil {
		a.pprofserver.Shutdown(context.Background())
	}
	a.mainserver.Shutdown(context.Background())
}

func (a *httpApplicationContext) setupHandlers() {
	a.mux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		setupCorsResponse(&writer, request)
		if (request).Method == "OPTIONS" {
			return
		}
		a.stop()
	})
	a.mux.HandleFunc("/reset", ResetHandler(a.indexingService))
	a.mux.HandleFunc("/stat/", StatHandler(a.args, a.indexingService))
	a.mux.HandleFunc("/text", IndexHandler(a.indexingService))
	a.mux.HandleFunc("/index", IndexHandler(a.indexingService))
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

type HttpHandler func(writer http.ResponseWriter, request *http.Request)

func ResetHandler(indexer *IndexingService) HttpHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		setupCorsResponse(&writer, request)
		if (request).Method == "OPTIONS" {
			return
		}
		indexer.Reset()
		writer.Header().Set("Content-Type", "application/json")
		data, _ := json.MarshalIndent(struct {
			Op    string `json:"op"`
			State string `json:"state"`
		}{"reset", "complete"}, "", "    ")
		writer.Write(data)
	}
}

func IndexHandler(indexer *IndexingService) HttpHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		setupCorsResponse(&writer, request)
		if (request).Method == "OPTIONS" {
			return
		}
		jsonpart := extractJsonFromRequest(request)
		data := struct {
			Op    string             `json:"op"`
			State string             `json:"state"`
			Data  index.JsonTextPart `json:"data"`
		}{"text", "success", jsonpart}
		statusCode := http.StatusOK
		switch {
		case jsonpart.Error != nil:
			data.State = "error"
			statusCode = http.StatusInternalServerError
		case jsonpart.Text == "":
			data.State = "empty"
			statusCode = http.StatusBadRequest
		default:
			indexer.Index(jsonpart.Number, jsonpart.Text)
		}
		out, _ := json.MarshalIndent(data, "", "    ")
		writer.WriteHeader(statusCode)
		writer.Write(out)
	}
}

func StatHandler(config *TextAnalyzerArgs, indexer *IndexingService) HttpHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		setupCorsResponse(&writer, request)
		if (request).Method == "OPTIONS" {
			return
		}
		data := struct {
			Op     string               `json:"op"`
			State  string               `json:"state"`
			Size   int                  `json:"size"`
			Error  error                `json:"error"`
			Filter *index.TermFilterDto `json:"filter"`
			Data   []index.TermStatDto  `json:"data"`
		}{"stat", "success", 0, nil, nil, nil}
		statusCode := http.StatusOK
		path := request.URL.Path
		parts := strings.Split(path, "/")
		size := -1
		switch {
		case len(parts) == 3 && parts[2] == "": // /stat ""+"stat" + ""
			size = config.Size()
		case len(parts) == 3: // /stat/10 "" + "stat"+ "10"
			_size, err := strconv.Atoi(parts[2])
			if err != nil {
				data.Error = err
				data.State = "error"
			}
			size = _size
		case len(parts) > 3: // /stat/10/xxx
			data.Error = errors.New("too many config in path")
			data.State = "error"
		}
		data.Size = size
		if size > 0 {
			filter := config.GetStatisticsFilter()
			filterDto := filter.ToDto()
			data.Filter = &filterDto
			terms := indexer.Find(size, filter)
			data.Data = make([]index.TermStatDto, len(terms))
			for i, v := range terms {
				data.Data[i] = v.ToDto()
			}
		} else {
			data.State = "empty"
		}
		out, _ := json.MarshalIndent(data, "", "    ")
		writer.WriteHeader(statusCode)
		writer.Write(out)
	}
}

func extractJsonFromRequest(r *http.Request) index.JsonTextPart {
	result := index.JsonTextPart{}
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		err := json.NewDecoder(r.Body).Decode(&result)
		if err != nil {
			result = index.JsonTextPart{Error: err}
			return result
		}
		return result
	}
	r.ParseForm()
	var params map[string][]string
	if len(r.Form) != 0 {
		params = r.Form
	} else {
		params = r.URL.Query()
	}
	_part, ok := params["number"]
	if !ok || len(_part) == 0 {
		_part = []string{"0"}
	}
	part, err := strconv.Atoi(_part[0])
	if err != nil {
		result.Error = err
		return result
	}
	_text, ok := params["text"]
	if !ok || len(_text) == 0 {
		_text = []string{""}
	}
	text := _text[0]
	result = index.JsonTextPart{
		Number: part,
		Text:   text,
	}
	return result
}

func (a *httpApplicationContext) start() {
	if a.args.PprofHttpMode() == PPROF_MAIN {
		fmt.Println("start listen main with pprof on " + strconv.Itoa(a.args.Http()))
	} else {
		fmt.Println("start listen main on " + strconv.Itoa(a.args.Http()))
	}
	fmt.Println(a.mainserver.ListenAndServe())
}
