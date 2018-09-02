// This is command line tool to convert image to specified format.
package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai2/matsumatsu20/cli"
)

func main() {
	cli := &cli.CLT{
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
