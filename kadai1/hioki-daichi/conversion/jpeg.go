package conversion

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/hioki-daichi/myfileutil"
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
	return myfileutil.IsJpeg(fp)
}
