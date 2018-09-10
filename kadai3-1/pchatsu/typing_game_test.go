package typinggame

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestNewTypingGame(t *testing.T) {
	type args struct {
		limitSec int
		r        io.Reader
		w        io.Writer
	}
	tests := []struct {
		name string
		args args
		want *TypingGame
	}{
		{"#1 normal case", args{5, bytes.NewBufferString("a"), &bytes.Buffer{}}, &TypingGame{limitSec: 5, r: bytes.NewBufferString("a"), w: &bytes.Buffer{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := NewTypingGame(tt.args.limitSec, tt.args.r, w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTypingGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAnswer(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"#1 normal case", args{bytes.NewBufferString("foo\nbar\nbaz\n")}, []string{"foo", "bar", "baz"}},
		{"#2 normal case", args{bytes.NewBufferString("foo\nbar\nbaz")}, []string{"foo", "bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan string)
			go readAnswer(tt.args.r, ch)
			for _, w := range tt.want {
				if got := <-ch; got != w {
					t.Errorf("channel got = %v, want %v", got, w)
				}
			}
		})
	}
}

func TestTypingGame_Play(t *testing.T) {
	// for tests
	questions = []string{"foo", "bar", "baz"}

	tests := []struct {
		name        string
		r           io.Reader
		wantCorrect int
		wantW       string
	}{
		{"#1 normal case", bytes.NewBufferString("foo\nbar\n"), 2, "foo\nbar\nbaz\ntime over! correct:2"},
		{"#2 1 mistake case", bytes.NewBufferString("foo\nbaz\nbar"), 2, "foo\nbar\nbaz\ntime over! correct:2"},
		{"#3 loop case", bytes.NewBufferString("foo\nbar\nbaz\nfoo\n"), 4, "foo\nbar\nbaz\nfoo\nbar\ntime over! correct:4"},
	}
	for _, tt := range tests {
		w := &bytes.Buffer{}
		tg := NewTypingGame(1, tt.r, w)
		t.Run(tt.name, func(t *testing.T) {
			tg.Play()
			if gotCorrect := tg.correct; gotCorrect != tt.wantCorrect {
				t.Errorf("correct got = %v, want %v", gotCorrect, tt.wantW)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("io.Writer got = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
