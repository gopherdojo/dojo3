package main

import (
	"os"
)

func main() {
	cli := &CLI{outStraem: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
