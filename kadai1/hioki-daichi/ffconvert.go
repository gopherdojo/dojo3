package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/hioki-daichi/myfileutil"
)

// FileFormat provides file formats like JPEG, PNG, GIF
type FileFormat int

const (
	// Jpeg is JPEG format
	Jpeg FileFormat = iota

	// Png is PNG format
	Png
)

func main() {
	JOpt := flag.Bool("J", false, "Convert from JPEG")
	POpt := flag.Bool("P", false, "Convert from PNG")
	jOpt := flag.Bool("j", false, "Convert to JPEG")
	pOpt := flag.Bool("p", false, "Convert to PNG")

	flag.Parse()

	var in FileFormat
	switch {
	case *JOpt:
		in = Jpeg
	case *POpt:
		in = Png
	default:
		in = Jpeg
	}

	var out FileFormat
	switch {
	case *jOpt:
		out = Jpeg
	case *pOpt:
		out = Png
	default:
		out = Png
	}

	ok := execute(flag.Args(), in, out)
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func execute(dirnames []string, in FileFormat, out FileFormat) (ok bool) {
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

			var isApplicable bool
			switch in {
			case Jpeg:
				isApplicable = myfileutil.IsJpeg(fp)
			case Png:
				isApplicable = myfileutil.IsPng(fp)
			}
			if !isApplicable {
				return nil
			}

			var extname string
			switch out {
			case Jpeg:
				extname = "jpg"
			case Png:
				extname = "png"
			}

			dstName := myfileutil.DropExtname(path) + "." + extname

			if myfileutil.Exists(dstName) {
				return errors.New("File already exists: " + dstName)
			}

			var img image.Image
			switch in {
			case Jpeg:
				img, err = jpeg.Decode(fp)
			case Png:
				img, err = png.Decode(fp)
			}
			if err != nil {
				return err
			}

			dstFile, err := os.Create(dstName)
			if err != nil {
				return err
			}

			switch out {
			case Jpeg:
				err = jpeg.Encode(dstFile, img, &jpeg.Options{Quality: 100})
			case Png:
				err = png.Encode(dstFile, img)
			}
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
