package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	htmlString, err := readHtmlFromFile("ex1.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	
	traverseHTMLLinks(doc)
}

func traverseHTMLLinks(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val) // Href
				break
			}		
	}
	fmt.Println(n.FirstChild.Data) // Text
}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTMLLinks(c)
	}
}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
			return "", err
	}

	return string(bs), nil
}