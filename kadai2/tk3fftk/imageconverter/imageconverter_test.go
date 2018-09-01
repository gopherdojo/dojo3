package imageconverter

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type mockWriteCloser struct {
	bytes.Buffer
}

func (wc mockWriteCloser) Close() error {
	return nil
}

func TestNew(t *testing.T) {
	cases := []struct {
		name     string
		from     string
		to       string
		expected ImageConverter
	}{
		{"jpg->gif", "jpg", "gif", ImageConverter{"jpg", "gif"}},
		{"jpg->png", "jpg", "png", ImageConverter{"jpg", "png"}},
		{"jpg->txt", "jpg", "txt", ImageConverter{}},
		{"png->gif", "png", "gif", ImageConverter{"png", "gif"}},
		{"png->jpg", "png", "jpg", ImageConverter{"png", "jpg"}},
		{"png->txt", "png", "txt", ImageConverter{}},
		{"gif->png", "gif", "png", ImageConverter{"gif", "png"}},
		{"gif->jpg", "gif", "jpg", ImageConverter{"gif", "jpg"}},
		{"gif->txt", "gif", "txt", ImageConverter{}},
		{"txt->jpg", "txt", "jpg", ImageConverter{}},
		{"jpg->jpg", "jpg", "jpg", ImageConverter{}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			testNew(t, c.from, c.to, c.expected)
		})
	}
}

func testNew(t *testing.T, from, to string, expected ImageConverter) {
	t.Helper()

	actual, _ := New(from, to)

	if !cmp.Equal(expected, actual, cmp.AllowUnexported(ImageConverter{})) {
		t.Errorf("expected=%v, actual=%v\n", expected, actual)
	}
}

func TestConvertImage(t *testing.T) {
	cases := []struct {
		name    string
		from    string
		to      string
		path    string
		success bool
	}{
		{"jpg->gif", "jpg", "gif", "./testdata/testimage.jpg", true},
		{"jpg->png", "jpg", "png", "./testdata/testimage.jpg", true},
		{"png->gif", "png", "gif", "./testdata/testimage.png", true},
		{"png->jpg", "png", "jpg", "./testdata/testimage.png", true},
		{"gif->png", "gif", "png", "./testdata/testimage.gif", true},
		{"gif->jpg", "gif", "jpg", "./testdata/testimage.gif", true},
		{"directory should not be converted", "jpg", "png", "./testdata", false},
		{".txt should not be converted", "jpg", "png", "./testdata/testimage.txt", false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			testConvertImage(t, c.from, c.to, c.path, c.success)
		})
	}
}

func testConvertImage(t *testing.T, from, to, path string, success bool) {
	t.Helper()

	output := new(mockWriteCloser)
	ic, err := New(from, to)
	if err != nil {
		t.Fatal("should not come here")
	}
	err = ic.ConvertImage(path, output)

	if !success {
		if !cmp.Equal(0, len(output.Bytes())) {
			t.Errorf("expected=%v, actual=%v\n", 0, output)
		}
	} else {
		if err != nil {
			t.Errorf("coverting to %s from %s should be success", to, from)
		}
	}
}
