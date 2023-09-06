package main

import (
	"fmt"
	"net/http"
	"time"

	urlshort "github.com/mustafa-mun/go-exercises/url-shortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := urlshort.UrlPaths{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"":                "",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the JSONHandler using the mapHandler as the
	// fallback

	json := `[
		{"path":"/dogs","url":"https://www.nationalgeographic.com/animals/mammals/facts/domestic-dog"},
		{"path":"/cats","url":"https://www.nationalgeographic.com/animals/mammals/facts/domestic-cat"}
	]`

	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	// Build the YAMLHandler using the JSONHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), jsonHandler)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Starting the server on :8080")
	srv := &http.Server{
    ReadTimeout:       1 * time.Second,
    WriteTimeout:      1 * time.Second,
    IdleTimeout:       30 * time.Second,
    ReadHeaderTimeout: 2 * time.Second,
    Handler:           yamlHandler,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
