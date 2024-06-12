package main

import (
	"encoding/json"
	"net/http"
)

const version = "1.0"

func (app *application) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	/* Parse status to json JSON */
	stats := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	jsonData, err := json.Marshal(stats)
	if err != nil {
		app.logger.Printf("error: %v", err)
		http.Error(w, "internal error while checking health", http.StatusInternalServerError)
		return
	}

	/* Send response */
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
