package register

import (
	"http-server-clean/app"
	"net/http"
)

func SetupHandlers(mux *http.ServeMux, config *app.AppConfig) {
	mux.Handle("/ok", &app.App{Config: config, Handler: HandleOkRoute})
}
