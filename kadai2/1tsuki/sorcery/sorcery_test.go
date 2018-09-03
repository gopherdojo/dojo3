package sorcery

import (
	"testing"
	"bytes"
	"os"
	"io/ioutil"
	"io"
	"fmt"
)

var buffer *bytes.Buffer
func init() {
	buffer = &bytes.Buffer{}
}

func ExampleSorcery_Exec() {

}

func TestExec(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempDir)

	type checkFunc func(io.Writer, error) error
	check := func(fns ...checkFunc) []checkFunc { return fns }

	isSomeError := func() checkFunc {
		return func(writer io.Writer, err error) error {
			if err == nil {
				return fmt.Errorf("expected error but has not occured")
			}
			return nil
		}
	}

	isSuccess := func() checkFunc {
		return func(writer io.Writer, err error) error {
			if err != nil {
				return fmt.Errorf("unexpected error %v", err)
			}
			return nil
		}
	}

	tests := [...]struct{
		name string
		from imgExt
		to imgExt
		dir string
		checks []checkFunc
	}{
		{"success", Jpeg, Png, tempDir, check(isSuccess())},
		{"error with invalid from", end, Jpeg, tempDir, check(isSomeError())},
		{"error with invalid to", Jpeg, end, tempDir, check(isSomeError())},
		{"error with invalid dir", Jpeg, Png, "", check(isSomeError())},
	}

	s := Sorcery(buffer)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			err := s.Exec(tt.from, tt.to, tt.dir)
			for _, check := range tt.checks {
				if err := check(buffer, err); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
