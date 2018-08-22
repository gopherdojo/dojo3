package conversion

import (
	"image"
	"image/gif"
	"io"
	"path/filepath"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
)

// Gif https://en.wikipedia.org/wiki/GIF
type Gif struct {
	Options *gif.Options
}

// Encode encodes the specified file to GIF
func (g *Gif) Encode(w io.Writer, img image.Image) error {
	return gif.Encode(w, img, g.Options)
}

// Decode decodes the specified GIF file
func (g *Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

// Extname returns "gif"
func (g *Gif) Extname() string {
	return "gif"
}

// IsDecodable returns whether the file content is GIF
func (g *Gif) IsDecodable(rs io.ReadSeeker) bool {
	return fileutil.StartsContentsWith(rs, []byte("GIF87a")) || fileutil.StartsContentsWith(rs, []byte("GIF89a"))
}

// HasProcessableExtname returns whether the specified path has ".gif"
func (g *Gif) HasProcessableExtname(path string) bool {
	return filepath.Ext(path) == ".gif"
}
