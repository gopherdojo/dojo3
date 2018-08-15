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

// CLI has streams and command line options.
type CLI struct {
	OutStream, ErrStream io.Writer
	in                   FileFormat
	out                  FileFormat
	force                bool
	verbose              bool
}

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

// Execute executes main processing.
func (c *CLI) Execute(dirname string) (ok bool) {
	ok = true

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if c.verbose {
				fmt.Fprintf(c.OutStream, "Skipped because the path is directory: %q\n", path)
			}
			return nil
		}

		fp, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()

		var isApplicable bool
		switch c.in {
		case Jpeg:
			isApplicable = myfileutil.IsJpeg(fp)
		case Png:
			isApplicable = myfileutil.IsPng(fp)
		case Gif:
			isApplicable = myfileutil.IsGif(fp)
		}
		if !isApplicable {
			if c.verbose {
				fmt.Fprintf(c.OutStream, "Skipped because the file is not applicable: %q\n", path)
			}
			return nil
		}

		var extname string
		switch c.out {
		case Jpeg:
			extname = "jpg"
		case Png:
			extname = "png"
		case Gif:
			extname = "gif"
		}

		dstName := myfileutil.DropExtname(path) + "." + extname

		if !c.force && myfileutil.Exists(dstName) {
			return errors.New("File already exists: " + dstName)
		}

		var img image.Image
		switch c.in {
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

		switch c.out {
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

		if c.verbose {
			fmt.Fprintf(c.OutStream, "Converted: %q\n", dstName)
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(c.ErrStream, err)
		ok = false
	}

	return
}

var JOpt bool
var POpt bool
var GOpt bool
var jOpt bool
var pOpt bool
var gOpt bool
var fOpt bool
var vOpt bool

func init() {
	flag.BoolVar(&JOpt, "J", false, "Convert from JPEG")
	flag.BoolVar(&POpt, "P", false, "Convert from PNG")
	flag.BoolVar(&GOpt, "G", false, "Convert from GIF")
	flag.BoolVar(&jOpt, "j", false, "Convert to JPEG")
	flag.BoolVar(&pOpt, "p", false, "Convert to PNG")
	flag.BoolVar(&gOpt, "g", false, "Convert to GIF")
	flag.BoolVar(&fOpt, "f", false, "Overwrite when the converted file name duplicates.")
	flag.BoolVar(&vOpt, "v", false, "Verbose Mode")
}

func main() {
	flag.Parse()

	var in FileFormat
	switch {
	case JOpt:
		in = Jpeg
	case POpt:
		in = Png
	case GOpt:
		in = Gif
	default:
		in = Jpeg
	}

	var out FileFormat
	switch {
	case jOpt:
		out = Jpeg
	case pOpt:
		out = Png
	case gOpt:
		out = Gif
	default:
		out = Png
	}

	if in == out {
		fmt.Println("You must specify a different file format before and after conversion.")
		os.Exit(1)
	}

	cli := &CLI{OutStream: os.Stdout, ErrStream: os.Stderr, in: in, out: out, force: fOpt, verbose: vOpt}

	dirnames := flag.Args()

	if len(dirnames) == 0 {
		fmt.Fprintln(cli.ErrStream, "Specify filenames as an arguments")
		os.Exit(1)
	}

	ok := cli.Execute(dirnames[0])

	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
