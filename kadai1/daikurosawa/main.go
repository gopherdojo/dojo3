package main

import (
	"os"

	c "./cli"
	_ "./convert/gif"
	_ "./convert/jpg"
	_ "./convert/png"
)

func main() {
	cli := &c.Cli{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
