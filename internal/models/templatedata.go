package models

import "github.com/prashanth-gajula/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	error     string
	Warning   string
	Form      *forms.Form
}
