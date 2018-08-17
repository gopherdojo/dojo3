package conversion

import (
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
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

// IsDecodable returns whether the file content is PNG
func (p *Png) IsDecodable(fp *os.File) bool {
	return fileutil.StartsContentsWith(fp, []uint8{137, 80, 78, 71, 13, 10, 26, 10})
}

// HasProcessableExtname returns whether the specified path has ".png"
func (p *Png) HasProcessableExtname(path string) bool {
	return filepath.Ext(path) == ".png"
}
