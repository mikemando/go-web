package main

import (
	"log"
	"net/http"

	"github.com/mikemando/go-go/pkg/config"
	"github.com/mikemando/go-go/pkg/handlers"
	"github.com/mikemando/go-go/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)
}
