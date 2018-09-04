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

	cases := []struct {
		path     string
		expected bool
	}{
		{path: "foo.jpg", expected: true},
		{path: "foo.jpeg", expected: true},
		{path: "foo.png", expected: false},
		{path: "foo.gif", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run("", func(t *testing.T) {
			t.Parallel()

			actual := j.HasProcessableExtname(c.path)
			if actual != c.expected {
				t.Errorf(`expected="%t" actual="%t"`, c.expected, actual)
			}
		})
	}
}
