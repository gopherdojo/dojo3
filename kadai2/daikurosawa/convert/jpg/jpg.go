// Package jpg is encode and decode to image.
package jpg

import (
	"image"
	"io"

	"image/jpeg"

	"github.com/gopherdojo/dojo3/kadai2/daikurosawa/di"
)

// Jpg implements convert.Converter.
type Jpg struct{}

func init() {
	di.Register("jpg", Jpg{})
}

// Decode returns image and error.
func (Jpg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Encode return error.
func (Jpg) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}
