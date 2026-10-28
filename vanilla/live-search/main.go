package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var languages = []string {
	"golang",
	"javascript",
	"python",
	"java",
	"c++",
	"lua",
	"rust",
	"zig",
	"odin",
	"typescript",
	"dart",
	"flutter",
	"scala",
	"mojo",
	"groovy",
	"swift",
	"kotlin",
	"sql",
	"brainfuck",
	"awk",
}

func handleAPISearch(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var result []string
	for _, language := range languages {
		if strings.Contains(language, query) {
			result = append(result, language)
		}
	}

	w.Header().Set("content-type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(result); err != nil {
		http.Error(w, "JSON Encode: " + err.Error(), http.StatusInternalServerError)
		return
	}
} 

func main() {
	templ, err := template.ParseFiles("assets/index.html")
	if err != nil {
		log.Fatal("Template init: ", err)
	}

	mux := http.NewServeMux()	
	mux.HandleFunc("GET /api/search", handleAPISearch)

	fServer := http.FileServer(http.Dir("assets/"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err = templ.Execute(w, nil); err != nil {
			errMsg := "Template execute: " + err.Error()
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server run: ", err)
	}
}
