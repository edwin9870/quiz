package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/edwin9870/quiz/internal/util"
)

func Read(filePath string) util.Answers {

	file, err := os.Open(filePath)
	util.CheckIfError(err)

	reader := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)
	typedAnswers := make([]string, 0)
	correctAnswers := make([]string, 0)
	answer := make(chan string)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		correctAnswers = append(correctAnswers, record[1])
		fmt.Printf("\n%v? ", record[0])

		go func() {
			scanner.Scan()
			answer <- scanner.Text()
		}()

		select {
		case typedAnswer := <-answer:
			typedAnswers = append(typedAnswers, typedAnswer)

		case <-time.After(7 * time.Second):
			typedAnswers = append(typedAnswers, "bad")

		}
	}

	return util.CheckAnswers(correctAnswers, typedAnswers)
}
