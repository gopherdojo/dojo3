package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var q = question{}

func TestInput(t *testing.T) {
	t.Helper()
	file, err := os.Open("testdata/sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	result := q.input(file)
	expects := []string{"banana", "apple"}
	for _, expected := range expects {
		r := <-result
		if r != expected {
			t.Error(r)
			t.Error(expected)
		}
	}
}

type testQuestion struct{}

func (m *testQuestion) input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "banana"
	}()
	return ch
}

func TestRun(t *testing.T) {
	t.Helper()
	m := &testQuestion{}
	result := Run(words, m)
	if result != ExitCodeOK {
		t.Error(result)
	}
}
