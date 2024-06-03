package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	/* Checking out servers */
	// go createCustomHTTPServer()
	// go createDefaultHTTPServer()
	// for {

	// }

	/* Exploring url package */
	exploreURLPackage()
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
	setupCustomMuxHandlers(mux)

	/* Create a server instance on your own */
	server := http.Server{Addr: ":8081", Handler: mux}

	/* Create server - indicates a custom mux will be utilized */
	log.Fatal(server.ListenAndServe())
}

func setupCustomMuxHandlers(mux *http.ServeMux) {
	/* There are two ways to register a handler for a route */
	/* First way */
	mux.HandleFunc("/custommuxhttp", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Custom Mux: Handler Func\n")
	})

	/* Second way */
	mux.Handle("/custommuxhttp2", &CustomHandlerForCustomMux{})
}

func exploreURLPackage() {

	/* Let's parse it into a parsedURL object */
	rawURL := "https://www.example.com/path%20spaced/innerpath?name=yuvraj%20shivkumar&surname=chettri#footer%20spaced"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}

	/* Explore some basic parameters and see if they're what we expect */
	fmt.Printf("Our raw URL is: %s\n\nLet's explore the parsedURL object that was parsed by parsedURL.Parse\n\n", rawURL)
	fmt.Printf("Scheme: %s\n", parsedURL.Scheme)
	fmt.Printf("Host: %s\n", parsedURL.Host)
	fmt.Printf("Path: %s\n", parsedURL.Path)
	fmt.Printf("Raw Path: %s\n", parsedURL.RawPath) /* Quirk: This is the escaped path - don't rely on it, it might be empty, use EscapedPath() method */
	fmt.Printf("Escaped Path: %s\n", parsedURL.EscapedPath())
	fmt.Printf("Query: %s\n", parsedURL.Query())
	fmt.Printf("Raw Query: %s\n", parsedURL.RawQuery)             /* Reliable, same as Query().Encode(), %20 for spaces */
	fmt.Printf("Encoded Query: %s\n", parsedURL.Query().Encode()) /* Reliable, same as RawQuery, '+' for spaces */
	fmt.Printf("Fragment: %s\n", parsedURL.Fragment)
	fmt.Printf("Raw Fragment: %s\n", parsedURL.RawFragment) /* Quirk: This is the escaped fragment - don't rely on it, it might be empty, use EscapedFragment() method */
	fmt.Printf("Escaped Fragment: %s\n", parsedURL.EscapedFragment())

	/* Exploring the query separately */
	fmt.Println()
	query := parsedURL.Query()
	for k, vList := range query {
		fmt.Printf("Key: %s\nValues:\n", k)
		for i, v := range vList {
			fmt.Printf("%d. Unescaped: %s\n", i+1, v)
			fmt.Printf("%d. Escaped: %s\n", i+1, url.QueryEscape(v)) // corresponding path.escape also exists
		}
		fmt.Println()
	}

}
