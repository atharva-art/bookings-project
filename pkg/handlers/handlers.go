package handlers

import (
	"github.com/atharva-art/bookings-project/pkg/config"
	"github.com/atharva-art/bookings-project/pkg/models"
	"github.com/atharva-art/bookings-project/pkg/render"
	"net/http"
)

// Repo is the Repository used by handlers
var Repo *Repository

// Repository is a repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHandlers sets repository for handlers
func NewHandlers(r *Repository){
	Repo = r
}

//Home is the Home Page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is thr about page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
