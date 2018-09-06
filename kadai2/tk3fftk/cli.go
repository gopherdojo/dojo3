package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo3/kadai2/tk3fftk/imageconverter"
	"io"
	"os"
	"path/filepath"
)

const (
	ExitCodeOK = iota
	ExitCodeErr
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) exitWithUsage(flags flag.FlagSet) int {
	flags.Usage()
	return ExitCodeErr
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("imageConv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	dir := flags.String("d", "", "target directory of conversion")
	from := flags.String("f", "jpg", "extension of target file (jpg, jpeg, png, gif)")
	to := flags.String("t", "png", "extension of after conversion (jpg, jpeg, png, gif)")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeErr
	}

	if *dir == "" {
		fmt.Println("dir is empty")
		return c.exitWithUsage(*flags)
	}

	ic, err := imageconverter.New(*from, *to)
	if err != nil {
		fmt.Printf("%v", err)
		return c.exitWithUsage(*flags)
	}

	err = filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		return ic.ConvertImage(path, nil)
	})
	if err != nil {
		fmt.Errorf("%v", err)
		return ExitCodeErr
	}

	return ExitCodeOK
}
