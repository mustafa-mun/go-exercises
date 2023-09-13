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
	type Url struct{
		XMLName xml.Name `xml:"url"`
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
	func traverseLinks(link, baseURL string, seen map[string]bool, sitemapArray *[]string, wg *sync.WaitGroup) error {
		defer wg.Done()
		mu.Lock()

		if seen[link] {
			mu.Unlock()
			return nil
		}
		wg.Add(1)
		seen[link] = true
		mu.Unlock()
		fmt.Println(link)
		newUrl := Url{ // Create new XML URL
			Loc: link,
		}
		*sitemapArray = append(*sitemapArray, newUrl.Loc)

		htmlString, err := getResponseHTML(link)
		if err != nil{
			return err
		}
		linkArray, err := linkparser.Parse(htmlString)
		if err != nil {
			return err
		}

		for _, link := range linkArray{
			var target string 
		
			if !IsUrl(link.Href){
				if link.Href == "/" {
					target = baseURL
				} else {
					target = baseURL + link.Href 
				}
				
			} else {
				target = link.Href
			}

			linkHost, err := GetURLHost(target)
			if err != nil{	
				return err
			}

			baseHost, err := GetURLHost(baseURL)
			if err != nil{
				return err
			}
			if linkHost == baseHost && !seen[target] { // If link host is not same with baseURL, return
				go traverseLinks(target, baseURL, seen, sitemapArray, wg)
			}
			
		}
		return nil
	}

	// CreateSitemap will fetch the base domain and
	// create a sitemap with all links under the same
	// domain and return encoded XML with the sitemaps
	func CreateSitemap(baseURL string) ([]byte, error) {
		seen := make(map[string]bool)
		var wg sync.WaitGroup
		var sitemapArray []string
		err := traverseLinks(baseURL, baseURL, seen, &sitemapArray, &wg)
		if err != nil{
			return nil, err
		}
		out, err := xml.MarshalIndent(sitemapArray, "", " ")
		
		if err != nil {
			return nil, err
		}
		wg.Wait()
		return out, nil
	}
