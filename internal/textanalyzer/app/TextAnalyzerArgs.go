package app

import (
	"flag"
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"os"
	"strconv"
)

type TextAnalyzerArgs struct {
	size      *int
	minlen    *int
	useFirst  *bool
	useLast   *bool
	nonfreq   *bool
	json      *bool
	cpuprof   *string
	memprof   *string
	debug     *bool
	http      *int
	pprofhttp *string
}

type NonFlagsAnalyzerConfig struct {
	Size      int
	Minlen    int
	UseFirst  bool
	UseLast   bool
	Nonfreq   bool
	Json      bool
	Cpuprof   string
	Memprof   string
	Debug     bool
	Http      int
	Pprofhttp string
}

func NewTextAnalyzerArgsNF(config NonFlagsAnalyzerConfig) *TextAnalyzerArgs {
	result := &TextAnalyzerArgs{
		minlen:    &config.Minlen,
		useFirst:  &config.UseFirst,
		useLast:   &config.UseLast,
		nonfreq:   &config.Nonfreq,
		json:      &config.Json,
		cpuprof:   &config.Cpuprof,
		memprof:   &config.Memprof,
		debug:     &config.Debug,
		http:      &config.Http,
		pprofhttp: &config.Pprofhttp,
	}
	return result
}

func NewTextAnalyzerArgsF() *TextAnalyzerArgs {
	ENV_DEBUG, exists := os.LookupEnv("DEBUG")
	var env_debug bool
	if exists {
		env_debug = ENV_DEBUG == "true"
	}
	result := &TextAnalyzerArgs{
		size:      flag.Int("size", 10, "Find top SIZE frequent words"),
		minlen:    flag.Int("minlen", 4, "Min length for word in symmbols"),
		useFirst:  flag.Bool("first", false, "Include first words of statements"),
		useLast:   flag.Bool("last", false, "Include last words of statements"),
		nonfreq:   flag.Bool("nonfreq", false, "Less frequent, not more frequent"),
		json:      flag.Bool("json", false, "Treat file as JSON with standard schema, not as plain"),
		cpuprof:   flag.String("cpuprof", "", "File for storing CPU profiler info"),
		memprof:   flag.String("memprof", "", "File for storing Mem profiler info"),
		debug:     flag.Bool("debug", env_debug, "Debug mode - NOT CLEAR HOW TO USE"),
		http:      flag.Int("http", -1, "If set - port to listen in http mode"),
		pprofhttp: flag.String("pprofhttp", "no", "If set - port to listen for http prof - if `same` or `` same as http"),
	}
	result.Parse()
	return result
}

func (a *TextAnalyzerArgs) GetCollectorConfig() index.CollectConfig {
	return index.CollectConfig{Filter: a.GetStatisticsFilter(), Mode: a.GetReadMode()}
}

func (a *TextAnalyzerArgs) GetStatisticsFilter() *index.TermFilter {
	return index.NewTermFilter(
		index.TermFilterOptions{
			MinLen:       a.Minlen(),
			IncludeFirst: a.UseFirst(),
			IncludeLast:  a.UseLast(),
			ReverseFreq:  a.Nonfreq(),
		},
	)
}

func (a *TextAnalyzerArgs) GetReadMode() index.ReadMode {
	var mode index.ReadMode
	if a.Json() {
		mode = index.MODE_PARALLEL_JSON
	} else {
		mode = index.MODE_PLAIN
	}
	return mode
}

type PprofHttpMode int

const (
	PPROF_NONE PprofHttpMode = 0
	PPROF_SELF PprofHttpMode = 1
	PPROF_MAIN PprofHttpMode = 2
)

func (a *TextAnalyzerArgs) IsHttpMode() bool {
	return a.Http() > 0
}

func (a *TextAnalyzerArgs) PprofHttpMode() PprofHttpMode {
	switch {
	case !a.IsHttpMode() || a.Pprofhttp() <= 0:
		return PPROF_NONE
	case a.Http() != a.Pprofhttp() :
		return PPROF_SELF
	case a.Http() == a.Pprofhttp() :
		return PPROF_MAIN
	}
	return PPROF_NONE
}

func (a *TextAnalyzerArgs) Parse() {
	flag.Parse()
}

func (a TextAnalyzerArgs) Size() int {
	return *a.size
}

func (a TextAnalyzerArgs) Minlen() int {
	return *a.minlen
}

func (a TextAnalyzerArgs) UseFirst() bool {
	return *a.useFirst
}

func (a TextAnalyzerArgs) UseLast() bool {
	return *a.useLast
}

func (a TextAnalyzerArgs) Nonfreq() bool {
	return *a.nonfreq
}

func (a *TextAnalyzerArgs) Cpuprof() string {
	return *a.cpuprof
}

func (a *TextAnalyzerArgs) Memprof() string {
	return *a.memprof
}

func (a *TextAnalyzerArgs) Debug() bool {
	return *a.debug
}

func (a *TextAnalyzerArgs) Http() int {
	return *a.http
}

func (a *TextAnalyzerArgs) Pprofhttp() int {
	switch {
	case !a.IsHttpMode() || *a.pprofhttp == "no":
		return -1
	case *a.pprofhttp == "same" || *a.pprofhttp == "" :
		return a.Http()
	default:
		res,_ := strconv.Atoi(*a.pprofhttp)
		return res
	}
}

func (a *TextAnalyzerArgs) Json() bool {
	return *a.json
}
