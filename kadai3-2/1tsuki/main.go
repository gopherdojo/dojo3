package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo3/kadai3-2/1tsuki/downloader"
	"io"
	"net/url"
	"os"
	"time"
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
		parallel int
		args     []string
	)
	flags := flag.NewFlagSet("pget", flag.ContinueOnError)
	flags.IntVar(&parallel, "p", 6, "number of download pipelines")
	flags.Parse(strArgs)
	args = flags.Args()

	rawUrl := args[0]
	url, err := url.Parse(rawUrl)
	if err != nil {
		printf("error parsing url: %v\n", err)
		return exitCodeInvalidOption
	}

	d := downloader.NewDownloader(writer)
	if err := d.Download(url, parallel, 1*time.Minute); err != nil {
		printf("error downloading file: %v\n", err)
		return exitCodeError
	}

	printf("downloaded file\n")
	return exitCodeOK
}

func printf(format string, a ... interface{}) {
	fmt.Fprintf(writer, format, a...)
}
