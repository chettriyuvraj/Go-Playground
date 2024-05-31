package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	proxy()
}

func proxy() {
	url, err := url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
	}

	resp, err := client.Get("http://localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func redirect() {
	redirectPolicyFunc := func(req *http.Request, via []*http.Request) error {
		fmt.Println("redirect list")
		for _, req := range via {
			header := req.Header
			loc := header.Get("Location")
			fmt.Println(loc)
		}
		return nil
	}

	client := http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Get("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
