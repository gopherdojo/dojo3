package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/nguyengiabk/opt"
)

var wordList []string

const dataFileName = "wordlist.txt"

func init() {
	path := filepath.Join(dataFileName)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		wordList = append(wordList, s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func makeInputChannel(r io.Reader) <-chan string {
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

func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return wordList[rand.Intn(len(wordList))]
}

func start(w io.Writer, params *opt.Parameter, inputChan <-chan string, wordChan <-chan string) {
	var num, successNum, failNum int
	timeout := time.After(time.Duration(params.Timeout) * time.Second)
	var printResult = func() {
		fmt.Fprintf(w, "\nSuccess: %d, Fail: %d\n", successNum, failNum)
	}
	for {
		currentWord, ok := <-wordChan
		if !ok {
			printResult()
			return
		}
		num++
		fmt.Fprintf(w, "%d. %s\n> ", num, currentWord)
		select {
		case input, ok := <-inputChan:
			if !ok {
				printResult()
				return
			}
			if input == currentWord {
				successNum++
			} else {
				fmt.Fprintln(w, "\nWrong answer")
				failNum++
			}
		case <-timeout:
			fmt.Fprintln(w, "\nTime up!")
			printResult()
			return
		}
	}
}

func main() {
	params, err := opt.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------TYPING GAME--------")
	fmt.Println("Type the words that showup")

	wordChan := make(chan string)
	go func() {
		for {
			wordChan <- getRandomWord()
		}
	}()

	inputChan := makeInputChannel(os.Stdin)
	for {
		start(os.Stdout, params, inputChan, wordChan)
	ASK_LOOP:
		for {
			fmt.Println("Do you want to play again? (Y/N)")
			answer, ok := <-inputChan
			if !ok {
				close(wordChan)
				return
			}
			switch answer {
			case "Y":
				fallthrough
			case "y":
				break ASK_LOOP
			case "N":
				fallthrough
			case "n":
				close(wordChan)
				return
			}
		}
	}
}
