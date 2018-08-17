package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/asuforce/converter"
)

var (
	version string
	path    string
)

func init() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&showVersion, "version", false, "Show version")
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
	err := filepath.Walk(c.Path, c.CrawlFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, image := range c.Files {
		err := converter.Convert(image)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
