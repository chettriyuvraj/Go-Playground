package main

import (
	"http-server-clean/app"
	"http-server-clean/handlers"
	"http-server-clean/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	/* Setup config, routers and middleware */
	config := app.AppConfig{Logger: log.New(os.Stdout, "[Yuvraj's personal logger:] ", log.Default().Flags())}
	mux := http.NewServeMux()
	handlers.Register(mux, &config)
	wrappedMux := middleware.Register(mux, &config)

	/* Launch server */
	server := http.Server{Addr: ":8081", Handler: wrappedMux}
	server.ListenAndServe()

}
