package main

import (
	"flag"
	"os"

	"github.com/mustafa-mun/go-exercises/sitemap-builder/pkg/sitemap"
)

func main() {
	urlPtr := flag.String("url", "https://gobyexample.com/", "Base URL for sitemap builder")
	depthPtr := flag.Int("depth", 9, "defines the maximum number of links to follow")
	flag.Parse()

	stmp, err := sitemap.CreateSitemap(*urlPtr, *depthPtr)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("sitemap.xml", stmp, 0644) // Write the sitemap XML to a file
	if err != nil {
		panic(err)
	}

}
