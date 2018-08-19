package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/abemotion/convimg/img"
)

const (
	ExitCodeOK        int = iota // 0
	ExitCodeError                // 1
	ExitCodeFileError            // 2
)

func main() {
	cli := &CLI{os.Stdout, os.Stderr}
	os.Exit(cli.Run(os.Args))
}

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var dir, from, to string
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	fs.StringVar(&dir, "d", "./test/", "directory for converting imgs")
	fs.StringVar(&from, "f", "jpg", "extension to convert")
	fs.StringVar(&to, "t", "png", "extension after converting")
	fs.Parse(args[1:])

	if err := img.Convert(dir, from, to); err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	return ExitCodeOK
}
