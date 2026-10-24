package internal

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

// App holds the application state to run the tailwind battle playground.
type App struct {
	assets embed.FS
	templ  *template.Template
}

// NewApp initializes an App struct and attempts to load html templates.
func NewApp(assets embed.FS) (*App, error) {
	templ, err := template.ParseFS(assets, "assets/html/index.html")
	if err != nil {
		return nil, fmt.Errorf("template parse: %w", err)
	}

	return &App{assets: assets, templ: templ}, nil
}

// router handles the various endpoints
func (app *App) router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.FileServerFS(app.assets))
	mux.HandleFunc("GET /", app.handleAPIHome)
	mux.HandleFunc("POST /api/target", app.handleAPIEncodeTargetImage)
	mux.HandleFunc("POST /api/diff", app.handleAPIDiffImages)
	return mux
}

// Serve spins up the playground on input addr and returns err on failure
func (app *App) Serve(addr string) error {
	err := http.ListenAndServe(addr, app.router())
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
