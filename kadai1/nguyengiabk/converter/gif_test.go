package converter_test

import (
	"bytes"
	"image"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/nguyengiabk/converter"
)

func TestGIF(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	buf := &bytes.Buffer{}
	gif := &converter.GIF{NumColors: 256}
	if err := gif.Encode(buf, img); err != nil {
		t.Errorf("Cannot encode GIF")
	}
	if _, err := gif.Decode(buf); err != nil {
		t.Errorf("Cannot decode GIF")
	}
	if gif.GetExt() != ".gif" {
		t.Errorf("GIF type must have .gif extension")
	}
}
