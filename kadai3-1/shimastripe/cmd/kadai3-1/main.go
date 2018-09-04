package main

import (
	"os"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-1/shimastripe"
)

func main() {
	cli := &shimastripe.CLI{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr, Interval: 1 * time.Minute}
	os.Exit(cli.Run(os.Args))
}
