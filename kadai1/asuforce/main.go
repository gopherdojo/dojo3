package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/asuforce/converter"
)

var (
	version = "1.0.0"
	path    string
	fromExt string
	toExt   string
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
	var c converter.Converter

	c.Path = path
	c.FromExt = fromExt
	c.ToExt = toExt
	err := filepath.Walk(c.Path, c.CrawlFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, image := range c.Files {
		err := c.Convert(image)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
