package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	i := input(os.Stdin)

	bc := context.Background()
	t := 10000 * time.Millisecond
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	result := 0
	q := makeQuestion()
	fmt.Println("Type it:", q)
	for {
		select {
		case answer := <-i:
			if answer == q {
				result++
			}
			fmt.Println("Your score:", result)
			q = makeQuestion()
			fmt.Println("Type it:", q)
		case <-ctx.Done():
			fmt.Println("Time up !")
			fmt.Println("Your final score:", result)
			return
		}
	}
}

func makeQuestion() string {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	return q
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		std := bufio.NewScanner(r)
		for std.Scan() {
			ch <- std.Text()
		}
		close(ch)
	}()
	return ch
}
