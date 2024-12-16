package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

// Create a new form request
func New(urls url.Values) *Form {
	form := Form{
		urls,
		errors(map[string][]string{}),
	}

	return &form
}

// Just check if a post request have a given value
func (f *Form) Has(field string, r *http.Request) bool {

	found := r.Form.Get(field)

	if found != "" {
		return true
	}

	return false
}
