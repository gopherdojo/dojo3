package main

import (
	"errors"
	"flag"
	"log"

	"github.com/gopherdojo/dojo3/kadai1/nguyengiabk/converter"
)

// Parameter holds options to run command
type Parameter struct {
	inputType    string
	outputType   string
	jpgQuality   int
	gifNumColors int
	path         []string
}

const jpgType = "jpg"
const jpegType = "jpeg"
const pngType = "png"
const gifType = "gif"
const minJpgQuaility = 1
const maxJpgQuality = 100
const minGifNumColors = 1
const maxGifNumColors = 256

var supportedTypes = []string{jpgType, jpegType, pngType, gifType}

func isValidType(t string) bool {
	for _, val := range supportedTypes {
		if t == val {
			return true
		}
	}
	return false
}

func parse() (*Parameter, error) {
	inputType := flag.String("i", jpgType, "Input file type")
	outputType := flag.String("o", pngType, "Output file type")
	jpgQuality := flag.Int("q", maxJpgQuality, "JPG Quality, ranges from 1 to 100, (only used for encoding jpg)")
	gifNumColors := flag.Int("n", maxGifNumColors, "Maximum number of colors, ranges from 1 to 256, (only used for encoding gif)")
	flag.Parse()

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
	if len(flag.Args()) <= 0 {
		return nil, errors.New("No input specified")
	}
	return &Parameter{
		inputType:    *inputType,
		outputType:   *outputType,
		jpgQuality:   *jpgQuality,
		gifNumColors: *gifNumColors,
		path:         flag.Args(),
	}, nil
}

func run(params Parameter) error {
	var decoder converter.Decoder
	var encoder converter.Encoder

	switch params.inputType {
	case jpgType:
		fallthrough
	case jpegType:
		decoder = &converter.JPEG{}
	case pngType:
		decoder = &converter.PNG{}
	case gifType:
		decoder = &converter.GIF{}
	}
	switch params.outputType {
	case jpgType:
		fallthrough
	case jpegType:
		encoder = &converter.JPEG{Quality: params.jpgQuality}
	case pngType:
		encoder = &converter.PNG{}
	case gifType:
		encoder = &converter.GIF{NumColors: params.gifNumColors}
	}

	converter := converter.Converter{Decoder: decoder, Encoder: encoder}
	for _, path := range params.path {
		if err := converter.Run(path); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	params, err := parse()
	if err != nil {
		log.Fatal(err)
	}
	if err = run(*params); err != nil {
		log.Fatal(err)
	}
}
