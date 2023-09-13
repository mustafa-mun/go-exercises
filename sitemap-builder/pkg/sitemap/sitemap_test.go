package sitemap

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// getResponseHTML tests //
func TestValidURL(t *testing.T) {
	_, err := getResponseHTML("https://www.google.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestUnvalidURL(t *testing.T) {
	_, err := getResponseHTML("nonvalid")
	if err == nil {
		t.Error("couldn't find an error")
	} 
}
func TestHTML(t *testing.T)  {
	htmlStr, err := getResponseHTML("https://www.google.com/")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	_, err = html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		t.Errorf("Want %v, got %v", nil, err)
	}
}


