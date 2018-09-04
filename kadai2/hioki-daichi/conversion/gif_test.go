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

	cases := map[string]struct {
		path     string
		expected bool
	}{
		"foo.gif":  {path: "foo.gif", expected: true},
		"foo.jpg":  {path: "foo.jpg", expected: false},
		"foo.jpeg": {path: "foo.jpeg", expected: false},
		"foo.png":  {path: "foo.png", expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			actual := g.HasProcessableExtname(c.path)
			if actual != c.expected {
				t.Errorf(`expected="%t" actual="%t"`, c.expected, actual)
			}
		})
	}
}
