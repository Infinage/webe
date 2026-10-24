package internal

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

// datastarShowErrPopup is a helper utility to send error popup to the frontend
func datastarShowErrPopup(w http.ResponseWriter, r *http.Request, errMessage string) {
	sse := datastar.NewSSE(w, r)
	signals := map[string]any{"showError": true, "errorMessage": errMessage, "score": 0.}
	sse.MarshalAndPatchSignals(signals)
}

func truncateURL(u string, maxLen int) string {
	truncated := u
	if len(u) > maxLen {
		half := (maxLen - 2) / 2
		truncated = u[:half] + "..." + u[len(u)-half:]
	}
	return truncated
}

func processTargetImageURL(ustr string, width, height int) (image.Image, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(ustr)
	if err != nil {
		return nil, fmt.Errorf("read target URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get target NOT OK")
	}

	if !strings.HasPrefix(resp.Header.Get("content-type"), "image/") {
		return nil, fmt.Errorf("target url returned non image")
	}

	target, _, err := image.Decode(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return nil, fmt.Errorf("decode target image: %w", err)
	}

	dstImg := rescaleImage(target, width, height)
	return dstImg, nil
}

// handleAPIHome returns the homepage for tailwind-battle playground
func (app *App) handleAPIHome(w http.ResponseWriter, r *http.Request) {
	if err := app.templ.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleAPIEncodeTargetImage fetches the image from URL provided and
// returns back base64 version of it to cache in the front end.
func (app *App) handleAPIEncodeTargetImage(w http.ResponseWriter, r *http.Request) {
	var targetURL string
	err := datastar.ReadSignals(r, &targetURL)
	if err != nil || targetURL == "" {
		log.Println("Read target URL:", err)
		datastarShowErrPopup(w, r, "Read signals")
		return
	}

	if _, ok := app.cache.Get(targetURL); ok {
		log.Printf("Target image [%s] already cached", truncateURL(targetURL, 30))
		return
	}

	dstImg, err := processTargetImageURL(targetURL, 400, 300)
	if err != nil {
		log.Println("Process target image URL:", err)
		datastarShowErrPopup(w, r, "Process target image URL")
		return
	}

	app.cache.Set(targetURL, dstImg)
	log.Printf("Target image [%s] encoded successfully", truncateURL(targetURL, 30))
}

// handleAPIDiffImages compares the inbound input and target image and
// scores based on an inverse weighted frequency.
func (app *App) handleAPIDiffImages(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Source    string `json:"source"`
		TargetURL string `json:"targetURL"`
	}

	err := datastar.ReadSignals(r, &data)
	if err != nil {
		log.Println("Read signals:", err)
		datastarShowErrPopup(w, r, "Read signals")
		return
	}

	source, err := loadImageFromBase64Blob(data.Source)
	if err != nil {
		log.Println("Source image load:", err)
		datastarShowErrPopup(w, r, "Load source image from blob")
		return
	}

	if _, ok := app.cache.Get(data.TargetURL); !ok {
		var target image.Image
		if data.TargetURL == "" {
			targetRaw, err := app.assets.ReadFile("assets/images/default.png")
			if err != nil {
				log.Println("Read default source:", err)
				datastarShowErrPopup(w, r, "Read default source")
				return
			}

			target, _, err = image.Decode(bytes.NewReader(targetRaw))
			if err != nil {
				log.Println("Decode default source:", err)
				datastarShowErrPopup(w, r, "Decode default source")
				return
			}

		} else {
			target, err = processTargetImageURL(data.TargetURL, 400, 300)
			if err != nil {
				log.Println("Process target image:", err)
				datastarShowErrPopup(w, r, "Process target image")
				return
			}
		}

		log.Printf("Cache updated for: %q", truncateURL(data.TargetURL, 30))
		app.cache.Set(data.TargetURL, target)
	}

	timg, ok := app.cache.Get(data.TargetURL)
	if !ok {
		log.Println("Target image missing from cache after set")
		datastarShowErrPopup(w, r, "Internal cache error")
		return
	}

	score, err := diffImages(source, timg.Image, timg.pixelWeights, timg.totalWeight)
	if err != nil {
		log.Println("Diff image:", err)
		datastarShowErrPopup(w, r, "Diff image")
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.MarshalAndPatchSignals(map[string]any{"score": score})
}
