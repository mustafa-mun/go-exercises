package main

import (
	"encoding/csv"
	"os"
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



