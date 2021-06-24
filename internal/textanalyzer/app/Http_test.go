package app_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func prepareTestMux() *http.ServeMux {
	args := app.NewTextAnalyzerArgsNF(app.NonFlagsAnalyzerConfig{Minlen: 4, UseFirst: false, UseLast: false})
	indexer := app.NewIndexService(args)
	return app.NewIndexerMux(indexer)
}
func TestHttp_Index(t *testing.T) {
	mux := prepareTestMux()
	indexcontent := `{"number":1, "text":"some simple test"}`
	req := httptest.NewRequest("POST", "http://myhost/index", strings.NewReader(indexcontent))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h, _ := mux.Handler(req)
	h.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type") )
	assert.Equal(t, `{
    "op": "text",
    "state": "success",
    "data": {
        "Number": 1,
        "Text": "some simple test",
        "Error": null
    }
}`, string(body))
}
func TestHttp_Reset(t *testing.T) {
	mux := prepareTestMux()
	req := httptest.NewRequest("GET", "http://myhost/reset",nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h, _ := mux.Handler(req)
	h.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type") )
	assert.Equal(t, `{
    "op": "reset",
    "state": "success"
}`, string(body))
}

func TestHttp_Stat(t *testing.T) {
	mux := prepareTestMux()
	indexcontent := `{"number":1, "text":"some simple test with some word repeat 3 times and word the word (word) repeat four times in text"}`
	req := httptest.NewRequest("POST", "http://myhost/index", strings.NewReader(indexcontent))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h, _ := mux.Handler(req)
	h.ServeHTTP(w, req)

	req = httptest.NewRequest("GET", "http://myhost/stat/3", strings.NewReader(indexcontent))
	w = httptest.NewRecorder()
	h, _ = mux.Handler(req)
	h.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type") )
	assert.Equal(t, `{
    "op": "stat",
    "state": "success",
    "size": 3,
    "error": null,
    "filter": {
        "minlen": 4,
        "include_first": false,
        "include_last": false,
        "reverse_freq": false
    },
    "data": [
        {
            "value": "word",
            "first_part": 1,
            "first_index": 5,
            "len": 4,
            "count": 4,
            "first_count": 0,
            "last_count": 0
        },
        {
            "value": "repeat",
            "first_part": 1,
            "first_index": 6,
            "len": 6,
            "count": 2,
            "first_count": 0,
            "last_count": 0
        },
        {
            "value": "times",
            "first_part": 1,
            "first_index": 7,
            "len": 5,
            "count": 2,
            "first_count": 0,
            "last_count": 0
        }
    ]
}`, string(body))
}
