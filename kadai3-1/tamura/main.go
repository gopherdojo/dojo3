package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var score int
	done := make(chan struct{}, 1)
	TimeUp(5*time.Second, done)

	ch := Run(os.Stdin, os.Stdout)

	go func() {
		for {
			fmt.Print(">")
			score = <-ch
		}
	}()

	<-done
	fmt.Printf("Your score is %v\n", score)
}
