package jpg

import (
	"image"
	"io"

	"image/jpeg"

	"../../convert"
)

// Jpg implements convert.Converter
type Jpg struct {
}

func init() {
	convert.Register("jpg", Jpg{})
}

// Decode returns error
func (j Jpg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Encode returns error
func (j Jpg) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}
