package converter

import (
	"image"
	"image/gif"
	"io"
)

// Gif is gif format struct
type Gif struct{}

// Encode is encoding image without gif
func (g *Gif) Encode(file io.Writer, img image.Image) error {
	return gif.Encode(file, img, nil)
}

// Decode is decoding image without gif
func (g *Gif) Decode(file io.Reader) (image.Image, error) {
	return gif.Decode(file)
}

// GetExt return gif
func (g *Gif) GetExt() string {
	return "gif"
}
