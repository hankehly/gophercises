package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"gophercises/m/v2/cyoa"
)

func main() {
	port := flag.Int("port", 3030, "CYOA web app port")
	flag.Parse()
	http.HandleFunc("/", index)

	fmt.Printf("Started the web application on http://localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func index(rw http.ResponseWriter, r *http.Request) {
	storyPath := "/Users/hankehly/Projects/gophercises/cyoa/story.json"

	f, err := os.Open(storyPath)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(rw, fmt.Sprintln("Internal server error"), http.StatusInternalServerError)
	}

	chapters, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(rw, fmt.Sprintln("Internal server error"), http.StatusInternalServerError)
	}

	// Argument to New must be base name of template file (if using a file)
	// https://pkg.go.dev/text/template@go1.17.6#Template.ParseFiles
	// pathTemplate := "/Users/hankehly/Projects/gophercises/cyoa/template.html"
	// baseName := filepath.Base(pathTemplate)
	// tmpl, err := template.New(baseName).ParseFiles(pathTemplate)

	tmpl := template.Must(template.New("doesNotMatter").Parse(cyoa.Template))

	path := strings.TrimSpace(r.URL.Path)
	parts := strings.Split(path, "/")
	key := parts[len(parts)-1]
	if key == "" {
		key = "intro"
	}

	if chapter, ok := chapters[key]; ok {
		err = tmpl.Execute(rw, chapter)
		if err != nil {
			fmt.Printf("%v", err)
			http.Error(rw, fmt.Sprintln("Internal server error"), http.StatusInternalServerError)
		}
		return
	}
	http.Error(rw, "Invalid chapter name", http.StatusNotFound)
}
