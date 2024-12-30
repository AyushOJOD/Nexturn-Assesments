package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Question struct to hold question details
type Question struct {
	QuestionText string
	Options      [4]string
	CorrectIndex int
}

var questionBank = []Question{
	{
		QuestionText: "What is the capital of France?",
		Options:      [4]string{"Berlin", "Paris", "Madrid", "Rome"},
		CorrectIndex: 1,
	},
	{
		QuestionText: "Which programming language is known as 'Go'?",
		Options:      [4]string{"Python", "Java", "Golang", "Ruby"},
		CorrectIndex: 2,
	},
	{
		QuestionText: "What is 5 + 3?",
		Options:      [4]string{"6", "7", "8", "9"},
		CorrectIndex: 2,
	},
}

func takeQuiz() int {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questionBank {
		fmt.Printf("Question %d: %s\n", i+1, question.QuestionText)
		for j, option := range question.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		for {
			fmt.Print("Enter your answer (1-4 or 'exit' to quit): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				fmt.Println("Exiting the quiz...")
				return score
			}

			selectedOption, err := strconv.Atoi(input)
			if err != nil || selectedOption < 1 || selectedOption > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			if selectedOption-1 == question.CorrectIndex {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong. The correct answer is:", question.Options[question.CorrectIndex])
			}
			break
		}
	}

	return score
}

func classifyPerformance(score int) {
	fmt.Printf("\nYour final score is: %d/%d\n", score, len(questionBank))

	if score == len(questionBank) {
		fmt.Println("Excellent!")
	} else if score >= len(questionBank)/2 {
		fmt.Println("Good.")
	} else {
		fmt.Println("Needs Improvement.")
	}
}

func main() {
	fmt.Println("Welcome to the Online Examination System")
	score := takeQuiz()
	classifyPerformance(score)
}
