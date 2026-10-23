package main

import (
	"bytes"
	"embed"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/starfederation/datastar-go/datastar"
	"golang.org/x/image/draw"
)

//go:embed assets/*
var assets embed.FS

// loadImageFromBase64Blob takes in a raw base64 encoded blob strips the
// prefix, base64 decodes it and parses it to png image.
func loadImageFromBase64Blob(raw string) (image.Image, error) {
	header, imgEncoded, ok := strings.Cut(raw, ",")
	if !ok || header != "data:image/png;base64" {
		return nil, fmt.Errorf("source not png blob")
	}

	data, err := base64.StdEncoding.DecodeString(imgEncoded)
	if err != nil {
		return nil, fmt.Errorf("base64 decode: %w", err)
	}

	image, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return image, nil
}

func diffImages(source, target image.Image) (float64, error) {
	bounds := source.Bounds()
	if bounds != target.Bounds() {
		return 0, fmt.Errorf("image size mismatch: [%v, %v]", source.Bounds(),
			target.Bounds())
	}

	var score uint64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := source.At(x, y).RGBA()
			r2, g2, b2, _ := target.At(x, y).RGBA()
			if r1 == r2 && g1 == g2 && b1 == b2 {
				score++
			}
		}
	}

	scoreF := 100. * float64(score) / float64(bounds.Dx()*bounds.Dy())
	return scoreF, nil
}

// datastarShowErrPopup is a helper utility to send error popup to the frontend
func datastarShowErrPopup(w http.ResponseWriter, r *http.Request, errMessage string) {
	sse := datastar.NewSSE(w, r)
	signals := map[string]any{"showError": true, "errorMessage": errMessage}
	sse.MarshalAndPatchSignals(signals)
}

func main() {
	templ, err := template.ParseFS(assets, "assets/html/index.html")
	if err != nil {
		log.Fatal("template parse: ", err)
	}

	mux := http.NewServeMux()

	assetsServer := http.FileServerFS(assets)
	mux.Handle("GET /assets/", assetsServer)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err = templ.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/target", func(w http.ResponseWriter, r *http.Request) {
		var targetURL string
		err := datastar.ReadSignals(r, &targetURL)
		if err != nil || targetURL == "" {
			datastarShowErrPopup(w, r, "Read target URL")
			return
		}

		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(targetURL)
		if err != nil {
			datastarShowErrPopup(w, r, "Read target URL")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			datastarShowErrPopup(w, r, "Get target NOT OK")
			return
		}

		target, _, err := image.Decode(io.LimitReader(resp.Body, 2<<20))
		if err != nil {
			datastarShowErrPopup(w, r, "Decode target image")
			return
		}

		dstImg := image.NewRGBA(image.Rect(0, 0, 400, 300))
		draw.BiLinear.Scale(dstImg, dstImg.Bounds(), target, target.Bounds(), draw.Over, nil)

		var buf bytes.Buffer
		err = png.Encode(&buf, dstImg)
		if err != nil {
			datastarShowErrPopup(w, r, "Re-encode target image")
			return
		}

		targetEncoded := base64.StdEncoding.EncodeToString(buf.Bytes())
		signals := map[string]any{"targetImgString": targetEncoded}

		sse := datastar.NewSSE(w, r)
		sse.MarshalAndPatchSignals(signals)
	})

	mux.HandleFunc("POST /api/diff", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Source       string `json:"source"`
			TargetString string `json:"targetString"`
		}

		err = datastar.ReadSignals(r, &data)
		if err != nil {
			datastarShowErrPopup(w, r, "read signals")
			return
		}

		source, err := loadImageFromBase64Blob(data.Source)
		if err != nil {
			log.Println("Source image load: ", err)
			datastarShowErrPopup(w, r, "load source image from blob")
			return
		}

		var targetRaw []byte
		if data.TargetString == "" {
			targetRaw, err = assets.ReadFile("assets/images/default.png")
			if err != nil {
				datastarShowErrPopup(w, r, "read default source")
				return
			}
		} else {
			targetRaw, err = base64.StdEncoding.DecodeString(data.TargetString)
			if err != nil {
				datastarShowErrPopup(w, r, "read target string")
				return
			}
		}

		target, _, err := image.Decode(bytes.NewReader(targetRaw))
		if err != nil {
			log.Println("Target image decode: ", err)
			datastarShowErrPopup(w, r, "load target image from string")
			return
		}

		score, err := diffImages(source, target)
		if err != nil {
			log.Println("diff image:", err)
			datastarShowErrPopup(w, r, "diff image")
			return
		}

		score = math.Round(score*100) / 100

		sse := datastar.NewSSE(w, r)
		sse.MarshalAndPatchSignals(map[string]any{"score": score})
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Print("serve fail: ", err)
	}
}
