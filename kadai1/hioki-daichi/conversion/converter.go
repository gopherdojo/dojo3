package conversion

import (
	"errors"
	"image"
	"io"
	"os"
	"path/filepath"
)

// Encoder can encode images.
type Encoder interface {
	Encode(*os.File, image.Image) error
	Extname() string
}

// Decoder can decode images.
type Decoder interface {
	Decode(*os.File) (image.Image, error)
	HasProcessableExtname(string) bool
	IsDecodable(*os.File) bool
}

// Converter has Encoder and Decoder.
type Converter struct {
	Encoder   Encoder
	Decoder   Decoder
	OutStream io.Writer
}

// Convert converts the specified path from own Decoder to own Encoder.
func (c *Converter) Convert(path string, force bool) (*os.File, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	img, err := c.Decoder.Decode(fp)
	if err != nil {
		return nil, err
	}

	dstPath := path[:len(path)-len(filepath.Ext(path))] + "." + c.Encoder.Extname()

	if !force {
		_, err := os.Stat(dstPath)
		if err == nil {
			return nil, errors.New("File already exists: " + dstPath)
		}
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}

	err = c.Encoder.Encode(dstFile, img)
	if err != nil {
		return nil, err
	}

	return dstFile, nil
}
