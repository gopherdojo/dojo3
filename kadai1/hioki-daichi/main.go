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

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/gathering"
	"github.com/hioki-daichi/myfileutil"
)

const usage = `USAGE: ffconvert [-JPGjpgfv] [dirname]

-J
    Input file format is JPEG
-P
    Input file format is PNG
-G
    Input file format is GIF
-j
    Output file format is JPEG
-p
    Output file format is PNG
-g
    Output file format is GIF
-f
    Overwrite when the converted file name duplicates.
-v
    Verbose Mode
`

// CLI has streams and command line options.
type CLI struct {
	OutStream, ErrStream io.Writer
	in                   FileFormat
	out                  FileFormat
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

func (c *CLI) convert(path string) error {
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

	if !cliopt.Force && myfileutil.Exists(dstName) {
		return errors.New("File already exists: " + dstName)
	}

	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

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

	if cliopt.Verbose {
		fmt.Fprintf(c.OutStream, "Converted: %q\n", dstName)
	}

	return nil
}

func init() {
	flag.BoolVar(&cliopt.FromJpeg, "J", false, "Convert from JPEG")
	flag.BoolVar(&cliopt.FromPng, "P", false, "Convert from PNG")
	flag.BoolVar(&cliopt.FromGif, "G", false, "Convert from GIF")
	flag.BoolVar(&cliopt.ToJpeg, "j", false, "Convert to JPEG")
	flag.BoolVar(&cliopt.ToPng, "p", false, "Convert to PNG")
	flag.BoolVar(&cliopt.ToGif, "g", false, "Convert to GIF")
	flag.BoolVar(&cliopt.Force, "f", false, "Overwrite when the converted file name duplicates.")
	flag.BoolVar(&cliopt.Verbose, "v", false, "Verbose Mode")
}

func main() {
	outStream := os.Stdout
	errStream := os.Stderr

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(outStream, usage)
		os.Exit(0)
	}

	cli := &CLI{OutStream: outStream, ErrStream: errStream, in: inputFileFormat(), out: outputFileFormat()}
	err := cli.execute(args[0])
	if err != nil {
		fmt.Fprintln(cli.ErrStream, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func (c *CLI) execute(dirname string) error {
	outStream := c.OutStream
	decoder := deriveDecoder()

	gatherer := &gathering.Gatherer{Decoder: decoder, OutStream: outStream}
	paths, err := gatherer.Gather(dirname)
	if err != nil {
		return err
	}

	for _, path := range paths {
		err = c.convert(path)

		if err != nil {
			return err
		}
	}

	return nil
}

func deriveDecoder() conversion.Decoder {
	switch {
	case cliopt.FromJpeg:
		return &conversion.Jpeg{}
	case cliopt.FromPng:
		return &conversion.Png{}
	case cliopt.FromGif:
		return &conversion.Gif{}
	default:
		return &conversion.Jpeg{}
	}
}

func inputFileFormat() FileFormat {
	switch {
	case cliopt.FromJpeg:
		return Jpeg
	case cliopt.FromPng:
		return Png
	case cliopt.FromGif:
		return Gif
	default:
		return Jpeg
	}
}

func outputFileFormat() FileFormat {
	switch {
	case cliopt.ToJpeg:
		return Jpeg
	case cliopt.ToPng:
		return Png
	case cliopt.ToGif:
		return Gif
	default:
		return Png
	}
}
