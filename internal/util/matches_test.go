package util

import "testing"

func TestCheckAnswers(t *testing.T) {
	correctAnswers := []string{"1", "Edwin", "Adios"}
	userAnswers := []string{"1", "Cindy", "Adios"}

	result := CheckAnswers(correctAnswers, userAnswers)

	if result.Correct != 2 || result.Invalid != 1 {
		t.Errorf("Invalid calculation, received: %v", result)
	}

}
