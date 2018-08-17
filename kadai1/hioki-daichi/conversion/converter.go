package conversion

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/cliopt"
	"github.com/hioki-daichi/myfileutil"
)

// Encoder can encode images.
type Encoder interface {
	Encode(*os.File, image.Image) error
	Extname() string
}

// Decoder can decode images.
type Decoder interface {
	Decode(*os.File) (image.Image, error)
	IsDecodable(*os.File) bool
}

// Converter has Encoder and Decoder.
type Converter struct {
	Encoder   Encoder
	Decoder   Decoder
	OutStream io.Writer
}

// Convert converts the specified path from own Decoder to own Encoder.
func (c *Converter) Convert(path string) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	img, err := c.Decoder.Decode(fp)
	if err != nil {
		return err
	}

	dstPath := path[:len(path)-len(filepath.Ext(path))] + "." + c.Encoder.Extname()

	if !cliopt.Force && myfileutil.Exists(dstPath) {
		return errors.New("File already exists: " + dstPath)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	err = c.Encoder.Encode(dstFile, img)
	if err != nil {
		return err
	}

	if cliopt.Verbose {
		fmt.Fprintf(c.OutStream, "Converted: %q\n", dstPath)
	}

	return nil
}
