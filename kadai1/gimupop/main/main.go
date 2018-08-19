package main

import (
	"os"
	"github.com/gopherdojo/dojo3/kadai1/gimupop"
)
func main() {
	cli := &gimupop.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}