/*
Package form provides methods to generate models from Form and Form from Request.
*/
package form

import (
	"net/http"

	"github.com/gopherdojo/dojo3/kadai4/hioki-daichi/fortune"
	"github.com/gopherdojo/dojo3/kadai4/hioki-daichi/person"
)

// RootForm has name.
type RootForm struct {
	name string
}

// NewRootForm returns a form for the route of "/".
func NewRootForm(req *http.Request) *RootForm {
	f := &RootForm{}
	nameParam := req.URL.Query().Get("name")
	if nameParam != "" {
		f.name = nameParam
	} else {
		f.name = "Gopher"
	}
	return f
}

// NewPerson generates a person according to the content of form.
func (f *RootForm) NewPerson(ftn fortune.Fortune) *person.Person {
	return person.NewPerson(f.name, ftn)
}
