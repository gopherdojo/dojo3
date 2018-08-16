package conversion

import (
	"image"
	"image/png"
	"os"

	"github.com/hioki-daichi/myfileutil"
)

// Png https://en.wikipedia.org/wiki/Portable_Network_Graphics
type Png struct{}

// Encode encodes the specified file to PNG
func (p *Png) Encode(fp *os.File, img image.Image) error {
	return png.Encode(fp, img)
}

// Decode decodes the specified PNG file
func (p *Png) Decode(fp *os.File) (image.Image, error) {
	return png.Decode(fp)
}

// Extname returns "png"
func (p *Png) Extname() string {
	return "png"
}

// IsValid returns whether the file content is PNG
func (p *Png) IsValid(fp *os.File) bool {
	return myfileutil.IsPng(fp)
}
