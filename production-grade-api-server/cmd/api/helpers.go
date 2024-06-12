package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(req *http.Request) (id int64, err error) {
	params := httprouter.ParamsFromContext(req.Context())

	id, err = strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		return -1, fmt.Errorf("error parsing movie id: %v", err)
	}
	if id < 1 {
		return -1, fmt.Errorf("invalid id")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, data interface{}, statusCode int, headers http.Header) error {
	/* Marshal to json */
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	/* Set any headers */
	for k, v := range headers {
		w.Header()[k] = v
	}

	/* Send response */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
	return nil
}
