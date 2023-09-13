package sitemap

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mustafa-mun/go-exercises/html-link-parser/pkg/linkparser"
)

// XML Sitemap structure (for decoding XML)
type Sitemap struct{}


// getResponseHTML will will fetch an URL and 
// return the response HTML as a string
func getResponseHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// Convert the response body to a string (HTML)
	return string(htmlBytes), nil
}

// traverseLinks will take a link as an input
// and traverse the whole link path recursively
// then return an array of Sitemaps as an output
func traverseLinks(link, baseURL string, seen map[string]bool) error{
	_, ok := seen[link]
	if ok { // If link is seen, return
		return nil
	}

	linkHost,err := GetURLHost(link)
	if err != nil{
		return err
	}
	baseHost, err := GetURLHost(baseURL)
	if err != nil{
		return err
	}

	if linkHost != baseHost { // If link host is not same with baseURL, return
		return nil
	}
	
	htmlString, err := getResponseHTML(link)
	if err != nil{
		return err
	}
	linkArray, err := linkparser.Parse(htmlString)
	if err != nil {
		return err
	}
	fmt.Println(linkArray)
	for _, link := range linkArray{
		seen[link.Href] = true // Add link to seen list
		traverseLinks(link.Href, baseURL, seen)
	}
	return nil
}

// CreateSitemap will fetch the base domain and
// create a sitemap with all links under the same
// domain and return encoded XML with the sitemaps
func CreateSitemap(baseURL string) ([]byte, error) {
	seen := make(map[string]bool)
	err := traverseLinks(baseURL, baseURL, seen)
	if err != nil{
		return nil, err
	}
	return nil, nil
}

func GetURLHost(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return response.Request.URL.Host, nil
}