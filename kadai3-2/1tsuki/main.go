package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo3/kadai3-2/1tsuki/downloader"
	"io"
	"net/url"
	"os"
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
		pararrel int
		// args     []string
	)
	flags := flag.NewFlagSet("pget", flag.ContinueOnError)
	flags.IntVar(&pararrel, "p", 6, "number of download pipelines")
	flags.Parse(strArgs)
	// args = flags.Args()

	// rawUrl := args[0]
	rawUrl := "https://www.recruit.co.jp/index"
	url, err := url.Parse(rawUrl)
	if err != nil {
		printf("error parsing url: %v", err)
		return exitCodeInvalidOption
	}

	filepath, err := downloader.Download(url)
	if err != nil {
		printf("error downloading file: %v", err)
		return exitCodeError
	}

	printf("downloaded file: %s", filepath)

	return exitCodeOK
}

func printf(format string, a ... interface{}) {
	fmt.Fprintf(writer, format, a...)
}
