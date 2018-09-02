package main

import (
	"os"
)

func main() {
	cli := &CLI{os.Stdout, os.Stderr}
	os.Exit(cli.Run(os.Args))
}
