package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/edwin9870/quiz/internal/util"
)

func Read(filePath string) util.Answers {

	file, err := os.Open(filePath)
	util.CheckIfError(err)

	reader := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)
	typedAnswers := make([]string, 0)
	correctAnswers := make([]string, 0)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		correctAnswers = append(correctAnswers, record[1])
		fmt.Printf("%v? ", record[0])

		scanner.Scan()
		typedAnswers = append(typedAnswers, scanner.Text())
	}

	return util.CheckAnswers(correctAnswers, typedAnswers)
}
