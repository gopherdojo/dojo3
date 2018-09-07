package main

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestMain_execute(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	err := execute(&buf, 1, "./weapons.txt")
	if err != nil {
		t.Fatalf("err %s", err)
	}
}

func TestMain_execute_NonExistentPath(t *testing.T) {
	t.Parallel()

	var err error
	var buf bytes.Buffer
	err = execute(&buf, 1, "./non-existent-path")
	actual := err.Error()
	expected := "open ./non-existent-path: no such file or directory"
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestMain_play(t *testing.T) {
	expected := map[rune]int{97: 1}

	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond) // waiting for start
		ch <- "b"                          // incorrect
		ch <- "a"                          // correct
		time.Sleep(200 * time.Millisecond) // waiting for finish
	}()

	var buf bytes.Buffer
	actual := play(&buf, ch, []string{"a"}, 200*time.Millisecond)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%v" actual="%v"`, expected, actual)
	}

	cases := map[string]struct {
		s        string
		expected bool
	}{
		"print ok a": {s: "Splatted a!", expected: true},
		"print ng a": {s: "Splatted by a!", expected: true},
		"print ok b": {s: "Splatted b!", expected: false},
		"print ng b": {s: "Splatted by b!", expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			r := regexp.MustCompile(c.s)
			actual := r.MatchString(buf.String())
			expected := c.expected
			if actual != expected {
				t.Errorf(`expected="%t" actual="%t"`, expected, actual)
			}
		})
	}
}

func TestMain_getInputChannel(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	buf.Write([]byte("a"))
	expected := "a"
	actual := <-getInputChannel(&buf)
	if actual != expected {
		t.Errorf(`expected="%s" actual="%s"`, expected, actual)
	}
}

func TestMain_printScore(t *testing.T) {
	var buf bytes.Buffer
	printScore(&buf, map[rune]int{97: 1, 98: 2}, 10)

	cases := map[string]struct {
		s string
	}{
		"print Finish!":            {s: "Finish!"},
		"print Hits/sec: 0.300000": {s: "Hits/sec: 0.300000"},
		"print Hits/key:":          {s: "Hits/key:"},
		"print - [a] : *":          {s: `- \[a\] : \*`},
		"print - [b] : **":         {s: `- \[b\] : \*\*`},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			r := regexp.MustCompile(c.s)
			matched := r.MatchString(buf.String())
			if !matched {
				t.Errorf("%q should match %q", buf.String(), c.s)
			}
		})
	}
}
