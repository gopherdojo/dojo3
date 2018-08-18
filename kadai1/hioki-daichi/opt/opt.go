package opt

import (
	"errors"
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/conversion"
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
`

// Options has command line options.
type Options struct {
	Decoder conversion.Decoder
	Encoder conversion.Encoder
	Force   bool
}

// Parse parses and returns command line options.
func Parse() (string, *Options, error) {
	fromJpeg := flag.Bool("J", false, "Convert from JPEG")
	fromPng := flag.Bool("P", false, "Convert from PNG")
	fromGif := flag.Bool("G", false, "Convert from GIF")

	toJpeg := flag.Bool("j", false, "Convert to JPEG")
	toPng := flag.Bool("p", false, "Convert to PNG")
	toGif := flag.Bool("g", false, "Convert to GIF")

	force := flag.Bool("f", false, "Overwrite when the converted file name duplicates.")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return "", nil, errors.New(usage)
	}

	options := &Options{
		Decoder: deriveDecoder(fromJpeg, fromPng, fromGif),
		Encoder: deriveEncoder(toJpeg, toPng, toGif),
		Force:   *force,
	}

	return args[0], options, nil
}

func deriveDecoder(fromJpeg *bool, fromPng *bool, fromGif *bool) conversion.Decoder {
	switch {
	case *fromPng:
		return &conversion.Png{}
	case *fromGif:
		return &conversion.Gif{}
	case *fromJpeg:
		fallthrough
	default:
		return &conversion.Jpeg{}
	}
}

func deriveEncoder(toJpeg *bool, toPng *bool, toGif *bool) conversion.Encoder {
	switch {
	case *toPng:
		return &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}
	case *toGif:
		return &conversion.Gif{Options: &gif.Options{NumColors: 256}}
	case *toJpeg:
		fallthrough
	default:
		return &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}
	}
}
