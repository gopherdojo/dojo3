package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Start Typing Game.")
	fmt.Println("Type the word appearing.")

	var correct int
	var words = []string{
		"archive",
		"tar",
		"zip",
		"bufio",
		"builtin",
		"bytes",
		"compress",
		"bzip2",
		"flate",
		"gzip",
		"lzw",
		"zlib",
		"container",
		"heap",
		"list",
		"ring",
		"context",
		"crypto",
		"aes",
		"cipher",
	}

	ch := input(os.Stdin)
	timeout := time.After(10 * time.Second)

	for {
		w := word(words)
		fmt.Println(w)

		select {
		case <-timeout:
			fmt.Println("Finish.")
			fmt.Printf("Your score: %d\n", correct)
			return
		case answer := <-ch:
			if w == answer {
				correct++
			}
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func word(s []string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s))]
}
