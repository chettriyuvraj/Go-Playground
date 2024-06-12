package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the createMovie route..")
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, req)
		return
	}

	fmt.Fprintf(w, "The movie id is %d\n", id)
}
