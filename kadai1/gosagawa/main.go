package main

import (
	"flag"
	"fmt"
	_ "image/gif"
)

func main() {

	var (
		inputImageType  = flag.String("i", "jpg", "変換対象の画像形式(jpg|gif|png)")
		outputImageType = flag.String("o", "png", "変換語の画像形式(jpg|gif|png)")
	)
	flag.Parse()
	args := flag.Args()

	fmt.Println(*inputImageType)
	fmt.Println(*outputImageType)
	fmt.Println(args[0])
}
