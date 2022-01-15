package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hankehly/gophercises/pkg/cyoa"
)

func main() {
	jsonStoryPath := flag.String("jsonStoryPath", "data/cyoastory.json", "Path to JSON story")
	flag.Parse()

	f, err := os.Open(*jsonStoryPath)
	if err != nil {
		log.Fatal(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	chapter := story["intro"]
	for {
		show(chapter)
		if len(chapter.ChapterOptions) == 0 {
			fmt.Printf("You have reached the end of the story!\n\n")
			os.Exit(0)
		}
		arc := next(chapter)
		chapter = story[arc]
	}
}

func show(chapter cyoa.Chapter) {
	fmt.Printf("-------------------------------\n")
	fmt.Printf("Chapter: %s\n\n", chapter.Title)
	for _, paragraph := range chapter.Story {
		fmt.Printf("%s\n\n", paragraph)
	}
}

func next(chapter cyoa.Chapter) string {
	fmt.Println("Please choose an option and press [Enter] or type 'exit' to quit")

	for i, opt := range chapter.ChapterOptions {
		fmt.Printf(" %d) %s\n", i+1, opt.Text)
	}

	var a string
	fmt.Printf("\n>>> ")
	fmt.Scanf("%s\n", &a)

	if a == "exit" {
		fmt.Printf("Cya!\n")
		os.Exit(0)
	}

	aI, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("Invalid input: %s\n", a)
	} else if aI > len(chapter.ChapterOptions) {
		log.Fatalf("Invalid choice: %d\n", aI)
	}

	return chapter.ChapterOptions[aI-1].Arc
}
