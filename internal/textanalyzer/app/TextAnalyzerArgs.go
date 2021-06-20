package app

import "flag"

type TextAnalyzerArgs struct {
	size     *int
	minlen   *int
	useFirst *bool
	useLast  *bool
	nonfreq  *bool
	json     *bool
	cpuprof  *string
	memprof  *string
}



func (t *TextAnalyzerArgs) Json() bool {
	return *t.json
}

func NewTextAnalyzerArgsF() *TextAnalyzerArgs {
	return &TextAnalyzerArgs{
		size:     flag.Int("size", 10, "Collect top SIZE frequent words"),
		minlen:   flag.Int("minlen", 4, "Min length for word in symmbols"),
		useFirst: flag.Bool("first", false, "Include first words of statements"),
		useLast:  flag.Bool("last", false, "Include last words of statements"),
		nonfreq:  flag.Bool("nonfreq", false, "Less frequent, not more frequent"),
		json:     flag.Bool("json", false, "Treat file as JSON with standard schema, not as plain"),
		cpuprof:  flag.String("cpuprof", "", "File for storing CPU profiler info"),
		memprof:  flag.String("memprof", "", "File for storing Mem profiler info"),
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