package shimastripe

import (
	"bytes"
	"testing"
	"time"
)

func TestActionComplete(t *testing.T) {
	input := bytes.NewBufferString("a\nb\nc")
	output := new(bytes.Buffer)
	errput := new(bytes.Buffer)
	interval := time.Duration(3 * time.Second)
	wordList := []string{"a", "b", "c"}

	cli := &CLI{InStream: input, OutStream: output, ErrStream: errput, Interval: interval, WordList: wordList}

	if status := cli.Run([]string{}); status != success {
		t.Errorf("Status code: %v\n", status)
	}

	result := bytes.Split(output.Bytes(), []byte("\n"))
	lastLine := string(result[len(result)-1])

	if lastLine != "Result: 3" {
		t.Errorf("Last line: %v, expected: %v\n", lastLine, completeMessage)
	}
}

func TestActionTimeup(t *testing.T) {
	input := bytes.NewBufferString("a\nb")
	output := new(bytes.Buffer)
	errput := new(bytes.Buffer)
	interval := time.Duration(3 * time.Second)
	wordList := []string{"a", "b", "c"}

	cli := &CLI{InStream: input, OutStream: output, ErrStream: errput, Interval: interval, WordList: wordList}

	if status := cli.Run([]string{}); status != success {
		t.Errorf("Status code: %v\n", status)
	}

	result := bytes.Split(output.Bytes(), []byte("\n"))
	lastLine := string(result[len(result)-1])

	if lastLine != "Result: 2" {
		t.Errorf("Last line: %v, expected: %v\n", lastLine, completeMessage)
	}
}
