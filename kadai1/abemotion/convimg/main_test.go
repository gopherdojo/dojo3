package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream, errStream}
	args := strings.Split("convimg -d ./test/", " ")

	code := cli.Run(args)
	if errStream.String() != "" {
		t.Errorf("expected %q to eq %q", errStream.String(), "")
	}

	if code != ExitCodeOK {
		t.Errorf("expected %q to eq %q", code, ExitCodeOK)
	}
}
