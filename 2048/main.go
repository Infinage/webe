package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/starfederation/datastar-go/datastar"
)

func renderBoard(b Board) string {
	var buf bytes.Buffer
	buf.WriteString(`<div id="board" class="grid grid-cols-4 grid-rows-4 gap-3 p-3 
		bg-[#9C8B7C] h-1/2 aspect-square rounded-xl"
	>`)

	for _, cell := range b {
		buf.WriteString(`<div class="bg-[#BDAC97] rounded-xl">`)
		buf.WriteString(strconv.FormatUint(uint64(cell), 10))
		buf.WriteString(`</div>`)
	}

	buf.WriteString(`<div class="bg-[#BDAC97] rounded-xl"></div></div>`)
	return buf.String()
}

func handleHome(templ *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templ.ExecuteTemplate(w, "Home", nil); err != nil {
			log.Printf("Template exec: %v", err)
		}
	}
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	sse.MarshalAndPatchSignals(map[string]any{"score": 0})
	b := NewBoard()
	sse.PatchElements(renderBoard(b))
}

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
	mux.HandleFunc("GET /", handleHome(templ))
	mux.HandleFunc("GET /api/new", handleNew)

	// Start the http server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	// 0, 2, 4, 8, 16, 32, 64, 128, 512, 1024, 2048
}
