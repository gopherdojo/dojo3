package converter

import (
	"image"
	"image/png"
	"io"
)

// PNG wraps image/png functions to implement Decoder and Encoder interface
type PNG struct{}

// Decode reads data from io.Reader and returns an image
func (image *PNG) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode write data from an image to io.Writer
func (image *PNG) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

// GetExt returns file extention of GIF format
func (image *PNG) GetExt() string {
	return ".png"
}
