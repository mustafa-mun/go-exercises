package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello world!")
}

func normalize(phoneNumber string) string {
	var output string

	for _, r := range phoneNumber {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			output += string(r)
		}
	}

	return output
}
