# Tailwind Battle

A real-time, lightweight CSS code-golfing platform built specifically for Tailwind CSS.

Tailwind Battle allows users to recreate target UI reference images using pure HTML and Tailwind utility classes, with a backend that calculates pixel-perfect diff scores on every keystroke.

Live: [tailwind-battle.infinage.space](https://tailwind-battle.infinage.space)

## Features

- **Live Pixel-Perfect Diffing:** Compares your markup against the target image in real-time, calculating a percentage score based on pixel accuracy.
- **Interactive Visualizers:** Toggle between an interactive slider or a hardware-accelerated difference blend to pinpoint exact misalignments.
- **Zero-Bloat Architecture:** No React, no heavy IDEs. Built with Datastar and CodeJar to keep the frontend bundle under 15kB.

## Tech Stack

- **Backend:** Go (Standard Library `net/http`) + `golang.org/x/image/draw`
- **Frontend State/SSE:** [Datastar](https://data-star.dev/)
- **Editor:** [CodeJar](https://medv.io/codejar/) (2kB) + PrismJS for syntax highlighting
- **Styling Engine:** Tailwind CSS v4 (via browser CDN for live sandbox rendering)
- **DOM Capture:** [snapdom](https://github.com/romainmenke/snapdom)

## How It Works

1. **The Sandbox:** User HTML is injected into an isolated `iframe` with fixed boundaries to prevent layout leakage. 
2. **The Capture:** `snapdom` takes a snapshot of the iframe's DOM and writes it to an offscreen canvas.
3. **The Proxy:** The Go backend proxies external reference images, scaling them safely to exactly 400x300 to eliminate browser CORS taint errors and Retina DPI inflation.
4. **The Diff:** The frontend passes the Base64 image string to Go via Datastar. The Go backend compares the RGBA values of both images, calculates the score, and pushes it back via Server-Sent Events (SSE).

## Running Locally

Requires Go 1.22+ installed on your machine.

```bash
# Download dependencies
go mod download

# Run the server
go run main.go
