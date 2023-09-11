package main

import (
	"flag"
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

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
			return "", err
	}

	return string(bs), nil
}


func main() {
	htmlPtr := flag.String("HTML", "html/ex5.html", "HTML file path to parse the links")
	flag.Parse() // Parse the flags

	htmlString, err := readHtmlFromFile(*htmlPtr)
	
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(strings.NewReader(htmlString)) // Parse the html file
	if err != nil {
		log.Fatal(err)
	}
	var linkArray []Link // Create linkArray to store the links
	traverseHTMLLinks(doc, &linkArray) // Get links from HTML
	for _, value := range linkArray { // Print the links
		fmt.Println("Href: " + value.Href + " |" + " Text: " + value.Text)
	}
}
// traverseHTMLLinks function will traverse the whole HTML
// to find the a tags, then it will take their href and text
// and crete Links, then it will return an array of Links
func traverseHTMLLinks(n *html.Node, linkArray *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" { // If node is an a tag
		for _, a := range n.Attr { // loop through the attributes of a tag to find the href
			if a.Key == "href" { // We found the href
				var linkText string
				traverseLinkChilds(n.FirstChild, &linkText) // Take the text from a tag
				newLink := Link{ // Create new Link
					Href: a.Val,
					Text: strings.TrimSpace(linkText),
				}
				*linkArray = append(*linkArray, newLink) // Append link to the linkArray
				break
			}		
		}
	}
	// Traverse the whole HTML
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		traverseHTMLLinks(c, linkArray)
	}
}

// traverseLinkChildren function will traverse all of the 
// children of an a tag and extract the text 
func traverseLinkChilds(n *html.Node, linkText *string) string { 
	if n != nil {
		if n.Type == html.TextNode{ // Take the text only if node is a TextNode
			*linkText = *linkText + " "+ strings.TrimSpace(n.Data) // Add the trimmed text to the output(linkText)
		}
		if n.FirstChild != nil { // If node has children, traverse the childs
			traverseLinkChilds(n.FirstChild, linkText)
		}
		traverseLinkChilds(n.NextSibling, linkText) // Traverse the nextsibling
	}
	return *linkText
}
