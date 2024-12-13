package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Import embed to resolve the build issue.
//
//go:embed index.html
var indexHTML string

type PageData struct {
	Version     string
	Environment map[string]string
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	envVars := os.Environ()
	envMap := make(map[string]string)
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	data := PageData{
		Version:     os.Getenv("VERSION"),
		Environment: envMap,
	}

	t, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", serveIndex)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
