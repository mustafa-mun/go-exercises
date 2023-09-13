package main

import (
	"github.com/mustafa-mun/go-exercises/sitemap-builder/pkg/sitemap"
)

func main()  {
	_, err := sitemap.CreateSitemap("https://www.calhoun.io/")
	if err != nil{
		panic(err)
	}	
}