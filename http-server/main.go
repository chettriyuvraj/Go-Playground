package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"time"
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
	// exploreURLPackage()

	/* Exploring request struct */
	// exploreRequestStruct()

	/* Demo a streaming server - launch and send a get req to /stream at :8082 */
	demoStreamingServer()
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
	query := parsedURL.Query() /* Has Add/Get/Del methods */
	for k, vList := range query {
		fmt.Printf("Key: %s\nValues:\n", k)
		for i, v := range vList {
			fmt.Printf("%d. Unescaped: %s\n", i+1, v)
			fmt.Printf("%d. Escaped: %s\n", i+1, url.QueryEscape(v)) // corresponding path.escape also exists
		}
		fmt.Println()
	}

}

/* Quickly taking a look at the Request struct */
func exploreRequestStruct() {
	req := httptest.NewRequest("GET", "http://www.example.com/test-req", io.NopCloser(bytes.NewReader([]byte("This is a test body"))))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "test/header")
	fmt.Println()
	fmt.Printf("Lets check out the request struct:\n\n")
	fmt.Printf("URL: %s\n", req.URL)
	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("Proto: %s\n", req.Proto)
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	fmt.Printf("Body: %s\n", buf)
	fmt.Printf("Headers: \n")
	headers := req.Header /* Has Add/Get/Del methods */
	for k, vList := range headers {
		fmt.Printf("Key: %s\n", k)
		fmt.Printf("Values:\n")
		for i, v := range vList {
			fmt.Printf("%d.  %s\n", i, v)
		}
	}

	/* Other stuff includes: Trailer, Response which contains redirect Response, Form, Multipart which is handled in another package */

}

/* Let's see how a server serves data in 'streaming' fashion with a small delay */
func demoStreamingServer() {

	/* Initialize the server and single handler that handles the streaming */
	mux := http.NewServeMux()
	streamHandleFunc := func(w http.ResponseWriter, req *http.Request) {
		var wg sync.WaitGroup

		/* Define both sides of streaming */
		writeToPipeWithDelay := func(w *io.PipeWriter) {
			defer w.Close()
			defer wg.Done()
			w.Write([]byte("Hey!"))
			time.Sleep(time.Second * 1)
			w.Write([]byte(" This"))
			time.Sleep(time.Second * 1)
			w.Write([]byte(" is"))
			time.Sleep(time.Second * 1)
			w.Write([]byte(" a"))
			time.Sleep(time.Second * 1)
			w.Write([]byte(" streamed reply :)"))
		}

		readFromPipeAndWriteToResponse := func(r *io.PipeReader, w http.ResponseWriter) {
			defer wg.Done()

			/* Initialize */
			defer r.Close()
			data := make([]byte, 200)

			/* Check if flush supported */
			f, flushSupported := w.(http.Flusher)

			/* Set headers, second one to ensure no buffering on client side */
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("X-Content-Type-Options", "nosniff")

			for {
				/* Read data */
				n, err := r.Read(data)
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}

				/* Write to response writer */
				_, err = w.Write(data[:n])
				if err != nil {
					log.Fatal(err)
				}

				/* Flush if possible */
				if flushSupported {
					f.Flush()
				}
			}
		}

		/* Create pipe */
		pipeReader, pipeWriter := io.Pipe()
		wg.Add(2)
		go readFromPipeAndWriteToResponse(pipeReader, w)
		go writeToPipeWithDelay(pipeWriter)
		wg.Wait()
	}

	/* Register with mux */
	mux.HandleFunc("/stream", streamHandleFunc)

	/* Create server and serve */
	server := http.Server{Addr: ":8082", Handler: mux}
	log.Fatal(server.ListenAndServe())
}
