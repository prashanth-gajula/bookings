package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/prashanth-gajula/bookings/internal/config"
	"github.com/prashanth-gajula/bookings/internal/forms"
	models2 "github.com/prashanth-gajula/bookings/internal/models"
	//"github.com/prashanth-gajula/bookings/internal/models"
	renders2 "github.com/prashanth-gajula/bookings/internal/renders"
	"log"

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

	renders2.RenderTemplate(w, r, "home.page.html", &models2.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	renders2.RenderTemplate(w, r, "about.page.html", &models2.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyreservation models2.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyreservation
	renders2.RenderTemplate(w, r, "make-reservation.page.html", &models2.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models2.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)
	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		renders2.RenderTemplate(w, r, "make-reservation.page.html", &models2.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", "reservation")
	http.Redirect(w, r, "/Reservation-Summary", http.StatusSeeOther)

}

// generals renders to rooms
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, r, "generals.page.html", &models2.TemplateData{})
}

//majors renders to rooms

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, r, "majors.page.html", &models2.TemplateData{})
}

// renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, r, "search-availability.html", &models2.TemplateData{})
}

// renders the post search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	//every data we get from a form is a string we need to type cast it based on our requirement.
	w.Write([]byte(fmt.Sprintf("start data is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// Availabilityjson handles request for availability and send json responses.
func (m *Repository) Availabilityjson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal("error in marshaling file:", err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// renders the contact page page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renders2.RenderTemplate(w, r, "contact.html", &models2.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models2.Reservation)
	if !ok {
		fmt.Println("cannot get data out of session")
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = reservation

	renders2.RenderTemplate(w, r, "reservation-summary.page.html", &models2.TemplateData{
		Data: data,
	})
}
