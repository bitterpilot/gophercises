package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv",
		"A csv file organized as question, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	check(err, "Failed to open csv file")

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	check(err, "Failed to read csv. Is it formatted correctly")

	problems := parseLines(lines)
	var correct uint
	for i, problem := range problems {
		fmt.Printf("Question %d:\t%s = ", i+1, problem.q)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == problem.a {
			fmt.Println("Correct!")
			correct++
		}
	}
	fmt.Printf("%d out of %d correct\n", correct, len(problems))
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
