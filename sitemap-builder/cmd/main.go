package main

import (
	"flag"
	"os"

	"github.com/mustafa-mun/go-exercises/sitemap-builder/pkg/sitemap"
)

func main()  {
	urlPtr := flag.String("url", "https://gobyexample.com/", "Base URL for sitemap builder")
	flag.Parse() 

	stmp, err := sitemap.CreateSitemap(*urlPtr)
	if err != nil{
		panic(err)
	}	
	err = os.WriteFile("sitemap.xml", stmp, 0644)
	if err != nil{
		panic(err)
	}	
	
}