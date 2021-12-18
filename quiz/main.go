package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvpath := flag.String("csv", "problems.csv", "A path to a CSV file containing records of format (question, answer)")
	flag.Parse()

	f, err := os.Open(*csvpath)

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
	for i, p := range problems {
		var a string
		fmt.Printf("%d) %s ", i+1, p.q)
		fmt.Scanf("%s\n", &a)
		if a == p.a {
			correct++
		}
	}

	total := len(records)
	score := float64(correct) / float64(total) * 100
	fmt.Printf("You scored %d/%d (%.2f%%)\n", correct, total, score)
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
