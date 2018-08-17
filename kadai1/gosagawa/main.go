package main

import (
	"flag"
	"fmt"
	"os"

	"./converter"
)

func main() {

	var (
		inType  = flag.String("i", "jpeg", "変換対象の画像形式(jpeg|gif|png)")
		outType = flag.String("o", "png", "変換語の画像形式(jpeg|gif|png)")
		dispLog = flag.Bool("v", false, "詳細なログを表示")
	)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if !isValidInput(*inType, *outType, args) {
		usage()
		os.Exit(2)
	}

	dir := args[0]
	c := converter.NewConverter(*inType, *outType, dir, *dispLog)
	c.ConvertImage()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] dir_path\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  -i string\n")
	fmt.Fprintf(os.Stderr, "    	変換対象の画像形式(jpeg|gif|png) (default \"jpeg\")\n")
	fmt.Fprintf(os.Stderr, "  -o string\n")
	fmt.Fprintf(os.Stderr, "    	変換語の画像形式(jpeg|gif|png) (default \"png\")\n")
	fmt.Fprintf(os.Stderr, "  -v	詳細なログを表示\n")
}

func isValidInput(inType string, outType string, args []string) bool {
	if len(args) != 1 {
		return false
	}
	if !isValidType(inType) {
		return false
	}
	if !isValidType(outType) {
		return false
	}
	return true
}

//XXX もっとスマートなやり方があるのでは
func isValidType(imageType string) bool {
	var isValidType bool
	switch imageType {
	case "jpeg":
		isValidType = true
	case "png":
		isValidType = true
	case "gif":
		isValidType = true
	default:
		isValidType = false
	}
	return isValidType
}
