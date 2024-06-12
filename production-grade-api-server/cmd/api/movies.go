package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chettriyuvraj/go-playground/production-grade-api-server/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the createMovie route..")
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	/* Parse named parameter id */
	id, err := app.readIDParam(req)
	if err != nil {
		app.logger.Printf("error: %v", err)
		http.NotFound(w, req)
		return
	}

	/* Create dummy movie */
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Date(2011, time.August, 2, 3, 4, 5, 0, time.Local),
		Title:     "Udaan",
		Year:      2033,
		Runtime:   32,
		Genres:    []string{"Coming of age"},
		Version:   1,
	}

	/* Write dummy movie as json to client */
	headers := map[string][]string{
		"Content-Type": {"application/json"},
	}
	err = app.writeJSON(w, envelope{"movie": movie}, http.StatusOK, headers)
	if err != nil {
		app.logger.Printf("error: %v", err)
		http.Error(w, "internal server error while retreiving movie", http.StatusInternalServerError)
		return
	}
}
