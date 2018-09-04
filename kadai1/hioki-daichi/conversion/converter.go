/*
Package conversion has the necessary processing to convert.

File is divided for each file format.
*/
package conversion

import (
	"errors"
	"image"
	"io"
	"os"
	"path/filepath"
)

// Converter represents encodable and decodable.
type Converter struct {
	Encoder Encoder
	Decoder Decoder
}

// Encoder configures encode-needed settings.
type Encoder interface {
	Encode(io.Writer, image.Image) error
	Extname() string
}

// Decoder configures decode-needed settings.
type Decoder interface {
	Decode(io.Reader) (image.Image, error)
	HasProcessableExtname(string) bool
	MagicBytesSlice() [][]byte
}

// Convert opens the file, decodes it, creates a file with a different extension, and writes the encoded result.
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
		_, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL, 0)
		if os.IsExist(err) {
			return nil, errors.New("File already exists: " + dstPath)
		}
		os.Remove(dstPath)
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
