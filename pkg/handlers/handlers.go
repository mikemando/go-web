package handlers

import (
	"net/http"

	"github.com/mikemando/go-go/pkg/config"
	"github.com/mikemando/go-go/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers/
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.html")
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.html")
}
