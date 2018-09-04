package opt

import (
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai2/hioki-daichi/conversion"
)

// Decoder
var jpegD conversion.Jpeg
var pngD conversion.Png
var gifD conversion.Gif

// Encoder
var jpegE = &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}
var pngE = &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}
var gifE = &conversion.Gif{Options: &gif.Options{NumColors: 256}}

func TestOpt_Parse(t *testing.T) {
	cases := map[string]struct {
		args    []string
		dirname string
		options *Options
		err     error
	}{
		"no argument": {args: []string{}, dirname: "", options: nil, err: errors.New("you must specify a directory")},

		"dirname only": {args: []string{"./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: false}, err: nil},

		"with -f option": {args: []string{"-f", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: true}, err: nil},

		// by format
		"JPEG to PNG": {args: []string{"-J", "-p", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: false}, err: nil},
		"JPEG to GIF": {args: []string{"-J", "-g", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: gifE, Force: false}, err: nil},
		"PNG to JPEG": {args: []string{"-P", "-j", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: jpegE, Force: false}, err: nil},
		"PNG to GIF":  {args: []string{"-P", "-g", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: gifE, Force: false}, err: nil},
		"GIF to JPEG": {args: []string{"-G", "-j", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &gifD, Encoder: jpegE, Force: false}, err: nil},
		"GIF to PNG":  {args: []string{"-G", "-p", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &gifD, Encoder: pngE, Force: false}, err: nil},

		// quality option
		"--quality=0":   {args: []string{"-P", "-j", "--quality=0", "./testdata/"}, dirname: "", options: nil, err: errors.New("--quality must be greater than or equal to 1")},
		"--quality=1":   {args: []string{"-P", "-j", "--quality=1", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 1}}, Force: false}, err: nil},
		"--quality=100": {args: []string{"-P", "-j", "--quality=100", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: false}, err: nil},
		"--quality=101": {args: []string{"-P", "-j", "--quality=101", "./testdata/"}, dirname: "", options: nil, err: errors.New("--quality must be less than or equal to 100")},

		// num-colors option
		"--num-colors=0":   {args: []string{"-J", "-g", "--num-colors=0", "./testdata/"}, dirname: "", options: nil, err: errors.New("--num-colors must be greater than or equal to 1")},
		"--num-colors=1":   {args: []string{"-J", "-g", "--num-colors=1", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 1}}, Force: false}, err: nil},
		"--num-colors=256": {args: []string{"-J", "-g", "--num-colors=256", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: false}, err: nil},
		"--num-colors=257": {args: []string{"-J", "-g", "--num-colors=257", "./testdata/"}, dirname: "", options: nil, err: errors.New("--num-colors must be less than or equal to 256")},

		// compression-level option
		"--compression-level=default":          {args: []string{"-J", "-p", "--compression-level=default", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: false}, err: nil},
		"--compression-level=no":               {args: []string{"-J", "-p", "--compression-level=no", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}, Force: false}, err: nil},
		"--compression-level=best-speed":       {args: []string{"-J", "-p", "--compression-level=best-speed", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.BestSpeed}}, Force: false}, err: nil},
		"--compression-level=best-compression": {args: []string{"-J", "-p", "--compression-level=best-compression", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.BestCompression}}, Force: false}, err: nil},
		"--compression-level=foo":              {args: []string{"-J", "-p", "--compression-level=foo", "./testdata/"}, dirname: "", options: nil, err: errors.New("--compression-level is not included in the list: \"default\", \"no\", \"best-speed\", \"best-compression\"")},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			dirname, options, err := Parse(c.args...)

			if c.err == nil { // If it is expected that no error will occur
				if err != nil {
					t.Fatalf("err %s", err)
				}
			} else {
				expected := c.err.Error()
				actual := err.Error()
				if actual != expected {
					t.Errorf(`expected="%s" actual="%s"`, expected, actual)
				}
			}

			if dirname != c.dirname {
				t.Errorf(`expected="%s" actual="%s"`, c.dirname, dirname)
			}

			if options != c.options {
				if !reflect.DeepEqual(options.Decoder, c.options.Decoder) {
					t.FailNow()
				}

				if !reflect.DeepEqual(options.Encoder, c.options.Encoder) {
					t.FailNow()
				}

				if options.Force != c.options.Force {
					t.FailNow()
				}
			}
		})
	}
}
