package main

import (
	"os"

	c "github.com/gopherdojo/dojo3/kadai3-1/daikurosawa/cli"
)

func main() {
	cli := &c.CLI{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
