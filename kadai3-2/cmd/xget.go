package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo3/kadai3-2"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	procs = flag.Int("p", runtime.NumCPU(), "the number of parallel workers")
)

type Cmd struct {
	r io.Reader
	w io.Writer
}

func (cmd *Cmd) Run(args []string) error {
	if len(args) != 1 {
		err := errors.New("invalid arguments")
		fmt.Fprintln(cmd.w, err.Error())
		return err
	}

	url := args[0]
	opt := xget.Option{Procs: *procs}
	c, err := xget.NewClient(url, opt)
	if err != nil {
		fmt.Fprintln(cmd.w, err.Error())
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case s := <-sigCh:
				fmt.Fprintln(cmd.w, s)
				cancel()
			default:
			}
		}
	}()

	if err := c.Run(ctx); err != nil {
		fmt.Fprintln(cmd.w, err.Error())
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	cmd := Cmd{os.Stdin, os.Stdout}
	if err := cmd.Run(flag.Args()); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
