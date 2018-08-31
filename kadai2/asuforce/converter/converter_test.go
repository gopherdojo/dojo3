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

func TestConverter_encodeImage(t *testing.T) {
}

func TestConverter_decodeImage(t *testing.T) {
}
