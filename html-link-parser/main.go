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

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)

	if err != nil {
			return "", err
	}

	return string(bs), nil
}


func main() {
	htmlString, err := readHtmlFromFile("ex4.html")
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
		fmt.Println("Href: " + value.Href + " Text:" + value.Text)
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
					Text: linkText,
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
		if n.Type != html.ElementNode && n.Type != html.CommentNode{ // Take the text only if node is not an element node or comment
			*linkText = *linkText + " "+ strings.TrimSpace(n.Data) // Add text to the link text
		}
		if n.FirstChild != nil { // If node has children, traverse the childs
			traverseLinkChilds(n.FirstChild, linkText)
		}
		traverseLinkChilds(n.NextSibling, linkText) // Traverse the nextsibling
	}
	return *linkText
}