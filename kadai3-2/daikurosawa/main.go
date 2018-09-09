package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai3-2/daikurosawa/cli"
)

func main() {
	cli := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
