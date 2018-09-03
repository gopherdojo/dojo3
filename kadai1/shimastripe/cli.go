package shimastripe

import (
	"flag"
	"fmt"
	"io"
)

// CLI struct for main
type CLI struct {
	OutStream, ErrStream io.Writer
}

const (
	success = iota
	flagError
	traverseError
)

// Run command
func (c *CLI) Run(args []string) int {
	var fromExt, toExt string

	flag.StringVar(&fromExt, "from", "", "Input file format. (Ex. jpg, png, gif)\n")
	flag.StringVar(&toExt, "to", "", "Output file format. (Ex. jpg, png, gif)\n")

	flag.Parse()

	if fromExt == "" || toExt == "" {
		fmt.Fprintf(c.ErrStream, "ParameterError: fromExt: %v, toExt: %v.\n", fromExt, toExt)
		return flagError
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintf(c.ErrStream, "Specify one directory.\n")
		return flagError
	}

	srcDir := flag.Arg(0)

	if err := Traverse(srcDir, fromExt, toExt); err != nil {
		fmt.Fprintf(c.ErrStream, "Traverse error: %v", err)
		return traverseError
	}

	return success
}
