package app

import "net/http"

type App struct {
	Config  *AppConfig
	Handler func(w http.ResponseWriter, req *http.Request, config *AppConfig)
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.Handler(w, req, a.Config)
}
