package convert

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"strings"
)

// Converter Convert is interface that has Decode and Encode function.
type Converter interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
}

var converts = map[string]Converter{}

// Register sets command to commands
func Register(key string, convert Converter) {
	converts[key] = convert
}

// Convert has command line options.
type Convert struct {
	Path          string
	FromExtension string
	ToExtension   string
}

// Convert image.
func (c *Convert) Convert() error {
	convert, err := getConverter(c.FromExtension)
	if err != nil {
		return err
	}
	image, err := decode(c.Path, convert)
	if err != nil {
		return err
	}

	if err := encode(image, convert, c); err != nil {
		return err
	}
	fmt.Print(c.Path + " -> to " + c.ToExtension)
	fmt.Printf("\x1b[32m%s\x1b[0m", "\tDONE\n")
	return nil
}

func decode(path string, converter Converter) (image.Image, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return converter.Decode(file)
}

func encode(image image.Image, converter Converter, c *Convert) error {
	output, err := os.Create(strings.TrimSuffix(c.Path, c.FromExtension) + c.ToExtension)
	if err != nil {
		return err
	}
	defer output.Close()
	return converter.Encode(output, image)
}

func getConverter(extension string) (Converter, error) {
	convert, ok := converts[extension]
	if ok {
		return convert, nil
	} else {
		return nil, errors.New("unsupported extension")
	}
}
