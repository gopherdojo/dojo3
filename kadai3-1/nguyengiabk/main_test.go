package main

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/nguyengiabk/opt"
)

var testMakeInputChannelFixtures = map[string]struct {
	input    string
	expected []string
}{
	"test normal case": {
		input:    "test\ntyping\ngame\n",
		expected: []string{"test", "typing", "game"},
	},
	"test empty case": {
		input:    "",
		expected: nil,
	},
	"test only one word case": {
		input:    "test",
		expected: []string{"test"},
	},
}

func TestMakeInputChannel(t *testing.T) {
	for name, tc := range testMakeInputChannelFixtures {
		t.Run(name, func(t *testing.T) {
			buf := bytes.Buffer{}
			ch := makeInputChannel(&buf)
			buf.WriteString(tc.input)
			var output []string
			for s := range ch {
				output = append(output, s)
			}
			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("Channel I/O mismatch, input = %v, output = %v", tc.expected, output)
			}
		})
	}
}

var testStartFixtures = map[string]struct {
	words     []string
	sleepTime int
	input     string
	expected  string
}{
	"test normal case": {
		words:     []string{"worker", "world", "worry"},
		sleepTime: 0,
		input:     "worker\nworld\nworry",
		expected:  "1. worker\n> 2. world\n> 3. worry\n> \nSuccess: 3, Fail: 0\n",
	},
	"test wrong answer": {
		words:     []string{"worker", "world", "worry"},
		sleepTime: 0,
		input:     "wrong\nworld\nworry",
		expected:  "1. worker\n> \nWrong answer\n2. world\n> 3. worry\n> \nSuccess: 2, Fail: 1\n",
	},
	"test EOF of input": {
		words:     []string{"worker", "world", "worry", "more"},
		sleepTime: 0,
		input:     "worker\nworld\nworry",
		expected:  "1. worker\n> 2. world\n> 3. worry\n> 4. more\n> \nSuccess: 3, Fail: 0\n",
	},
	"test timeout": {
		words:     []string{"worker", "world", "worry"},
		sleepTime: 2,
		input:     "worker\nworld\nworry",
		expected:  "1. worker\n> \nTime up!\n\nSuccess: 0, Fail: 0\n",
	},
}

func TestStart(t *testing.T) {
	for name, tc := range testStartFixtures {
		t.Run(name, func(t *testing.T) {
			inputBuffer := bytes.Buffer{}
			inputBuffer.WriteString(tc.input)
			inputChan := makeInputChannelWithSleep(&inputBuffer, tc.sleepTime)
			outputBuffer := bytes.Buffer{}

			wordChan := make(chan string)
			go func() {
				for _, word := range tc.words {
					wordChan <- word
				}
				close(wordChan)
			}()

			start(&outputBuffer, &opt.Parameter{Timeout: 1}, inputChan, wordChan)
			if tc.expected != outputBuffer.String() {
				t.Errorf("Game result is unexpected, expected = %v, actual = %v", tc.expected, outputBuffer.String())
			}
		})
	}
}

func makeInputChannelWithSleep(r io.Reader, interval int) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			time.Sleep(time.Duration(interval) * time.Second)
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
