package middleware

import (
	"http-server-clean/app"
	"net/http"
	"time"
)

func LoggingMiddleware(handler http.Handler, config *app.AppConfig) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			timeStart := time.Now()
			handler.ServeHTTP(w, req)
			timeTaken := time.Since(timeStart)
			config.Logger.Printf("(Path: %s; Method: %s, Time taken: %s)", req.URL.EscapedPath(), req.Method, timeTaken)
		},
	)
}
