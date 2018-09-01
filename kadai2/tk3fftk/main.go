package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai2/tk3fftk/imageconverter"
)

var (
	dir  string
	from string
	to   string
)

func init() {
	flag.StringVar(&dir, "d", "", "target directory of conversion")
	flag.StringVar(&from, "f", "jpg", "extension of target file (jpg, jpeg, png, gif)")
	flag.StringVar(&to, "t", "png", "extension of after conversion (jpg, jpeg, png, gif)")
}

func exitWithUsage() {
	flag.Usage()
	os.Exit(2)
}

func main() {
	flag.Parse()
	if dir == "" {
		fmt.Println("dir is empty")
		exitWithUsage()
	}

	ic, err := imageconverter.New(from, to)
	if err != nil {
		fmt.Printf("%v", err)
		exitWithUsage()
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		return ic.ConvertImage(path, nil)
	})
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
