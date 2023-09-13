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
func traverseLinks(link string) ([]Sitemap, error) {
	return nil, nil
}

// CreateSitemap will fetch the base domain and
// create a sitemap with all links under the same
// domain and return encoded XML with the sitemaps
func CreateSitemap(baseURL string) ([]byte, error) {
	htmlString, err := getResponseHTML(baseURL)
	if err != nil{
		return nil, err
	}
	linkArray, err := linkparser.Parse(htmlString)
	if err != nil {
		return nil, err
	}
	for _, link := range linkArray{
		siteMap, err := traverseLinks(link.Href)
		if err != nil {
			return nil, err
		}
		fmt.Println(siteMap)
	}
	return nil, nil
}
