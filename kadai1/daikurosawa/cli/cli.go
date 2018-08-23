// Package cli is image convert cli tool.
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
	"golang.org/x/sync/errgroup"
)

// Exit code.
const (
	exitCodeOK = iota
	ExitCodeParseFlagError
	exitCodeProcessError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run command.
func (c *CLI) Run(args []string) int {

	var from, to string
	flags := flag.NewFlagSet("awesome-cli", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&from, "from", "jpg", "Input file extension.")
	flag.StringVar(&to, "to", "png", "Output file extension.")
	flag.Parse()

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	dirName := flags.Arg(0)

	if fileInfo, err := os.Stat(dirName); err != nil || fileInfo.IsDir() == false {
		fmt.Fprintln(os.Stderr, err.Error())
		return exitCodeProcessError
	}

	option := &option.Option{DirName: dirName, FromExtension: from, ToExtension: to}
	convert := convert.NewConvert(option)

	if err := walkDirectory(dirName, from, convert); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return exitCodeProcessError
	}

	return exitCodeOK
}

func walkDirectory(dirName string, fromExtension string, convert convert.Convert) error {
	eg := errgroup.Group{}

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if strings.TrimPrefix(filepath.Ext(path), ".") == fromExtension {
			eg.Go(func() error {
				return convert.Convert(path)
			})
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
