package png

import (
	"image"
	"io"

	"image/png"

	"../../convert"
)

// Png implements convert.Converter
type Png struct {
}

func init() {
	convert.Register("png", Png{})
}

// Decode returns error
func (p Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode returns error
func (p Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
