package cli

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/gopherdojo/dojo3/kadai3-2/daikurosawa/download"
)

// Exit code.
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeParseUrlError
	ExitCodeProcessError
)

// Command line tool struct
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run command.
func (c *CLI) Run(args []string) int {
	var (
		parallel int64
		timeout  time.Duration
	)
	flags := flag.NewFlagSet("range-get", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.Int64Var(&parallel, "parallel", 5, "parallel count")
	flags.DurationVar(&timeout, "timeout", 300*time.Second, "time out")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeParseFlagError
	}

	rawUrl := flags.Arg(0)
	url, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Fprintln(c.ErrStream, err.Error())
		return ExitCodeParseUrlError
	}

	dl := download.NewDownloader(c.OutStream, url, parallel, timeout)
	if err := dl.Download(); err != nil {
		return ExitCodeProcessError
	}

	return ExitCodeOK
}
