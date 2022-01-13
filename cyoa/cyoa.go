package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

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
	chapters, err := parseChapters("/Users/hankehly/Projects/gophercises/cyoa/gopher.json")

	if err != nil {
		log.Fatalln(err)
	}

	// Argument to New must be base name of template file
	// https://pkg.go.dev/text/template@go1.17.6#Template.ParseFiles
	tmpl, err := template.New("template.html").ParseFiles("/Users/hankehly/Projects/gophercises/cyoa/template.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.Execute(os.Stdout, chapters["intro"])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tmpl)
}

// Test comment 123
// Read and unmarshal json file containing chapters data
func parseChapters(jsonPath string) (map[string]Chapter, error) {
	var chapters map[string]Chapter
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &chapters)
	if err != nil {
		return nil, err
	}
	return chapters, nil
}
