package imageconverter

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNew(t *testing.T) {
	expected := ImageConverter{
		".jpg",
		".gif",
	}
	actual := New("jpg", "gif")

	if !cmp.Equal(expected, actual, cmp.AllowUnexported(ImageConverter{})) {
		t.Errorf("expected=%v, actual=%v\n", expected, actual)
	}
}

func TestOpenImage(t *testing.T) {
	cases := []struct{path string; successs bool}{
		{"./testdata/testimage.jpg", true},
		{"./testdata/notexist.jpg", false},
	}
	ic := &ImageConverter{
		".jpg",
		".gif",
	}

	for _, c := range cases {
		c := c
		_, err := ExportImageConverterOpenImage(ic, c.path)
		if err != nil && c.successs {
			t.Errorf("expected=%v, got '%v'\n", c, err)
		}
	}
}

func TestPrepareImage(t *testing.T) {

}