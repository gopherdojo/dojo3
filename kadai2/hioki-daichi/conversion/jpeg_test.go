package conversion

import (
	"reflect"
	"testing"
)

func TestConversion_Jpeg_MagicBytesSlice(t *testing.T) {
	t.Parallel()

	expected := [][]byte{[]byte("\xFF\xD8\xFF")}

	j := Jpeg{}

	actual := j.MagicBytesSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestConversion_Jpeg_HasProcessableExtname(t *testing.T) {
	j := Jpeg{}

	cases := map[string]struct {
		path     string
		expected bool
	}{
		"foo.jpg":  {path: "foo.jpg", expected: true},
		"foo.jpeg": {path: "foo.jpeg", expected: true},
		"foo.png":  {path: "foo.png", expected: false},
		"foo.gif":  {path: "foo.gif", expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			actual := j.HasProcessableExtname(c.path)
			if actual != c.expected {
				t.Errorf(`expected="%t" actual="%t"`, c.expected, actual)
			}
		})
	}
}
