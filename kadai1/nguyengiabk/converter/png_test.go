package converter_test

import (
	"bytes"
	"image"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/nguyengiabk/converter"
)

func TestPNG(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	buf := &bytes.Buffer{}
	png := &converter.PNG{}
	if err := png.Encode(buf, img); err != nil {
		t.Errorf("Cannot encode PNG")
	}
	if _, err := png.Decode(buf); err != nil {
		t.Errorf("Cannot decode PNG")
	}
	if png.GetExt() != ".png" {
		t.Errorf("PNG type must have .png extension")
	}
}
