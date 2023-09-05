package main

import (
	"fmt"
	"strings"
)

// printPrompt displays the repl prompt at the start of each loop
func printPrompt(cliName string) {
	fmt.Print(cliName, ": ")
}
// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
 
