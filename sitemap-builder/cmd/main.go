package main

import (
	"os"

	"github.com/mustafa-mun/go-exercises/sitemap-builder/pkg/sitemap"
)

func main()  {
	stmp, err := sitemap.CreateSitemap("https://recipe-share-cmv8.onrender.com")
	if err != nil{
		panic(err)
	}	
	err = os.WriteFile("sitemap.xml", stmp, 0644)
	if err != nil{
		panic(err)
	}	
	
}