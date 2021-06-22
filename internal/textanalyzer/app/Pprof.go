package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
)

// запускает мониторинг CPU если в аргументах указан файл для хранения результата
func PprofStartCpuIfRequired(args *TextAnalyzerArgs) {
	if args.Cpuprof()!="" {
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
}

// выполняет мониторинг памяти в конце работы приложения если настроен соответствующий файл
func PprofWriteMemoryIfRequired(args *TextAnalyzerArgs) {
	if args.Memprof() != "" {
		f, err := os.Create(args.Memprof())
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

// стартует сервер, позволяет вызвать колбак после окончания его работы
func PprofStartStandaloneServer(port int, oncomplete func()) *http.Server {
	pprofserver := &http.Server{Addr: "127.0.0.1:"+strconv.Itoa(port)}
	// можно бы было зарегистрировать и как ShutdownHook, но мы именно
	// фиксируем точку закрытия горутины
	go func() {
		if oncomplete!=nil {
			defer oncomplete()
		}
		fmt.Println("start listen pprof on "+strconv.Itoa(port))
		fmt.Println(pprofserver.ListenAndServe())
	}()
	return pprofserver
}
