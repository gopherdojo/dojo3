// Command to convert several images extension.

package main

import (
	"os"
	"fmt"
	"io"
	"flag"
	"path/filepath"
	"github.com/gopherdojo/dojo3/kadai1/1tsuki/sorcery"
)

var writer io.Writer

const (
	exitCodeOK = iota
	exitCodeInvalidOption
	exitCodeError
)

func init() {
	writer = os.Stdout
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(strArgs []string) int {
	var (
		fromStr string
		toStr   string
		dirs    []string
	)

	flags := flag.NewFlagSet("Sorcery", flag.ContinueOnError)
	flags.StringVar(&fromStr, "from", "jpg", "source image extension.")
	flags.StringVar(&toStr, "to", "png", "target image extension.")
	flags.Parse(strArgs)
	dirs = flags.Args()

	from, err := sorcery.ImgExt(fromStr)
	if err != nil {
		fmt.Fprintf(writer, "%v", err)
		return exitCodeInvalidOption
	}

	to, err := sorcery.ImgExt(toStr)
	if err != nil {
		fmt.Fprintf(writer, "%v", err)
		return exitCodeInvalidOption
	}

	if len(dirs) <= 0 {
		currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Fprintf(writer, "%v", err)
			return exitCodeError
		}
		dirs = []string{currentDir}
	}

	s := sorcery.Sorcery(writer)
	for _, dir := range dirs {
		err = s.Exec(from, to, dir)
		if err != nil {
			fmt.Fprintf(writer, "%v", err)
			return exitCodeError
		}
	}

	return exitCodeOK
}
