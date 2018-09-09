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
	ch := input(os.Stdin)

	questions := []string{"facebook", "google", "apple", "uber", "yahoo"}
	rand.Seed(time.Now().UnixNano())

	score := 0

	t := time.After(time.Duration(10) * time.Second)

	for {
		s := questions[rand.Intn(len(questions))]
		fmt.Printf(">%s\n", s)
		select {
		case v := <-ch:
			if s == v {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong!")
			}
		case <-t:
			fmt.Println("Time up!")
			fmt.Println("Your score:", score)
			return
		}
	}
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
