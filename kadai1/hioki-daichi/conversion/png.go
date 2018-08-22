package conversion

import (
	"image"
	"image/png"
	"io"
	"path/filepath"
)

// Png https://en.wikipedia.org/wiki/Portable_Network_Graphics
type Png struct {
	Encoder *png.Encoder
}

// Encode encodes the specified file to PNG
func (p *Png) Encode(w io.Writer, img image.Image) error {
	return p.Encoder.Encode(w, img)
}

// Decode decodes the specified PNG file
func (p *Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

// Extname returns "png"
func (p *Png) Extname() string {
	return "png"
}

// MagicBytesSlice returns the magic bytes slice of PNG
func (p *Png) MagicBytesSlice() [][]byte {
	return [][]byte{[]byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")}
}

// HasProcessableExtname returns whether the specified path has ".png"
func (p *Png) HasProcessableExtname(path string) bool {
	return filepath.Ext(path) == ".png"
}
