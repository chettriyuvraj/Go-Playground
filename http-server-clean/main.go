package main

import (
	"http-server-clean/app"
	"http-server-clean/handlers"
	"http-server-clean/middleware"
	"log"
	"net/http"
	"os"
)

const (
	LOGGER_PREFIX = "[Yuvraj's personal logger:] "
)

var LOGGER_DEFAULT_FLAGS = log.Default().Flags()

func main() {
	/* Setup config, routers and middleware */
	config := app.AppConfig{Logger: log.New(os.Stdout, LOGGER_PREFIX, LOGGER_DEFAULT_FLAGS)}
	mux := http.NewServeMux()
	handlers.Register(mux, &config)
	wrappedMux := middleware.Register(mux, &config)

	/* Launch server */
	server := http.Server{Addr: ":8081", Handler: wrappedMux}
	server.ListenAndServe()

}
