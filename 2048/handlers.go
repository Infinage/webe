package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/starfederation/datastar-go/datastar"
)

func handleHome(templ *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if err := templ.ExecuteTemplate(w, "Home", nil); err != nil {
			log.Printf("Template exec: %v", err)
		}
	}
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	b := NewBoard()
	sse.PatchElements(renderBoard(b))

	signals := map[string]any{"scores": []int{0}, "boards": []Board{b}}
	sse.MarshalAndPatchSignals(signals)
}

func handleSlide(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key    string  `json:"Key"`
		Boards []Board `json:"boards"`
		Scores []int   `json:"scores"`
		HScore int     `json:"hscore"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	N := len(req.Boards)
	if N == 0 || N != len(req.Scores) {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	var dir Direction
	switch req.Key {
	case "k", "ArrowUp":
		dir = DirectionUp
	case "j", "ArrowDown":
		dir = DirectionDown
	case "h", "ArrowLeft":
		dir = DirectionLeft
	case "l", "ArrowRight":
		dir = DirectionRight
	}

	currB := req.Boards[len(req.Boards)-1]
	b, score, ok := currB.Slide(dir)
	b = b.Spawn()

	sse := datastar.NewSSE(w, r)
	if !ok {
		sse.DispatchCustomEvent("shake", nil)
		return
	}

	sse.PatchElements(renderBoard(b))

	// Retain just the last 10 entries
	boards := append(req.Boards, b)
	scores := append(req.Scores, req.Scores[N-1]+score)
	if N > 10 {
		boards = boards[N-9:]
		scores = scores[N-9:]
	}

	hscore := max(req.HScore, scores[len(scores)-1])
	gameOver := b.Status() == StatusLose
	signals := map[string]any{
		"scores": scores, "boards": boards, 
		"_hscore": hscore, "_gameOver": gameOver,
	}

	sse.MarshalAndPatchSignals(signals)
}

func handleUndo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Boards []Board `json:"boards"`
		Scores []int   `json:"scores"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	N := len(req.Boards)
	if N <= 1 || N != len(req.Scores) { // Expect atleast 2 elements
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	req.Boards = req.Boards[:N-1]
	req.Scores = req.Scores[:N-1]

	signals := map[string]any{
		"boards":    req.Boards,
		"scores":    req.Scores,
		"_gameOver": false,
	}

	sse := datastar.NewSSE(w, r)
	sse.MarshalAndPatchSignals(signals)
	sse.PatchElements(renderBoard(req.Boards[N-2]))
}
