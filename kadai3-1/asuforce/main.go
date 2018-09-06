package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var questions = []string{"hoge", "fuga", "piyo"}

func main() {
	i := input(os.Stdin)

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
		case <-time.After(5 * time.Second):
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
