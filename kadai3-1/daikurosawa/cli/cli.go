package cli

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/word"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeProcessError
	green = "\x1b[32m%s\x1b[0m\n"
	red   = "\x1b[31m%s\x1b[0m\n"
)

// Command line tool struct
type CLI struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer
	word.Word
}

// Run command
func (c *CLI) Run(args []string) int {
	var limit time.Duration
	flags := flag.NewFlagSet("type-game", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.DurationVar(&limit, "limit", 30*time.Second, "time limit")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeParseFlagError
	}

	path := flags.Arg(0)
	c.Word = word.NewWordFile(path)

	println(limit)

	if err := c.play(path, limit); err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeProcessError
	}
	return ExitCodeOK
}

func (c *CLI) play(path string, limit time.Duration) error {

	if err := c.Generate(); err != nil {
		return err
	}
	ch := c.scan()
	cnt := 0
	word, err := c.GetWord()
	if err != nil {
		return err
	}
	c.print(word)
	timer := time.NewTimer(limit)

LOOP:
	for {
		select {
		case input := <-ch:
			next, err := c.verify(input, word, &cnt)
			if err != nil {
				return err
			}
			word = next
			c.print(word)
		case <-timer.C:
			fmt.Fprintln(c.OutStream, "Time out!")
			break LOOP
		}
	}

	fmt.Fprintf(c.OutStream, "Result: %v\n", cnt)
	return nil
}

func (c *CLI) scan() <-chan string {
	ch := make(chan string)
	scan := bufio.NewScanner(c.InStream)

	go func() {
		for scan.Scan() {
			ch <- scan.Text()
		}
	}()

	return ch
}

func (c *CLI) print(word string) {
	fmt.Fprintf(c.OutStream, "> %s\n", word)
}

func (c *CLI) verify(actual, answer string, count *int) (string, error) {
	if actual == answer {
		fmt.Fprintf(c.OutStream, green, "Success!")
		*count++
	} else {
		fmt.Fprintf(c.OutStream, red, "Failed")
	}
	return c.GetWord()
}
