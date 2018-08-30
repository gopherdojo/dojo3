package converter

import (
	"testing"
)

func TestNewImageSuccess(t *testing.T) {
	i, err := NewImage(path)
	if err != nil {
		t.Errorf("failed test\ngot: %v", err)
	}

	actualPath := i.path
	actualName := i.name
	actualExt := i.ext

	expectedPath := "../testdata/test.jpeg"
	expectedName := "test"
	expectedExt := ".jpeg"

	if actualPath != expectedPath {
		t.Errorf("got: %v\nwant: %v", actualPath, expectedPath)
	}

	if actualName != expectedName {
		t.Errorf("got: %v\nwant: %v", actualName, expectedName)
	}

	if actualExt != expectedExt {
		t.Errorf("got: %v\nwant: %v", actualExt, expectedExt)
	}

}

func TestNewImageFailed(t *testing.T) {
	i, err := NewImage("")
	if err == nil {
		t.Fatal("failed test")
	}

	expected := Image{}
	if i != expected {
		t.Fatal("failed test")
	}
}

func TestGetFileNameSuccess(t *testing.T) {
	i, _ := NewImage(path)

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
	i, _ := NewImage(path)

	actual, err := i.GetFileName("")
	if err == nil {
		t.Fatal("failed test")
	}

	if actual != "" {
		t.Fatal("failed test")
	}
}
