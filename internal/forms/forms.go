package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

// form creates custom form struct embeds url.values object
type Form struct {
	url.Values
	Errors errors
}

// valid returns true if there are no errors otherwise false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// new initilize the form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// checks for required fields.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// checks wether the required field is available  or not in the form.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		//f.Errors.Add(field, "this field cannot be blank")
		return false
	}
	return true
}

// checks wether the fields are of required lenhgth or not.
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field should be minimum of length %d", length))
		return false
	}
	return true
}

// checks wether the given value is valid email or not
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Not valid email address")
	}
}
