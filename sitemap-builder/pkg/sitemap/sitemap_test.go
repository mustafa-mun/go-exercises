package sitemap

import (
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// fetchURL tests //
func TestValidURL(t *testing.T) {
	resp, err := fetchURL("https://www.google.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Want %v, got %v", http.StatusOK, resp.StatusCode)
	}
}

func TestUnvalidURL(t *testing.T) {
	_, err := fetchURL("nonvalid")
	if err == nil {
		t.Errorf("Want %v, got %v", nil, err)
	} 
}

// getResponseHTML tests //
func TestHTML(t *testing.T)  {
	resp, err := fetchURL("https://www.google.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	htmlStr, err := getResponseHTML(resp)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	_, err = html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		t.Errorf("Want %v, got %v", nil, err)
	}
}

// getLinks tests //

func TestgetLinks(t *testing.T) {}