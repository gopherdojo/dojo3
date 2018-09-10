package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	ExitCodeOK = iota
)

var words = []string{"apple", "banana", "strawberry", "kiwi", "cherry", "melon", "plum", "pear", "pineapple", "grape", "peach"}

type inputer interface {
	input(io.Reader) <-chan string
}

type question struct {
}

type Vocabulary []string

type Score struct {
	correctInput int
	totalInput   int
}

func (w Vocabulary) chooseWord() string {
	index := rand.Intn(len(w))
	return w[index]
}

func (q *question) input(r io.Reader) <-chan string {
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

func Run(candidates []string, i inputer) int {
	var score Score
	ch := i.input(os.Stdin)
	timeout := time.After(30 * time.Second)
	fmt.Println("---------Typing Start!! ---------")
	for {
		answer := Vocabulary.chooseWord(candidates)
		fmt.Println(answer)
		fmt.Print("> ")
		select {
		case input := <-ch:
			if answer == input {
				score.correctInput++
				fmt.Printf("correct!!\n\n")
			} else {
				fmt.Printf("failed!\n\n")
			}
			score.totalInput++
		case <-timeout:
			fmt.Printf("\n\n---------Timed Out---------\n\n")
			fmt.Printf("total_input : %d  correct_answer : %d\n", score.totalInput, score.correctInput)
			return ExitCodeOK
		}
	}
}

func main() {
	q := &question{}
	os.Exit(Run(words, q))
}
