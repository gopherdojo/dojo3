package main

import (
	"flag"
	"fmt"
	"io"
	"os"

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

	err := collect.CollectPath(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	c := &converter.Converter{
		Encoder: cli.switchEncoder(toExt),
		Decoder: cli.switchDecoder(fromExt),
	}
	for _, i := range collect.Paths {
		err := c.Convert(i)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return ExitCodeError
		}
	}

	return ExitCodeOK
}

func (cli *CLI) switchEncoder(ext string) converter.Encoder {
	switch ext {
	case "jpg", "jpeg":
		return &converter.Jpg{}
	case "gif":
		return &converter.Gif{}
	default:
		return &converter.Png{}
	}
}

func (cli *CLI) switchDecoder(ext string) converter.Decoder {
	switch ext {
	case "png":
		return &converter.Png{}
	case "gif":
		return &converter.Gif{}
	default:
		return &converter.Jpg{}
	}
}
