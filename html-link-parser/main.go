package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	htmlString, err := readHtmlFromFile("ex1.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}
	var linkArray []Link
	traverseHTMLLinks(doc, &linkArray)
	for _, value := range linkArray {
		fmt.Println("Href: " + value.Href + " Text: " + value.Text)
	}
}

func traverseHTMLLinks(n *html.Node, linkArray *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				var linkText string
				traverseLinkChildren(n.FirstChild, &linkText)
				newLink := Link{
					Href: a.Val,
					Text: linkText,
				}
				*linkArray = append(*linkArray, newLink)
				break
			}		
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTMLLinks(c, linkArray)
	}
}

func traverseLinkChildren(n *html.Node, linkText *string) string { // Take the text from links childs
	if n != nil {
		*linkText = *linkText + " "+ strings.TrimSpace(n.Data)
		if n.FirstChild != nil { // If node has children, traverse the childs
			traverseLinkChildren(n.FirstChild, linkText)
		}
		traverseLinkChildren(n.NextSibling, linkText) // Traverse the nextsibling
	}
	return *linkText
}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
			return "", err
	}

	return string(bs), nil
}
