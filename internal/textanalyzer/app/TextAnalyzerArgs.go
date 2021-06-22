package app

import (
	"flag"
	"os"
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
	pprofhttp *int
}



func NewTextAnalyzerArgsF() *TextAnalyzerArgs {
	ENV_DEBUG, exists := os.LookupEnv("DEBUG")
	var env_debug bool
	if exists {
		env_debug = ENV_DEBUG == "true"
	}
	return &TextAnalyzerArgs{
		size:      flag.Int("size", 10, "Collect top SIZE frequent words"),
		minlen:    flag.Int("minlen", 4, "Min length for word in symmbols"),
		useFirst:  flag.Bool("first", false, "Include first words of statements"),
		useLast:   flag.Bool("last", false, "Include last words of statements"),
		nonfreq:   flag.Bool("nonfreq", false, "Less frequent, not more frequent"),
		json:      flag.Bool("json", false, "Treat file as JSON with standard schema, not as plain"),
		cpuprof:   flag.String("cpuprof", "", "File for storing CPU profiler info"),
		memprof:   flag.String("memprof", "", "File for storing Mem profiler info"),
		debug:     flag.Bool("debug", env_debug, "Debug mode - NOT CLEAR HOW TO USE"),
		http:      flag.Int("http", -1, "If set - port to listen in http mode"),
		pprofhttp: flag.Int("pprofhttp", -1, "If set - port to listen for http prof"),
	}
}

func (t *TextAnalyzerArgs) Parse() {
	flag.Parse()
}

func (t TextAnalyzerArgs) Size() int {
	return *t.size
}

func (t TextAnalyzerArgs) Minlen() int {
	return *t.minlen
}

func (t TextAnalyzerArgs) UseFirst() bool {
	return *t.useFirst
}

func (t TextAnalyzerArgs) UseLast() bool {
	return *t.useLast
}

func (t TextAnalyzerArgs) Nonfreq() bool {
	return *t.nonfreq
}

func (t *TextAnalyzerArgs) Cpuprof() string {
	return *t.cpuprof
}

func (t *TextAnalyzerArgs) Memprof() string {
	return *t.memprof
}

func (t *TextAnalyzerArgs) Debug() bool {
	return *t.debug
}

func (t *TextAnalyzerArgs) Http() int {
	return *t.http
}

func (t *TextAnalyzerArgs) Pprofhttp() int {
	return *t.pprofhttp
}

func (t *TextAnalyzerArgs) Json() bool {
	return *t.json
}