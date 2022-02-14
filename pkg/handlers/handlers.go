package handlers

import (
	"net/http"

	"github.com/mikemando/go-go/pkg/render"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.html")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.html")
}
