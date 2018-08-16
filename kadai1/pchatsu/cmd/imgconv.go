package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/pchatsu"
)

var (
	validFormat = map[string]struct{}{"gif": {}, "jpeg": {}, "png": {}}
)

var (
	srcDir = flag.String("d", "./", "target directory")
	srcExt = flag.String("from", "jpeg", "source extension")
	dstExt = flag.String("to", "png", "number lines")
)

func main() {
	flag.Parse()
	if err := Run(*srcDir, *srcExt, *dstExt); err != nil {
		fmt.Fprintln(os.Stderr, "imgconv:", err.Error())
		os.Exit(1)
	}
}

func Run(path string, src string, dst string) error {
	if err := validate(path, src, dst); err != nil {
		return err
	}
	if err := imgconv.Convert(path, src, dst); err != nil {
		return err
	}

	return nil
}

func validate(path string, srcExt string, dstExt string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("%s no such directory", path)
	}

	if !isValidFormatType(srcExt) || !isValidFormatType(dstExt) {
		return errors.New("available formats are gif, jpeg and png")
	}

	if srcExt == dstExt {
		return errors.New("can't convert to the same format")
	}

	return nil
}

func isValidFormatType(f string) bool {
	_, ok := validFormat[f]
	return ok
}
