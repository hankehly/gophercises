package main

import (
	"flag"
	"fmt"
	"gophercises/m/v2/urlshort"
	"net/http"
	"os"
)

// Eg. go run urlshort/main/main.go -path urlshort/urlshort.yaml
func main() {
	path := flag.String("path", "urlshort.yaml", "A path to a yaml file containing an array of items of struct (path: string, url: string)")
	flag.Parse()

	yamlData, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
