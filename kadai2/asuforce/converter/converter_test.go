package converter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var path = "../testdata/test.jpeg"

func TestNewConverterSuccess(t *testing.T) {
	c, err := NewConverter(path, "jpeg", "png")
	if err != nil {
		t.Errorf("failed test: %v", err)
	}

	actualPath := c.Path
	actualFromExt := c.FromExt
	actualToExt := c.ToExt

	expectedPath := "../testdata/test.jpeg"
	expectedFromExt := "jpeg"
	expectedToExt := "png"

	if actualPath != expectedPath {
		t.Errorf("got: %v\nwant: %v", actualPath, expectedPath)
	}

	if actualFromExt != expectedFromExt {
		t.Errorf("got: %v\nwant: %v", actualFromExt, expectedFromExt)
	}

	if actualToExt != expectedToExt {
		t.Errorf("got: %v\nwant: %v", actualToExt, expectedToExt)
	}
}

func TestNewConverterFailed(t *testing.T) {
	c, err := NewConverter("", "jpeg", "png")
	if err == nil {
		t.Errorf("failed test: %v", err)
	}

	expected := Converter{}
	if cmp.Equal(c, expected) {
		t.Fatal("failed test: ")
	}
}
func TestConverter_FetchConverter(t *testing.T) {
}

func TestConverter_Convert(t *testing.T) {
}

func TestConverter_CrawlFile(t *testing.T) {
}

func Test_checkExtension(t *testing.T) {
}

func TestConverter_encodeImage(t *testing.T) {
}

func TestConverter_decodeImage(t *testing.T) {
}
