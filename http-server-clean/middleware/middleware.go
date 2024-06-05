package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			timeStart := time.Now()
			handler.ServeHTTP(w, req)
			timeTaken := time.Since(timeStart)
			log.Printf("Path: %s; Method: %s, Time taken: %s", req.URL.EscapedPath(), req.Method, timeTaken)
		},
	)
}
