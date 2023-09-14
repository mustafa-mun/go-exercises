package sitemap

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/mustafa-mun/go-exercises/html-link-parser/pkg/linkparser"
)

// XML Sitemap structure (for decoding XML)
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func GetURLHost(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return response.Request.URL.Host, nil
}
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

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

var mu sync.Mutex // Mutex to protect the 'seen' map

// traverseLinks will take a link as an input
// and traverse the whole link path recursively
// then return an array of Sitemaps as an output
func traverseLinks(link, baseURL string, depth int, seen map[string]bool, sitemapArray *[]URL, wg *sync.WaitGroup) error {
	defer wg.Done() // Decrase the waitgroup right before the return
	mu.Lock()       // Lock the mutex
	wg.Add(1)       // Add one to waitgroup

	if depth < 0 { // If max depth has reached, return
		return nil
	}

	if seen[link] { // If link is seen, return
		mu.Unlock() // Unlock the mutex
		return nil
	}

	seen[link] = true // Set link as seen
	mu.Unlock()
	fmt.Println(link)

	*sitemapArray = append(*sitemapArray, URL{Loc: link}) // Add link to sitemapArray

	htmlString, err := getResponseHTML(link) // Parse the html from a link
	if err != nil {
		return err
	}
	linkArray, err := linkparser.Parse(htmlString) // Get links array from HTML string
	if err != nil {
		return err
	}
	// For each link in the link array
	for _, link := range linkArray {
		var target string // This will store the target link
		// Check if link is a valid URL
		if !IsUrl(link.Href) {
			// Not a valid URL, ex. (/about #home)
			if link.Href == "/" {
				target = baseURL
			} else {
				// Add baseURL to path
				target = baseURL + link.Href
			}
		} else {
			// Link is a valid URL
			target = link.Href
		}

		linkHost, err := GetURLHost(target) // Get targets host
		if err != nil {
			return err
		}

		baseHost, err := GetURLHost(baseURL) // Get base host
		if err != nil {
			return err
		}
		// If target link's host is same with baseURL and target is not seen and depth is not reached
		if linkHost == baseHost && !seen[target] && depth > 0 {
			// Start a goroutine with the target link and one less depth
			go traverseLinks(target, baseURL, depth-1, seen, sitemapArray, wg)
		}
	}
	return nil
}

// CreateSitemap will fetch the base domain and
// create a sitemap with all links under the same
// domain and return encoded XML with the sitemaps
func CreateSitemap(baseURL string, depth int) ([]byte, error) {
	seen := make(map[string]bool)
	var wg sync.WaitGroup
	var sitemapArray []URL
	fmt.Println("Creating Sitemap for URL: " + baseURL)
	err := traverseLinks(baseURL, baseURL, depth, seen, &sitemapArray, &wg)
	if err != nil {
		return nil, err
	}
	sitemap := Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  sitemapArray,
	}
	out, err := xml.MarshalIndent(sitemap, "", " ")

	if err != nil {
		return nil, err
	}
	wg.Wait()
	return out, nil
}

// TraverseLinksBFS is a BFS version of traverseLinks
func TraverseLinksBFS(baseURL string, seen map[string]bool,sitemapArray []URL,) ([]URL, error){
	queue := []string{baseURL}

	for len(queue) > 0 {
		array, first := PopStart[string](queue)		
		queue = array
		sitemapArray = append(sitemapArray, URL{Loc: first}) // Add link to sitemapArray

		htmlString, err := getResponseHTML(first) // Parse the html from a link
		if err != nil {
			return nil, err
		}
		linkArray, err := linkparser.Parse(htmlString) // Get links array from HTML string
		if err != nil {
			return nil, err
		}
		for _, link := range linkArray {
			var target string // This will store the target link
		// Check if link is a valid URL
			if !IsUrl(link.Href) {
				// Not a valid URL, ex. (/about #home)
				if link.Href == "/" {
					target = baseURL
				} else {
					// Add baseURL to path
					target = baseURL + link.Href
				}
			} else {
				// Link is a valid URL
				target = link.Href
			}

			linkHost, err := GetURLHost(target) // Get targets host
			if err != nil {
				return nil, err
			}

			baseHost, err := GetURLHost(baseURL) // Get base host
			if err != nil {
				return nil, err
			}
			// If target link's host is same with baseURL and target is not seen and depth is not reached
			if linkHost == baseHost && !seen[target] {
				seen[target] = true
				queue = append(queue, target)
			}
		}
	}
	return sitemapArray, nil
}

func PopStart[T any](array []T) ([]T, T) {
	var output []T = make([]T, len(array) - 1)
	first := array[0]
	for i := 0; i < cap(output); i++ {
		output[i] = array[i + 1]
	}
	return output, first
}
