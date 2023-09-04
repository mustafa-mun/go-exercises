package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)


func processInput(data [][]string, wg *sync.WaitGroup, correctAnswerCount *int) {
	defer wg.Done() // Close the waitgroup when the repl loop is over
	
	var questionNumber int = 1
	var pointer int = 0
	var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)

	questionNumberString := strconv.Itoa(questionNumber)
	printPrompt("Quesion #"+ questionNumberString + " "+ data[pointer][0])
	// Begin the repl loop

	for reader.Scan() {
		if pointer == len(data) - 1 { // Stop the loop when user is answered all questions
			fmt.Println("You scored", strconv.Itoa(*correctAnswerCount), "out of", strconv.Itoa(len(data)), "questions.")
			return
		}
		text := cleanInput(reader.Text()) // Get the input
		answer := data[pointer][1] 
		// Check if answer is correct
		if text == answer {
			// Answer is correct
			*correctAnswerCount += 1
		}
		// Move to the next question
		pointer += 1
		questionNumber += 1 
		printPrompt("Quesion #"+ strconv.Itoa(questionNumber) + " "+ data[pointer][0])
	}
	
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}