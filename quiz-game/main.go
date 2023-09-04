package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {	
	csvPtr := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse() // Parse the flags

	data, err := ReadCsvFile(*csvPtr) // Read the csv file
	if err != nil {
		panic(err)
	}
	// Create a waitgroup
	// This is for waiting till all the goroutines are done
	var wg sync.WaitGroup 
	var correctAnswerCount = 0

	wg.Add(1) // Start the waitgroup

	fmt.Println("You have", strconv.Itoa(*limitPtr), "seconds for the exam")
	fmt.Println("Press Enter to start the exam...")
	fmt.Scanln() // Wait for the Enter key

	go func() {	// Create a goroutine to start the repl loop
		processInput(data, &wg, &correctAnswerCount)  
	}()

	timer := time.NewTimer(time.Duration(*limitPtr) * time.Second) // Create timer with limit flag
	go func() { // Create a goroutine to fire the timer
		<-timer.C
		// Time is out
		fmt.Println("You scored", strconv.Itoa(correctAnswerCount), "out of", strconv.Itoa(len(data)), "questions.")
		wg.Done() // Close the waitgroup
	}()
	// Wait until all goroutines are over
	wg.Wait()
}

