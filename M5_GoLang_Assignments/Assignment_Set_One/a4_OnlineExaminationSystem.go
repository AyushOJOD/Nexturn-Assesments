package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Questions struct {
	Question string
	Options  []string
	Answer   int
}

const (
	ExcellentScore = 80
	GoodScore      = 60
	PassScore      = 40
)

var QuestionsList = []Questions{
	{
		Question: "What is the capital of India?",
		Options:  []string{"Mumbai", "Delhi", "Kolkata", "Chennai"},
		Answer:   1,
	},
	{
		Question: "What is the capital of USA?",
		Options:  []string{"New York", "Washington DC", "Los Angeles", "Chicago"},
		Answer:   1,
	},
	{
		Question: "What is the capital of UK?",
		Options:  []string{"Manchester", "Birmingham", "London", "Liverpool"},
		Answer:   2,
	},
}

func startQuiz(questions []Questions) int {
	score := 0
	reader := bufio.NewReader(os.Stdin)
	for i, ques := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, ques.Question)
		for j, option := range ques.Options {
			fmt.Printf("%d, %s\n", j+1, option)
		}

		fmt.Print("Enter your answer (or 'exit'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting quiz")
			break
		}

		answer, err := strconv.Atoi(input)
		if err != nil || answer < 1 || answer > len(ques.Options) {
			fmt.Println("Invalid input")
			continue
		}

		if answer-1 == ques.Answer {
			fmt.Println("Correct answer!")
			score++
		} else {
			fmt.Println("Wrong answer!")
		}
	}
	return score
}

func evaluateScore(score, totalQuestions int) {
	percentage := float64(score) / float64(totalQuestions) * 100

	if percentage >= ExcellentScore {
		fmt.Println("Excellent Performance!")
	} else if percentage >= GoodScore {
		fmt.Println("Good Performance!")
	} else if percentage >= PassScore {
		fmt.Println("Pass!")
	} else {
		fmt.Println("Failed!")
	}
}

// func main() {
// 	score := startQuiz(QuestionsList)
// 	fmt.Printf("\nScore: %d/%d\n", score, len(QuestionsList))

// 	evaluateScore(score, len(QuestionsList))
// }
