package main

import (
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/hioki-daichi/myfileutil"
)

func main() {
	flag.Parse()
	ok := execute(flag.Args())
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func execute(dirnames []string) (ok bool) {
	ok = true
	if len(dirnames) == 0 {
		fmt.Println("Specify filenames as an arguments")
		ok = false
		return
	}

DIRNAMES_LOOP:
	for _, dirname := range dirnames {
		err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fp.Close()

			if !myfileutil.IsJpeg(fp) {
				return nil
			}

			extname := "png"

			dstName := myfileutil.DropExtname(path) + "." + extname

			if myfileutil.Exists(dstName) {
				return errors.New("File already exists: " + dstName)
			}

			img, err := jpeg.Decode(fp)
			if err != nil {
				return err
			}

			dstFile, err := os.Create(dstName)
			if err != nil {
				return err
			}

			err = png.Encode(dstFile, img)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
			ok = false
			break DIRNAMES_LOOP
		}
	}
	return
}
