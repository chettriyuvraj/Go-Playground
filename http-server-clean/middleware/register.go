package middleware

import (
	"http-server-clean/app"
	"net/http"
)

func Register(mux *http.ServeMux, config *app.AppConfig) http.HandlerFunc {
	return LoggingMiddleware(mux, config)
}
