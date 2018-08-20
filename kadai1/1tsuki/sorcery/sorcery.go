// sorcery converts an image file with specific extension into another, contained in specified directory.
// Use Sorcery method to create new instance, and then call Exec method with options.
// See imgExt for supported image file extensions.

package sorcery

import (
	"fmt"
	"io"
	"os"
)

var (
	// ErrUnsupportedExtension is returned when invalid imgExt was passed
	ErrUnsupportedExtension = fmt.Errorf("unsupported extension specified")
)

// sorcery convert image formats, and print result to writer
type sorcery struct {
	writer io.Writer
}

// Sorcery creates instance of sorcery
func Sorcery(writer io.Writer) *sorcery {
	return &sorcery{writer}
}

// Exec is a method to start image format conversions
// Be aware that it automatically search subdirectories
func (s *sorcery) Exec(from imgExt, to imgExt, dir string) error {
	if !from.isValid() || !to.isValid() {
		return ErrUnsupportedExtension
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}

	return scan(dir, from, func(in string) error {
		c := &converter{Path: in}
		out, err := c.convert(to)
		if err != nil {
			return err
		}

		fmt.Fprintf(s.writer, "file converted: %s to %s\n", in, out)
		return nil
	})
}
