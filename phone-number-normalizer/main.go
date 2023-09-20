package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello world!")
}

func normalize(phoneNumber string) string {
	var output bytes.Buffer

	for _, r := range phoneNumber {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			output.WriteRune(r)
		}
	}

	return output.String()
}
