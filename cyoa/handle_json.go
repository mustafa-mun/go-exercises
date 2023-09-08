package main

import (
	"encoding/json"
	"io"
	"os"
)

func parseJSONToStory(jsn *os.File) (Story, error) {
	var storyMap Story
	byteValue, err := io.ReadAll(jsn)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &storyMap)
	if err != nil {
		return nil, err
	}
	return storyMap, nil
}

func openJSON(jsn string) (*os.File, error) {
	jsonPointer, err := os.Open(jsn)
	if err != nil {
		return nil, err
	}
	return jsonPointer, nil
}
