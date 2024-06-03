package main

import (
	"fmt"
	"log"
	"net/http"
)

type CustomHandlerForDefaultMux struct{}

func (handler *CustomHandlerForDefaultMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Default Mux: Handler Struct\n")
}

type CustomHandlerForCustomMux struct{}

func (handler *CustomHandlerForCustomMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Custom Mux: Handler Struct\n")
}

func main() {
	go createCustomHTTPServer()
	go createDefaultHTTPServer()
	for {

	}
}

/* Use the default mux to create a server */
func createDefaultHTTPServer() {
	/* There are two ways to register a handler for a route */

	/* First way */
	http.HandleFunc("/defaultmuxhttp", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Default Mux: Handler Func\n")
	})

	/* Second way */
	http.Handle("/defaultmuxhttp2", &CustomHandlerForDefaultMux{})

	/* Create server - the nil indicates that the http package's default ServeMux will be utilized, ServeMux can be considered to be the 'router' */
	log.Fatal(http.ListenAndServe(":8080", nil))

}

/* Use the non-default mux to create a server */
func createCustomHTTPServer() {

	mux := http.NewServeMux()
	/* There are two ways to register a handler for a route */

	/* First way */
	mux.HandleFunc("/newmuxhttp", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Custom Mux: Handler Func\n")
	})

	/* Second way */
	mux.Handle("/newmuxhttp2", &CustomHandlerForCustomMux{})

	/* Create server - indicates a custom mux will be utilized */
	log.Fatal(http.ListenAndServe(":8081", mux))

}
