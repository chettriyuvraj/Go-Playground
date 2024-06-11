package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type config struct {
	port int
	env  string
}

type app struct {
	config *config
	logger *log.Logger
}

func main() {

	/* Initialize configurations and logger */
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	/* Initialize app */
	app := app{config: &cfg, logger: logger}

	/* Configure the mux */
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	/* Launch server */
	server := &http.Server{
		Addr:         net.JoinHostPort("localhost", strconv.Itoa(cfg.port)),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Print("Launching server..")
	err := server.ListenAndServe()
	log.Fatal(err)

}
