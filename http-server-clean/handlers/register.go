package handlers

import (
	"http-server-clean/app"
	"net/http"
)

func Register(mux *http.ServeMux, config *app.AppConfig) {
	mux.Handle("/ok", &app.App{Config: config, Handler: HandleOkRoute})
}
