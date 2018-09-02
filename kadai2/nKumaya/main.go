package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dojo3/kadai2/nKumaya/format"
	"github.com/dojo3/kadai2/nKumaya/imgconv"
)

const (
	// ExitCodeOK 期待する終了コード
	ExitCodeOK = iota
	// ExitCodeParseFlagError フラグParse時のエラーコード
	ExitCodeParseFlagError
	// ExitCodeParseArgError 対象ディレクトリが指定されていない場合のエラーコード
	ExitCodeParseArgError
	// ExitCodeParseArgError 対象ディレクトリが指定されていない場合のエラーコード
	ExitCodeSearchError
	// ExitCodeFormatError サポートされているフォーマット以外の場合のエラーコード
	ExitCodeFormatError
	// ExitCodeConvertError 画像変換に失敗した場合のエラーコード
	ExitCodeConvertError
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Execute(args []string) int {
	var from, to string
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&from, "f", "jpg", "from")
	flags.StringVar(&to, "t", "png", "to")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if flags.NArg() == 0 {
		fmt.Fprintf(c.ErrStream, "not set directoryPath\n")
		return ExitCodeParseArgError
	}

	format, err := format.NewFormat(from, to)
	if err != nil {
		return ExitCodeFormatError
	}

	dir := flags.Arg(0)

	if err := c.convertFiles(dir, "."+from, "."+to, format); err != nil {
		fmt.Fprintf(c.ErrStream, "convert err\n")
		return ExitCodeConvertError
	}

	return ExitCodeOK
}

func (c *CLI) convertFiles(dir, from, to string, format *format.Format) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == from {
			newFile := path[:len(path)-len(filepath.Ext(path))] + to
			converter := imgconv.NewConverter(path, newFile, *format)
			if err := converter.Convert(); err != nil {
				return err
			}
			fmt.Fprintf(c.OutStream, "success from %s  to %s\n", path, newFile)
		}
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func main() {
	cli := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Execute(os.Args))
}
