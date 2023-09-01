package handlecsv

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ReadCsvFile(csvFile string) ([][]string, error){
	file, err := os.Open(csvFile) // Open csv file
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file) // Create reader
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll() // Read file
  if err != nil {
		return nil, err     
  }

	defer file.Close()
	return data,nil // Return file data
}

func GetMathAnswerFromString(str string) (int, error) {
	operator := string(str[1]) // find the operator of question
	strArr := strings.Split(str, operator)

	numOne, err := strconv.Atoi(strArr[0])
	if err != nil {
		return 0, err
	}
	numTwo, err := strconv.Atoi(strArr[1])
	if err != nil {
		return 0, err
	}
	answer,err := Calculate(operator, numOne, numTwo)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func Calculate(operator string, numOne,numTwo int) (int, error) {
	switch operator {
	case "+":
		return numOne + numTwo, nil
	case "-":
		return numOne - numTwo, nil
	case "*":
		return numOne * numTwo, nil
	case "/":
		return numOne / numTwo, nil
	}
	return 0, errors.New("operation is not valid")  
}


