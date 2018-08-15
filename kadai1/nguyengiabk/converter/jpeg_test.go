package converter_test

import (
	"bytes"
	"image"
	"testing"

	"github.com/gopherdojo/dojo3/kadai1/nguyengiabk/converter"
)

func TestJPG(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	buf := &bytes.Buffer{}
	jpeg := &converter.JPEG{Quality: 100}
	if err := jpeg.Encode(buf, img); err != nil {
		t.Errorf("Cannot encode JPEG")
	}
	if _, err := jpeg.Decode(buf); err != nil {
		t.Errorf("Cannot decode JPEG")
	}
	if jpeg.GetExt() != ".jpg" {
		t.Errorf("JPEG type must have .jpg extension")
	}
}
