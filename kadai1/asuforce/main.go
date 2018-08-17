package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/asuforce/converter"
)

var path string

func init() {
	flag.Parse()
	path = flag.Arg(0)
}

func main() {
	err := converter.Convert(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
