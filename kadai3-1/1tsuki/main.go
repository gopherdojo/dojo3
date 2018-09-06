// Command to convert several images extension.
package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo3/kadai3-1/1tsuki/typing_lesson"
	"io"
	"os"
	"time"
)

var writer io.Writer

const (
	exitCodeOK = iota
)

func init() {
	writer = os.Stdout
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(strArgs []string) int {
	var (
		dict string
	)
	flags := flag.NewFlagSet("TypingLesson", flag.PanicOnError)
	flags.StringVar(&dict, "d", "typing_lesson/_dictionaries/technologic.dic", "specify your original dictionary file")
	flags.Parse(strArgs)

	f, err := os.OpenFile(dict, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(writer, "Error occured during file read: %v", err)
	}
	defer f.Close()

	t, err := typing_lesson.NewTypingLesson(f, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(writer, "Error initializing TypingLesson: %v", err)
	}
	ok, ng := t.Start(10 * time.Second)
	fmt.Fprintf(writer, "Your score was %d / %d", ok, ok+ng)

	return exitCodeOK
}
