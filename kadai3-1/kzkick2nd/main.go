package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter a string")

	ch := input(os.Stdin)
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Time Over")
			return
		case answer := <-ch:
			fmt.Println(answer)
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
