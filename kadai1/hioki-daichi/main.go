package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/gathering"
)

const usage = `USAGE: ffconvert [-JPGjpgfv] [dirname]

-J
    Input file format is JPEG (default)
-P
    Input file format is PNG
-G
    Input file format is GIF
-j
    Output file format is JPEG
-p
    Output file format is PNG (default)
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
	Decoder              conversion.Decoder
	Encoder              conversion.Encoder
	Force                bool
	Verbose              bool
}

func main() {
	fromJpeg := flag.Bool("J", false, "Convert from JPEG")
	fromPng := flag.Bool("P", false, "Convert from PNG")
	fromGif := flag.Bool("G", false, "Convert from GIF")
	toJpeg := flag.Bool("j", false, "Convert to JPEG")
	toPng := flag.Bool("p", false, "Convert to PNG")
	toGif := flag.Bool("g", false, "Convert to GIF")
	force := flag.Bool("f", false, "Overwrite when the converted file name duplicates.")
	verbose := flag.Bool("v", false, "Verbose Mode")

	outStream := os.Stdout
	errStream := os.Stderr

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(outStream, usage)
		os.Exit(0)
	}

	decoder := deriveDecoder(fromJpeg, fromPng, fromGif)
	encoder := deriveEncoder(toJpeg, toPng, toGif)

	cli := &CLI{OutStream: outStream, ErrStream: errStream, Decoder: decoder, Encoder: encoder, Force: *force, Verbose: *verbose}
	err := cli.execute(args[0])
	if err != nil {
		fmt.Fprintln(cli.ErrStream, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func (c *CLI) execute(dirname string) error {
	gatherer := &gathering.Gatherer{Decoder: c.Decoder, OutStream: c.OutStream}
	paths, err := gatherer.Gather(dirname)
	if err != nil {
		return err
	}

	converter := &conversion.Converter{Decoder: c.Decoder, Encoder: c.Encoder, OutStream: c.OutStream}

	for _, path := range paths {
		fp, err := converter.Convert(path, c.Force)
		if err != nil {
			return err
		}

		if c.Verbose {
			fmt.Fprintf(c.OutStream, "Converted: %q\n", fp.Name())
		}
	}

	return nil
}

func deriveDecoder(fromJpeg *bool, fromPng *bool, fromGif *bool) conversion.Decoder {
	switch {
	case *fromJpeg:
		return &conversion.Jpeg{}
	case *fromPng:
		return &conversion.Png{}
	case *fromGif:
		return &conversion.Gif{}
	default:
		return &conversion.Jpeg{}
	}
}

func deriveEncoder(toJpeg *bool, toPng *bool, toGif *bool) conversion.Encoder {
	switch {
	case *toJpeg:
		return &conversion.Jpeg{}
	case *toPng:
		return &conversion.Png{}
	case *toGif:
		return &conversion.Gif{}
	default:
		return &conversion.Jpeg{}
	}
}
