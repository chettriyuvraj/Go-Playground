package handlers

import (
	"fmt"
	"http-server-clean/app"
	"net/http"
)

func HandleOkRoute(w http.ResponseWriter, req *http.Request, config *app.AppConfig) {
	/* Maybe you ordinarily would check the method as well */
	fmt.Fprintf(w, "This is an ok response!")
}
