package cli

import (
	"os"
	"testing"
)

type AssertFn func(err error)

func TestValidateArgs(t *testing.T) {
	if err := os.Mkdir("./dir", 0777); err != nil {
		t.Fatal("Failed to create directory.")
	}

	f, err := os.Create("./dir/file")
	if err != nil {
		t.Fatal("Failed to create file.")
	}
	defer f.Close()

	noError := func(err error) {
		if err != nil {
			t.Errorf("expected no error")
		}
	}

	withError := func(err error) {
		if err == nil {
			t.Errorf("expected returning error")
		}
	}

	cases := []struct {
		dir          string
		inputFormat  string
		outputFormat string
		assertFn     AssertFn
	}{
		{dir: "dir", inputFormat: "jpeg", outputFormat: "png", assertFn: noError},
		{dir: "hoge", inputFormat: "jpeg", outputFormat: "png", assertFn: withError},
		{dir: "file", inputFormat: "jpeg", outputFormat: "png", assertFn: withError},
		{dir: "dir", inputFormat: "fuga", outputFormat: "png", assertFn: withError},
		{dir: "dir", inputFormat: "png", outputFormat: "fuga", assertFn: withError},
	}

	for _, v := range cases {
		err := validateArgs(v.dir, v.inputFormat, v.outputFormat)
		v.assertFn(err)
	}

	if err := os.RemoveAll("./dir"); err != nil {
		t.Fatal("Failed to delete directory.")
	}
}
