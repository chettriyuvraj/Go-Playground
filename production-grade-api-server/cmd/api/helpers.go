package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

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
	// jsonData, err := json.MarshalIndent(data, "prefix", "\t")
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

func (app *application) readJSON(w http.ResponseWriter, req *http.Request, dst interface{}) error {
	err := json.NewDecoder(req.Body).Decode(&dst)
	if err != nil { /* Bad responses 400 = When there is an error during decoding  */
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("syntax error at offset %d", syntaxError.Offset)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON at offset %d", unmarshalTypeError.Offset)

		case errors.As(err, &invalidUnmarshalError): /* Will occur when we pass an invalid pointer to decode */
			panic(err)

		case errors.Is(err, io.ErrUnexpectedEOF): /* At times, decode returns this - we send a generic response here */
			return errors.New("badly formed JSON")
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		default:
			return err
		}
	}

	return nil
}
