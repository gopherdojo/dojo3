package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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

	// Gif is GIF format
	Gif
)

func main() {
	JOpt := flag.Bool("J", false, "Convert from JPEG")
	POpt := flag.Bool("P", false, "Convert from PNG")
	GOpt := flag.Bool("G", false, "Convert from GIF")
	jOpt := flag.Bool("j", false, "Convert to JPEG")
	pOpt := flag.Bool("p", false, "Convert to PNG")
	gOpt := flag.Bool("g", false, "Convert to GIF")

	fOpt := flag.Bool("f", false, "Overwrite when the converted file name duplicates.")
	vOpt := flag.Bool("v", false, "Verbose Mode")

	flag.Parse()

	var in FileFormat
	switch {
	case *JOpt:
		in = Jpeg
	case *POpt:
		in = Png
	case *GOpt:
		in = Gif
	default:
		in = Jpeg
	}

	var out FileFormat
	switch {
	case *jOpt:
		out = Jpeg
	case *pOpt:
		out = Png
	case *gOpt:
		out = Gif
	default:
		out = Png
	}

	if in == out {
		fmt.Println("You must specify a different file format before and after conversion.")
		os.Exit(1)
	}

	ok := execute(os.Stdout, flag.Args(), in, out, *fOpt, *vOpt)
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func execute(w io.Writer, dirnames []string, in FileFormat, out FileFormat, force bool, verbose bool) (ok bool) {
	ok = true
	if len(dirnames) == 0 {
		fmt.Fprintln(w, "Specify filenames as an arguments")
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
				if verbose {
					fmt.Fprintf(w, "Skipped because the path is directory: %q\n", path)
				}
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
			case Gif:
				isApplicable = myfileutil.IsGif(fp)
			}
			if !isApplicable {
				if verbose {
					fmt.Fprintf(w, "Skipped because the file is not applicable: %q\n", path)
				}
				return nil
			}

			var extname string
			switch out {
			case Jpeg:
				extname = "jpg"
			case Png:
				extname = "png"
			case Gif:
				extname = "gif"
			}

			dstName := myfileutil.DropExtname(path) + "." + extname

			if !force && myfileutil.Exists(dstName) {
				return errors.New("File already exists: " + dstName)
			}

			var img image.Image
			switch in {
			case Jpeg:
				img, err = jpeg.Decode(fp)
			case Png:
				img, err = png.Decode(fp)
			case Gif:
				img, err = gif.Decode(fp)
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
			case Gif:
				err = gif.Encode(dstFile, img, &gif.Options{NumColors: 256})
			}
			if err != nil {
				return err
			}

			if verbose {
				fmt.Fprintf(w, "Converted: %q\n", dstName)
			}

			return nil
		})

		if err != nil {
			fmt.Fprintln(w, err)
			ok = false
			break DIRNAMES_LOOP
		}
	}
	return
}
