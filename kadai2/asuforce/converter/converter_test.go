package converter

import (
	"testing"
)

var path = "../testdata/test.jpeg"

func TestConverter_FetchConverter(t *testing.T) {
}

func TestConverter_Convert(t *testing.T) {
}

func TestConverter_CrawlFile(t *testing.T) {
}

func Test_getFileName(t *testing.T) {
	c := &Converter{Encoder: &Png{}}

	actual, err := c.getFileName(path)
	if err != nil {
		t.Errorf("failed test\ngot: %v", err)
	}

	expected := "test.png"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual, err = c.getFileName("")
	if err == nil {
		t.Errorf("failed test\ngot: %v", err)
	}

	if actual != "" {
		t.Fatal("failed test")
	}
}

func TestConverter_encodeImage(t *testing.T) {
}

func TestConverter_decodeImage(t *testing.T) {
}
