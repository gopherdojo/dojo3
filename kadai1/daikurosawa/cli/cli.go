// Package cli is image convert cli tool.
package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
	"golang.org/x/sync/errgroup"
)

// Exit code.
const (
	ExitCodeOK = iota
	ExitCodeProcessError
)

// Cil is interface that has Run function.
type Cil interface {
	Run() int
}

type cli struct {
	convert convert.Convert
	option  *option.Option
}

// NewCli is Cli interface constructor.
func NewCli(convert convert.Convert, option *option.Option) Cil {
	return &cli{convert: convert, option: option}
}

// Run command.
func (c *cli) Run() int {

	if fileInfo, err := os.Stat(c.option.DirName); err != nil || fileInfo.IsDir() == false {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeProcessError
	}

	if err := walkDirectory(c.option.DirName, c.option.FromExtension, c.convert); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeProcessError
	}

	return ExitCodeOK
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
