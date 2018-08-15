// Package converter provides functions and type to convert images inside a directory to other format
package converter

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)

// Decoder defines interface that has image decode function
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Encoder defines interface that has image encode function. It also has a function to return file extension of this image type.
type Encoder interface {
	Encode(w io.Writer, m image.Image) error
	GetExt() string
}

// Converter converts images inside a directory from input type to output type
type Converter struct {
	Decoder Decoder
	Encoder Encoder
}

// Run recursively processes all files inside a directory and converts all images has specified type
func (converter *Converter) Run(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
			return converter.processFile(path, info)
		}
	})
	return err
}

func (converter *Converter) processFile(path string, info os.FileInfo) error {
	img, err := converter.readImage(path)
	if err != nil {
		return err
	}
	// in case decode fail
	if img == nil {
		return nil
	}
	outputFilePath := path[0:len(path)-len(filepath.Ext(path))] + converter.Encoder.GetExt()
	return converter.writeImage(outputFilePath, img)
}

func (converter *Converter) readImage(path string) (image.Image, error) {
	fmt.Println("Reading file: " + path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, err := converter.Decoder.Decode(file)
	// if decode fail, just print log and continue for other files
	if err != nil {
		fmt.Println(err)
	}
	return image, nil
}

func (converter *Converter) writeImage(path string, image image.Image) error {
	fmt.Println("Writing file: " + path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return converter.Encoder.Encode(file, image)
}
