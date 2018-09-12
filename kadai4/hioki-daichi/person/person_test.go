package person

import (
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai4/hioki-daichi/fortune"
)

func TestPerson_NewPerson(t *testing.T) {
	expected := "*person.Person"

	p := NewPerson("Gopher", fortune.Daikichi)

	actual := reflect.TypeOf(p).String()
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestPerson_Validate(t *testing.T) {
	expected := "Name is too long (maximum is 32 characters)"

	p := NewPerson("123456789012345678901234567890123", fortune.Daikichi)
	p.Validate()

	actual := p.Errors[0]
	if actual != expected {
		t.Errorf(`unexpected : expected: "%s" actual: "%s"`, expected, actual)
	}
}
