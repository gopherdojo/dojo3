package converter

import (
	"image"
	"image/png"
	"io"
)

// Png is png format struct
type Png struct{}

// Encode is encoding image without png
func (p *Png) Encode(file io.Writer, img image.Image) error {
	return png.Encode(file, img)
}

// Decode is decoding image without png
func (p *Png) Decode(file io.Reader) (image.Image, error) {
	return png.Decode(file)
}

// GetExt return png
func (p *Png) GetExt() string {
	return "png"
}
