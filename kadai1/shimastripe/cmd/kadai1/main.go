package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai1"
)

func main() {
	cli := &kadai1.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
