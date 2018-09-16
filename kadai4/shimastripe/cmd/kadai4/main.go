package main

import (
	"os"

	"github.com/gopherdojo/dojo3/kadai4/shimastripe"
)

func main() {
	cli := &shimastripe.CLI{}
	os.Exit(cli.Run(os.Args))
}
