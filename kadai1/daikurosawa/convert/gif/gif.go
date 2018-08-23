// Package gif is encode and decode to image.
package gif

import (
	"image"
	"image/gif"
	"io"

	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/di"
)

// Gif implements convert.Converter
type Gif struct{}

func init() {
	di.Register("gif", Gif{})
}

// Decode returns image and error.
func (Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// Encode return error.
func (Gif) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
