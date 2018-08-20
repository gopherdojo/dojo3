package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai1/shimastripe"
)

func main() {
	cli := &shimastripe.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
