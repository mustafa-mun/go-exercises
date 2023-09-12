package sitemap

import (
	"net/http"

	"github.com/mustafa-mun/go-exercises/html-link-parser/pkg/linkparser"
)

// XML Sitemap structure (for decoding XML)
type Sitemap struct{}

// fetchURL will fetch an URL and return a
// pointer to http Response
func fetchURL(url string) *http.Response {
	return nil
}

// getResponseHTML will take a pointer to HTML
// Response and return the HTML as a string
func getResponseHTML(resp *http.Response) string {
	return ""
}

// getLinks will take a HTML string as an input
// and return an output of all links as an array
// of Link's
func getLinks(html string) []linkparser.Link {
	return nil
}

// traverseLinks will take a link as an input
// and traverse the whole link path recursively
// or iteratively (DFS or BFS) then return an
// array of Sitemaps as an output
func traverseLinks(link string) []Sitemap {
	return nil
}

// CreateSitemap will fetch the base domain and
// create a sitemap with all links under the same
// domain and return encoded XML with the sitemaps
func CreateSitemap(baseURL string) []byte {
	return nil
}
