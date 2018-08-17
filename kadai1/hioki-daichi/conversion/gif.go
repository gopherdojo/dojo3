package conversion

import (
	"image"
	"image/gif"
	"os"

	"github.com/gopherdojo/dojo3/kadai1/hioki-daichi/fileutil"
)

// Gif https://en.wikipedia.org/wiki/GIF
type Gif struct{}

// Encode encodes the specified file to GIF
func (g *Gif) Encode(fp *os.File, img image.Image) error {
	return gif.Encode(fp, img, &gif.Options{NumColors: 256})
}

// Decode decodes the specified GIF file
func (g *Gif) Decode(fp *os.File) (image.Image, error) {
	return gif.Decode(fp)
}

// Extname returns "gif"
func (g *Gif) Extname() string {
	return "gif"
}

// IsDecodable returns whether the file content is GIF
func (g *Gif) IsDecodable(fp *os.File) bool {
	return fileutil.StartsContentsWith(fp, []uint8{71, 73, 70, 56, 55, 97}) || fileutil.StartsContentsWith(fp, []uint8{71, 73, 70, 56, 57, 97})
}
