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

	// TODO: validate value from 1 to 100
	quality := flag.Int("quality", 100, "JPEG Quality to be used with '-j' option")

	// TODO: validate value from 1 to 256
	numColors := flag.Int("num-colors", 256, "Maximum number of colors used in the image to be used with '-g' option")

	// TODO: validate inclusion in "default", "no", "best-speed", "best-compression"
	humanCompressionLevel := flag.String("compression-level", "default", "(selected from 'default', 'no', 'best-speed', 'best-compression') Options to specify the compression level of PNG to be used with '-p' option")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return "", nil, errors.New(usage)
	}

	options := &Options{
		Decoder: deriveDecoder(fromJpeg, fromPng, fromGif),
		Encoder: deriveEncoder(toJpeg, toPng, toGif, quality, numColors, humanCompressionLevel),
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

func deriveEncoder(toJpeg *bool, toPng *bool, toGif *bool, quality *int, numColors *int, humanCompressionLevel *string) conversion.Encoder {
	switch {
	case *toPng:
		return &conversion.Png{Encoder: &png.Encoder{CompressionLevel: toCompressionLevel(humanCompressionLevel)}}
	case *toGif:
		return &conversion.Gif{Options: &gif.Options{NumColors: *numColors}}
	case *toJpeg:
		fallthrough
	default:
		return &conversion.Jpeg{Options: &jpeg.Options{Quality: *quality}}
	}
}

func toCompressionLevel(humanCompressionLevel *string) png.CompressionLevel {
	switch *humanCompressionLevel {
	case "no":
		return png.NoCompression
	case "best-speed":
		return png.BestSpeed
	case "best-compression":
		return png.BestCompression
	case "default":
		fallthrough
	default:
		return png.DefaultCompression
	}
}
