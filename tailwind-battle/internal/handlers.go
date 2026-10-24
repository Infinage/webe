package internal

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

// datastarShowErrPopup is a helper utility to send error popup to the frontend
func datastarShowErrPopup(w http.ResponseWriter, r *http.Request, errMessage string) {
	sse := datastar.NewSSE(w, r)
	signals := map[string]any{"showError": true, "errorMessage": errMessage}
	sse.MarshalAndPatchSignals(signals)
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
		datastarShowErrPopup(w, r, "Read target URL")
		return
	}

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(targetURL)
	if err != nil {
		log.Println("Read target URL:", err)
		datastarShowErrPopup(w, r, "Read target URL")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Get target NOT OK")
		datastarShowErrPopup(w, r, "Get target NOT OK")
		return
	}

	target, _, err := image.Decode(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		log.Println("Decode target image:", err)
		datastarShowErrPopup(w, r, "Decode target image")
		return
	}

	dstImg := rescaleImage(target, 400, 300)

	// Write into temporary memory buffer to encode it
	var buf bytes.Buffer
	if err := png.Encode(&buf, dstImg); err != nil {
		log.Println("Re-encode target image: ", err)
		datastarShowErrPopup(w, r, "Re-encode target image: ")
		return
	}

	targetEncoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	targetWeights, targetTotalWeight := getWeights(dstImg)

	signals := map[string]any{
		"targetImgString":      targetEncoded,
		"targetImgWeights":     targetWeights,
		"targetImgTotalWeight": targetTotalWeight,
	}

	sse := datastar.NewSSE(w, r)
	sse.MarshalAndPatchSignals(signals)

	// Truncate the url for logging
	truncatedURL := targetURL
	if maxLen := 30; len(targetURL) > maxLen {
		half := (maxLen - 2) / 2
		truncatedURL = targetURL[:half] + "..." + targetURL[len(targetURL)-half:]
	}

	log.Printf("Target image [%s] encoded successfully", truncatedURL)
}

// handleAPIDiffImages compares the inbound input and target image and
// scores based on an inverse weighted frequency.
func (app *App) handleAPIDiffImages(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Source            string             `json:"source"`
		TargetString      string             `json:"targetString"`
		TargetWeights     map[uint16]float64 `json:"targetWeights"`
		TargetTotalWeight float64            `json:"targetTotalWeight"`
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

	var targetRaw []byte
	if data.TargetString == "" {
		targetRaw, err = app.assets.ReadFile("assets/images/default.png")
		if err != nil {
			log.Println("Read default source:", err)
			datastarShowErrPopup(w, r, "Read default source")
			return
		}
	} else {
		targetRaw, err = base64.StdEncoding.DecodeString(data.TargetString)
		if err != nil {
			log.Println("Read target string:", err)
			datastarShowErrPopup(w, r, "Read target string")
			return
		}
	}

	target, _, err := image.Decode(bytes.NewReader(targetRaw))
	if err != nil {
		log.Println("Target image decode:", err)
		datastarShowErrPopup(w, r, "Load target image from string")
		return
	}

	// If not set recalculate it
	targetWeights, targetTotalWeight := data.TargetWeights, data.TargetTotalWeight
	if len(data.TargetWeights) == 0 || data.TargetTotalWeight <= 0. {
		targetWeights, targetTotalWeight = getWeights(target)
	}

	score, err := diffImages(source, target, targetWeights, targetTotalWeight)
	if err != nil {
		log.Println("Diff image:", err)
		datastarShowErrPopup(w, r, "Diff image")
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.MarshalAndPatchSignals(map[string]any{
		"score":                score,
		"targetImgWeights":     targetWeights,
		"targetImgTotalWeight": targetTotalWeight,
	})
}
