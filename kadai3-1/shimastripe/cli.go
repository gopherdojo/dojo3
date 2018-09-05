package shimastripe

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// CLI struct for main
type CLI struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer
	Interval             time.Duration
	WordList             []string
}

const (
	success = iota
	wordListError
)

const (
	completeMessage = "You answered all words correctly."
)

// Run command
func (c *CLI) Run(args []string) int {
	if len(c.WordList) < 1 {
		fmt.Fprintf(c.ErrStream, "Word list is insufficient.")
		return wordListError
	}

	c.Action(c.WordList)
	return success
}

func (c *CLI) Input() <-chan string {
	ch := make(chan string)

	go func() {
		sc := bufio.NewScanner(c.InStream)

		for sc.Scan() {
			ch <- sc.Text()
		}

		close(ch)
	}()

	return ch
}

func (c *CLI) Action(wordList []string) error {
	answerWord := wordList[0]
	wordList = wordList[1:]
	fmt.Fprintf(c.OutStream, "TIME: %v minutes\n", c.Interval.Minutes())
	fmt.Fprintln(c.OutStream, answerWord)
	ch := c.Input()
	counter := 0

L:
	for {
		select {
		case userInput := <-ch:
			if userInput == answerWord {
				counter += 1

				if len(wordList) < 1 {
					fmt.Fprintln(c.OutStream, completeMessage)
					break L
				}

				answerWord = wordList[0]
				wordList = wordList[1:]
				fmt.Fprintln(c.OutStream, answerWord) // new word
			} else {
				fmt.Fprintln(c.OutStream, "wrong!")
			}

		case <-time.After(c.Interval):
			break L
		}
	}

	fmt.Fprintf(c.OutStream, "Result: %v", counter)
	return nil
}
