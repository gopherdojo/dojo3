package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/gathering"
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

	cli := &CLI{OutStream: outStream, ErrStream: errStream}
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

	encoder := deriveEncoder()
	converter := &conversion.Converter{Decoder: decoder, Encoder: encoder, OutStream: outStream}

	for _, path := range paths {
		err = converter.Convert(path)
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

func deriveEncoder() conversion.Encoder {
	switch {
	case cliopt.ToJpeg:
		return &conversion.Jpeg{}
	case cliopt.ToPng:
		return &conversion.Png{}
	case cliopt.ToGif:
		return &conversion.Gif{}
	default:
		return &conversion.Jpeg{}
	}
}
