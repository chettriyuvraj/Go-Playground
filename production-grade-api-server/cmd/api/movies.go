package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chettriyuvraj/go-playground/production-grade-api-server/internal/data"
	"github.com/chettriyuvraj/go-playground/production-grade-api-server/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, req, &input)
	if err != nil { /* Bad responses 400 = When there is an error during decoding  */
		app.errorResponse(w, req, http.StatusBadRequest, err.Error())
		return
	}

	movie := data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()
	data.ValidateMovie(v, &movie)
	if !v.Valid() {
		app.failedValidationResponse(w, req, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input) /* TODO: cant we use writeJSON here directly? */

}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	/* Parse named parameter id */
	id, err := app.readIDParam(req)
	if err != nil {
		app.notFoundResponse(w, req)
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
		app.serverErrorResponse(w, req, err)
	}
}
