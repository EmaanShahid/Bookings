package handlers

import (
	"net/http"

	"github.com/EmaanShahid/bookings/pkg/config"
	"github.com/EmaanShahid/bookings/pkg/models"
	"github.com/EmaanShahid/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// home is the about page handler

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// about is the about page handler

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	//send the data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
