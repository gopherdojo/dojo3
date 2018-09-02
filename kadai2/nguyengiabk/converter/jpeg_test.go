package converter_test

import (
	"bytes"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter"
)

func TestJPG(t *testing.T) {
	img := createImage(t)
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

var jpegCheckExtTestFixtures = []struct {
	path   string
	result bool
}{
	{"image.jpg", true},
	{"image.jpeg", true},
	{"image.png", false},
	{"image.doc", false},
}

func TestJpegCheckExt(t *testing.T) {
	jpeg := &converter.JPEG{Quality: 100}
	for _, tc := range jpegCheckExtTestFixtures {
		actual := jpeg.CheckExt(tc.path)
		if actual != tc.result {
			t.Errorf("CheckExt(%v) failed, actual = %v, expected = %v", tc.path, actual, tc.result)
		}
	}
}
