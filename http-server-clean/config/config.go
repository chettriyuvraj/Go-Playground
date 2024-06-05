package config

import (
	"log"
	"net/http"
)

type AppConfig struct {
	Logger  *log.Logger /* An example of an 'injection', we could also add things like a DB connection, etc. */
	Handler func(*http.ResponseWriter, *http.Request, *log.Logger)
}

func (appConfig *AppConfig) ServeHTTP(w *http.ResponseWriter, req *http.Request) {
	appConfig.Handler(w, req, appConfig.Logger)
}
