package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the createMovie route..")
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	id, err := app.readIDParam(req)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	fmt.Fprintf(w, "The movie id is %d\n", id)
}
