package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai2/asuforce/converter"
)

// Exit codes are int values.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	outStraem, errStream io.Writer
}

// Run invokes the CLI
func (cli *CLI) Run(args []string) int {
	var (
		version bool
		path    string
		fromExt string
		toExt   string
	)

	flag.StringVar(&fromExt, "f", "jpg", "Specify input image extension")
	flag.StringVar(&fromExt, "from", "jpg", "Specify input image extension")
	flag.StringVar(&toExt, "t", "png", "Specify output image extension")
	flag.StringVar(&toExt, "to", "png", "Specify output image extension")

	flag.BoolVar(&version, "v", false, "Show version")
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()

	if version {
		fmt.Fprintf(cli.errStream, "Version: %s\n", Version)
		return ExitCodeOK
	}

	path = flag.Arg(0)

	collect := &converter.Collect{FromExt: fromExt}

	err := filepath.Walk(path, collect.CollectPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	c := &converter.Converter{FromExt: fromExt, ToExt: toExt}
	for _, i := range collect.Paths {
		err := c.Convert(i)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return ExitCodeError
		}
	}

	return ExitCodeOK
}
