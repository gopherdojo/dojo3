package main

import (
	"log"
	"os"

	"github.com/gopherdojo/dojo3/kadai3-2/nguyengiabk/gget"
	"github.com/gopherdojo/dojo3/kadai3-2/nguyengiabk/opt"
)

func main() {
	params, err := opt.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	g, err := gget.NewGGet(params)
	if err != nil {
		log.Fatal(err)
	}
	if err := g.Process(); err != nil {
		log.Fatal(err)
	}
}
