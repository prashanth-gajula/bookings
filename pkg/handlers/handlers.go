package handlers

import (
	config2 "github.com/prashanth-gajula/bookings/pkg/config"
	models2 "github.com/prashanth-gajula/bookings/pkg/models"
	renders2 "github.com/prashanth-gajula/bookings/pkg/renders"
	//"github/prashanth-gajula/go-course/pkg/models"
	//"github/prashanth-gajula/go-course/pkg/renders"
	"net/http"
)

// TemplateData will hold the data that will be loaded into templates
var Repo Repository

// Repository id Repository type
type Repository struct {
	App *config2.AppConfig
}

// creates new repository
func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// creates new handlers
func NewHandlers(r *Repository) {
	Repo = *r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	renders2.RenderTemplate(w, "home.page.html", &models2.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	renders2.RenderTemplate(w, "about.page.html", &models2.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, "make-reservation.page.html", &models2.TemplateData{})
}

// generals renders to rooms
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, "generals.page.html", &models2.TemplateData{})
}

//majors renders to rooms

func (m *Repository) majors(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, "majors.page.html", &models2.TemplateData{})
}
