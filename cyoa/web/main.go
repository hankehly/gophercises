package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gophercises/m/v2/cyoa"
)

func main() {
	// Todo: What is a more elegant way of handling null user input?
	unset := "UNSET"

	port := flag.Int("port", 3030, "CYOA web app port")
	jsonStoryPath := flag.String("jsonStoryPath", unset, "Path to JSON story")

	flag.Parse()

	if *jsonStoryPath == unset {
		log.Fatalln("jsonStoryPath is required")
	}
	f, err := os.Open(*jsonStoryPath)
	if err != nil {
		log.Fatal(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	// Set a custom template using function options like this
	// tmpl := template.Must(template.New("notImportant").Parse("hello world"))
	// setTemplate := cyoa.WithTemplate(*tmpl)
	// h := cyoa.NewHandler(story, setTemplate)
	//
	// Or use default values (set inside NewHandler)
	h := cyoa.NewHandler(story)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Printf("Started the web application on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}
