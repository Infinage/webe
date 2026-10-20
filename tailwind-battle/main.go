package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

func main() {
	templ, err := template.ParseFiles("assets/html/index.html")
	if err != nil {
		log.Fatal("template parse: ", err)
	}

	mux := http.NewServeMux()

	assetsServer := http.FileServer(http.Dir("assets/"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetsServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = templ.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Print("serve fail: ", err)
	}
}
