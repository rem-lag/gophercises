package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type QuizGame interface {
	runQuiz(i int, correct *int, t *time.Timer)
}

type GameRound struct {
	game QuizGame
}

type qa struct {
	q string
	a string
}

func (prob qa) runQuiz(i int, correct *int, t *time.Timer) {

	fmt.Printf("Question #%d: %s = \n", i+1, prob.q)

	ansCh := make(chan string)
	go func() {
		var a string
		fmt.Scanf("%s\n", &a)
		ansCh <- a
	}()

	select {
	case <-t.C:
		exit(fmt.Sprintf("Your score was %d correct questions", *correct), 0)
	case ans := <-ansCh:
		if ans == prob.a {
			*correct++
		}
	}
}

func parseQa(lines [][]string) []GameRound {
	ret := make([]GameRound, len(lines))

	for i, val := range lines {
		q := qa{
			q: strings.TrimSpace(val[0]),
			a: strings.TrimSpace(val[1]),
		}
		ret[i] = GameRound{
			game: q,
		}
	}

	return ret
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file with questions and answers")
	timeLimit := flag.Int("limit", 30, "the time limit in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file %s\n", *csvFile), 1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Could not parse provided CSV", 1)
	}

	correct := 0
	problems := parseQa(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, prob := range problems {
		prob.game.runQuiz(i, &correct, timer)
	}
}
