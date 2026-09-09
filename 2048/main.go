package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Parse html templates
	templ, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("Template parse fail: %v", err)
	}

	// Initialize the router
	mux := http.NewServeMux()
	handleStatic := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	mux.Handle("GET /static/", handleStatic)
	mux.HandleFunc("GET /api/new", handleNew)
	mux.HandleFunc("POST /api/slide", handleSlide)
	mux.HandleFunc("POST /api/undo", handleUndo)
	mux.HandleFunc("GET /", handleHome(templ))

	// Start the http server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
