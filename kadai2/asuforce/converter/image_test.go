package converter

import (
	"testing"
)

func TestGetFileNameSuccess(t *testing.T) {
	i := Image{path: path}

	actual, err := i.GetFileName("png")
	if err != nil {
		t.Errorf("failed test\ngot: %v", err)
	}

	expected := "test.png"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestGetFileNameFailed(t *testing.T) {
	i := Image{path: path}

	actual, err := i.GetFileName("")
	if err == nil {
		t.Fatal("failed test")
	}

	if actual != "" {
		t.Fatal("failed test")
	}
}
