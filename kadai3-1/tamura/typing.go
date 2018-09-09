package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var dictionary = []string{
	"book",
	"word",
	"dog",
	"woman",
	"good",
	"coffee",
	"tea",
	"clean",
	"yakitori",
}

func TimeUp(d time.Duration, ch chan struct{}) {
	go func() {
		time.Sleep(d)
		fmt.Println("\nTime is up!")
		close(ch)
	}()
}

func Run(r io.Reader, w io.Writer) <-chan int {
	var score int
	ch := make(chan int, 1)
	update := make(chan bool, 1)
	update <- true
	q := make(chan string, 1)

	go func() {
		for {
			q <- Question(update, w)
		}
	}()

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ans := s.Text()
			if ans == <-q {
				fmt.Fprintln(w, "ok")
				update <- true
				score++
			} else {
				fmt.Fprintln(w, "ng")
			}
			ch <- score
		}
		if err := s.Err(); err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		close(ch)
	}()

	return ch
}

func Question(update <-chan bool, w io.Writer) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	select {
	case <-update:
		d := dictionary[r.Intn(len(dictionary))]
		fmt.Fprintf(w, "%s\n", d)
		return d
	}
}
