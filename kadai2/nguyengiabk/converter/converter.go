// Package converter provides functions and type to convert images inside a directory to other format
package converter

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)

// Decoder defines interface that has image decode function
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
	CheckExt(path string) bool
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

// SupportedTypes defines file types that can be converted
var SupportedTypes = map[string]bool{"jpg": true, "jpeg": true, "png": true, "gif": true}

// Run recursively processes all files inside a directory and converts all images has specified type
func (converter *Converter) Run(path string, w io.Writer) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
			// just print log and continue for other files
			if err := converter.processFile(path, info); err != nil {
				fmt.Fprintf(w, "%s, continue processing\n", err.Error())
			}
			return nil
		}
	})
	return err
}

func (converter *Converter) processFile(path string, info os.FileInfo) error {
	if !converter.Decoder.CheckExt(path) {
		return nil
	}
	img, err := converter.readImage(path)
	if err != nil {
		return err
	}
	outputFilePath := path[0:len(path)-len(filepath.Ext(path))] + converter.Encoder.GetExt()
	return converter.writeImage(outputFilePath, img)
}

func (converter *Converter) readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	image, err := converter.Decoder.Decode(file)
	if err != nil {
		return nil, errors.New("Cannot decode file " + path)
	}
	return image, nil
}

func (converter *Converter) writeImage(path string, image image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return converter.Encoder.Encode(file, image)
}
