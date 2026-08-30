package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type RipplesApp struct {
	templ   *template.Template // Holds the front end views
	uEvents sync.Map           // Each user gets their own channel to send UI events
	FPS     int                // Frames per second (determines the SSE send interval)
}

// NewRipplesApp initalizes the views and inits a sync map
func NewRipplesApp() (*RipplesApp, error) {
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		return nil, fmt.Errorf("init failed: %w", err)
	}
	return &RipplesApp{templ: templ, FPS: 60}, nil
}

// Routes returns an http.ServeMux with preconfigured routes.
//
// Routes:
//   - GET / — Returns the homepage.
//   - POST /loop — Establishes an SSE connection, spins up a new goroutine,
//     and sends painted frames.
//   - POST /resize — Every resize on the UI triggers a board.Resize.
//   - POST /click — Every click on the UI registers the event here.
func (app *RipplesApp) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.handleHome)
	mux.HandleFunc("GET /loop", app.handleLoop)
	mux.HandleFunc("POST /resize", app.handleResize)
	mux.HandleFunc("POST /click", app.handleClick)
	mux.Handle("GET /static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	return mux
}

// handleHome tags the incoming client and renders the home page
func (app *RipplesApp) handleHome(w http.ResponseWriter, r *http.Request) {
	id := generateId()
	ch := make(chan Event, 1)
	app.uEvents.Store(id, ch)

	data := map[string]string{"id": id}
	if err := app.templ.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
