package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"gophercises/m/v2/cyoa"
)

func main() {
	http.HandleFunc("/chapters/", index)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	path := "/Users/hankehly/Projects/gophercises/cyoa/story.json"

	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	chapters, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatalln(err)
	}
	// Argument to New must be base name of template file
	// https://pkg.go.dev/text/template@go1.17.6#Template.ParseFiles
	// pathTemplate := "/Users/hankehly/Projects/gophercises/cyoa/template.html"
	// baseName := filepath.Base(pathTemplate)
	// tmpl, err := template.New(baseName).ParseFiles(pathTemplate)

	tmpl, err := template.New("test").Parse(cyoa.Template)
	if err != nil {
		log.Fatalln(err)
	}

	parts := strings.Split(r.URL.Path, "/")
	key := parts[len(parts)-1]

	if chapter, ok := chapters[key]; ok {
		err = tmpl.Execute(rw, chapter)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
