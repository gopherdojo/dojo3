package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/gopherdojo/dojo3/kadai2/pchatsu/cli"
)

var (
	srcDir = flag.String("d", "./", "target directory")
	srcExt = flag.String("from", "jpeg", "source extension")
	dstExt = flag.String("to", "png", "number lines")
)

func main() {
	flag.Parse()
	if err := cli.Run(*srcDir, *srcExt, *dstExt); err != nil {
		if err == cli.ErrSameFormat {
			os.Exit(0)
		} else {
			fmt.Fprintln(os.Stderr, "imgconv:", err.Error())
			os.Exit(1)
		}
	}
}

