package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"dir is empty", "main", 1},
		{"from and to are same", "main -d ./testdata/testimage.jpg -f jpg -t jpg", 1},
		{"success walking", "main -d ./testdata/testimage.jpg -f gif", 0},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			testRun(t, c.input, c.expected)
		})
	}
}

func testRun(t *testing.T, input string, expected int) {
	t.Helper()

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream, errStream}
	args := strings.Split(input, " ")
	status := cli.Run(args)
	if status != expected {
		t.Errorf("expected: %v, actual: %v", expected, status)
	}
}
