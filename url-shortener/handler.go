package main

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type urlPaths map[string]string // Map type for path to urls

type URLMapping struct { // Yaml file structure
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]
		if ok {
			// Path found in map
			http.Redirect(w,r, url, http.StatusSeeOther)
		}
		// Path not found in map
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil,err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// parseYAML function will parse the string into URLMapping
// struct and return an array of URLMappings
func parseYAML(yamalData []byte) ([]URLMapping, error){
	var mappings []URLMapping
	err := yaml.Unmarshal(yamalData, &mappings)
	if err != nil {
		return nil, err
	}
	return mappings, nil
}

// buildMap function will take an array of URLMappings and
// convert them into urlPaths map
func buildMap(urlArray []URLMapping) urlPaths {
	output := urlPaths{}
	for _, value := range urlArray {
		output[value.Path] = value.URL
	}
	return output
}