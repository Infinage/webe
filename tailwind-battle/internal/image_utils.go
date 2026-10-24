package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"math"
	"strings"

	"golang.org/x/image/draw"
)

// colorKey is used in scoring the user created image. It quantizes the
// pixel colors to ensure minor color stirations / noise don't impact scoring.
func colorKey(img image.Image, x, y int) uint16 {
	r, g, b, _ := img.At(x, y).RGBA()
	return uint16((r>>11)<<10 | (g>>11)<<5 | (b >> 11))
}

// getWeights computes color quantized weights for each pixel in the input image.
func getWeights(img image.Image) (map[uint16]float64, float64) {
	bounds := img.Bounds()

	// Count colors across the target
	pixelCounts := make(map[uint16]int)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelCounts[colorKey(img, x, y)]++
		}
	}

	// Weight of a color = 1 / sqrt(freq)
	var totalWeight float64
	weights := make(map[uint16]float64)
	for color, count := range pixelCounts {
		weight := 1 / math.Sqrt(float64(count))
		weights[color] = weight
		totalWeight += float64(count) * weight
	}

	return weights, totalWeight
}

// diffImages uses inverse frequency weighted score comparing
// pixels across both images to derive at a score.
func diffImages(source, target image.Image, targetWeights map[uint16]float64,
	totalWeight float64) (float64, error) {
	bounds := source.Bounds()
	if bounds != target.Bounds() {
		return 0, fmt.Errorf("image size mismatch: [%v, %v]", source.Bounds(),
			target.Bounds())
	}

	var score float64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			key := colorKey(target, x, y)
			if key == colorKey(source, x, y) {
				score += targetWeights[key]
			}
		}
	}

	scoreF := 100. * score / totalWeight
	return math.Round(scoreF*100) / 100, nil
}

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

// rescaleImage redraws the input image into the required shape.
func rescaleImage(img image.Image, width, height int) image.Image {
	dstImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(dstImg, dstImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dstImg
}
