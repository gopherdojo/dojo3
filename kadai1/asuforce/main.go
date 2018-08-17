package main

import (
	"flag"

	"github.com/gopherdojo/dojo3/kadai1/asuforce/converter"
)

var path string

func init() {
	flag.Parse()
	path = flag.Arg(0)
}

func main() {
	converter.Convert(path)
}
