package main

import (
	"log"
	"os"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/cmd"
)

func main() {
	params, err := cmd.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Run(*params, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
