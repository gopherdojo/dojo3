package typing

import (
	"io"
	"time"
	"fmt"
	"bufio"
)

func Exec(r io.Reader, w io.Writer, words []string, timeoutSec time.Duration) {
	count := 0
	success := 0

	ch := input(r)
	timeout := time.After(timeoutSec * time.Second)

	printHeader(w)
	word := words[0]
	fmt.Fprintln(w, word)

while:
	for {
		select {
		case typed := <-ch:
			if typed == word {
				success++
			}
			count++
			word = words[count%len(words)]
			fmt.Fprintln(w, word)
		case <-timeout:
			printTimeOver(w)
			break while
		}
	}
	printResult(w, success, count, timeoutSec)
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
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
	fmt.Println(w)
	fmt.Println(w)
	fmt.Println(w, "--- Time Over!! ---")
}

func printResult(w io.Writer, success int, count int, timeoutSec time.Duration) {
	fmt.Fprintln(w, "--- Result ---")
	fmt.Fprintf(w, "Typed %d words\n", success)
	fmt.Fprintf(w, "Success Rate %d%%\n", 100*success/count)
	fmt.Fprintf(w, "Average %.4f words/sec\n", float64(success)/float64(timeoutSec))
	fmt.Fprintln(w, "--------------")
}
