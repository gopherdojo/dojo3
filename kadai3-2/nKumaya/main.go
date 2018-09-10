package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/dojo3/kadai3-2/nKumaya/kget"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeInvalidUrlError
	ExitCodeCreateHTTPClient
	ExitCodeErrorDownload
	ExitCodeErrorCansel
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("kget", flag.ContinueOnError)
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	url := args[1]
	fmt.Fprintln(c.outStream, "Checking now", url)
	response, err := http.Get(url)
	if response.StatusCode != 200 {
		return ExitCodeInvalidUrlError
	}
	client, err := kget.NewClient(url)
	if err != nil {
		return ExitCodeCreateHTTPClient
	}
	bc := context.Background()
	ctx, cancel := context.WithCancel(bc)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	defer func() {
		signal.Stop(ch)
		cancel()
	}()
	fmt.Fprintln(c.outStream, "Download start", url)
	go func() error {
		select {
		case <-ch:
			cancel()
			return nil
		case <-ctx.Done():
			if err = client.DeleteFiles(); err != nil {
				return err
			}
			return nil
		}
	}()
	err = client.Download(ctx)
	if err != nil {
		return ExitCodeErrorDownload
	}
	fmt.Fprintln(c.outStream, "Complite")
	return ExitCodeOK
}

func main() {
	cli := &CLI{os.Stdout, os.Stderr}
	os.Exit(cli.Run(os.Args))
}
