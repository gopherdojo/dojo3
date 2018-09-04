package conversion

import (
	"reflect"
	"testing"
)

func TestConversion_Png_MagicBytesSlice(t *testing.T) {
	t.Parallel()

	expected := [][]byte{[]byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")}

	p := Png{}

	actual := p.MagicBytesSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestConversion_Png_HasProcessableExtname(t *testing.T) {
	p := Png{}

	cases := map[string]struct {
		path     string
		expected bool
	}{
		"foo.png":  {path: "foo.png", expected: true},
		"foo.jpg":  {path: "foo.jpg", expected: false},
		"foo.jpeg": {path: "foo.jpeg", expected: false},
		"foo.gif":  {path: "foo.gif", expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			actual := p.HasProcessableExtname(c.path)
			if actual != c.expected {
				t.Errorf(`expected="%t" actual="%t"`, c.expected, actual)
			}
		})
	}
}
