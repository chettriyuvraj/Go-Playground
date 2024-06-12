package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the createMovie route..")
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	movie := params.ByName("id")
	fmt.Fprintf(w, "The movie id is %s\n", movie)
}
