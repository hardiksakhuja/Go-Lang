package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	CsvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*CsvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s\n", *CsvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	// timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChan <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, ques := range lines {
		ret[i] = problem{
			q: ques[0],
			a: strings.TrimSpace(ques[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
