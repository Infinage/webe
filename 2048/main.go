package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	// 0, 2, 4, 8, 16, 32, 64, 128, 512, 1024, 2048
}
