package main

import (
	"os"

	c "./cli"
	_ "./convert/gif"
	_ "./convert/jpg"
	_ "./convert/png"
	_ "github.com/dojo3/kadai1/daikurosawa/convert/gif"
)

func main() {
	cli := &c.Cli{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
