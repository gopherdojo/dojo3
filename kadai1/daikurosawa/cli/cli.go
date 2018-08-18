package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"strings"

	"../convert"
	"golang.org/x/sync/errgroup"
)

// Exit code.
const (
	ExitCodeOK = iota
	ExitCodeProcessError
)

// Cli has Stdout and Stderr streams.
type Cli struct {
	OutStream, ErrStream io.Writer
}

// Option has command line options.
type Option struct {
	FromExtension string
	ToExtension   string
}

// Run command.
func (c *Cli) Run(args []string) int {
	var (
		from = flag.String("from", "jpg", "Input file extension.")
		to   = flag.String("to", "png", "Output file extension.")
	)
	flag.Parse()

	dirName := flag.Arg(0)
	if fileInfo, err := os.Stat(dirName); err != nil || fileInfo.IsDir() == false {
		fmt.Fprintf(os.Stderr, err.Error())
		return ExitCodeProcessError
	}

	option := &Option{FromExtension: *from, ToExtension: *to}

	if err := option.walkDirectory(dirName); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return ExitCodeProcessError
	}

	return ExitCodeOK
}

func (o *Option) walkDirectory(dirName string) error {
	eg := errgroup.Group{}

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if strings.TrimPrefix(filepath.Ext(path), ".") == o.FromExtension {
			eg.Go(func() error {
				convert := &convert.Convert{Path: path,
					FromExtension: o.FromExtension,
					ToExtension:   o.ToExtension}
				return convert.Convert()
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
