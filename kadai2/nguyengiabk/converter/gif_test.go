package converter_test

import (
	"bytes"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter"
)

func TestGIF(t *testing.T) {
	img := createImage(t)
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

var gifCheckExtTestFixtures = []struct {
	path   string
	result bool
}{
	{"image.gif", true},
	{"image.png", false},
	{"image.doc", false},
}

func TestGifCheckExt(t *testing.T) {
	gif := &converter.GIF{NumColors: 256}
	for _, tc := range gifCheckExtTestFixtures {
		actual := gif.CheckExt(tc.path)
		if actual != tc.result {
			t.Errorf("CheckExt(%v) failed, actual = %v, expected = %v", tc.path, actual, tc.result)
		}
	}
}
