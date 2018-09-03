package main

//
import (
	"flag"
	"fmt"
	"os"

	"github.com/dojo3/kadai2/gosagawa/converter"
)

//
func main() {

	var (
		inType  = flag.String("i", "jpeg", "変換対象の画像形式(jpeg|gif|png)")
		outType = flag.String("o", "png", "変換語の画像形式(jpeg|gif|png)")
		dispLog = flag.Bool("v", false, "詳細なログを表示")
	)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if err := IsValidInput(*inType, *outType, args); err != nil {
		usage()
		os.Exit(2)
	}

	dir := args[0]
	c := converter.NewConverter(*inType, *outType, dir, *dispLog)
	c.ConvertImage()
}

//
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] dir_path\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  -i string\n")
	fmt.Fprintf(os.Stderr, "    	変換対象の画像形式(jpeg|gif|png) (default \"jpeg\")\n")
	fmt.Fprintf(os.Stderr, "  -o string\n")
	fmt.Fprintf(os.Stderr, "    	変換語の画像形式(jpeg|gif|png) (default \"png\")\n")
	fmt.Fprintf(os.Stderr, "  -v	詳細なログを表示\n")
}

// Validate Input
func IsValidInput(inType string, outType string, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Invalid Args len \n")
	}
	if err := converter.IsValidImageType(inType); err != nil {
		return err
	}
	if err := converter.IsValidImageType(outType); err != nil {
		return err
	}
	return nil
}
