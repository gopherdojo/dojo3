/*
Package person is a package that manages processing around person.
*/
package person

import "github.com/gopherdojo/dojo3/kadai4/hioki-daichi/fortune"

// Person has Name and Fortune.
type Person struct {
	Name    string          `json:"name"`
	Fortune fortune.Fortune `json:"fortune"`
	Errors  []string        `json:"-"`
}

// NewPerson generates a new person.
func NewPerson(n string, f fortune.Fortune) *Person {
	return &Person{
		Name:    n,
		Fortune: f,
	}
}

// Validate validates its own fields.
func (p *Person) Validate() {
	if len(p.Name) > 32 {
		p.Errors = append(p.Errors, "Name is too long (maximum is 32 characters)")
	}
}
