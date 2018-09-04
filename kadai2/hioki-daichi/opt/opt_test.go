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
	cases := []struct {
		args    []string
		dirname string
		options *Options
		err     error
	}{
		// when there is no argument
		{args: []string{}, dirname: "", options: nil, err: errors.New("you must specify a directory")},

		// when only the directory name is specified
		{args: []string{"./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: false}, err: nil},

		// when -f option is specified
		{args: []string{"-f", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: true}, err: nil},

		// by format
		{args: []string{"-J", "-p", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: pngE, Force: false}, err: nil},
		{args: []string{"-J", "-g", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: gifE, Force: false}, err: nil},
		{args: []string{"-P", "-j", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: jpegE, Force: false}, err: nil},
		{args: []string{"-P", "-g", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: gifE, Force: false}, err: nil},
		{args: []string{"-G", "-j", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &gifD, Encoder: jpegE, Force: false}, err: nil},
		{args: []string{"-G", "-p", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &gifD, Encoder: pngE, Force: false}, err: nil},

		// quality option
		{args: []string{"-P", "-j", "--quality=0", "./testdata/"}, dirname: "", options: nil, err: errors.New("--quality must be greater than or equal to 1")},
		{args: []string{"-P", "-j", "--quality=1", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 1}}, Force: false}, err: nil},
		{args: []string{"-P", "-j", "--quality=100", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &pngD, Encoder: &conversion.Jpeg{Options: &jpeg.Options{Quality: 100}}, Force: false}, err: nil},
		{args: []string{"-P", "-j", "--quality=101", "./testdata/"}, dirname: "", options: nil, err: errors.New("--quality must be less than or equal to 100")},

		// num-colors option
		{args: []string{"-J", "-g", "--num-colors=0", "./testdata/"}, dirname: "", options: nil, err: errors.New("--num-colors must be greater than or equal to 1")},
		{args: []string{"-J", "-g", "--num-colors=1", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 1}}, Force: false}, err: nil},
		{args: []string{"-J", "-g", "--num-colors=256", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Gif{Options: &gif.Options{NumColors: 256}}, Force: false}, err: nil},
		{args: []string{"-J", "-g", "--num-colors=257", "./testdata/"}, dirname: "", options: nil, err: errors.New("--num-colors must be less than or equal to 256")},

		// compression-level option
		{args: []string{"-J", "-p", "--compression-level=default", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}, Force: false}, err: nil},
		{args: []string{"-J", "-p", "--compression-level=no", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}, Force: false}, err: nil},
		{args: []string{"-J", "-p", "--compression-level=best-speed", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.BestSpeed}}, Force: false}, err: nil},
		{args: []string{"-J", "-p", "--compression-level=best-compression", "./testdata/"}, dirname: "./testdata/", options: &Options{Decoder: &jpegD, Encoder: &conversion.Png{Encoder: &png.Encoder{CompressionLevel: png.BestCompression}}, Force: false}, err: nil},
		{args: []string{"-J", "-p", "--compression-level=foo", "./testdata/"}, dirname: "", options: nil, err: errors.New("--compression-level is not included in the list: \"default\", \"no\", \"best-speed\", \"best-compression\"")},
	}

	for _, c := range cases {
		c := c
		t.Run("", func(t *testing.T) {
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
