package converter

import (
	"image"
	"image/jpeg"
	"io"
)

// Jpg is jpg format struct
type Jpg struct{}

// Encode is encoding image without jpg
func (j *Jpg) Encode(file io.Writer, img image.Image) error {
	return jpeg.Encode(file, img, nil)
}

// Decode is decoding image without jpg
func (j *Jpg) Decode(file io.Reader) (image.Image, error) {
	return jpeg.Decode(file)
}

// GetExt return jpg
func (j *Jpg) GetExt() string {
	return "jpg"
}
