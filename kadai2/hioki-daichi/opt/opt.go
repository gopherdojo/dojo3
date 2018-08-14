/*
Package opt is a package for parsing the command line option and building necessary information.
*/
package opt

import (
	"errors"
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/conversion"
)

// Options sets Decoder, Encoder and Force.
type Options struct {
	Decoder conversion.Decoder
	Encoder conversion.Encoder
	Force   bool
}

// Parse parses the command line option, validates it, constructs the necessary information for the later conversion process and return it.
func Parse(args ...string) (string, *Options, error) {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	fromJpeg := flg.Bool("J", false, "Convert from JPEG")
	fromPng := flg.Bool("P", false, "Convert from PNG")
	fromGif := flg.Bool("G", false, "Convert from GIF")
	toJpeg := flg.Bool("j", false, "Convert to JPEG")
	toPng := flg.Bool("p", false, "Convert to PNG")
	toGif := flg.Bool("g", false, "Convert to GIF")
	force := flg.Bool("f", false, "Overwrite when the converted file name duplicates.")
	quality := flg.Int("quality", 100, "JPEG Quality to be used with '-j' option. You can specify 1 to 100.")
	numColors := flg.Int("num-colors", 256, "Maximum number of colors used in the GIF image to be used with '-g' option. You can specify 1 to 256.")
	humanCompressionLevel := flg.String("compression-level", "default", "Options to specify the compression level of PNG to be used with '-p' option. You can specify from 'default', 'no', 'best-speed', 'best-compression'.")

	flg.Parse(args)

	if *toJpeg {
		if *quality < 1 {
			return "", nil, errors.New("--quality must be greater than or equal to 1")
		} else if *quality > 100 {
			return "", nil, errors.New("--quality must be less than or equal to 100")
		}
	}

	if *toGif {
		if *numColors < 1 {
			return "", nil, errors.New("--num-colors must be greater than or equal to 1")
		} else if *numColors > 256 {
			return "", nil, errors.New("--num-colors must be less than or equal to 256")
		}
	}

	if *toPng {
		switch *humanCompressionLevel {
		case "default", "no", "best-speed", "best-compression":
		default:
			return "", nil, errors.New("--compression-level is not included in the list: \"default\", \"no\", \"best-speed\", \"best-compression\"")
		}
	}

	dirnames := flg.Args()
	if len(dirnames) == 0 {
		return "", nil, errors.New("you must specify a directory")
	}

	options := &Options{
		Decoder: deriveDecoder(fromJpeg, fromPng, fromGif),
		Encoder: deriveEncoder(toJpeg, toPng, toGif, quality, numColors, humanCompressionLevel),
		Force:   *force,
	}

	return dirnames[0], options, nil
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
	case *toJpeg:
		return &conversion.Jpeg{Options: &jpeg.Options{Quality: *quality}}
	case *toGif:
		return &conversion.Gif{Options: &gif.Options{NumColors: *numColors}}
	case *toPng:
		fallthrough
	default:
		return &conversion.Png{Encoder: &png.Encoder{CompressionLevel: toCompressionLevel(humanCompressionLevel)}}
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
