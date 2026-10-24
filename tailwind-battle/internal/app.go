package internal

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// App holds the application state to run the tailwind battle playground.
type App struct {
	assets embed.FS
	templ  *template.Template
	cache  *targetCache
}

// NewApp initializes an App struct and attempts to load html templates.
func NewApp(assets embed.FS) (*App, error) {
	templ, err := template.ParseFS(assets, "assets/html/index.html")
	if err != nil {
		return nil, fmt.Errorf("template parse: %w", err)
	}

	cache := newTargetCache(time.Hour * 6)
	return &App{assets: assets, templ: templ, cache: cache}, nil
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if app.cache.evictOlderThan > 0 {
		ticker := time.NewTicker(app.cache.evictOlderThan)
		defer ticker.Stop()
		go func() {
			for {
				select {
				case <-ticker.C:
					app.cache.Clean()
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	err := http.ListenAndServe(addr, app.router())
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
