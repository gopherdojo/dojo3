package typing

import (
	"context"
	"time"

	"fmt"

	"bufio"
	"os"

	"math/rand"

	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/color"
	"github.com/gopherdojo/dojo3/kadai3-1/shuntaka9576/words"
)

const (
	red    = "\x1b[31m%s\x1b[0m"
	green  = "\x1b[32m%s\x1b[0m"
	yellow = "\x1b[32m%s\x1b[0m"
)

type result struct {
	correct, uncorrect int
}

type typing struct {
	words  words.Words
	result result
}

func New() *typing {
	return &typing{}
}

func (t *typing) SetWords(ws words.Words) *typing {
	return &typing{ws, result{0, 0}}
}

func (t *typing) Run(ctx context.Context, duration time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	inputCh := make(chan string)

	go t.typingGame(ctx, inputCh)

	select {
	case <-ctx.Done():
		fmt.Println("Finished!!!")
		printColor("************* result *************",color.Blue)
		printColor(fmt.Sprintf("%v/%v", t.result.correct, t.result.correct+t.result.uncorrect), color.Blue)
		printColor("**********************************",color.Blue)
	}
}

func (t *typing) typingGame(ctx context.Context, inputCh chan string) {
	for {
		word := t.words[rand.Intn(len(t.words))]
		fmt.Fprintln(os.Stdout, word.Input)
		go func(chan string) {
			scan := bufio.NewScanner(os.Stdin)
			for scan.Scan() {
				inputCh <- scan.Text()
			}
		}(inputCh)

		select {
		case text := <-inputCh:
			if text == word.Expected {
				printColor("Success!", color.Green)
				t.result.correct++
			} else {
				printColor("Fail!", color.Red)
				t.result.uncorrect++
			}
		}
	}
}

func printColor(stdout string, color color.Color) {
	fmt.Printf("\x1b["+color.Code()+"m"+"%s"+"\x1b[0m\n", stdout)
}
