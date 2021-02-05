package util

type Answers struct {
	Correct int8
	Invalid int8
}

func CheckAnswers(correctAnswers []string, userAnswers []string) Answers {

	answers := Answers{}

	for i, v := range correctAnswers {
		if v == userAnswers[i] {
			answers.Correct++
		} else {
			answers.Invalid++
		}
	}

	return answers
}
