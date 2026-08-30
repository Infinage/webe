package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	app, err := NewRipplesApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if err := http.ListenAndServe(":8080", app.Routes()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
