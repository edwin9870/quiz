package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/edwin9870/quiz/internal/csv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quizz",
	Short: "Quiz is a quizz generator based on csv",
	Long:  `Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello world")
		answers := csv.Read("/home/eramirez/workspace/go/quiz/problems.csv")
		fmt.Printf("Number of correct answers: %v\nNumber of incorrect answers: %v\nTotal numbers of questions: %v", answers.Correct, answers.Invalid, answers.Correct+answers.Invalid)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
