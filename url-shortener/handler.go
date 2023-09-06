package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// Map type for path to urls
type UrlPaths map[string]string 

// YAML file structure
type YAMLUrlMapping struct { 
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// JSON file structure
type JSONUrlMapping struct { 
	Path string `json:"path"`
	URL  string `json:"url"`
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
	pathMap := buildYAMLMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// parseYAML function will parse the string into YAMLUrlMapping
// struct and return an array of YAMLUrlMappings
func parseYAML(yamalData []byte) ([]YAMLUrlMapping, error){
	var mappings []YAMLUrlMapping
	err := yaml.Unmarshal(yamalData, &mappings)
	if err != nil {
		return nil, err
	}
	return mappings, nil
}

// buildYAMLdMap function will take an array of YAMLUrlMapping's and
// convert them into UrlPaths map
func buildYAMLMap(urlArray []YAMLUrlMapping) UrlPaths {
	output := UrlPaths{}
	for _, value := range urlArray {
		output[value.Path] = value.URL
	}
	return output
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     {"path":"/some-path", "url":"https://www.some-url.com/demo"}`
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn []byte, fallback http.Handler)  (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(jsn)
	if err != nil {
		return nil,err
	}
	pathMap := buildJSONMap(parsedJSON)
	return MapHandler(pathMap, fallback), nil
}

// parseJSON function will parse the string into JSONUrlMapping
// struct and return an array of JSONUrlMappings
func parseJSON(jsonData []byte) ([]JSONUrlMapping, error) {
	var mappings []JSONUrlMapping
	err := json.Unmarshal(jsonData, &mappings)
	if err != nil {
		return mappings, err
	}
	fmt.Println(mappings)
	return mappings, nil
}

// buildJSONMap function will take an array of JSONUrlMapping's and
// convert them into UrlPaths map
func buildJSONMap(urlArray []JSONUrlMapping) UrlPaths {
	output := UrlPaths{}
	for _, value := range urlArray {
		output[value.Path] = value.URL
	}
	return output
}
