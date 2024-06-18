package main

import (
	"net/http"
)

/*
Generic handler to log errors for server
TODO: Modify for structured logging
*/
func (app *application) logError(req *http.Request, err error) {
	app.logger.Println(err)
}

/*
Generic handler to return error responses with a particular response code
TODO: Modify for structured logging
*/
func (app *application) errorResponse(w http.ResponseWriter, req *http.Request, statusCode int, message interface{}) {
	err := app.writeJSON(w, envelope{"error": message}, statusCode, http.Header{})
	/* If err returned, fall back to a standard 500 */
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter, req *http.Request) {
	message := "this resource does not exist on the server"
	app.errorResponse(w, req, http.StatusNotFound, message)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.logError(req, err)

	message := "this resource does not exist on the server"
	app.errorResponse(w, req, http.StatusInternalServerError, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, req *http.Request) {
	message := "this method is not supported"
	app.errorResponse(w, req, http.StatusMethodNotAllowed, message)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, req *http.Request, errors map[string]string) {
	app.errorResponse(w, req, http.StatusUnprocessableEntity, errors)
}
