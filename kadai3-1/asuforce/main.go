package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	words := []string{"hoge", "fuga", "piyo"}
	stdin := bufio.NewScanner(os.Stdin)
	score := 0

	rand.Seed(time.Now().UnixNano())

	for {
		theme := words[rand.Intn(3)]
		fmt.Println("Type it:", theme)

		for stdin.Scan() {
			answer := stdin.Text()
			if theme == answer {
				score++
				fmt.Println("ok")
			} else {
				fmt.Printf("got: %v, want: %v\n", answer, theme)
			}
			fmt.Println("Your score:", score)

			break
		}
	}
}
