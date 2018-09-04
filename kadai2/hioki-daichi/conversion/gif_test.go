package conversion

import (
	"reflect"
	"testing"
)

func TestConversion_Gif_MagicBytesSlice(t *testing.T) {
	t.Parallel()

	expected := [][]byte{[]byte("GIF87a"), []byte("GIF89a")}

	g := Gif{}

	actual := g.MagicBytesSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestConversion_Gif_HasProcessableExtname(t *testing.T) {
	g := Gif{}

	cases := []struct {
		path     string
		expected bool
	}{
		{path: "foo.gif", expected: true},
		{path: "foo.jpg", expected: false},
		{path: "foo.jpeg", expected: false},
		{path: "foo.png", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run("", func(t *testing.T) {
			t.Parallel()

			actual := g.HasProcessableExtname(c.path)
			if actual != c.expected {
				t.Errorf(`expected="%t" actual="%t"`, c.expected, actual)
			}
		})
	}
}
