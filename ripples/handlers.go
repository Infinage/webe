package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type userMeta struct {
	id            string
	x, y          uint
	height, width uint
	ch            chan Event
}

// extractUserMeta parses the incoming 'post' request and extracts
// fields as 'x-www-form-urlencoded'. It extracts following:
// 'id', 'x', 'y', 'height', 'width'. Throws an error on missing
// id or when the floats are less than 0.
func extractUserMeta(r *http.Request, m *sync.Map) (userMeta, error) {
	if err := r.ParseForm(); err != nil {
		return userMeta{}, err
	}

	id := r.Form.Get("id")
	eventsCh, ok := m.Load(id)
	if !ok {
		return userMeta{}, fmt.Errorf("id not found: %s", id)
	}

	extract := func(name string) (uint, error) {
		valStr := r.Form.Get(name)
		if valStr == "" {
			return 0, nil
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil || val < 0 {
			return 0, fmt.Errorf("invalid %q value: %s", name, r.Form.Get(name))
		} else {
			return uint(math.Round(val)), nil
		}
	}

	values := make(map[string]uint, 4)
	for _, name := range []string{"x", "y", "height", "width"} {
		v, err := extract(name)
		if err != nil {
			return userMeta{}, err
		}
		values[name] = v
	}

	return userMeta{
		id:     id,
		x:      values["x"],
		y:      values["y"],
		height: values["height"],
		width:  values["width"],
		ch:     eventsCh.(chan Event),
	}, nil
}

// handleResize is an auto generated event that resizes the water board
func (app *RipplesApp) handleResize(w http.ResponseWriter, r *http.Request) {
	meta, err := extractUserMeta(r, &app.uEvents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Blocking send
	meta.ch <- &EventResize{height: meta.height, width: meta.width}
}

// handleClick registers a click event and sends the event into the appropriate
// channel for the goroutine to handle
func (app *RipplesApp) handleClick(w http.ResponseWriter, r *http.Request) {
	meta, err := extractUserMeta(r, &app.uEvents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Non blocking send
	select {
	case meta.ch <- &EventClick{x: meta.x, y: meta.y}:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotModified)
	}
}

// handleLoop inits a board, establises an SSE connection with the client
// and sends the board over the connection for every frame
func (app *RipplesApp) handleLoop(w http.ResponseWriter, r *http.Request) {
	meta, err := extractUserMeta(r, &app.uEvents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Cleanup user on exit
	fmt.Printf("Starting loop for %q\n", meta.id)
	defer func() {
		app.uEvents.Delete(meta.id)
		fmt.Printf("Ending loop for %q\n", meta.id)
	}()

	// Set header for Server sent events (SSE)
	w.Header().Set("Content-Type", "text/event-stream")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Initialze the water board
	board, err := NewBoard(meta.height, meta.width, 0.99)
	if err != nil {
		http.Error(w, fmt.Sprintf("Board init: %v", err), http.StatusInternalServerError)
		return
	}

	ticker := time.NewTicker(time.Second / time.Duration(app.FPS))
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return

		case event, ok := <-meta.ch:
			if !ok {
				return
			}
			event.Do(&board)

		case <-ticker.C:
			res := make([]byte, board.rows*board.cols*4)
			for x, row := range board.Step() {
				for y, cell := range row {
					idx := (x*int(board.cols) + y) * 4
					r, g, b := waveColor(cell)
					copy(res[idx:], []byte{r, g, b, 255})
				}
			}
			encoded := base64.StdEncoding.EncodeToString(res)
			w.Write([]byte(fmt.Sprintf("data: %s\n\n", encoded)))
			flusher.Flush()
		}
	}
}

func waveColor(v float32) (byte, byte, byte) {
	const gain = 10.0
	v *= gain

	// Clamp and boost subtle waves.
	v = float32(math.Max(-1, math.Min(1, float64(v/255))))
	mag := float32(math.Pow(math.Abs(float64(v)), 0.5))

	// Dark blue water base.
	const (
		baseR = 8.0
		baseG = 18.0
		baseB = 35.0
	)

	if v >= 0 {
		// Positive waves: blue -> yellow/white.
		return byte(baseR + 247*mag),
			byte(baseG + 237*mag),
			byte(baseB + 25*mag)
	}

	// Negative waves: blue -> violet.
	return byte(baseR + 90*mag),
		byte(baseG + 70*mag),
		byte(baseB + 220*mag)
}
