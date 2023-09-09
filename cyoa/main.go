package main

import (
	"net/http"
	"text/template"
)

type Story map[string]Arc

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic("Error: " + err.Error())
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := openJSON("gopher.json")
	if err != nil {
		panic("Error: " + err.Error())
	}
	story, err := parseJSONToStory(jsonFile)
	if err != nil {
		panic("Error: " + err.Error())
	}
	// Parse the html template
	tmpl := template.Must(template.ParseFiles("story.html")) 
	// Take url path in order to decide the chapter 
	var chapter string = r.URL.Path[1:]
	if chapter == "" { // Render the intro if path is "/"
		chapter = "intro"
	}
	// Check if url path is in the story map
	_, ok := story[chapter]
	if ok {
		// Path(chapter) found in map
		tmpl.Execute(w, story[chapter])
	} else {
		// Path(chapter) not found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Chapter not found."))
	}
}
