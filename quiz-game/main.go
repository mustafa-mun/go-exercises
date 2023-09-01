package main

import (
	"flag"
	"fmt"
)

func main() {

    csvPtr := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
    limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
  
    flag.Parse()

		csv := *csvPtr
		readCsvFile(csv)
		// limit := *limitPtr
    fmt.Println("csv:", *csvPtr)
    fmt.Println("limit:", *limitPtr)
		_ = getMathAnswerFromString("5+5")
}