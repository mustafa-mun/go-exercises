package main

import (
	"flag"
	"fmt"

	"github.com/mustafa-mun/go-exercises/html-link-parser/internal/linkparser"
)

func main() {
	htmlPtr := flag.String("HTML", "html/ex2.html", "HTML file path to parse the links")
	flag.Parse() // Parse the flags

	linkArray, err := linkparser.Parse(htmlPtr)
	if err != nil {
		panic(err)
	}
	for _, value := range linkArray { // Print the links
		fmt.Println("Href: " + value.Href + " |" + " Text: " + value.Text)
	}
}
