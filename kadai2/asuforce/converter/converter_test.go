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

func Test_checkExtension(t *testing.T) {
	actual, err := checkExtension(".jpeg")

	expected := ".jpg"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual, err = checkExtension("")
	if err == nil {
		t.Errorf("failed test: %v", err)
	}

	if actual != "" {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

}

func Test_getFileName(t *testing.T) {
	c := &Converter{ToExt: "png"}

	actual, err := c.getFileName(path)
	if err != nil {
		t.Errorf("failed test\ngot: %v", err)
	}

	expected := "test.png"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	c.ToExt = ""
	actual, err = c.getFileName("")
	if err == nil {
		t.Fatal("failed test")
	}

	if actual != "" {
		t.Fatal("failed test")
	}
}

func TestConverter_encodeImage(t *testing.T) {
}

func TestConverter_decodeImage(t *testing.T) {
}
