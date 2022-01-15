package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	aliastest "github.com/hankehly/gophercises/pkg/cyoa"
	"github.com/hankehly/gophercises/pkg/cyoa/cyoasubpkg"
)

func main() {
	port := flag.Int("port", 3030, "CYOA web app port")
	jsonStoryPath := flag.String("jsonStoryPath", "data/cyoastory.json", "Path to JSON story")

	flag.Parse()

	f, err := os.Open(*jsonStoryPath)
	if err != nil {
		log.Fatal(err)
	}
	story, err := aliastest.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	// Set a custom template using function options like this
	// tmpl := template.Must(template.New("notImportant").Parse("hello world"))
	// setTemplate := cyoa.WithTemplate(*tmpl)
	// h := cyoa.NewHandler(story, setTemplate)
	//
	// Or use default values (set inside NewHandler)
	h := cyoasubpkg.NewHandler(story)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Printf("Started the web application on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}
