package main

import (
	"fmt"
	"net/http"
)

const version = "1.0"

func (a *app) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", a.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
