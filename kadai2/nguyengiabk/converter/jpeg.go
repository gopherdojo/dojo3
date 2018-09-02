package converter

import (
	"image"
	"image/jpeg"
	"io"
	"path/filepath"
)

// JPEG wraps image/jpeg functions to implement Decoder and Encoder interface
type JPEG struct {
	Quality int
}

// Decode reads data from io.Reader and returns an image
func (image *JPEG) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// CheckExt checks extension of file-to-be-decoded is valid or not
func (image *JPEG) CheckExt(path string) bool {
	return filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg"
}

// Encode write data from an image to io.Writer
func (image *JPEG) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: image.Quality})
}

// GetExt returns file extention of GIF format
func (image *JPEG) GetExt() string {
	return ".jpg"
}
