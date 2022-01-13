package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Story map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	http.HandleFunc("/chapters/", index)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	pathStory := "/Users/hankehly/Projects/gophercises/cyoa/story.json"
	pathTemplate := "/Users/hankehly/Projects/gophercises/cyoa/template.html"

	chapters, err := parseStory(pathStory)
	if err != nil {
		log.Fatalln(err)
	}
	// Argument to New must be base name of template file
	// https://pkg.go.dev/text/template@go1.17.6#Template.ParseFiles
	baseName := filepath.Base(pathTemplate)
	tmpl, err := template.New(baseName).ParseFiles(pathTemplate)
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

// Test comment 123
// Read and unmarshal json file containing chapters data
func parseStory(jsonPath string) (map[string]Chapter, error) {
	var story Story
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
