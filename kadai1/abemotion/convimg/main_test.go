package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		Args     []string
		WantCode int
	}{
		{
			strings.Split("convimg -d ./test/", " "),
			ExitCodeOK,
		},
		{
			strings.Split("convimg -d ./test/ -f jpg -t png", " "),
			ExitCodeOK,
		},
	}

	for _, tc := range cases {
		testRun(t, tc.Args, tc.WantCode)
	}
}

func testRun(t *testing.T, in []string, expectedCode int) {
	t.Helper()

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream, errStream}

	code := cli.Run(in)
	if errStream.String() != "" {
		t.Errorf("expected %q to eq %q", errStream.String(), "")
	}

	if code != expectedCode {
		t.Errorf("expected %q to eq %q", code, ExitCodeOK)
	}
}
