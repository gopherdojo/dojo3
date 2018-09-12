package form

import (
	"net/http/httptest"
	"testing"

	"github.com/gopherdojo/dojo3/kadai4/hioki-daichi/fortune"
)

func TestForm_NewRootForm(t *testing.T) {
	cases := map[string]struct {
		nameParam string
		expected  string
	}{
		"/?name=hioki-daichi": {nameParam: "hioki-daichi", expected: "hioki-daichi"},
		"/":                   {nameParam: "", expected: "Gopher"},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			nameParam := c.nameParam
			expected := c.expected

			req := httptest.NewRequest("GET", "/", nil)
			q := req.URL.Query()
			q.Add("name", nameParam)
			req.URL.RawQuery = q.Encode()
			f := NewRootForm(req)

			actual := f.name
			if actual != expected {
				t.Errorf(`unexpected name: expected: "%s" actual: "%s"`, expected, actual)
			}
		})
	}
}

func TestForm_NewPerson(t *testing.T) {
	name := "foo"

	f := RootForm{name: name}
	p := f.NewPerson(fortune.Daikichi)

	actual := p.Name
	if actual != name {
		t.Errorf(`unexpected name: expected: "%s" actual: "%s"`, name, actual)
	}
}
