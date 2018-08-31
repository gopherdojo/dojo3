package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

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
		wg      sync.WaitGroup
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

	c, err := converter.NewConverter(path, fromExt, toExt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	err = filepath.Walk(c.Path, c.CrawlFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	queue := make(chan converter.Image)
	for _, image := range c.Files {
		wg.Add(1)
		go c.FetchConverter(queue, &wg)
		queue <- image
	}

	close(queue)
	wg.Wait()

	return ExitCodeOK
}
