package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cmd"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/opt"
)

func main() {
	dirname, options, err := opt.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	runner := &cmd.Runner{
		OutStream: os.Stdout,
		Decoder:   options.Decoder,
		Encoder:   options.Encoder,
		Force:     options.Force,
		Verbose:   options.Verbose,
	}
	err = runner.Run(dirname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
