package gathering

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/conversion"
)

func TestGathering_Gather(t *testing.T) {
	cases := map[string]struct {
		decoder  conversion.Decoder
		expected []string
	}{
		"JPEG": {decoder: jpegDecoder(t), expected: []string{"../testdata/jpeg/sample1.jpg", "../testdata/jpeg/sample2.jpg", "../testdata/jpeg/sample3.jpeg"}},
		"PNG":  {decoder: pngDecoder(t), expected: []string{"../testdata/png/sample1.png", "../testdata/png/sample2.png"}},
		"GIF":  {decoder: gifDecoder(t), expected: []string{"../testdata/gif/sample1.gif"}},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			g := Gatherer{Decoder: c.decoder}

			actual, err := g.Gather("../testdata/")
			if err != nil {
				t.Fatalf("err %s", err)
			}
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf(`expected="%s" actual="%s"`, c.expected, actual)
			}
		})
	}
}

func TestGathering_Gather_Nonexistence(t *testing.T) {
	t.Parallel()

	expected := "lstat nonexistent_path: no such file or directory"

	g := Gatherer{Decoder: jpegDecoder(t)}

	_, err := g.Gather("nonexistent_path")
	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_Unopenable(t *testing.T) {
	t.Parallel()

	tempdir, err := ioutil.TempDir("", "imgconv")
	if err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.RemoveAll(tempdir)

	path := filepath.Join(tempdir, "unopenable.jpg")
	if _, err := os.OpenFile(path, os.O_CREATE, 000); err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.Remove(path)

	expected := "open " + path + ": permission denied"

	g := Gatherer{Decoder: jpegDecoder(t)}

	_, err = g.Gather(tempdir)

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_FailedToCheckDecodable(t *testing.T) {
	t.Parallel()

	expected := "EOF"

	g := Gatherer{Decoder: jpegDecoder(t)}

	_, err := g.Gather("./testdata/empty.jpg")

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_Undecodable(t *testing.T) {
	t.Parallel()

	g := Gatherer{Decoder: jpegDecoder(t)}

	_, err := g.Gather("./testdata/undecodable.jpg")
	if err != nil {
		t.Fatalf("err %s", err)
	}
}

func jpegDecoder(t *testing.T) *conversion.Jpeg {
	t.Helper()
	var d conversion.Jpeg
	return &d
}

func pngDecoder(t *testing.T) *conversion.Png {
	t.Helper()
	var d conversion.Png
	return &d
}

func gifDecoder(t *testing.T) *conversion.Gif {
	t.Helper()
	var d conversion.Gif
	return &d
}
