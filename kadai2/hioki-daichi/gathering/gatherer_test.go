package gathering

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/conversion"
)

// Decoder
var jpegD conversion.Jpeg
var pngD conversion.Png
var gifD conversion.Gif

// Encoder
var jpegE = &conversion.Jpeg{Options: &jpeg.Options{Quality: 1}}
var pngE = &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}
var gifE = &conversion.Gif{Options: &gif.Options{NumColors: 1}}

func TestGathering_Gather(t *testing.T) {
	t.Parallel()

	cases := []struct {
		decoder  conversion.Decoder
		expected []string
	}{
		{decoder: &jpegD, expected: []string{"../testdata/jpeg/sample1.jpg", "../testdata/jpeg/sample2.jpg", "../testdata/jpeg/sample3.jpeg"}},
		{decoder: &pngD, expected: []string{"../testdata/png/sample1.png", "../testdata/png/sample2.png"}},
		{decoder: &gifD, expected: []string{"../testdata/gif/sample1.gif"}},
	}

	for _, c := range cases {
		c := c
		t.Run("", func(t *testing.T) {
			g := Gatherer{Decoder: c.decoder}

			actual, _ := g.Gather("../testdata/")
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf(`expected="%s" actual="%s"`, c.expected, actual)
			}
		})
	}
}

func TestGathering_Gather_Nonexistence(t *testing.T) {
	t.Parallel()

	expected := "lstat nonexistent_path: no such file or directory"

	g := Gatherer{Decoder: &jpegD}

	_, err := g.Gather("nonexistent_path")
	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_Unopenable(t *testing.T) {
	t.Parallel()

	tempdir, _ := ioutil.TempDir("", "imgconv")
	defer os.RemoveAll(tempdir)

	path := filepath.Join(tempdir, "unopenable.jpg")
	if _, err := os.OpenFile(path, os.O_CREATE, 000); err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.Remove(path)

	expected := "open " + path + ": permission denied"

	g := Gatherer{Decoder: &jpegD}

	_, err := g.Gather(tempdir)

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_FailedToCheckDecodable(t *testing.T) {
	t.Parallel()

	expected := "EOF"

	g := Gatherer{Decoder: &jpegD}

	_, err := g.Gather("./testdata/empty.jpg")

	actual := err.Error()
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestGathering_Gather_Undecodable(t *testing.T) {
	t.Parallel()

	g := Gatherer{Decoder: &jpegD}

	_, err := g.Gather("./testdata/undecodable.jpg")
	if err != nil {
		t.Fatalf("err %s", err)
	}
}
