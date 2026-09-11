package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed static
var FS embed.FS

func main() {
	// Parse html templates
	templ, err := template.ParseFS(FS, "static/index.html")
	if err != nil {
		log.Fatalf("Template parse fail: %v", err)
	}

	// Initialize the router
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(FS))
	mux.HandleFunc("GET /api/new", handleNew)
	mux.HandleFunc("POST /api/slide", handleSlide)
	mux.HandleFunc("POST /api/undo", handleUndo)
	mux.HandleFunc("GET /", handleHome(templ))

	// Start the http server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
