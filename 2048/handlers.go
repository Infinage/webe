package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	// pop in effect for new spawn
	var animations [16]string
	for idx := range 16 {
		if b[idx] != 0 {
			animations[idx] = "animate-pop-in-once"
		}
	}

	sse.PatchElements(renderBoard(b, animations))
	signals := map[string]any{"scores": []int{0}, "boards": []Board{b}, "_gameOver": false}
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
	var slideAnimation string
	switch req.Key {
	case "k", "ArrowUp":
		dir = DirectionUp
		slideAnimation = "animate-slide-from-y-"
	case "j", "ArrowDown":
		dir = DirectionDown
		slideAnimation = "-animate-slide-from-y-"
	case "h", "ArrowLeft":
		dir = DirectionLeft
		slideAnimation = "animate-slide-from-x-"
	case "l", "ArrowRight":
		dir = DirectionRight
		slideAnimation = "-animate-slide-from-x-"
	}

	prevB := req.Boards[len(req.Boards)-1]
	currB, deltas, score, ok := prevB.Slide(dir)
	currB = currB.Spawn()

	// Create animation effects
	var animations [16]string
	for idx := range 16 {
		switch {
		case deltas[idx] != 0:
			animations[idx] = slideAnimation + strconv.Itoa(int(deltas[idx]))
		case prevB[idx] != currB[idx] && currB[idx] != 0:
			animations[idx] = "animate-pop-in-once"
		}
	}

	sse := datastar.NewSSE(w, r)
	if !ok {
		sse.DispatchCustomEvent("shake", nil)
		return
	}

	sse.PatchElements(renderBoard(currB, animations))

	// Retain just the last 10 entries
	boards := append(req.Boards, currB)
	scores := append(req.Scores, req.Scores[N-1]+score)
	if N > 10 {
		boards = boards[N-9:]
		scores = scores[N-9:]
	}

	hscore := max(req.HScore, scores[len(scores)-1])
	gameOver := currB.Status() == StatusLoss
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
	sse.PatchElements(renderBoard(req.Boards[N-2], [16]string{}))
}
