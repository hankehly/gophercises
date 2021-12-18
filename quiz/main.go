package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvpath := flag.String("csv", "problems.csv", "The path to the CSV file")
	flag.Parse()

	f, err := os.Open(*csvpath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.ReuseRecord = true

	var answer, n_correct, n_total int

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		expected, err := strconv.Atoi(record[1])

		if err != nil {
			log.Fatal(err)
		}

		n_total++

		fmt.Printf("%s ", record[0])
		fmt.Scan(&answer)

		if answer == expected {
			n_correct++
		}
	}

	score_pct := float64(n_correct) / float64(n_total) * 100
	fmt.Printf("Score: %d/%d (%.2f%%)\n", n_correct, n_total, score_pct)
}
