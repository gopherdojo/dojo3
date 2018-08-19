package main

import (
	"./convert"
	"flag"
	"log"
	"os"
	"path/filepath"
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

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}
		if filepath.Ext(path) != "."+inputExt {
			return nil
		}
		converter := convert.Converter{
			Src:       path,
			OutputExt: outputExt,
		}
		errConvert := converter.Convert()
		if errConvert != nil {
			log.Fatal(errConvert)
			return errConvert
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
