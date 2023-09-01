package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func readCsvFile(csvFile string) ([][]string, error){
	// Open csv file
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
  if err != nil {
		return nil, err     
  }

	// Print the CSV data
	for _, row := range data {
		fmt.Printf("Question:%s, Answer: %s ", row[0], row[1])
		fmt.Println()
	}

	defer file.Close()
	return data,nil
}

func getMathAnswerFromString(str string) int {
	operator := string(str[1]) // find the operator of question
	strArr := strings.Split(str, operator)

	for _, value := range strArr {
		fmt.Println(value)
	}
	return 0
}




