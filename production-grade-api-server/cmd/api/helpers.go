package main

import (
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
		return -1, fmt.Errorf("invalid id", err)
	}

	return id, nil
}
