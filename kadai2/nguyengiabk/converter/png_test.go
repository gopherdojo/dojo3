package converter_test

import (
	"bytes"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter"
)

func TestPNG(t *testing.T) {
	img := createImage(t)
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

var pngCheckExtTestFixtures = []struct {
	path   string
	result bool
}{
	{"image.png", true},
	{"image.jpg", false},
	{"image.doc", false},
}

func TestPngCheckExt(t *testing.T) {
	png := &converter.PNG{}
	for _, tc := range pngCheckExtTestFixtures {
		actual := png.CheckExt(tc.path)
		if actual != tc.result {
			t.Errorf("CheckExt(%v) failed, actual = %v, expected = %v", tc.path, actual, tc.result)
		}
	}
}
