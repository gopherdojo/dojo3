package converter

import (
	"image"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

// Encoder interface
type Encoder interface {
	Encode(io.Writer, image.Image) error
	GetExt() string
}

// Decoder interface
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

// Converter struct
type Converter struct {
	Encoder Encoder
	Decoder Decoder
}

// Convert image functon
func (c *Converter) Convert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := c.Decoder.Decode(file)
	if err != nil {
		return err
	}

	fileName, err := c.getFileName(path)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = c.Encoder.Encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

func (c *Converter) getFileName(path string) (string, error) {
	ext := c.Encoder.GetExt()
	if path == "" {
		return "", errors.New("path must not be empty")
	}

	imageExt := filepath.Ext(path)
	rep := regexp.MustCompile(imageExt + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	if imageExt == ext {
		return "", errors.New("input and output ext is same")
	}
	return name + ext, nil
}
