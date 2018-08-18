package png

import (
	"image"
	"io"

	"image/png"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
)

// Png implements convert.Converter
type Png struct {
}

func init() {
	convert.Register("png", Png{})
}

// Decode returns error
func (Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Encode returns error
func (Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
