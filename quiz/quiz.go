package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "A CSV file containing records in format (question, answer)")
	seconds := flag.Int("limit", 30, "Stop the quiz as soon as the this limit has exceeded")

	flag.Parse()

	fmt.Printf("You have %d seconds to answer all of the questions. Press [enter] to start the quiz ", *seconds)
	fmt.Scanf("\n")

	f, err := os.Open(*csvPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	problems := parseRecords(records)
	correct := 0
	timer := time.NewTimer(time.Duration(*seconds) * time.Second)

	for i, p := range problems {
		fmt.Printf("%d) %s ", i+1, p.q)
		aC := make(chan string)
		go func() {
			var a string
			fmt.Scanf("%s\n", &a)
			aC <- a
		}()
		select {
		case <-timer.C:
			fmt.Print("\nSorry, you ran out of time.\n")
			total := len(records)
			score := float64(correct) / float64(total) * 100
			fmt.Printf("You scored %d out of %d. That's %.2f%%.\n", correct, total, score)
			return
		case a := <-aC:
			if a == p.a {
				correct++
			}
		}
	}
	total := len(records)
	score := float64(correct) / float64(total) * 100
	fmt.Printf("You scored %d out of %d. That's %.2f%%.\n", correct, total, score)
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, record := range records {
		ret[i] = problem{
			strings.TrimSpace(record[0]),
			strings.TrimSpace(record[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
