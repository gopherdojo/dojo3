// Package convert is convert image extension.
package convert

import (
	"errors"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
)

// Convert is interface that has Convert function.
type Convert interface {
	Convert(path string) error
}

type convert struct {
	option *option.Option
}

// NewConvert is Convert interface constructor.
func NewConvert(option *option.Option) Convert {
	return &convert{option: option}
}

// Convert image.
func (c *convert) Convert(path string) error {
	fromConverter, err := getConverter(c.option.FromExtension)
	if err != nil {
		return err
	}
	image, err := decode(path, fromConverter)
	if err != nil {
		return err
	}

	toConverter, err := getConverter(c.option.ToExtension)
	if err != nil {
		return err
	}
	if err := encode(image, toConverter, path, c); err != nil {
		return err
	}
	fmt.Print(path + " -> to " + c.option.ToExtension)
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

func encode(image image.Image, converter Converter, path string, c *convert) error {
	output, err := os.Create(strings.TrimSuffix(path, c.option.FromExtension) + c.option.ToExtension)
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
