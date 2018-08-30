package converter

import "testing"

func TestNewImageSuccess(t *testing.T) {
	i := NewImage("../testdata/test.jpeg")

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
