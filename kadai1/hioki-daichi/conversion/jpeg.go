package conversion

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
)

// Jpeg https://en.wikipedia.org/wiki/JPEG
type Jpeg struct{}

// Encode encodes the specified file to JPEG
func (j *Jpeg) Encode(fp *os.File, img image.Image) error {
	return jpeg.Encode(fp, img, &jpeg.Options{Quality: 100})
}

// Decode decodes the specified JPEG file
func (j *Jpeg) Decode(fp *os.File) (image.Image, error) {
	return jpeg.Decode(fp)
}

// Extname returns "jpg"
func (j *Jpeg) Extname() string {
	return "jpg"
}

// IsDecodable returns whether the file content is JPEG
func (j *Jpeg) IsDecodable(fp *os.File) bool {
	return fileutil.StartsContentsWith(fp, []uint8{255, 216, 255})
}

// HasProcessableExtname returns whether the specified path has ".jpg" or ".jpeg"
func (j *Jpeg) HasProcessableExtname(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".jpg" || ext == ".jpeg"
}
