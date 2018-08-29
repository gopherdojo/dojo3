package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gopherdojo/dojo3/kadai2/asuforce/converter"
)

var (
	version = "1.0.0"
	path    string
	fromExt string
	toExt   string
	wg      sync.WaitGroup
)

func init() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&fromExt, "f", "jpg", "Specify input image extension")
	flag.StringVar(&fromExt, "from", "jpg", "Specify input image extension")
	flag.StringVar(&toExt, "t", "png", "Specify output image extension")
	flag.StringVar(&toExt, "to", "png", "Specify output image extension")
	flag.Parse()

	if showVersion {
		fmt.Println("Version: ", version)
		os.Exit(0)
	}

	path = flag.Arg(0)
}

func main() {
	var c = converter.NewConverter(path, fromExt, toExt)

	err := filepath.Walk(c.Path, c.CrawlFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	queue := make(chan converter.Image)
	for _, image := range c.Files {
		wg.Add(1)
		go c.FetchConverter(queue, &wg)
		queue <- image
	}

	close(queue)
	wg.Wait()
}
