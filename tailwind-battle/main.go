package main

import (
	"embed"
	"log"

	"github.com/infinage/webe/tailwind-battle/internal"
)

//go:embed assets/*
var assets embed.FS

func main() {
	app, err := internal.NewApp(assets)
	if err != nil {
		log.Fatal("App init: ", err)
	}

	if err = app.Serve(":8080"); err != nil {
		log.Fatal("Serve fail: ", err)
	}
}
