package main

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"

	"github.com/starfederation/datastar-go/datastar"
)

//go:embed assets/*
var assets embed.FS

// density in increasing order
const densityStr = " _.,-=+:;cba!?0123456789$W#@"

func handleHome(templ *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.ExecuteTemplate(w, "index.html", nil)
	}
}

func handleAPIAsciify(w http.ResponseWriter, r *http.Request) {
	var output struct {
		Size    string `json:"_size"`
		Content string `json:"_output"`
	}
	output.Size = "16px"

	defer func() {
		datastar.NewSSE(w, r).MarshalAndPatchSignals(output)
	}()

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		output.Content = fmt.Sprintf("form parse: %v", err)
		return
	}

	downscaleStr := r.FormValue("downscale")
	downscale, err := strconv.ParseUint(downscaleStr, 10, 16)
	if err != nil || downscale == 0 || downscale > 100 {
		output.Content = fmt.Sprintf("Not a valid downscale factor (%q) [1 - 100]: %v",
			downscaleStr, err)
		return
	}

	var imgReader io.Reader
	file, _, err := r.FormFile("image")
	switch {
	// File has been uploaded
	case err == nil:
		defer file.Close()
		imgReader = file

	// File is missing, check for URL
	case errors.Is(err, http.ErrMissingFile):
		u, err := url.Parse(r.FormValue("url"))
		if err != nil {
			output.Content = fmt.Sprintf("Invalid url: %v", err)
			return
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			output.Content = fmt.Sprintf("Only http/https protocols allowed, got %q", u.Scheme)
			return
		}
		if host := u.Hostname(); host == "localhost" || host == "127.0.0.1" {
			output.Content = "Access to internall addresses are restricted."
			return
		}

		client := http.Client{Timeout: time.Second * 3}
		resp, err := client.Get(u.String())
		if err != nil {
			output.Content = fmt.Sprintf("Fetch failed: %v", err)
			return
		}
		defer resp.Body.Close()
		imgReader = resp.Body

	// Unexpected error
	default:
		output.Content = fmt.Sprintf("form parse (unexpected): %v", err)
		return
	}

	img, _, err := image.Decode(io.LimitReader(imgReader, 5e6))
	if err != nil {
		output.Content = fmt.Sprintf("Image read fail: %v", err)
		return
	}

	output.Content = asciify(img, int(downscale))
	output.Size = fmt.Sprintf("calc(100vw/%d)", img.Bounds().Max.X/int(downscale))
}

func asciify(img image.Image, downscale int) string {
	density := []rune(densityStr)
	bounds := img.Bounds()
	var buf strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y += downscale * 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x += downscale {
			grey := avgPool(img, x, y, downscale)
			idx := int(grey * float32(len(density)-1))
			buf.WriteRune(density[idx])
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func avgPool(img image.Image, x, y, window int) float32 {
	var R, G, B, count float32
	bounds := img.Bounds()
	for i := x; i < min(bounds.Max.X, x+window); i++ {
		for j := y; j < min(bounds.Max.Y, y+window); j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			R += float32(r)
			G += float32(g)
			B += float32(b)
			count++
		}
	}

	grey := (R + G + B) / (0xFFFF * count * 3)
	return grey
}

func main() {
	templ, err := template.ParseFS(assets, "assets/index.html")
	if err != nil {
		log.Fatalf("Failed to load assets: %v", err)
	}

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleHome(templ))
	mux.HandleFunc("POST /api/asciify", handleAPIAsciify)
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	// Start the server
	addr := ":8080"
	log.Printf("Starting ascii-art sever on %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server err: %v", err)
	}
}
