package typing

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

func Exec(r io.Reader, w io.Writer, timeout <-chan time.Time, words []string) {
	var stat stat
	ch := input(r)

	printHeader(w)
	word := words[0]
	fmt.Fprintln(w, word)

while:
	for {
		select {
		case typed := <-ch:
			if typed == word {
				stat.Succeed()
			} else {
				stat.Fail()
			}
			word = words[stat.Count()%len(words)]
			fmt.Fprintln(w, word)
		case <-timeout:
			printTimeOver(w)
			break while
		}
	}
	fmt.Fprintln(w, stat.String())
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		//defer close(ch)
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func printHeader(w io.Writer) {
	fmt.Fprintln(w, "--- Typing Game Start!! ---")
}

func printTimeOver(w io.Writer) {
	fmt.Fprintln(w, "\n\n--- Time Over!! ---")
}
