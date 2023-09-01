package main

import (
	"flag"
	"fmt"
	"github.com/mustafa-mun/go-exercises/quiz-game/internal/handlecsv"
)

func main() {

    csvPtr := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
    limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
  
    flag.Parse()

		csv := *csvPtr
		handlecsv.ReadCsvFile(csv)
		// limit := *limitPtr
    fmt.Println("csv:", *csvPtr)
    fmt.Println("limit:", *limitPtr)
		
	
}