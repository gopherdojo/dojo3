package main

import (
	"flag"
	"log"

	"github.com/gopherdojo/dojo3/kadai2/kzkick2nd/cli"
)

func main() {
	var dir string
	var inputExt string
	var outputExt string

	// FIXME validationどうやる？（表記揺れ、拡張子間違い、PATH間違い、複数Dir）

	flag.StringVar(&inputExt, "i", "jpg", "input format (jpg | png)")
	flag.StringVar(&outputExt, "o", "png", "output format (jpg | png)")
	flag.Parse()
	dir = flag.Arg(0)

	// FIXME 並列処理にできない？

	w := cli.Worker{
		Input:  inputExt,
	  Output: outputExt,
	}
	err := w.Run(dir)
	if err != nil {
		log.Fatal(err)
	}

}
