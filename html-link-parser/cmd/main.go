package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mustafa-mun/go-exercises/html-link-parser/pkg/linkparser"
)

func main() {
	htmlPtr := flag.String("HTML", "html/ex2.html", "HTML file path to parse the links")
	flag.Parse() // Parse the flags

	htmlString, err := readHtmlFromFile(*htmlPtr)
	if err != nil {
		panic(err)
	}

	linkArray, err := linkparser.Parse(htmlString)
	if err != nil {
		panic(err)
	}
	for _, value := range linkArray { // Print the links
		fmt.Println("Href: " + value.Href + " |" + " Text: " + value.Text)
	}
}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}
