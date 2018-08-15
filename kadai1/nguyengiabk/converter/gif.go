package converter

import (
	"image"
	"image/gif"
	"io"
)

// GIF wraps image/gif functions to implement Decoder and Encoder interface
type GIF struct {
	NumColors int
}

// Decode reads data from io.Reader and returns an image
func (image *GIF) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// Encode write data from an image to io.Writer
func (image *GIF) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, &gif.Options{NumColors: image.NumColors})
}

// GetExt returns file extention of GIF format
func (image *GIF) GetExt() string {
	return ".gif"
}
