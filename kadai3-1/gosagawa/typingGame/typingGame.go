package typingGame

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var questions []string

const playTime int = 10

func init() {
	questions = getQuestions()
}

func Start() {
	var num int
	var endGame bool
	inputCh := input(os.Stdin)

	startDisplay()
	<-inputCh

	timeCh := time.After(time.Duration(playTime) * time.Second)

	for {
		word := chooseWord()
		displayWord(word)

		select {
		case input := <-inputCh:
			isCorrect := check(word, input)
			displayResult(isCorrect)
			if isCorrect {
				num++
			}

		case <-timeCh:
			endDisplay(num)
			endGame = true
		}

		if endGame {
			break
		}
	}
}

func getQuestions() []string {
	var List = []string{"white", "yellow", "orange", "red", "pink", "purple", "blue", "green", "brown", "grey", "black"}
	return List
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

func startDisplay() {
	fmt.Println("Please type same words as much as possible ...")
	fmt.Println("Hit any key to start..")
	fmt.Println("")

}
func chooseWord() string {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	return q

	return "test"
}

func displayWord(word string) {
	fmt.Println(word)
	fmt.Printf(">")
}

func check(word string, input string) bool {
	if word == input {
		return true
	}
	return false
}

func displayResult(isCorrect bool) {
	if isCorrect {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Miss...")
	}
	fmt.Println("")
}

func endDisplay(correctNumber int) {
	fmt.Println("")
	fmt.Println("Time's up !!!")
	fmt.Printf("Your score is %v\n", correctNumber)
}
