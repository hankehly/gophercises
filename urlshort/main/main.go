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
	// yamlPath := flag.String("yaml", "urlshort.yaml", "A path to a YAML file containing an array of items of struct (path: string, url: string)")
	jsonPath := flag.String("json", "urlshort.json", "A path to a JSON file containing an array of items of struct (path: string, url: string)")

	flag.Parse()

	// yamlData, err := os.ReadFile(*yamlPath)
	jsonData, err := os.ReadFile(*jsonPath)
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
	// yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
	jsonHandler, err := urlshort.JSONHandler(jsonData, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
