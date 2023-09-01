package handlecsv

import (
	"strconv"
	"testing"
)

func TestAnswer(t *testing.T) {
	data, _ := ReadCsvFile("problem.csv")

	for _, value := range data {
		question := value[0]
		answer, _ := strconv.Atoi(value[1])

		calculation, _ := GetMathAnswerFromString(question)
		if calculation != answer {
			t.Errorf("Result was incorrect, got: %d, want: %d.", calculation, answer)
		}
	}
}