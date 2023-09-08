package main

import (
	"fmt"
)

type Story map[string]Arc

type Arc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

func main()  {
	jsonFile, err := openJSON("gopher.json")
	if err != nil {
		panic(err)
	}
	story, err := parseJSONToStory(jsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(story)
}
