package main

import (
	"errors"
	"net/http"
)

func main() {
	println("Server listening on :8080..")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		println("Serve fail: ", err.Error())
	}
}
