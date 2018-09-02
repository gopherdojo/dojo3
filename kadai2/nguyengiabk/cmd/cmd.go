// Package cmd provides functions and type to parse ther arguments and run the command
package cmd

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/gopherdojo/dojo3/kadai2/nguyengiabk/converter"
)

// Parameter holds options to run command
type Parameter struct {
	InputType    string
	OutputType   string
	JpgQuality   int
	GifNumColors int
	Path         []string
}

const jpgType = "jpg"
const jpegType = "jpeg"
const pngType = "png"
const gifType = "gif"
const minJpgQuaility = 1
const maxJpgQuality = 100
const minGifNumColors = 1
const maxGifNumColors = 256

func isValidType(t string) bool {
	_, ok := converter.SupportedTypes[t]
	return ok
}

// Parse parses the commandline arguments to Parameter
func Parse(args []string) (*Parameter, error) {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	inputType := flg.String("i", jpgType, "Input file type")
	outputType := flg.String("o", pngType, "Output file type")
	jpgQuality := flg.Int("q", maxJpgQuality, "JPG Quality, ranges from 1 to 100, (only used for encoding jpg)")
	gifNumColors := flg.Int("n", maxGifNumColors, "Maximum number of colors, ranges from 1 to 256, (only used for encoding gif)")
	flg.Parse(args)

	if !isValidType(*inputType) {
		return nil, errors.New("Invalid input type")
	}
	if !isValidType(*outputType) {
		return nil, errors.New("Invalid output type")
	}
	if *jpgQuality < minJpgQuaility || *jpgQuality > maxJpgQuality {
		return nil, errors.New("Invalid JPG quality")
	}
	if *gifNumColors < minGifNumColors || *gifNumColors > maxGifNumColors {
		return nil, errors.New("Invalid maximum number of colors for GIF encoding")
	}
	if len(flg.Args()) <= 0 {
		return nil, errors.New("No input specified")
	}
	return &Parameter{
		InputType:    *inputType,
		OutputType:   *outputType,
		JpgQuality:   *jpgQuality,
		GifNumColors: *gifNumColors,
		Path:         flg.Args(),
	}, nil
}

// Run runs the converter based on the parameters that passed in
func Run(params Parameter, w io.Writer) error {
	var decoder converter.Decoder
	var encoder converter.Encoder

	switch params.InputType {
	case jpgType:
		fallthrough
	case jpegType:
		decoder = &converter.JPEG{}
	case pngType:
		decoder = &converter.PNG{}
	case gifType:
		decoder = &converter.GIF{}
	}
	switch params.OutputType {
	case jpgType:
		fallthrough
	case jpegType:
		encoder = &converter.JPEG{Quality: params.JpgQuality}
	case pngType:
		encoder = &converter.PNG{}
	case gifType:
		encoder = &converter.GIF{NumColors: params.GifNumColors}
	}

	converter := converter.Converter{Decoder: decoder, Encoder: encoder}
	for _, path := range params.Path {
		if err := converter.Run(path, w); err != nil {
			return err
		}
	}
	return nil
}
