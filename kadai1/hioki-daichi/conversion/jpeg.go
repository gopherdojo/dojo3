package conversion

import (
	"image"
	"image/jpeg"
	"io"
	"path/filepath"
)

// Jpeg https://en.wikipedia.org/wiki/JPEG
type Jpeg struct {
	Options *jpeg.Options
}

// Encode encodes the specified file to JPEG
func (j *Jpeg) Encode(w io.Writer, img image.Image) error {
	return jpeg.Encode(w, img, j.Options)
}

// Decode decodes the specified JPEG file
func (j *Jpeg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

// Extname returns "jpg"
func (j *Jpeg) Extname() string {
	return "jpg"
}

// MagicBytesSlice returns the magic bytes slice of JPEG
func (j *Jpeg) MagicBytesSlice() [][]byte {
	return [][]byte{[]byte("\xFF\xD8\xFF")}
}

// HasProcessableExtname returns whether the specified path has ".jpg" or ".jpeg"
func (j *Jpeg) HasProcessableExtname(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".jpg" || ext == ".jpeg"
}
