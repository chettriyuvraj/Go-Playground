package main

import (
	"net/http"
)

const version = "1.0"

func (app *application) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	/* Construct dummy stats */
	stats := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	/* Add headers */
	headers := map[string][]string{
		"Content-Type": {"application/json"},
	}

	/* Write json */
	err := app.writeJSON(w, envelope{"stats": stats}, http.StatusOK, headers)
	if err != nil {
		app.serverErrorResponse(w, req, err)
	}

}
