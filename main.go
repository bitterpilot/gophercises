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
	csvFilename := flag.String("csv", "problems.csv",
		"A csv file organized as question, answer")
	duration := flag.Int("timer", 30, "the allowed time for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	check(err, "Failed to open csv file")

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	check(err, "Failed to read csv. Is it formatted correctly")

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	var correct uint
	for i, problem := range problems {

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		fmt.Printf("Question %d:\t%s = ", i+1, problem.q)
		select {
		case <-timer.C:
			fmt.Printf("\n\n%d out of %d correct\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				// fmt.Println("Correct!")
				correct++
			}
		}
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{q: line[0], a: strings.TrimSpace(line[1])}
	}
	return problems
}

func check(err error, msg string) {
	if err != nil {
		fmt.Printf("%s:\t%s\n", msg, err)
		os.Exit(1)
	}
}
