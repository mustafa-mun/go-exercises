package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)


func main() {
	
	csvPtr := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	csv := *csvPtr
	data, err := ReadCsvFile(csv)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		processInput(data, &wg)
	}()

	timer := time.NewTimer(time.Duration(*limitPtr) * time.Second)
	
	go func() {
		<-timer.C
		fmt.Println("Time is out")
		wg.Done()
	}()

	wg.Wait()
}


func processInput(data [][]string, wg *sync.WaitGroup) {
	defer wg.Done() // Close the waitgroup when the repl loop is over

	var questionNumber int = 1
	var pointer int = 0
	var correctAnswerCount = 0
	var examLength string = strconv.Itoa(len(data))
	reader := bufio.NewScanner(os.Stdin)
	
	questionNumberString := strconv.Itoa(questionNumber)
	printPrompt("Quesion #"+ questionNumberString + " "+ data[pointer][0])
	// Begin the repl loop

	for reader.Scan() {
		if pointer == len(data) - 1 { // Stop the loop if user answers all questions correctly
			fmt.Println("You scored", examLength, "out of", examLength, "questions.")
			return
		}
		text := cleanInput(reader.Text()) // Get the input
		answer := data[pointer][1] 
		// Check if answer is correct
		if text != answer {
			// Answer is not correct 
			fmt.Println("You scored", strconv.Itoa(correctAnswerCount), "out of", examLength, "questions.")
			return
		}
		// Answer is correct
		correctAnswerCount += 1
		pointer += 1
		questionNumber += 1 
		printPrompt("Quesion #"+ strconv.Itoa(questionNumber) + " "+ data[pointer][0])

	}
	
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}