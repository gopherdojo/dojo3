package gif

import (
	"image"
	"image/gif"
	"io"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
)

// Gif implements convert.Converter
type Gif struct {
}

func init() {
	convert.Register("gif", Gif{})
}

// Decode returns error
func (g Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// Encode returns error
func (g Gif) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
