package conversion

import (
	"image"
	"image/jpeg"
	"io"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
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

// IsDecodable returns whether the file content is JPEG
func (j *Jpeg) IsDecodable(rs io.ReadSeeker) bool {
	return fileutil.StartsContentsWith(rs, []uint8{255, 216, 255})
}

// HasProcessableExtname returns whether the specified path has ".jpg" or ".jpeg"
func (j *Jpeg) HasProcessableExtname(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".jpg" || ext == ".jpeg"
}
