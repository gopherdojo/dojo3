package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/hioki-daichi/color"
	"github.com/gopherdojo/dojo3/kadai3-1/hioki-daichi/envutil"
	"github.com/gopherdojo/dojo3/kadai3-1/hioki-daichi/wording"
)

func main() {
	err := execute(os.Stdout, 15, "./weapons.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func execute(w io.Writer, defaultTimeout int, path string) error {
	timeout, err := envutil.GetIntEnvOrElse("TIMEOUT", defaultTimeout)
	if err != nil {
		return err
	}

	words, err := wording.NewWorder(path).Words()
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	score := play(w, getInputChannel(os.Stdin), words, time.Duration(timeout)*time.Second)

	printScore(w, score, timeout)
	return nil
}

func play(w io.Writer, inputChannel <-chan string, words []string, timeout time.Duration) (score map[rune]int) {
	timer := time.NewTimer(timeout)

	score = make(map[rune]int)

	for {
		word := words[rand.Intn(len(words))]
		printPrompt(w, word)

	LOOP_BY_WORD:
		for {
			select {
			case input := <-inputChannel:
				if word == input {
					break LOOP_BY_WORD
				}
				printIncorrect(w, word)
				printPrompt(w, word)
			case <-timer.C:
				return
			}
		}

		printCorrect(w, word)

		for _, c := range word {
			score[c]++
		}
	}
}

func getInputChannel(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func printScore(w io.Writer, hitsOfKey map[rune]int, totalTime int) {
	printWithBalloon(w, color.Cyan, "Finish!")

	var totalHits int
	for k := range hitsOfKey {
		totalHits += hitsOfKey[k]
	}

	fmt.Fprintf(w, "%-9s %d\n", "Hits:", totalHits)
	fmt.Fprintf(w, "%-9s %f\n", "Hits/sec:", float64(totalHits)/float64(totalTime))
	fmt.Fprintf(w, "Hits/key:\n")

	keys := make([]rune, 0, len(hitsOfKey))
	for k := range hitsOfKey {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		fmt.Fprintf(w, "  - [%s] : %s\n", string(k), strings.Repeat("*", hitsOfKey[k]))
	}
}

func printPrompt(w io.Writer, s string) {
	fmt.Fprintln(w, "> "+s)
}

func printCorrect(w io.Writer, s string) {
	printWithBalloon(w, color.Green, "Splatted "+s+"!")
}

func printIncorrect(w io.Writer, s string) {
	printWithBalloon(w, color.Red, "Splatted by "+s+"!")
}

func printWithBalloon(w io.Writer, c color.Color, s string) {
	l := len(s)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "\x1b["+c.Code()+"m%s\x1b[0m\n", "＿人"+strings.Repeat("人", l/2)+"人＿")
	fmt.Fprintf(w, "\x1b["+c.Code()+"m%s\x1b[0m\n", "＞　"+s+"　＜")
	fmt.Fprintf(w, "\x1b["+c.Code()+"m%s\x1b[0m\n", "￣Y^"+strings.Repeat("Y^", l/2)+"Y^￣")
	fmt.Fprintln(w)
}
