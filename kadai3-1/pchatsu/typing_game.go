package typinggame

import (
	"bufio"
	"io"
	"strconv"
	"time"
)

type TypingGame struct {
	limitSec int
	correct  int
	r        io.Reader
	w        io.Writer
}

func NewTypingGame(limitSec int, r io.Reader, w io.Writer) *TypingGame {
	tg := TypingGame{
		limitSec: limitSec,
		r:        r,
		w:        w,
	}
	return &tg
}

func readAnswer(r io.Reader, ch chan<- string) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		ch <- sc.Text()
	}
}

func (t *TypingGame) Play() {
	ch := make(chan string)
	go readAnswer(t.r, ch)

	timer := time.After(time.Duration(t.limitSec) * time.Second)
	question := questions[t.correct]

	t.w.Write([]byte(question + "\n"))
LOOP:
	for {
		select {
		case s := <-ch:
			if question == s {
				t.correct++
				question = questions[t.correct%len(questions)]
				t.w.Write([]byte(question + "\n"))
			}
		case <-timer:
			t.w.Write([]byte("time over! correct:" + strconv.Itoa(t.correct)))
			break LOOP
		}
	}
}
