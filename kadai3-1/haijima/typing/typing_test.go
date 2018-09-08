package typing

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	ch := time.After(100 * time.Millisecond)
	words := []string{"hello", "world"}
	r := strings.NewReader("hello\nmundo\n")

	w := &bytes.Buffer{}
	stop := make(chan bool)
	go func() {
		Exec(r, w, ch, words)
		stop <- true
	}()
	<-stop
	want := "--- Typing Game Start!! ---\nhello\nworld\nhello\n\n\n--- Time Over!! ---\n" +
		"--- Result ---\n" +
		"Typed 1 words\n" +
		"Succeed Rate 50.0%\n" +
		"--------------\n"
	if got := w.String(); got != want {
		t.Errorf("Should print result. got: %v, want: %v", got, want)
	}
}

func Test_input(t *testing.T) {
	r := &bytes.Buffer{}
	ch := input(r)
	r.Write([]byte("Hello "))
	r.Write([]byte("world\n"))
	r.Write([]byte("gopher\n"))

	str := <-ch
	if str != "Hello world" {
		t.Errorf("ch = %v, want Hello world", str)
	}
	str = <-ch
	if str != "gopher" {
		t.Errorf("ch = %v, want gopher", str)
	}
}

func Test_printHeader(t *testing.T) {
	want := "--- Typing Game Start!! ---\n"
	w := &bytes.Buffer{}
	printHeader(w)
	if got := w.String(); got != want {
		t.Errorf("printHeader() = %v, want %v", got, want)
	}
}

func Test_printTimeOver(t *testing.T) {
	want := "\n\n--- Time Over!! ---\n"
	w := &bytes.Buffer{}
	printTimeOver(w)
	if got := w.String(); got != want {
		t.Errorf("printTimeOver() = %v, want %v", got, want)
	}
}
