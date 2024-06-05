package handlers

import (
	"fmt"
	"http-server-clean/app"
	"net/http"
)

func HandleOkRoute(w http.ResponseWriter, req *http.Request, config *app.AppConfig) {
	fmt.Fprintf(w, "This is an ok response!")
}
