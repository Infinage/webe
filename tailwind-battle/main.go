package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/starfederation/datastar-go/datastar"
)

func main() {
	templ, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		log.Fatal("template parse: ", err)
	}

	mux := http.NewServeMux()

	assetsServer := http.FileServer(http.Dir("assets/"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetsServer))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err = templ.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/diff", func(w http.ResponseWriter, r *http.Request) {
		var data struct { Image string `json:"image"` }
		err = datastar.ReadSignals(r, &data)
		if err != nil {
			http.Error(w, "read signals", http.StatusBadRequest)
			return
		}

		fmt.Println(data.Image)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Print("serve fail: ", err)
	}
}
