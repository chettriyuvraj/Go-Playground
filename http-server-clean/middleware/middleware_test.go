package middleware

import (
	"bytes"
	"http-server-clean/app"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	LOGGER_PREFIX = "[Yuvraj's personal logger:] "
)

var LOGGER_DEFAULT_FLAGS = log.Default().Flags()

func TestLoggingMiddleware(t *testing.T) {
	/* Init structs required for test */
	logBuf := bytes.NewBuffer([]byte{})
	config := app.AppConfig{Logger: log.New(logBuf, LOGGER_PREFIX, LOGGER_DEFAULT_FLAGS)}
	mux := http.NewServeMux()

	/* Init test request and response recorder */
	req := httptest.NewRequest("GET", "/ok", nil)
	recorder := httptest.NewRecorder()

	/* Convert to middleware */
	wrappedMux := LoggingMiddleware(mux, &config)
	wrappedMux.ServeHTTP(recorder, req)

	/* Check output in buffer */
	keywords := []string{"Path:", "Method:", "GET", "Time taken"}
	logOutput := logBuf.String()
	for _, word := range keywords {
		if !strings.Contains(logOutput, word) {
			t.Errorf("Error in logging middleware output: Doesn't contain keyword %s", word)
		}
	}
}
